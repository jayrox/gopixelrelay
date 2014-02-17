package controllers

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"pixelrelay/db"
)

func TagImage(args martini.Params, r render.Render) {
	tag := args["name"]
	image := args["image"]
	d := db.InitDB()
	imagetag := db.TagImage(&d, tag, image)

	fmt.Println(imagetag)
	//r.HTML(200, "image_link", imageLinks)
}
