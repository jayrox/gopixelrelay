package models

type User struct {
	Id        int64  `json:"-"`
	Name      string `sql:"type:varchar(255);" form:"name"`
	UserName  string `sql:"type:varchar(255);" form:"username" json:"-"`
	Email     string `sql:"type:varchar(255);" form:"email" json:"-" attr:"type:email;placeholder:Email;label:Email;value:input;required" required`
	Password  string `sql:"type:varchar(255);" form:"password" json:"-" attr:"type:password;label:Password;value:input;required" required`
	Salt      string `sql:"type:varchar(255);" form:"salt" json:"-"`
	Timestamp int64  `json:"-"`
}
