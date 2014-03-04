package forms

import (
	"pixelrelay/models"
)

type Tag struct {
	models.Tag `form:"inherit"`
	Image      string `form:"image" sql:"-" attr:"type:hidden;value:input"`
	Submit     string `form:"submit" sql:"-" attr:"type:submit;value:Save"`
}
