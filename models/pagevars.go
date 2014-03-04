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

func (p *Page) SetTitle(title ...string) {
	if title[0] == "" {
		p.Title = utils.AppCfg.Title()
		return
	}
	ntitle := strings.Join(title, Seperator)
	p.Title = strings.Join([]string{utils.AppCfg.Title(), ntitle}, Seperator)
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
