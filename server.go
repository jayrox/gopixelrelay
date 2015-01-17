package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/3d0c/martini-contrib/config"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/controllers"
	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/forms"
	"pixelrelay/middleware"
	"pixelrelay/models"
	"pixelrelay/test"
	"pixelrelay/utils"
)

var (
	flagInit    *bool
	flagMigrate *bool
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	config.Init("./pixelrelay.json")
	config.LoadInto(utils.AppCfg)
	config.LoadInto(utils.DbCfg)
	config.LoadInto(utils.ImageCfg)

	flagInit = flag.Bool("init", false, "Initial Setup")
	flagMigrate = flag.Bool("migrate", false, "Migrate Database Changes")
	flag.Parse()
}

func main() {
	m := martini.Classic()

	// Gzip all
	m.Use(gzip.All())

	// Create sessions cookie store
	store := sessions.NewCookieStore([]byte(utils.AppCfg.SecretKey()))
	m.Use(sessions.Sessions("pixelrelay", store))

	// Setup render options
	m.Use(render.Renderer(render.Options{
		Directory: "templates",                            // Specify what path to load the templates from.
		Layout:    "layout",                               // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Charset:   "UTF-8",                                // Sets encoding for json and html content-types.
		Delims:    render.Delims{Left: "{[", Right: "]}"}, // Sets delimiters to the specified strings.
	}))

	// Setup DB
	dbh := db.Init(&db.Dbh{})
	m.Map(dbh)

	// Setup static file handling
	opts := martini.StaticOptions{SkipLogging: false, Expires: utils.ExpiresHeader}
	m.Use(martini.Static("static", opts))

	// Auth user and assign to session
	m.Use(middleware.UserAuth(models.User{}, dbh))

	// Encoder for .html or .json encoding
	m.Use(encoder.MapEncoder)

	// Setup Page
	p := models.InitPage(&models.Page{})
	m.Map(p)

	// Set up routes
	m.Get("/", controllers.Index)

	// Images
	m.Get("/image/:name", middleware.VerifyFile, controllers.ImagePage)
	m.Get("/o/:name", middleware.VerifyFile, controllers.Image)
	m.Get("/i/:name", middleware.VerifyFile, controllers.ImageModified)
	m.Get("/t/:name", middleware.VerifyFile, controllers.Thumb)
	m.Get("/image/trash/:album/:name", middleware.AuthRequired, controllers.ImageTrash)
	m.Get("/image/recover/:album/:name", middleware.AuthRequired, controllers.ImageRecover)

	// Albums
	m.Get("/albums", controllers.Albums)
	m.Get("/album/:name", controllers.Album)
	m.Get("/album/:name/qr", controllers.QR)
	m.Get("/:user/albums", controllers.Albums)
	m.Get("/:user/album/:name", controllers.Album)
	m.Get("/album/:name/private/:state", controllers.AlbumPrivate)

	m.Get("/album/:name/:key", controllers.Album)

	m.Post("/album/create", middleware.AuthRequired, controllers.AlbumCreate)
	m.Post("/album/update", middleware.AuthRequired, controllers.AlbumUpdate)
	m.Post("/album/delete/:name", middleware.AuthRequired, controllers.AlbumDelete)
	m.Post("/album/recover/:name", middleware.AuthRequired, controllers.AlbumRecover)
	m.Post("/album/move", middleware.AuthRequired, controllers.AlbumMove)

	// Tag
	m.Get("/tags", controllers.Tags)
	m.Get("/tag/:tag", controllers.Tagged)
	m.Post("/tag", middleware.AuthRequired, controllers.TagImage)

	// Auth
	m.Get("/login", controllers.Login)
	m.Post("/login", binding.Bind(forms.Login{}), binding.ErrorHandler, controllers.LoginPost)
	m.Get("/logout", controllers.Logout)

	// Upload
	m.Post("/up", middleware.Verify, binding.MultipartForm(models.ImageUpload{}), controllers.UploadImage)

	// 404
	m.NotFound(controllers.NotFound)

	// Account
	m.Get("/account", middleware.AuthRequired, controllers.Account)

	// Profile
	m.Get("/profile/:name", controllers.Profile)

	// Testing
	m.Get("/test/hash", test.Hash)
	m.Get("/test/hash/:user_id/:image_id", test.Hash)
	m.Get("/test/list", middleware.AuthRequired, test.List)
	m.Get("/test/listdb", middleware.AuthRequired, test.ListDB)
	m.Get("/test/listt", middleware.AuthRequired, test.ListThumb)
	m.Get("/test/uploader", middleware.AuthRequired, test.Uploader)

	// Start server and begin listening for HTTP requests
	// Watch for non-https requests and redirect to https
	// If request method is POST and path is /up
	//   discard file and redirect to correct https route
	log.Printf("Listening for \x1b[32;1mHTTP\x1b[0m connections on \x1b[32;1m%s\x1b[0m\n", utils.AppCfg.ListenOn())
	go http.ListenAndServe(utils.AppCfg.ListenOn(), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		url := strings.Join([]string{utils.AppCfg.Url(), urlPath}, "")

		// Check if path is for the Upload controller.
		// Redirect client to HTTPS end point
		if urlPath == "/up" && r.Method == "POST" {
			ur := models.UploadResult{
				Code:   http.StatusFound,
				Result: "HTTPS Required",
				Name:   url,
			}
			log.Println(ur.String())

			// Move uploaded file to DevNull
			file, _, _ := r.FormFile("uploaded_file")
			defer file.Close()
			out, err := os.Create(os.DevNull)
			if err != nil {
				log.Println(err)
			}
			defer out.Close()
			_, err = io.Copy(out, file)
			if err != nil {
				log.Println(err)
			}

			// Respond with JSON that server requires HTTPS and correct path.
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("Expires", utils.ExpiresHeader())
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, ur)
			return
		}

		log.Println("Redirect: ", url)
		http.Redirect(w, r, url, http.StatusFound)
	}))

	// Start server and begin listening for TLS requests
	log.Printf("Listening for \x1b[32;1mTLS\x1b[0m connections on \x1b[32;1m%s\x1b[0m\n", utils.AppCfg.TLSListenOn())
	go http.ListenAndServeTLS(utils.AppCfg.TLSListenOn(), "./tls/cert.pem", "./tls/key.pem", m)

	/******************************************
	*	INITIAL SETUP
	*
	*   Creates the initial tables
	*   Populates the default admin user
	*
	*   Potential security risks are present
	*   if this mode is left running.
	*	restart server with the "-init" flag
	*   unset.
	*
	*   usage: -init
	 */
	if *flagInit {
		fmt.Println("\x1b[31;1mInitial Setup flag (-init) has been set to \x1b[32;1mTRUE\x1b[0m")
		fmt.Println("\x1b[31;1mOnce setup is complete please restart server with this flag disabled.\x1b[0m")

		// Add default tables
		dbh.AddTables()

		su := martini.Classic()

		store := sessions.NewCookieStore([]byte(utils.AppCfg.SecretKey()))
		su.Use(sessions.Sessions("pixelrelay", store))
		su.Use(render.Renderer(render.Options{
			Directory: "templates", // Specify what path to load the templates from.
			Layout:    "layout",    // Specify a layout template. Layouts can call {{ yield }} to render the current template.
			Charset:   "UTF-8",     // Sets encoding for json and html content-types.
		}))
		su.Get("/setup", controllers.SetupAdmin)
		su.Post("/setup", binding.Bind(forms.SetupAdmin{}), binding.ErrorHandler, controllers.SetupAdminPost)
		// Start server and begin listening for requests
		log.Printf("Listening for connections on \x1b[32;1m%s\x1b[0m\n", utils.AppCfg.ListenOnSetup())

		go http.ListenAndServe(utils.AppCfg.ListenOnSetup(), su)
	}

	/******************************************
	*	MIGRATE DATABASE UPDATES
	*
	*   Migrates changes to database tables
	*
	*   You should backup the database before
	*   migrating. As there is a potential risk
	*   of data loss
	*
	*   usage: -migrate
	 */
	if *flagMigrate {
		dbh.MigrateDB()
	}

	select {}
}
