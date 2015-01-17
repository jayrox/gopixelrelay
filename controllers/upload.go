/*
 Route:  /up

 Method: POST

 Return: JSON
  Params:
   - Error
   - Code
   - Name

 TODO:
  - Change return params on android to better describe their contents
*/

package controllers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/kr/pretty"
	"github.com/martini-contrib/render"

	"pixelrelay/db"
	"pixelrelay/models"
	"pixelrelay/utils"
)

func UploadImage(w http.ResponseWriter, upload models.ImageUpload, req *http.Request, r render.Render, dbh *db.Dbh) {
	ur := &models.UploadResult{}

	rEmail := upload.Email
	rAlbum := upload.Album
	rPrivateKey := upload.PrivateKey

	fiName := upload.File.Filename

	upload_time := time.Now().Unix()

	ur.SetCode(200)
	ur.SetResult("success")
	ur.SetName(fiName)

	tmp_file := utils.ImageCfg.Root() + ur.GetName()

	if Exists(tmp_file) {
		log.Println("Error: File exists. (" + tmp_file + ")")
		ur.SetCode(2)
		ur.SetResult("File exists")
		r.JSON(500, ur)
		return
	}

	out, err := os.Create(tmp_file)
	if err != nil {
		log.Println("Error: Unable to open file.")
		ur.SetCode(500)
		ur.SetResult("Failed to open the file for writing.")
		r.JSON(500, ur)
		return
	}
	defer out.Close()

	fi, err := upload.File.Open()
	if err != nil {
		log.Println("fi err: ", err)
		ur.SetCode(500)
		ur.SetResult(err.Error())
		r.JSON(500, ur)
		return
	}
	defer fi.Close()

	_, err = io.Copy(out, fi)
	if err != nil {
		log.Println("Error: Failed to copy file.")
		ur.SetCode(500)
		ur.SetResult("Failed to copy file to new location.")
		r.JSON(500, ur)
		return
	}

	log.Printf("tmp_file: %s\n", tmp_file)

	// Add image uploader to database
	dbh.AddUploader(models.Uploader{Email: rEmail, Timestamp: upload_time})

	// Setup hashid to create unique file name
	var hid models.HashID
	hid.Init(utils.AppCfg.SecretKey(), 10)

	// Get user id
	user := dbh.GetUserByEmail(rEmail)
	log.Println("user: ", user.Id)
	if user.Id > 0 {
		// Add user id to hashid - seg 1
		hid.AddId(int(user.Id))
	} else {
		hid.AddId(0)
	}

	// Get uploader id
	user2 := dbh.GetUploaderByEmail(rEmail)
	log.Println("uploader user: ", user2.Id)
	if user2.Id > 0 {
		// Add uploader id to hashid - seg 2
		hid.AddId(int(user2.Id))
	} else {
		hid.AddId(0)
	}

	// Create default album description
	album := models.Album{
		Name:       rAlbum,
		User:       user.Id,
		Privatekey: rPrivateKey,
		Private:    true,
		Timestamp:  upload_time}
	album.DefaultDescription()
	log.Println(album)

	// Add album
	dbh.AddAlbum(album)

	log.Println("album: ", album)

	nAlbum := dbh.GetAlbum(rAlbum)

	// Add image
	image := dbh.AddImage(models.Image{
		Name:      fiName,
		Album:     rAlbum,
		User:      user.Id,
		AlbumId:   nAlbum.Id,
		Timestamp: upload_time})

	// Add image id to hashid - seg 3
	hid.AddId(int(image.Id))

	// Add upload time to hashid - seg 4
	hid.AddId(int(upload_time))

	// Get file extension and create new file name
	extension := filepath.Ext(fiName)
	nname := hid.Encrypt() + extension
	log.Printf("New name: %s\n", nname)

	image.HashId = nname
	dbh.UpdateImage(image)

	// Rename file to new name
	hash_name := utils.ImageCfg.Root() + nname
	os.Rename(tmp_file, hash_name)

	ur.SetName(utils.AppCfg.Url() + "/image/" + nname)

	// Create Thumb
	tname := utils.ImageCfg.Thumbs() + nname

	if !Exists(string(tname)) {
		okc := make(chan bool, 1)
		utils.CreateThumb(okc, hash_name, tname, 150, 150)
		<-okc
	}

	log.Printf("%# v\n", pretty.Formatter(album))
	log.Printf("%# v\n", image)
	log.Printf("%# v\n", pretty.Formatter(ur))

	r.JSON(200, ur)
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
