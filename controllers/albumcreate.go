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
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/models"
)

func AlbumCreate(req *http.Request, r render.Render, su models.User, dbh *db.Dbh, p *models.Page) {
	// Default status
	var status string = "Permission Denied"
	var code int = 401

	var reader io.Reader = req.Body
	b, e := ioutil.ReadAll(reader)
	if e != nil {
		log.Println(e)
	}

	adata := string(b)

	params := strings.Split(adata, "&")

	var name, privatekey, description string
	for _, val := range params {
		vars := strings.Split(val, "=")
		k := strings.ToLower(vars[0])
		v := vars[1]
		switch k {
		case "name":
			name = v
		case "privatekey":
			privatekey = v
		case "description":
			description = v
		}

	}
	if len(description) == 0 {
		const layout = "Auto-created 2 January 2006"
		t := time.Now()
		description = t.Format(layout)
	}

	// Add album
	nAlbum := models.Album{Name: name, User: su.Id, Privatekey: privatekey, Private: true, Description: description, Timestamp: time.Now().Unix()}
	album := dbh.AddAlbum(nAlbum)
	if album.Id > 0 {
		code = 200
		status = "Success"
	}

	p.Data = models.AlbumResult{Status: status}
	encoder.Render("json", code, "", p, r)
}
