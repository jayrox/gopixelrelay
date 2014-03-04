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
