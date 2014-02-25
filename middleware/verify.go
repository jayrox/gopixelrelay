package middleware

import (
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
		http.Error(res, "Invalid Private Key", http.StatusUnauthorized)
		r.JSON(http.StatusUnauthorized, Response{"error": http.StatusUnauthorized, "code": "Invalid Private Key", "name": a})
	}

	return
}
