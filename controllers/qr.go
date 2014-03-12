package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"code.google.com/p/rsc/qr"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"pixelrelay/db"
	"pixelrelay/models"
	"pixelrelay/utils"
)

func QR(args martini.Params, su models.User, dbh *db.Dbh, session sessions.Session, r render.Render, res http.ResponseWriter, req *http.Request) {
	name := args["name"]

	album := dbh.GetAlbum(name)

	key := album.Privatekey
	private := album.Private

	if private && su.Id != album.User || album.Id == 0 {
		session.Set("flash", "Login Required")
		r.Redirect(strings.Join([]string{utils.AppCfg.Url(), "login"}, "/"), http.StatusFound)
		return
	}

	log.Printf("name: %s key: %s private: %t\n", name, key, private)

	file := createQR(name, key)

	dir := http.Dir(utils.ImageCfg.QR())
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

	res.Header().Set("X-Content-Type-Options", "nosniff")
	res.Header().Set("Expires", utils.ExpiresHeader())
	res.Header().Add("Content-Type", "image/png")
	http.ServeContent(res, req, file, fi.ModTime(), f)
}

// PixelRelay://scan?host=HOSTURL&album=ALBUMNAME&privatekey=PRIVATEKEY
func createQR(album, key string) (qrname string) {

	qrname = strings.Join([]string{genQRName(album, key), "png"}, ".")
	qrpath := utils.ImageCfg.QR()
	qrtemp := strings.Join([]string{qrpath, qrname}, "")

	if Exists(qrtemp) {
		return
	}

	u := &url.URL{
		Scheme: "PixelRelay",
		Host:   "scan",
		Path:   "",
	}

	params := url.Values{
		"host":       {strings.Join([]string{utils.AppCfg.Url(), "up"}, "/")},
		"album":      {album},
		"privatekey": {key},
	}

	u.RawQuery = params.Encode()

	c, err := qr.Encode(u.String(), qr.M)

	if err != nil {
		log.Println("qr err: ", err)
	}

	out, err := os.Create(qrtemp)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = out.Write(c.PNG())
	if err != nil {
		panic(err)
	}

	return
}

func genQRName(album, key string) (sha string) {
	str := strings.Join([]string{album, key}, "//")
	bv := []byte(str)

	hasher := sha1.New()
	hasher.Write(bv)
	sha = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	sha = strings.Replace(sha, "=", "", -1)
	return
}
