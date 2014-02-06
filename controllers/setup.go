package controllers

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"pixelrelay/forms"
	"pixelrelay/utils"
)

func SetupAdmin(args martini.Params, r render.Render) {//string {
	form := &forms.SetupAdmin{}
	form.Email = "jayrox@gmail.com"
	form.Name = "Jay"
	form.Password = ""
	form.Confirm = ""
	
	
	genform := utils.GenerateForm(form, "/setup", "POST", nil)
	r.HTML(200, "setup", genform)
}

func SetupAdminPost(sa forms.SetupAdmin, args martini.Params, r render.Render) {
	errs := Validate(&sa)
	if len(errs) > 0 {
		fmt.Printf(`{"errors":"%v"}`, errs)
		fmt.Println("\n")
	}
	
	genform := utils.GenerateForm(&sa, "/setup", "POST", errs)
	r.HTML(200, "setup", genform)
}

func Validate(sa *forms.SetupAdmin) (map[string]string) {
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
