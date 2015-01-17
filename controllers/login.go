// Route: /login
// Route: /logout

package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/auth"
	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/forms"
	"pixelrelay/models"
	"pixelrelay/utils"
)

type LoginVars struct {
	Form template.HTML
}

func Login(session sessions.Session, su models.User, r render.Render, p *models.Page) {

	// Check if we are already logged in
	if su.Id > 0 {
		r.Redirect(strings.Join([]string{utils.AppCfg.Url(), "albums"}, "/"), http.StatusFound)
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

	p.SetUser(su)
	p.SetTitle("Login")
	p.Data = LoginVars{Form: genform}

	encoder.Render(p.Encoding, 200, "login", p, r)
}

func LoginPost(lu forms.Login, session sessions.Session, r render.Render, dbh *db.Dbh) {
	errs := ValidateLogin(&lu)
	if len(errs) > 0 {
		log.Printf("errors: %+v\n", errs)
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

		r.Redirect(strings.Join([]string{utils.AppCfg.Url(), "albums"}, "/"), http.StatusFound)
		return
	}

	session.Set("flash", "Invalid Email or Password")

	r.Redirect(strings.Join([]string{utils.AppCfg.Url(), "login"}, "/"), http.StatusFound)
}

func Logout(session sessions.Session, r render.Render, dbh *db.Dbh) {
	sessionkey := session.Get("key")
	uid := session.Get("uid")

	session.Set("loggedin", "false")
	session.Set("uid", nil)
	session.Set("email", nil)
	session.Set("key", nil)

	if uid != "" && uid != nil {
		dbh.DestroySession(uid.(int64), sessionkey.(string))
	}
	r.Redirect(utils.AppCfg.Url(), http.StatusFound)
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
