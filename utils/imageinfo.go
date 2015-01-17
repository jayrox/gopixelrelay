package utils

import (
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

type ImageInfo struct {
	FileName     string
	TempFileName string
	Orientation  int
}

func Load(ii *ImageInfo, ok chan bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered: ", r)
			return
		}
	}()

	fn, err := os.Open(ii.FileName)
	if err != nil {
		log.Println("ir fn err: ", err)
	}
	defer fn.Close()

	x, err := exif.Decode(fn)
	if err != nil {
		log.Printf("o err: %s\n", err)
	}

	ImageOrientation(ii, x)
	ok <- true
}

func ImageOrientation(ii *ImageInfo, x *exif.Exif) {
	orientation, err := x.Get(exif.Orientation)
	if err != nil {
		log.Println("o err: ", err)
		return
	}
	or := 0
	if orientation.Val[0] != 0 {
		or = int(orientation.Val[0])
	}
	if orientation.Val[1] != 0 {
		or = int(orientation.Val[1])
	}
	ii.Orientation = or
}
