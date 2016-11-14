package models

import (
	"strings"

	"pixelrelay/utils"
)

type Page struct {
	Encoding  string
	SiteTitle string
	Title     string
	Url       string
	User      User
	Data      interface{}
}

func (p *Page) SetUser(user User) {
	p.User = user
}

func (p *Page) SetUrl(url string) {
	p.Url = url
}

func (p *Page) SetTitle(title ...string) {
	if title[0] == "" {
		p.Title = utils.AppCfg.Title()
		return
	}
	ntitle := strings.Join(title, Separator)
	p.Title = strings.Join([]string{utils.AppCfg.Title(), ntitle}, Separator)
}

func (p *Page) SetSiteTitle(title string) {
	p.SiteTitle = title
}

func (p *Page) SetDefaults() {
	p.SiteTitle = utils.AppCfg.Title()
	p.Title = utils.AppCfg.Title()
	p.Url = utils.AppCfg.Url()
	p.Encoding = "html"
}

func (p *Page) SetEncoding(encoding string) {
	p.Encoding = encoding
}

func InitPage(p *Page) *Page {
	p.SetDefaults()
	return p
}

var (
	// Separator is used to separate the parts of the title
	// Set this to whatever value you prefer; default is " :: "
	Separator = string(" :: ")
)
