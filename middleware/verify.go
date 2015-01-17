package middleware

import (
	"log"
	"net/http"

	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/utils"
)

type Response map[string]interface{}

func Verify(res http.ResponseWriter, req *http.Request, r render.Render, dbh *db.Dbh) {
	pk := req.FormValue("user_private_key")
	a := req.FormValue("file_album")

	album := dbh.GetAlbum(a)

	if album.Id > 0 && album.Privatekey == pk {
		return
	}

	if pk == "" || pk != utils.ImageCfg.SecretKey() {
		log.Printf("Upload failed: Invalid Private Key (%s)\n", pk)
		r.JSON(http.StatusUnauthorized, Response{"error": http.StatusUnauthorized, "code": "Invalid Private Key", "name": a})
		return
	}

	if a == "" {
		log.Printf("Upload failed: Invalid Album Name (%s)\n", a)
		r.JSON(http.StatusUnauthorized, Response{"error": http.StatusUnauthorized, "code": "Invalid Album Name", "name": a})
		return
	}

	log.Println("verify: good")
	return
}
