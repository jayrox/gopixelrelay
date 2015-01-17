/*
 Route:  /album/update

 Method: POST

 Return:
  - JSON
  - Params:
   - status string
*/

package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
)

func AlbumUpdate(req *http.Request, r render.Render, su models.User, dbh *db.Dbh, p *models.Page) {
	p.SetUser(su)

	var reader io.Reader = req.Body
	b, e := ioutil.ReadAll(reader)
	if e != nil {
		log.Println(e)
	}

	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Println(err)
	}

	m := f.(map[string]interface{})
	md := m["data"].(map[string]interface{})

	// Default status
	var status string = "Permission Denied"
	var code int = 401

	mi := int64(md["Id"].(float64))
	mo := int64(md["Owner"].(float64))
	name := md["Name"].(string)

	album := dbh.GetAlbumById(mi)

	if album.Id == mi && album.User == mo && su.Id == mo {
		mAlbum := models.Album{
			Id:          mi,
			Description: md["Description"].(string),
			Name:        md["Name"].(string),
			User:        mo,
			Poster:      md["Poster"].(string),
			Private:     md["Private"].(bool),
			Privatekey:  md["Privatekey"].(string)}

		log.Printf("mAlbum: %+v\n", mAlbum)
		uAlbum := dbh.AlbumUpdate(mAlbum)
		log.Printf("uAlbum: %+v\n", uAlbum)
		status = "Success"
		code = 200
	}

	log.Println("Album Update: ", "Name: ", name, " Album Id: ", album.Id, "=", mi, " Album User: ", album.User, "=", mo, " Session: ", su.Id, "=", mo, " Status: ", status)

	p.Data = models.AlbumResult{Status: status}
	encoder.Render("json", code, "", p, r)
}
