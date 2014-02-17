package middleware

import (
	"github.com/codegangsta/martini"
	"net/http"
	"strings"
)

func VerifyFile(args martini.Params, res http.ResponseWriter, req *http.Request) {
	file := args["name"]

	// Verify file is not a "dotfile" and contains an extension
	if !strings.Contains(file, ".") || strings.HasPrefix(file, ".") || file == ".." {
		http.NotFound(res, req)
		return
	}

	return
}
