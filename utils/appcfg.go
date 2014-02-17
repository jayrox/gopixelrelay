package utils

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
		ListenOnSetup       string `json:"listen_on_setup"`
		SecretKey       	string `json:"secretkey"`
		Url                 string `json:"url"`
	} `json:"application"`
}

type DBScheme struct {
	DB struct {
		Host string `json:"host"`
		Name string `json:"name"`
		User string `json:"user"`
		Pass string `json:"pass"`
	} `json:"database"`
}

type ImageScheme struct {
	Image struct {
		Root      string `json:"root"`
		Secretkey string `json:"secretkey"`
		Thumbs    string `json:"thumbs"`
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

func (this *ConfigScheme) ListenOnSetup() string {
	return this.App.ListenOnSetup
}

func (this *ConfigScheme) Url() string {
	return this.App.Url
}

func (this *ConfigScheme) SecretKey() string {
	return this.App.SecretKey
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

// Image Storage Config
func (this *ImageScheme) Root() string {
	return this.Image.Root
}

func (this *ImageScheme) Thumbs() string {
	return this.Image.Thumbs
}

func (this *ImageScheme) SecretKey() string {
	return this.Image.Secretkey
}
