package utils

// https://github.com/3d0c/skeleton/blob/master/utils/appcfg.go

var AppCfg *ConfigScheme

func init() {
        AppCfg = &ConfigScheme{}
}

type ConfigScheme struct {
        App struct {
                Secretkey string `json:"secretkey"`
                ListenOn  string `json:"listen_on"`
        } `json:"application"`
}

type DBScheme struct {
        App struct {
                Spec string `json:"spec"`
                Db_name  string `json:"db_name"`
				Db_user  string `json:"db_user"`
				Db_pass  string `json:"db_pass"`
        } `json:"database"`
}

func (this *ConfigScheme) SecretKey() string {
        return this.App.Secretkey
}

func (this *ConfigScheme) ListenOn() string {
        return this.App.ListenOn
}