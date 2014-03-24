package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"pixelrelay/db"
	"pixelrelay/models"
)

func AlbumUpdate(res http.ResponseWriter, req *http.Request, dbh *db.Dbh) {
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

	mi := int64(m["Id"].(float64))
	mo := int64(m["Owner"].(float64))

	mAlbum := models.Album{Id: mi,
		Description: m["Description"].(string),
		Name:        m["Name"].(string),
		User:        mo,
		Poster:      m["Poster"].(string),
		Private:     m["Private"].(bool),
		Privatekey:  m["Privatekey"].(string)}

	dbh.AlbumUpdate(mAlbum)
	return
}
