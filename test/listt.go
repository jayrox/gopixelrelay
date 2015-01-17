// Route: /test/listt

/*
 Was used to rebuild image thumbnails
*/

package test

import (
	"io/ioutil"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/encoder"
	"pixelrelay/models"
	"pixelrelay/utils"
)

func ListThumb(args martini.Params, su models.User, r render.Render, p *models.Page) {

	files, _ := ioutil.ReadDir(utils.ImageCfg.Root())

	var imageLinks []ImageLink

	for _, f := range files {
		if strings.Contains(f.Name(), ".") && !strings.HasPrefix(f.Name(), ".") {
			imageLinks = append(imageLinks, ImageLink{Title: f.Name(), FileName: f.Name()})
		}
	}

	p.SetUser(su)
	p.SetTitle("List", "Thumbs")
	p.Data = ListVars{Title: "All Images", Description: "Listing of all images in database", ImageLinks: imageLinks}

	encoder.Render(p.Encoding, 200, "test/thumbs", p, r)
}
