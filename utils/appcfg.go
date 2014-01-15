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
		Secretkey string `json:"secretkey"`
		ListenOn  string `json:"listen_on"`
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
		Root   string `json:"root"`
		Thumbs string `json:"thumbs"`
	} `json:"image"`
}

// App Config
func (this *ConfigScheme) SecretKey() string {
	return this.App.Secretkey
}

func (this *ConfigScheme) ListenOn() string {
	return this.App.ListenOn
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
