package controllers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/models"
	"pixelrelay/utils"
)

type UploadResult struct {
	Error int    `json:"error"`
	Code  string `json:"code"`
	Name  string `json:"name"`
}

func UploadImage(w http.ResponseWriter, req *http.Request, r render.Render, dbh *db.Dbh) {
	file, header, _ := req.FormFile("uploaded_file")
	defer file.Close()

	ur := &UploadResult{}

	rEmail := req.FormValue("user_email")
	rAlbum := req.FormValue("file_album")
	rPrivateKey := req.FormValue("user_private_key")

	log.Printf("header.Filename: %s\n", header.Filename)
	log.Printf("version: %s\n", req.FormValue("version"))
	log.Printf("user_email: %s\n", rEmail)
	log.Printf("user_private_key: %s\n", rPrivateKey)
	log.Printf("file_host: %s\n", req.FormValue("file_host"))
	log.Printf("file_album: %s\n", rAlbum)
	log.Printf("file_name: %s\n", req.FormValue("file_name"))
	log.Printf("file_mime: %s\n", req.FormValue("file_mime"))


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
		log.Println("fi err: ", err)
		ur.SetError(500)
		ur.SetCode(err.Error())
		r.JSON(500, ur)
		return
	}
	defer fi.Close()

	buf := make([]byte, 512)
	n, err := fi.Read(buf)
	if err != nil {
		log.Println("mime err: ", err)
		r.JSON(500, ur)
	}

	mime := http.DetectContentType(buf[:n])

	if mime != req.FormValue("file_mime") {
		ur.SetError(3)
		ur.SetCode("Invalid file type: " + mime)
		r.JSON(500, ur)
		return
	}

	log.Printf("tmp_file: %s\n", tmp_file)

	// Create Thumb
	tname := utils.ImageCfg.Thumbs() + ur.GetName()
	log.Printf("tname: %s\n", tname)

	if !Exists(string(tname)) && strings.Contains(tmp_file, "jpg") {
		okc := make(chan bool, 1)
		utils.CreateThumb(okc, tmp_file, tname)
		<-okc
	}

	// Add image to database
	dbh.AddUploader(models.Uploader{Email: rEmail, Timestamp: time.Now().Unix()})

	var user models.User
	user = dbh.GetUserByEmail(rEmail)
	log.Println("user: ", user)
	if user.Id == 0 {
		user = dbh.GetUploaderByEmail(rEmail)
		log.Println("uploader: ", user)
	} 
	log.Println("user: ", user)

	// Add image
	image := models.Image{Name: header.Filename, Album: rAlbum, User: user.Id, Timestamp: time.Now().Unix()}
	ai := dbh.AddImage(image)
	log.Println("ai: ", ai)

	// Add album
	album := models.Album{Name: rAlbum, User: user.Id, Privatekey: rPrivateKey, Private: true, Timestamp: time.Now().Unix()}
	dbh.AddAlbum(album)
	log.Println("album: ", album)

	ur.SetName(utils.AppCfg.Url() + "/i/" + header.Filename)
	log.Println("ur: ", ur)

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
