/*
 Route:  /t/:name

 Method: GET

 Params:
  - name string

 Return:
  - image/* thumb
*/

package controllers

import (
	"log"
	"net/http"

	"pixelrelay/utils"

	"github.com/go-martini/martini"
)

func Thumb(args martini.Params, res http.ResponseWriter, req *http.Request) {
	file := args["name"]

	org_dir := utils.ImageCfg.Root()
	temp_dir := utils.ImageCfg.Thumbs()

	org_file := org_dir + file
	temp_file := temp_dir + file

	if !Exists(temp_file) {
		okc := make(chan bool, 1)
		go utils.CreateThumb(okc, org_file, temp_file, 150, 150)
		<-okc
	}

	dir := http.Dir(temp_dir)

	f, err := dir.Open(file)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return
	}

	res.Header().Set("X-Content-Type-Options", "nosniff")
	res.Header().Set("Expires", utils.ExpiresHeader())
	http.ServeContent(res, req, file, fi.ModTime(), f)
}
