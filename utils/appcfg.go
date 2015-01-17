package utils

import "strings"

// https://github.com/3d0c/skeleton/blob/master/utils/appcfg.go

var AppCfg *ConfigScheme
var DbCfg *DBScheme
var ImageCfg *ImageScheme

func init() {
	AppCfg = &ConfigScheme{}
	DbCfg = &DBScheme{}
	ImageCfg = &ImageScheme{}
}

type ConfigScheme struct {
	App struct {
		MobileAlbumCreation string `json:"mobile_album_creation"`
		Debug               bool   `json:"debug"`
		ListenOn            string `json:"listen_on"`
		TLSListenOn         string `json:"tls_listen_on"`
		ListenOnSetup       string `json:"listen_on_setup"`
		SecretKey           string `json:"secretkey"`
		Url                 string `json:"url"`
		Title               string `json:"title"`
	} `json:"application"`
}

type DBScheme struct {
	DB struct {
		Host  string `json:"host"`
		Name  string `json:"name"`
		User  string `json:"user"`
		Pass  string `json:"pass"`
		Debug bool   `json:"debug"`
	} `json:"database"`
}

type ImageScheme struct {
	Image struct {
		Root      string `json:"root"`
		Secretkey string `json:"secretkey"`
		Thumbs    string `json:"thumbs"`
		QR        string `json:"qr"`
	} `json:"image"`
}

// App Config
func (this *ConfigScheme) MobileAlbumCreation() string {
	return this.App.MobileAlbumCreation
}

func (this *ConfigScheme) Debug() bool {
	return this.App.Debug
}

func (this *ConfigScheme) ListenOn() string {
	return this.App.ListenOn
}

func (this *ConfigScheme) TLSListenOn() string {
	return this.App.TLSListenOn
}

func (this *ConfigScheme) ListenOnSetup() string {
	return this.App.ListenOnSetup
}

func (this *ConfigScheme) Url() string {
	if strings.HasPrefix(this.App.Url, "http://") {
		this.App.Url = strings.Replace(this.App.Url, "http://", "https://", 1)
	}
	if !strings.HasPrefix(this.App.Url, "https://") {
		this.App.Url = strings.Join([]string{"https://", this.App.Url}, "")
	}
	return this.App.Url
}

func (this *ConfigScheme) SecretKey() string {
	return this.App.SecretKey
}

func (this *ConfigScheme) Title() string {
	return this.App.Title
}

// Database Config
func (this *DBScheme) Host() string {
	return this.DB.Host
}

func (this *DBScheme) Name() string {
	return this.DB.Name
}

func (this *DBScheme) User() string {
	return this.DB.User
}

func (this *DBScheme) Pass() string {
	return this.DB.Pass
}

func (this *DBScheme) Debug() bool {
	return this.DB.Debug
}

// Image Storage Config
func (this *ImageScheme) QR() string {
	return this.Image.QR
}

func (this *ImageScheme) Root() string {
	return this.Image.Root
}

func (this *ImageScheme) Thumbs() string {
	return this.Image.Thumbs
}

func (this *ImageScheme) SecretKey() string {
	return this.Image.Secretkey
}
