package encoder

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/models"
)

// The regex to check for the requested format (allows an optional trailing
// slash).
var rxExt = regexp.MustCompile(`(\.(?:html|json))\/?$`)

// MapEncoder intercepts the request's URL, detects the requested format,
// and injects the correct encoder dependency for this request. It rewrites
// the URL to remove the format extension, so that routes can be defined
// without it.
func MapEncoder(c martini.Context, w http.ResponseWriter, r *http.Request, p *models.Page) {
	// Get the format extension
	matches := rxExt.FindStringSubmatch(r.URL.Path)

	ft := ".html"
	if len(matches) > 1 {
		// Rewrite the URL without the format extension
		l := len(r.URL.Path) - len(matches[1])
		if strings.HasSuffix(r.URL.Path, "/") {
			l--
		}
		r.URL.Path = r.URL.Path[:l]
		ft = matches[1]
	}
	// Inject the requested encoder
	switch ft {
	case ".json":
		p.SetEncoding("json")
	default:
		p.SetEncoding("html")
	}
}

func Render(encoding string, code int, layout string, vars *models.Page, r render.Render) {
	switch encoding {

	case "json":
		r.JSON(code, vars.Data)
	default:
		r.HTML(code, layout, vars)
	}
}
