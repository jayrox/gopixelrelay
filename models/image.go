package models

type Image struct {
	Id int64
	Name string `sql:"not null;unique;type:varchar(255);"`
	Album string `sql:"type:varchar(255);"`
	User int64
	Timestamp int64
}