package utils

import (
	"code.google.com/p/graphics-go/graphics"
	"encoding/binary"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"image"
	"image/jpeg"
	"math"
	"os"
)

type ImageInfo struct {
	FileName     string
	TempFileName string
	Orientation  int
	Height       int
	Width        int
	Ratio        float32
}

func Load(ii *ImageInfo, ok chan bool) {
	fn, err := os.Open(ii.FileName)
	if err != nil {
		fmt.Println("ir fn err: ", err)
	}
	defer fn.Close()

	x, err := exif.Decode(fn)
	if err != nil {
		fmt.Printf("o err: %s\n", err)
	}

	ImageOrientation(ii, x)
	ImageWidth(ii, x)
	ImageHeight(ii, x)
	ImageRatio(ii)
	ok <- true
}

func ImageOrientation(ii *ImageInfo, x *exif.Exif) {
	orientation, err := x.Get(exif.Orientation)
	if err != nil {
		fmt.Println("o err: ", err)
		return
	}
	ii.Orientation = int(orientation.Val[1])
}

func ImageWidth(ii *ImageInfo, x *exif.Exif) {
	imageWidth, _ := x.Get(exif.ImageWidth)
	wu8 := imageWidth.Val
	wu32BE := binary.BigEndian.Uint32(wu8)

	ii.Width = int(wu32BE)
}

func ImageHeight(ii *ImageInfo, x *exif.Exif) {
	imageLength, _ := x.Get(exif.ImageLength)
	hu8 := imageLength.Val
	hu32BE := binary.BigEndian.Uint32(hu8)

	ii.Height = int(hu32BE)
}

func ImageRatio(ii *ImageInfo) {
	if ii.Height == 0 || ii.Width == 0 {
		ii.Ratio = 0
		return
	}
	if ii.Height > ii.Width {
		ii.Ratio = float32(ii.Height) / float32(ii.Width)
		return
	}
	if ii.Height < ii.Width {
		ii.Ratio = float32(ii.Width) / float32(ii.Height)
		return
	}
}

func ImageRotate(ii *ImageInfo, ok chan bool) {
	// 6 portrait			 	// 0
	// 1 landscape, left		// 90
	// 8 portrait upside down	// 180
	// 3 landscape, right		// 270

	fn, err := os.Open(ii.TempFileName)
	if err != nil {
		fmt.Println("ir fn err: ", err)
	}
	defer fn.Close()

	img, _, err := image.Decode(fn)
	if err != nil {
		fmt.Println("err: ", err)
	}

	r := 0.0
	switch ii.Orientation {
	case 6:
		r = math.Pi / 2
	case 3:
		r = math.Pi
	case 8:
		r = math.Pi * 1.5
	case 1:
		r = 0.0
	}

	if r == 0 {
		ok <- true
		return
	}

	h := 120
	w := 120
	if ii.Height > ii.Width {
		h = int(float32(h) * ii.Ratio)
	} else {
		w = int(float32(w) * ii.Ratio)
	}

	m := image.NewRGBA(image.Rect(0, 0, w, h))

	graphics.Rotate(m, img, &graphics.RotateOptions{r})

	toimg, err := os.Create(ii.TempFileName)
	if err != nil {
		fmt.Println("rotate err: ", err)
	}
	defer toimg.Close()

	jpeg.Encode(toimg, m, &jpeg.Options{Quality: 90})

	ok <- true
}
