package controllers

import (
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
	ii "pixelrelay/utils"
	"strings"
)

func Image(args martini.Params, res http.ResponseWriter, req *http.Request) {
	file := args["name"]
	fdir := "./tmp/"
	fname := fdir + file

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

	if strings.Contains(fname, "jpg") {
		ok := make(chan bool, 1)
		go ii.ImageOrientation(fname, ok)
		fmt.Printf("get orintation for %s\n", fname)
		<-ok
		fmt.Printf("got orientation for %s\n", fname)
	}

	http.ServeContent(res, req, file, fi.ModTime(), f)
}
