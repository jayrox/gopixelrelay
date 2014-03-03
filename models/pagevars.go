package models

import (
	"strings"

	"pixelrelay/utils"
)

type Page struct {
	Url       string
	User      User
	SiteTitle string
	Title     string
}

func (p *Page) SetUser(user User) {
	p.User = user
}

func (p *Page) SetUrl(url string) {
	p.Url = url
}

func (p *Page) SetTitle(title string) {
	p.Title = strings.Join([]string{utils.AppCfg.Title(), title}, " :: ")
}

func (p *Page) SetSiteTitle(title string) {
	p.SiteTitle = title
}

func (p *Page) SetDefaults() {
	p.SiteTitle = utils.AppCfg.Title()
	p.Title = utils.AppCfg.Title()
	p.Url = utils.AppCfg.Url()
}

func InitPage(p *Page) *Page {
	p.SetDefaults()
	return p
}
