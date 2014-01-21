package models

type Albums struct {
	Id int64
	Name string `sql:"not null;unique"`
	User int64
	Privatekey string
	Private bool
	Timestamp int64
}
