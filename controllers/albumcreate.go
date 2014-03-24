package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"pixelrelay/db"
	"pixelrelay/models"
	"pixelrelay/utils"
)

func AlbumCreate(res http.ResponseWriter, req *http.Request, su models.User, dbh *db.Dbh) {
	var reader io.Reader = req.Body
	b, e := ioutil.ReadAll(reader)
	if e != nil {
		log.Println(e)
	}

	log.Println(b)

	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Println(err)
	}

	m := f.(map[string]interface{})

	name := m["name"].(string)

	privatekey := m["privatekey"].(string)
	if len(privatekey) == 0 {
		privatekey = utils.ImageCfg.SecretKey()
	}

	description := m["description"].(string)
	if len(description) == 0 {
		const layout = "Auto-created 2 January 2006"
		t := time.Now()
		description = t.Format(layout)
	}

	// Add album
	album := models.Album{Name: name, User: su.Id, Privatekey: privatekey, Private: true, Description: description, Timestamp: time.Now().Unix()}
	log.Println(album)
	dbh.AddAlbum(album)
	return
}
