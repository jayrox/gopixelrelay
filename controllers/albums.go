package controllers

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"pixelrelay/db"
	"pixelrelay/models"
	//"reflect"
)

func Albums(args martini.Params, r render.Render) {

	d := db.InitDB()
		
	albums := db.GetAllAlbums(&d)
	
	var albumList []models.AlbumList
	
	for _, f := range albums {
		i := db.FirstImage(&d, f.Name)
		fmt.Println("i: ", i[0].Name)
		albumList = append(albumList, models.AlbumList{Name: f.Name, Poster: i[0].Name})
	}

	r.HTML(200, "albums", albumList)
}
