package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/db"
	"pixelrelay/models"
	"pixelrelay/utils"
)

func AlbumPrivate(args martini.Params, session sessions.Session, su models.User, r render.Render, res http.ResponseWriter, req *http.Request, dbh *db.Dbh) {
	name := args["name"]
	state, err := strconv.ParseBool(args["state"])
	if err != nil {
		log.Println("Invalid state: ", args["state"])
		return
	}
	log.Printf("uid: %d name: %s state: %t\n", su.Id, name, state)

	if su.Id > 0 {
		dbh.SetAlbumPrivacy(su.Id, name, state)
	}

	http.Redirect(res, req, strings.Join([]string{utils.AppCfg.Url(), "albums"}, "/"), http.StatusFound)
	return
}
