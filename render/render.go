package render

import (
	"log"

	"github.com/martini-contrib/render"

	"pixelrelay/models"
)

func Render(code int, layout string, vars models.Page, r render.Render) {
	log.Println("\ncode: ", code, "\nlayout: ", layout, "\nvars: ", vars)
	switch vars.Type {

	case "json":
		r.JSON(code, vars)
	default:
		r.HTML(code, layout, vars)
	}
}
