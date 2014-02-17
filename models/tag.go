package models

type Tag struct {
	Id   int64
	Name string `sql:"not null;unique;type:varchar(255);"`
}
