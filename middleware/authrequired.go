package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/martini-contrib/sessions"

	"pixelrelay/models"
	"pixelrelay/utils"
)

func AuthRequired(su models.User, session sessions.Session, res http.ResponseWriter, req *http.Request) {
	if su.Id > 0 {
		return
	}
	log.Println("Login Required")
	session.Set("flash", "Login Required")
	http.Redirect(res, req, strings.Join([]string{utils.AppCfg.Url(), "login"}, "/"), http.StatusFound)
}
