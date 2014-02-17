package middleware

import (
	"github.com/martini-contrib/render"
	"net/http"
	"pixelrelay/db"
	"pixelrelay/utils"
)

type Response map[string]interface{}

func Verify(res http.ResponseWriter, req *http.Request, r render.Render) {
	pk := req.FormValue("user_private_key")
	a := req.FormValue("file_album")

	d := db.InitDB()
	album := db.GetAlbum(&d, a)

	if album.Id > 0 && album.Privatekey == pk {
		return
	}

	if pk == "" || pk != utils.ImageCfg.SecretKey() {
		http.Error(res, "Invalid Private Key", http.StatusUnauthorized)
		r.JSON(http.StatusUnauthorized, Response{"error": http.StatusUnauthorized, "code": "Invalid Private Key", "name": a})
	}

	return
}
