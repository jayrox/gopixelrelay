package controllers

import (
	"net/http"

	"github.com/codegangsta/martini"

	"pixelrelay/utils"
)

func Image(args martini.Params, res http.ResponseWriter, req *http.Request) {
	file := args["name"]
	fdir := utils.ImageCfg.Root()

	dir := http.Dir(fdir)

	f, err := dir.Open(file)
	if err != nil {
		// discard the error?
		http.NotFound(res, req)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.NotFound(res, req)
		return
	}

	res.Header().Set("X-Content-Type-Options", "nosniff")
	res.Header().Set("Expires", utils.ExpiresHeader())
	http.ServeContent(res, req, file, fi.ModTime(), f)
}
