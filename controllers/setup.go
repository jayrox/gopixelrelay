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
	"time"
)

func SetupAdmin(args martini.Params, session sessions.Session, r render.Render) {
	form := &forms.SetupAdmin{}
	session.Set("setup", "true")

	genform := utils.GenerateForm(form, "/setup", "POST", nil)
	r.HTML(200, "setup", genform)
}

func SetupAdminPost(sa forms.SetupAdmin, args martini.Params, session sessions.Session, r render.Render, res http.ResponseWriter) {
	errs := Validate(&sa)
	if len(errs) > 0 {
		fmt.Printf(`{"errors":"%v"}`, errs)
		fmt.Println("\n")
	}

	v := session.Get("setup")
	if v != "true" {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	d := db.InitDB()
	user := db.GetUserByEmail(&d, sa.Email)

	if user.Id > 0 {
		fmt.Println("user already exists")
		session.Set("uid", user.Id)
	}

	if user.Id == 0 {
		fmt.Println("id: 0")
		hash, salt, err := auth.EncryptPassword(sa.Password)
		if err != nil {
			fmt.Println("hash err: ", err, "\n")
		}
		newuser := models.User{Name: sa.Name, Email: sa.Email, Password: hash, Salt: salt, Timestamp: time.Now().Unix()}
		db.InsertUser(&d, newuser)
		session.Set("uid", newuser.Id)
	}

	uid := session.Get("uid")

	fmt.Println("uid: ", uid)
	genform := utils.GenerateForm(&sa, "/setup", "POST", errs)
	r.HTML(200, "setup", genform)
}

func Validate(sa *forms.SetupAdmin) map[string]string {
	errs := make(map[string]string)

	if len(sa.Name) <= 0 {
		errs["name"] = "Name must be set"
	} else if len(sa.Name) > 255 {
		errs["name"] = "Too long; maximum 255 characters"
	}
	if len(sa.Email) <= 0 {
		errs["email"] = "Email must be at set"
	}
	if len(sa.Password) <= 0 {
		errs["password"] = "Password must be at set"
	}
	if len(sa.Confirm) <= 0 {
		errs["confirm"] = "Password confirmation must be at set"
	}
	if sa.Password != sa.Confirm {
		errs["password"] = "Passwords must match"
		errs["confirm"] = "Passwords must match"
	}

	return errs
}
