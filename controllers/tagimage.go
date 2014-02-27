package controllers

import (
	"fmt"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
)

func TagImage(args martini.Params, r render.Render, dbh *db.Dbh) {
	tag := args["name"]
	image := args["image"]

	imagetag := dbh.TagImage(tag, image)

	fmt.Println(imagetag)
	//r.HTML(200, "image_link", imageLinks)
}
