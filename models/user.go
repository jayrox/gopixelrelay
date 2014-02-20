package models

type User struct {
	Id        int64
	Name      string `sql:"type:varchar(255);"`
	Email     string `sql:"type:varchar(255);" form:"email" json:"email" attr:"type:email;placeholder:Email;label:Email;value:input;required" required`
	Password  string `sql:"type:varchar(255);" form:"password" json:"password" attr:"type:password;label:Password;value:input;required" required`
	Salt      string `sql:"type:varchar(255);"`
	Timestamp int64
}
