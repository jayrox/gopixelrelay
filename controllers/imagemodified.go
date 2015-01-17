/*
 Route:  /i/:name

 Method: GET

 Return:
  - image/* resized
*/

package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-martini/martini"

	"pixelrelay/utils"
)

func ImageModified(args martini.Params, res http.ResponseWriter, req *http.Request) {
	file := args["name"]

	size := "500"
	sfile := strings.Split(file, ".")
	ext := sfile[len(sfile)-1]
	sfile[len(sfile)-1] = size
	sfile = append(sfile, ext)

	fname := strings.Join(sfile, ".")

	log.Println(fname)

	odir := utils.ImageCfg.Root()
	org_file := odir + file
	temp_file := odir + fname

	if !Exists(temp_file) {
		okc := make(chan bool, 1)
		go utils.ResizeImage(okc, org_file, temp_file, 500, 0)
		<-okc
	}

	dir := http.Dir(odir)

	f, err := dir.Open(fname)
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
	http.ServeContent(res, req, fname, fi.ModTime(), f)
}
