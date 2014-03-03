package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/codegangsta/martini"

	"pixelrelay/db"
	"pixelrelay/utils"
)

func TagImage(args martini.Params, res http.ResponseWriter, req *http.Request, dbh *db.Dbh) {
	tag := args["name"]
	image := args["image"]

	imagetag := dbh.TagImage(tag, image)
	log.Println(imagetag)

	http.Redirect(res, req, strings.Join([]string{utils.AppCfg.Url(), "image", image}, "/"), http.StatusFound)
	return
}
