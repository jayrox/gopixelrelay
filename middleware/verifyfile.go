package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/codegangsta/martini"

	"pixelrelay/utils"
)

func VerifyFile(args martini.Params, res http.ResponseWriter, req *http.Request) {
	file := args["name"]

	// Verify file is not a "dotfile" and contains an extension
	if !strings.Contains(file, ".") || strings.HasPrefix(file, ".") || file == ".." {
		http.NotFound(res, req)
		return
	}

	fullpath := strings.Join([]string{utils.ImageCfg.Root(), file}, "")
	if !Exists(fullpath) {
		http.NotFound(res, req)
		return
	}

	return
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
