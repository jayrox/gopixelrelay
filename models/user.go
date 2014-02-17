package models

type User struct {
	Id        int64
	Name      string `sql:"type:varchar(255);"`
	Password  string `sql:"type:varchar(255);"`
	Salt      string `sql:"type:varchar(255);"`
	Timestamp int64
}
