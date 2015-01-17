package test

import (
	"log"

	"github.com/martini-contrib/render"

	"pixelrelay/db"
)

func Uploader(r render.Render, dbh *db.Dbh) {
	upper := dbh.GetAllUploaders()
	log.Println(upper)

	r.JSON(200, upper)
}
