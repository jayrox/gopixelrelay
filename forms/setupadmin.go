package forms

type SetupAdmin struct {
	Name     string `form:"name" json:"name" attr:"type:text;maxlength:255;label:Name;placeholder:Name;value:input;required" required`
	Email    string `form:"email" json:"email" attr:"type:email;placeholder:Email;label:Email;value:input;required" required`
	Password string `form:"password" json:"password" attr:"type:password;label:Password;value:input;required" required`
	Confirm  string `form:"confirm" json:"confirm" attr:"type:password;label:Verify Password;value:input;required" required`
	CSRF     string `form:"token" json:"token" attr:"type:hidden;required" required`
	Submit   string `form:"save" attr:"type:submit;value:Save"`
}

/*
func (sa forms.SetupAdmin) Validate(errors *binding.Errors) {   //, req *http.Request
	if len(sa.Name) <= 0 {
        errors.Fields["name"] = "Name must be set"
    } else if len(sa.Name) > 255 {
        errors.Fields["name"] = "Too long; maximum 255 characters"
    }
    if len(sa.Email) <= 0 {
        errors.Fields["email"] = "Email must be at set"
    }
    if len(sa.Password) <= 0 {
        errors.Fields["password"] = "Password must be at set"
    }
    if len(sa.Confirm) <= 0 {
        errors.Fields["confirm"] = "Password confirmation must be at set"
    }
	if sa.Password != sa.Confirm {
		errors.Fields["confirm"] = "Passwords must match"
	}
}
*/
