package auth

import (
	"github.com/martini-contrib/sessions"
	"pixelrelay/db"
	"pixelrelay/models"
)

func UserAuth(session sessions.Session) models.User {
	var user models.User

	loggedin := session.Get("loggedin")
	if loggedin != "true" {
		return user
	}

	d := db.InitDB()

	suid := session.Get("uid").(int64)
	
	sessionkey := session.Get("key").(string)
	if sessionkey == "" {
		return user
	}

	usk := db.GetUserByIdSessionKey(&d, suid, sessionkey)
	if usk.Active != true || usk.Id < 1{
		return user
	}

	u := db.GetUserById(&d, suid)

	email := session.Get("email").(string)
	if u.Email != email {
		return user
	}

	return u
}