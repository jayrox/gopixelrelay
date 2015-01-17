/*
 Route:  /tag

 Method: POST

 Params:
  - tag string
  - image string

 Return:
  - Redirect back to image.

 TODO:
  - Convert to Martini Bind and return JSON success message.
  - POST with redirect?
*/

package controllers

import (
	"log"
	"net/http"
	"strings"

	"pixelrelay/db"
	"pixelrelay/utils"
)

func TagImage(res http.ResponseWriter, req *http.Request, dbh *db.Dbh) {
	tag := req.FormValue("tag")
	image := req.FormValue("image")

	imagetag := dbh.TagImage(tag, image)
	log.Println(imagetag)

	http.Redirect(res, req, strings.Join([]string{utils.AppCfg.Url(), "image", image}, "/"), http.StatusFound)
	return
}
