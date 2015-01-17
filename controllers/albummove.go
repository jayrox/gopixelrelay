/*
 Route:  /album/create

 Method: POST

 Return:
  - JSON
  - Params
   - Status string
*/

package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
)

type Data struct {
	Md MoveData `json:"data"`
}

type MoveData struct {
	AlbumId int64       `json:"album"`
	Images  []MoveImage `json:"images"`
}

type MoveImage struct {
	Name string `json:"Name"`
	Id   string `json:"Id"`
}

func AlbumMove(req *http.Request, r render.Render, su models.User, dbh *db.Dbh, p *models.Page) {
	// Default status
	var status string = "Permission Denied"
	var code int = 401

	var reader io.Reader = req.Body
	b, e := ioutil.ReadAll(reader)
	if e != nil {
		log.Println(e)
	}

	var f Data
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Println(err)
	}
	log.Println(f)

	albumId := f.Md.AlbumId
	album := dbh.GetAlbumById(albumId)
	log.Println("album: ", albumId)

	if album.User == su.Id {
		status = "Success"
		code = 200

		var image models.Image
		for _, f := range f.Md.Images {
			imageId, err := strconv.ParseInt(f.Id, 10, 10)
			if err != nil {
				log.Println("invalid image id: ", f.Id)
			}
			log.Println("Moving: ", f.Name, imageId)

			image = dbh.GetImageById(imageId)
			if image.User != su.Id {
				status = "Permission Denied"
				code = 401
				continue
			}
			fromId := image.AlbumId
			image.AlbumId = albumId
			_ = dbh.UpdateImage(image)
			log.Printf("Moved %s from: %d to: %d status: %s code: %d\n", f.Name, fromId, albumId, status, code)
		}

	}
	p.Data = status
	encoder.Render("json", code, "", p, r)
}
