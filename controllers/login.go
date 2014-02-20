package controllers

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	"pixelrelay/auth"
	"pixelrelay/db"
	"pixelrelay/forms"
	"pixelrelay/models"
	"pixelrelay/utils"
	"strings"
	"time"

	"crypto/sha1"
	"encoding/base64"
	"strconv"
)

func Login(args martini.Params, session sessions.Session, res http.ResponseWriter, req *http.Request, ren render.Render) {
	session.Set("loggedin", "false")
	form := &forms.Login{}

	// Init error holder
	errs := make(map[string]string)

	// If email set, apply to form
	email := session.Get("email")
	if email != nil {
		form.Email = email.(string)
		session.Set("email", nil)
	}

	// If error_p is set, apply error test to password field
	err_pass := session.Get("error_p")
	if err_pass != nil {
		errs["password"] = err_pass.(string)
		session.Set("error_p", nil)
	}

	genform := utils.GenerateForm(form, "/login", "POST", errs)
	ren.HTML(200, "login", genform)
}

func LoginPost(lu forms.Login, args martini.Params, session sessions.Session, r render.Render, res http.ResponseWriter, req *http.Request) {
	errs := ValidateLogin(&lu)
	if len(errs) > 0 {
		fmt.Printf(`{"errors":"%v"}`, errs)
		fmt.Println("\n")
	}

	d := db.InitDB()
	user := db.GetUserByEmail(&d, lu.Email)

	match := auth.MatchPassword(lu.Password, user.Password, user.Salt)

	if match {
		sessionkey := SessionKey(user.Email, user.Password, user.Salt)

		session.Set("loggedin", "true")
		session.Set("uid", user.Id)
		session.Set("email", user.Email)
		session.Set("key", sessionkey)

		db.CreateSession(&d, models.UserSession{UserId: user.Id, Key: sessionkey, Active: true, Timestamp: time.Now().Unix()})
		
		http.Redirect(res, req, strings.Join([]string{utils.AppCfg.Url(), "albums"}, "/"), http.StatusMovedPermanently)
		return
	}

	session.Set("error_p", "Invalid Password")
	session.Set("email", lu.Email)
	http.Redirect(res, req, strings.Join([]string{utils.AppCfg.Url(), "login"}, "/"), http.StatusMovedPermanently)
}

func ValidateLogin(lu *forms.Login) map[string]string {
	errs := make(map[string]string)

	if len(lu.Email) <= 0 {
		errs["email"] = "Email must be at set"
	}
	if len(lu.Password) <= 0 {
		errs["password"] = "Password must be at set"
	}

	return errs
}

// Generate session key
func SessionKey(email string, password string, salt string) string {
	tnow := strconv.FormatInt(time.Now().Unix(), 10)
	str := strings.Join([]string{tnow, email, password, salt}, "//")
	bv := []byte(str)

	hasher := sha1.New()
	hasher.Write(bv)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha
}
