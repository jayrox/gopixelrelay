package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func CreateThumb(okc chan bool, fname string, tname string) {
	log.Printf("creating thumb for %s\n", fname)

	ok := make(chan bool, 1)
	go CreateThumbJpeg(ok, fname, tname, 150, 150)
	<-ok

	ii := &ImageInfo{FileName: fname, TempFileName: tname}
	Load(ii, ok)
	<-ok

	go ImageRotate(ii, ok)
	<-ok

	log.Printf("thumb created for %s\n", fname)
	okc <- true
}

func CreateThumbJpeg(ok chan bool, fname string, tname string, h uint, w uint) {
	file, err := os.Open(fname)
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

	out, err := os.Create(tname)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	ok <- true
}

func CreateThumbPng(ok chan bool, fname string, tname string, h uint, w uint) {
	file, err := os.Open(fname)
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

	out, err := os.Create(tname)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
	ok <- true
}
