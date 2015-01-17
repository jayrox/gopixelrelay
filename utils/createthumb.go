package utils

import (
	"image"
	"image/color"
	"log"
	"math"
	"runtime"
	"strings"

	"github.com/disintegration/imaging"
)

func CreateThumb(okc chan bool, fname, tname string, w, h int) {
	log.Printf("Creating thumb for %s\n", fname)

	ok := make(chan bool, 1)
	go CreateThumbJpeg(ok, fname, tname, w, h)
	<-ok

	if strings.Contains(fname, "jpg") || strings.Contains(fname, "jpeg") {
		ii := &ImageInfo{FileName: fname, TempFileName: tname}
		Load(ii, ok)
		<-ok

		go RotateImage(ok, tname, tname, ii)
		<-ok
	} else {
		log.Println("File is not of type JPEG. Skipped rotating.")
	}

	log.Printf("Created thumb for %s\n", fname)
	okc <- true
}

func CreateThumbJpeg(ok chan bool, fname, tname string, w, h int) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered: ", r)
			return
		}
	}()

	// use all CPU cores for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	img, err := imaging.Open(fname)
	if err != nil {
		panic(err)
	}
	log.Printf("Height: %d, Width: %d\n", h, w)
	thumb := imaging.Thumbnail(img, w, h, imaging.CatmullRom)

	dst := imaging.New(w, h, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, thumb, image.Pt(0, 0))
	err = imaging.Save(dst, tname)
	if err != nil {
		panic(err)
	}
	log.Printf("Thumb file: %s\n", tname)
	ok <- true
}

func RotateImage(ok chan bool, fname, tname string, ii *ImageInfo) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered: ", r)
			return
		}
	}()

	// 6 portrait			 	// 0
	// 1 landscape, left		// 90
	// 8 portrait upside down	// 180
	// 3 landscape, right		// 270

	img, err := imaging.Open(fname)
	if err != nil {
		panic(err)
	}

	switch ii.Orientation {
	case 6:
		img = imaging.Rotate270(img)
	case 3:
		img = imaging.Rotate180(img)
	case 8:
		img = imaging.Rotate90(img)
	case 1:
		//r = 0.0
	}
	err = imaging.Save(img, tname)
	if err != nil {
		panic(err)
	}
	ok <- true
}

func ResizeImage(okc chan bool, fname, tname string, w, h int) {
	log.Printf("Resizing image: %s\n", fname)

	ok := make(chan bool, 1)

	if strings.Contains(fname, "jpg") || strings.Contains(fname, "jpeg") {

		ii := &ImageInfo{FileName: fname, TempFileName: tname}
		Load(ii, ok)
		<-ok

		go RotateImage(ok, fname, tname, ii)
		<-ok

		go ResizeJpeg(ok, tname, tname, w, h)
		<-ok
	} else {
		go ResizeJpeg(ok, fname, tname, w, h)
		<-ok
	}

	log.Printf("Resized image: %s\n", tname)
	okc <- true
}

func ResizeJpeg(ok chan bool, fname, tname string, w, h int) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered: ", r)
			return
		}
	}()

	// use all CPU cores for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	img, err := imaging.Open(fname)
	if err != nil {
		panic(err)
	}
	log.Printf("Width: %d, Height: %d\n", w, h)
	srcW := img.Bounds().Max.X
	srcH := img.Bounds().Max.Y
	log.Println("srcW: ", srcW, " srcH: ", srcH)

	if w == 0 {
		tmpW := float64(h) * float64(srcW) / float64(srcH)
		w = int(math.Max(1.0, math.Floor(tmpW+0.5)))
	}
	if h == 0 {
		tmpH := float64(w) * float64(srcH) / float64(srcW)
		h = int(math.Max(1.0, math.Floor(tmpH+0.5)))
	}
	log.Println("dstW: ", w, " dstH: ", h)

	thumb := imaging.Resize(img, w, h, imaging.Lanczos)

	dst := imaging.New(w, h, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, thumb, image.Pt(0, 0))
	err = imaging.Save(dst, tname)
	if err != nil {
		panic(err)
	}
	log.Printf("Modified file: %s\n", tname)
	ok <- true
}
