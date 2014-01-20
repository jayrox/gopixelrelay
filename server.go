package main

import (
	"fmt"
	"github.com/3d0c/martini-contrib/config"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"log"
	"net/http"
	"pixelrelay/controllers"
	"pixelrelay/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

	sqlConnection := fmt.Sprintf("%s:%s@%s/%s?clientFoundRows=true&charset=UTF8", utils.DbCfg.User(), utils.DbCfg.Pass(), utils.DbCfg.Host(), utils.DbCfg.Name())
	db, err := sql.Open("mysql", sqlConnection)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	m.Map(db)
	
	//rows, err := db.Query("show databases;")
	//if err != nil {
	//	fmt.Println("query err: ", err)
	//}
	//fmt.Println("row: ", rows)
	
	
	m.Use(martini.Static("static"))

	m.Get("/", controllers.Index)

	m.Get("/i/:name", controllers.Image)
	m.Get("/t/:name", controllers.Thumb)

	m.Get("/list", controllers.List)

	m.Post("/up", utils.Verify, controllers.UploadImage)

	log.Printf("Listening for connections on %s\n", utils.AppCfg.ListenOn())

	if err := http.ListenAndServe(utils.AppCfg.ListenOn(), m); err != nil {
		log.Fatal(err)
	}
}
