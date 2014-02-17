package main

import (
	"flag"
	"fmt"
	"github.com/3d0c/martini-contrib/config"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"log"
	"net/http"
	"pixelrelay/controllers"
	"pixelrelay/db"
	"pixelrelay/forms"
	"pixelrelay/middleware"
	//"pixelrelay/models"
	"pixelrelay/utils"
)

var flagInit *bool
var flagMigrate *bool

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

	// Create sessions cookie store
	store := sessions.NewCookieStore([]byte(utils.AppCfg.SecretKey()))
    m.Use(sessions.Sessions("pixelrelay", store))
	
	// Setup render options
	m.Use(render.Renderer(render.Options{
		Directory: "templates", // Specify what path to load the templates from.
		Layout:    "layout",    // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Charset:   "UTF-8",     // Sets encoding for json and html content-types.
	}))

	// Setup static file handling
	m.Use(martini.Static("static"))

	// Init DB
	d := db.InitDB()

	// Set up routes
	m.Get("/", controllers.Index)
	m.Get("/i/:name", middleware.VerifyFile, controllers.Image)
	m.Get("/t/:name", middleware.VerifyFile, controllers.Thumb)
	m.Get("/list", controllers.List)
	m.Get("/albums", controllers.Albums)
	m.Get("/album/:name", controllers.Album)
	m.Get("/auth/:password", controllers.Auth)
	m.Get("/tags", controllers.Tags)
	m.Get("/tag/:name", controllers.Tagged)
	m.Get("/tag/:name/:image", controllers.TagImage)

	//binding.MaxMemory = int64(1024 * 1024 * 30)
	//binding.Bind(models.PostImage{}),
	m.Post("/up", middleware.Verify, controllers.UploadImage)

	// Start server and begin listening for requests
	log.Printf("Listening for connections on \x1b[32;1m%s\x1b[0m\n", utils.AppCfg.ListenOn())

	go http.ListenAndServe(utils.AppCfg.ListenOn(), m)

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
		fmt.Println("\x1b[31;1mInitial Setup flag (-init) has been set to\x1b[0m \x1b[32;1mTRUE\x1b[0m")
		fmt.Println("\x1b[31;1mOnce setup is complete please restart server with this flag disabled.\x1b[0m")

		// Add default tables
		db.AddTables(&d)

		su := martini.Classic()
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
		db.MigrateDB(&d)
	}
	
	select {}
}
