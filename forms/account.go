package forms

import (
	"pixelrelay/models"
)

type Account struct {
	models.User `form:"inherit"`
	Confirm     string `form:"confirm" attr:"type:password;label:Confirm Password;value:input;required" required`
	Submit      string `form:"save" attr:"type:submit;value:Save"`
}
