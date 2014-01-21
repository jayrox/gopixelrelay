package main

import (
	//"fmt"
	"github.com/3d0c/martini-contrib/config"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"log"
	"net/http"
	"pixelrelay/controllers"
	//"pixelrelay/models"
	"pixelrelay/utils"
	"pixelrelay/db"

	//"reflect"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	config.Init("./pixelrelay.json")
	config.LoadInto(utils.AppCfg)
	config.LoadInto(utils.DbCfg)
	config.LoadInto(utils.ImageCfg)
}

func main() {

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory: "templates", // Specify what path to load the templates from.
		Layout:    "layout",    // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Charset:   "UTF-8",     // Sets encoding for json and html content-types.
	}))

	m.Use(martini.Static("static"))
	
	d := db.InitDB()
	db.MigrateDB(&d)
	
	// Set up routes
	m.Get("/", controllers.Index)
	m.Get("/i/:name", controllers.Image)
	m.Get("/t/:name", controllers.Thumb)
	m.Get("/list", controllers.List)
	m.Post("/up", utils.Verify, controllers.UploadImage)


	// Start server and begin listening for requests
	log.Printf("Listening for connections on %s\n", utils.AppCfg.ListenOn())

	if err := http.ListenAndServe(utils.AppCfg.ListenOn(), m); err != nil {
		log.Fatal(err)
	}
}