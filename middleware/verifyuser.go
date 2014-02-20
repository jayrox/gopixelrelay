package middleware

import (
	"github.com/martini-contrib/render"
	"net/http"
)

func VerifyUser(res http.ResponseWriter, req *http.Request, r render.Render) string {
	return ""
}
