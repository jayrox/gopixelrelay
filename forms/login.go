package forms

type Login struct {
	Id       int64  `form:"-"`
	Email    string `form:"email" json:"email" attr:"type:email;placeholder:Email;label:Email;value:input;required" required`
	Password string `form:"password" json:"password" attr:"type:password;label:Password;value:input;required" required`
	Submit   string `form:"login" attr:"type:submit;value:Login"`
}
