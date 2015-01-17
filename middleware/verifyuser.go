package middleware

import (
	"net/http"
	"reflect"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"

	"pixelrelay/db"
	"pixelrelay/models"
)

func UserAuth(formStruct interface{}, dbh *db.Dbh, ifacePtr ...interface{}) martini.Handler {
	return func(context martini.Context, session sessions.Session, req *http.Request) {
		formStruct := reflect.New(reflect.TypeOf(formStruct))

		su := getSessionUser(session, dbh)
		assignSessionUser(formStruct, su)
		validateAndMap(formStruct, context, ifacePtr...)
	}
}

func getSessionUser(session sessions.Session, dbh *db.Dbh) (user models.User) {
	var email, sessionkey, loggedin string
	var suid int64

	sloggedin := session.Get("loggedin")
	if sloggedin != nil {
		loggedin = sloggedin.(string)
	}

	semail := session.Get("email")
	if semail != nil {
		email = semail.(string)
	}

	ssessionkey := session.Get("key")
	if ssessionkey != nil {
		sessionkey = ssessionkey.(string)
	}

	ssuid := session.Get("uid")
	if ssuid != nil {
		suid = ssuid.(int64)
	}

	if loggedin != "true" || sessionkey == "" {
		return
	}

	usk := dbh.GetUserByIdSessionKey(suid, sessionkey)
	if usk.Active != true || usk.Id < 1 {
		return
	}

	u := dbh.GetUserById(suid)
	if u.Email != email {
		return
	}

	return u
}

func assignSessionUser(formStruct reflect.Value, su models.User) {
	typ := formStruct.Elem().Type()
	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := formStruct.Elem().Field(i)
		switch typeField.Name {
		case "Id":
			structField.SetInt(su.Id)
		case "Name":
			structField.SetString(su.Name)
		case "UserName":
			structField.SetString(su.UserName)
		case "Email":
			structField.SetString(su.Email)
		}
	}
}

func validateAndMap(obj reflect.Value, context martini.Context, ifacePtr ...interface{}) {
	context.Map(obj.Elem().Interface())
	if len(ifacePtr) > 0 {
		context.MapTo(obj.Elem().Interface(), ifacePtr[0])
	}
}
