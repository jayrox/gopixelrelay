package controllers

import (
	"net/http"
	"github.com/codegangsta/martini"
    "github.com/nfnt/resize"
	"image"
    "image/jpeg"
	"image/png"
    "log"
    "os"
	"fmt"
	"strings"
)

func Thumb(args martini.Params, res http.ResponseWriter, req *http.Request) {
	file := args["name"]
	
	org_dir := "./tmp/"
	temp_dir := org_dir+"thumbs/"
	
	org_file := org_dir+file
	temp_file := temp_dir+file
	
	if ! Exists(temp_file) && strings.Contains(temp_file, "jpg") {
		ok := make(chan bool, 1)
		go createThumbJpeg(ok, file, 150, 150)
		fmt.Printf("creating thumb for %s\n", temp_file)
		<-ok
		fmt.Printf("thumb created for %s\n", temp_file)
	} else if ! strings.Contains(temp_file, "jpg") {
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
	
	http.ServeContent(res, req, file, fi.ModTime(), f)

}


func createThumbJpeg(ok chan bool, filename string, h uint, w uint) {

	dir := "./tmp/"
	temp_dir := dir+"thumbs/"
	org_file := dir+filename
	temp_file := temp_dir+filename
	
    file, err := os.Open(org_file)
    if err != nil {
        log.Fatal(err)
    }

	// decode jpeg into image.Image
    img, _, err := image.Decode(file)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    m := resize.Resize(w, h, img, resize.Bilinear)

    out, err := os.Create(temp_file)
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()

    // write new image to file
    jpeg.Encode(out, m, nil)
	ok <- true
}

func createThumbPng(ok chan bool, filename string, h uint, w uint) {

	dir := "./tmp/"
	temp_dir := dir+"thumbs/"
	org_file := dir+filename
	temp_file := temp_dir+filename
	
    file, err := os.Open(org_file)
    if err != nil {
        log.Fatal(err)
    }

	// decode jpeg into image.Image
    img, _, err := image.Decode(file)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    m := resize.Resize(w, h, img, resize.Bilinear)

    out, err := os.Create(temp_file)
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()

    // write new image to file
    png.Encode(out, m)
	ok <- true
}