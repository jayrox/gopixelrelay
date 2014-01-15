package utils

import (
	"code.google.com/p/graphics-go/graphics"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"image"
	"image/jpeg"
	"os"
)

type Or struct {
	Val []byte
}

func ImageOrientation(fname string, ok chan bool) {
	// values
	// 6 portrait			 	// 0
	// 1 landscape, left		// 90
	// 8 portrait upside down	// 180
	// 3 landscape, right		// 270

	f, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Orientation err: %s\n", err)
	}

	x, err := exif.Decode(f)
	defer f.Close()
	if err != nil {
		fmt.Printf("Orientation err: %s\n", err)
	}

	orientation, err := x.Get(exif.Orientation)
	fmt.Println("orientation: %s\n", orientation)
	fmt.Printf("Orientation err: %s\n", err)

	imageWidth, _ := x.Get(exif.ImageWidth)
	fmt.Printf("imagewidth: %s\n", imageWidth)

	imageLength, _ := x.Get(exif.ImageLength)
	fmt.Printf("imagelength: %s\n", imageLength)

	ok <- true
}

func ImageRotate(fname string, tname string, ok chan bool) {
	fImg1, _ := os.Open(fname)
	defer fImg1.Close()
	img1, _, _ := image.Decode(fImg1)

	m := image.NewRGBA(image.Rect(0, 0, 800, 600))
	graphics.Rotate(m, img1, &graphics.RotateOptions{3.5})

	toimg, _ := os.Create(tname)
	defer toimg.Close()

	jpeg.Encode(toimg, m, &jpeg.Options{jpeg.DefaultQuality})

	ok <- true
}
