/*
 Route:  /account

 Method: GET

 Return:
  - Current user profile
*/

package controllers

import (
	"html/template"
	"strconv"
	"strings"

	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/db"
	"pixelrelay/encoder"
	"pixelrelay/forms"
	"pixelrelay/models"
	"pixelrelay/utils"
)

type accountVars struct {
	Form template.HTML
}

func Account(su models.User, session sessions.Session, r render.Render, dbh *db.Dbh, p *models.Page) {

	// Init error holder
	errs := make(map[string]string)

	account := forms.Account{}
	account.Id = su.Id
	account.Email = su.Email
	account.Name = su.Name

	username := su.UserName
	if username == "" {
		username = genUserName(su.Email, dbh)
	}
	account.UserName = username
	genform := utils.GenerateForm(&account, "/account", "POST", errs)

	p.SetTitle("Account")
	p.SetUser(su)
	p.Data = accountVars{Form: genform}

	encoder.Render(p.Encoding, 200, "account", p, r)
}

func genUserName(email string, dbh *db.Dbh) string {
	name := strings.Split(email, "@")[0]

	var u models.User
	u = dbh.GetUserByUserName(name)
	if u.Id != 0 {
		var i = 0
		var tmp_name string
		for {
			i++
			tmp_name = strings.Join([]string{name, strconv.Itoa(i)}, "")
			u = dbh.GetUserByUserName(tmp_name)
			if u.Id == 0 {
				return tmp_name
			}
		}
	}
	return name
}
