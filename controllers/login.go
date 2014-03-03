package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/auth"
	"pixelrelay/db"
	"pixelrelay/forms"
	"pixelrelay/models"
	"pixelrelay/utils"
)

type LoginVars struct {
	LoginForm template.HTML
	Page      *models.Page
}

func Login(session sessions.Session, su models.User, r render.Render, res http.ResponseWriter, req *http.Request, p *models.Page) {

	// Check if we are already logged in
	if su.Id > 0 {
		http.Redirect(res, req, strings.Join([]string{utils.AppCfg.Url(), "albums"}, "/"), http.StatusFound)
		return
	}

	session.Set("loggedin", "false")

	// Init error holder
	errs := make(map[string]string)

	err_flash := session.Get("flash")
	if err_flash != nil {
		errs["flash"] = err_flash.(string)
		session.Set("flash", nil)
	}

	genform := utils.GenerateForm(&forms.Login{}, "/login", "POST", errs)

	var loginVars LoginVars
	loginVars.Page = p
	loginVars.Page.SetUser(su)
	loginVars.Page.SetTitle("Login")
	loginVars.LoginForm = genform

	r.HTML(200, "login", loginVars)
}

func LoginPost(lu forms.Login, session sessions.Session, res http.ResponseWriter, req *http.Request, dbh *db.Dbh) {
	errs := ValidateLogin(&lu)
	if len(errs) > 0 {
		fmt.Printf(`{"errors":"%v"}`, errs)
		fmt.Println("\n")
	}

	user := dbh.GetUserByEmail(lu.Email)

	match := auth.MatchPassword(lu.Password, user.Password, user.Salt)

	if match {
		sessionkey := SessionKey(user.Email, user.Password, user.Salt)

		session.Set("loggedin", "true")
		session.Set("uid", user.Id)
		session.Set("email", user.Email)
		session.Set("key", sessionkey)

		dbh.CreateSession(models.UserSession{UserId: user.Id, SessionKey: sessionkey, Active: true, Timestamp: time.Now().Unix()})

		http.Redirect(res, req, strings.Join([]string{utils.AppCfg.Url(), "albums"}, "/"), http.StatusFound)
		return
	}

	session.Set("flash", "Invalid Email or Password")
	http.Redirect(res, req, strings.Join([]string{utils.AppCfg.Url(), "login"}, "/"), http.StatusFound)
}

func Logout(session sessions.Session, res http.ResponseWriter, req *http.Request, dbh *db.Dbh) {
	sessionkey := session.Get("key")
	uid := session.Get("uid")

	session.Set("loggedin", "false")
	session.Set("uid", nil)
	session.Set("email", nil)
	session.Set("key", nil)

	dbh.DestroySession(uid.(int64), sessionkey.(string))
	http.Redirect(res, req, utils.AppCfg.Url(), http.StatusFound)
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
func SessionKey(email string, password string, salt string) (sha string) {
	tnow := strconv.FormatInt(time.Now().Unix(), 10)
	str := strings.Join([]string{tnow, email, password, salt}, "//")
	bv := []byte(str)

	hasher := sha1.New()
	hasher.Write(bv)
	sha = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return
}
