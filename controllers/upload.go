package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"io"
	"net/http"
	"os"
	"pixelrelay/utils"
)

type UploadResult struct {
	Error int    `json:"error"`
	Code  string `json:"code"`
	Name  string `json:"name"`
}

func UploadImage(w http.ResponseWriter, req *http.Request, r render.Render) {
	file, header, _ := req.FormFile("uploaded_file")
	defer file.Close()

	ur := &UploadResult{}

	fmt.Printf("header.Filename: %s\n", header.Filename)
	fmt.Printf("version: %s\n", req.FormValue("version"))
	fmt.Printf("user_email: %s\n", req.FormValue("user_email"))
	fmt.Printf("user_private_key: %s\n", req.FormValue("user_private_key"))
	fmt.Printf("file_host: %s\n", req.FormValue("file_host"))
	fmt.Printf("file_album: %s\n", req.FormValue("file_album"))
	fmt.Printf("file_name: %s\n", req.FormValue("file_name"))
	fmt.Printf("file_mime: %s\n", req.FormValue("file_mime"))

	ur.SetError(200)
	ur.SetCode("success")

	ur.SetName(header.Filename)

	tmp_file := utils.ImageCfg.Root() + ur.GetName()
	
	if Exists(tmp_file) {
		ur.SetError(2)
		ur.SetCode("file exists")
	} else {
		out, err := os.Create(tmp_file)
		if err != nil {
			ur.SetError(500)
			ur.SetCode("Failed to open the file for writing.")
			return
		}

		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			ur.SetError(500)
			ur.SetCode("Failed to copy file to new location.")
			fmt.Fprintln(w, err)
		}

		fImg1, _ := os.Open(tmp_file)
		defer fImg1.Close()

		ok := make(chan bool, 1)
		//go utils.ImageOrientation(fImg1) FIXME
		fmt.Printf("get orintation for %s\n", tmp_file)
		<-ok
		fmt.Printf("got orientation for %s\n", tmp_file)
	}

	bytesOfJSON, jerr := json.Marshal(ur)
	if jerr != nil {
		fmt.Println(jerr)
		ur.SetError(500)
		ur.SetCode("Unknown JSON error.")
		return
	}
	fmt.Fprintf(w, string(bytesOfJSON))
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
