package controllers

import (
	"github.com/codegangsta/martini"
	"net/http"
	"pixelrelay/utils"
	"strings"
)

func Thumb(args martini.Params, res http.ResponseWriter, req *http.Request) {
	file := args["name"]

	org_dir := utils.ImageCfg.Root()
	temp_dir := utils.ImageCfg.Thumbs()

	org_file := org_dir + file
	temp_file := temp_dir + file

	if !Exists(temp_file) && strings.Contains(temp_file, "jpg") {
		okc := make(chan bool, 1)
		utils.CreateThumb(okc, org_file, temp_file)
		<-okc
	} else if !strings.Contains(temp_file, "jpg") {
		temp_file = org_file
		temp_dir = org_dir
	}

	dir := http.Dir(temp_dir)

	f, err := dir.Open(file)
	if err != nil {
		// discard the error?
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return
	}

	res.Header().Set("X-Content-Type-Options", "nosniff")
	http.ServeContent(res, req, file, fi.ModTime(), f)
}
