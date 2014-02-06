package controllers

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"pixelrelay/auth"
)

func Auth(args martini.Params, r render.Render) {
	password := args["password"]
	
	hash, salt, err := auth.EncryptPassword(password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("pass: %s\nhash: %s\nsalt: %s\n", password, hash, salt)
	
	fmt.Println("match: ", auth.MatchPassword(password, hash, salt))
}
