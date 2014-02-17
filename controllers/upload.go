package controllers

import (
	"fmt"
	"github.com/martini-contrib/render"
	"io"
	"net/http"
	"os"
	"pixelrelay/db"
	"pixelrelay/models"
	"pixelrelay/utils"
	//"reflect"
	"strings"
	"time"
)

type UploadResult struct {
	Error int    `json:"error"`
	Code  string `json:"code"`
	Name  string `json:"name"`
}

//pimage models.PostImage,
func UploadImage(w http.ResponseWriter, req *http.Request, r render.Render) {
	file, header, _ := req.FormFile("uploaded_file")
	defer file.Close()

	//fmt.Println(reflect.TypeOf(file))
	//fmt.Println(reflect.TypeOf(req.FormFile("uploaded_file")))

	ur := &UploadResult{}
	d := db.InitDB()

	fmt.Printf("header.Filename: %s\n", header.Filename)
	fmt.Printf("version: %s\n", req.FormValue("version"))
	fmt.Printf("user_email: %s\n", req.FormValue("user_email"))
	fmt.Printf("user_private_key: %s\n", req.FormValue("user_private_key"))
	fmt.Printf("file_host: %s\n", req.FormValue("file_host"))
	fmt.Printf("file_album: %s\n", req.FormValue("file_album"))
	fmt.Printf("file_name: %s\n", req.FormValue("file_name"))
	fmt.Printf("file_mime: %s\n", req.FormValue("file_mime"))

	//fmt.Println("pimage: ", pimage)

	ur.SetError(200)
	ur.SetCode("success")

	ur.SetName(header.Filename)

	tmp_file := utils.ImageCfg.Root() + ur.GetName()

	if Exists(tmp_file) {
		ur.SetError(2)
		ur.SetCode("File exists")
		r.JSON(500, ur)
		return
	}

	out, err := os.Create(tmp_file)
	if err != nil {
		ur.SetError(500)
		ur.SetCode("Failed to open the file for writing.")
		r.JSON(500, ur)
		return
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		ur.SetError(500)
		ur.SetCode("Failed to copy file to new location.")
		r.JSON(500, ur)
		return
	}

	fi, err := os.Open(tmp_file)
	if err != nil {
		fmt.Println("fi err: ", err)
		ur.SetError(500)
		ur.SetCode(err.Error())
		r.JSON(500, ur)
		return
	}
	defer fi.Close()

	buf := make([]byte, 512)
	n, err := fi.Read(buf)
	if err != nil {
		fmt.Println("mime err: ", err)
		r.JSON(500, ur)
	}

	mime := http.DetectContentType(buf[:n])

	if mime != req.FormValue("file_mime") {
		ur.SetError(3)
		ur.SetCode("Invalid file type: " + mime)
		r.JSON(500, ur)
		return
	}
	
	fmt.Printf("tmp_file: %s\n", tmp_file)
	
	// Create Thumb
	tname := utils.ImageCfg.Thumbs() + ur.GetName()
	fmt.Printf("tname: %s\n", tname)
	
	if !Exists(string(tname)) && strings.Contains(tmp_file, "jpg") {
		fmt.Printf("creating thumb for %s\n", tmp_file)

		ok := make(chan bool, 1)
		go utils.CreateThumbJpeg(ok, tmp_file, tname, 150, 150)
		<-ok

		tname := utils.ImageCfg.Thumbs() + ur.GetName()
		ii := &utils.ImageInfo{FileName: tmp_file, TempFileName: tname}
		utils.Load(ii, ok)
		<-ok

		go utils.ImageRotate(ii, ok)
		<-ok

		fmt.Printf("thumb created for %s\n", tmp_file)
	}
	
	// Add image to database
	u := int64(1) //FIXME!!!

	// Add image
	image := models.Image{Name: header.Filename, Album: req.FormValue("file_album"), User: u, Timestamp: time.Now().Unix()}
	var ui models.Image
	ui = db.AddImage(&d, image)
	fmt.Println("ui: ", ui.Id)

	// Add album
	album := models.Album{Name: req.FormValue("file_album"), User: u, Privatekey: req.FormValue("user_private_key"), Private: true, Timestamp: time.Now().Unix()}
	db.AddAlbum(&d, album)

	if ui.Id > 0 {
		uploader := models.Uploader{User: u, Image: ui.Id, Email: req.FormValue("user_email"), Timestamp: time.Now().Unix()}
		db.AddUpload(&d, uploader)
		ur.SetName(utils.AppCfg.Url() + "/i/" + header.Filename)
	}
	r.JSON(200, ur)
}

func (ur *UploadResult) SetError(error int) {
	ur.Error = error
}

func (ur UploadResult) GetError() int {
	return ur.Error
}

func (ur *UploadResult) SetCode(code string) {
	ur.Code = code
}

func (ur UploadResult) GetCode() string {
	return ur.Code
}

func (ur *UploadResult) SetName(name string) {
	ur.Name = name
}

func (ur UploadResult) GetName() string {
	return ur.Name
}

// https://github.com/noll/mjau/blob/master/util/util.go#L42
// http://stackoverflow.com/a/12527546/24802
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
