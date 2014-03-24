package test

import (
	"log"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/models"
)

func ListDB(args martini.Params, su models.User, r render.Render, p *models.Page, dbh *db.Dbh) {

	images := dbh.GetAllImages()

	var album models.Album
	for _, image := range images {
		album = dbh.GetAlbum(image.Album)
		image.AlbumId = album.Id
		nImage := dbh.UpdateImage(image)
		log.Println(nImage)
	}

	return
}
