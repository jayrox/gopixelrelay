package models

type Album struct {
	Id int64
	Name string `sql:"not null;unique;type:varchar(255);"`
	User int64
	Privatekey string `sql:"type:varchar(255);"`
	Private bool
	Timestamp int64
}
