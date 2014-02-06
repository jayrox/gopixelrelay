package models

type Image struct {
	Id          int64
	User        int64
	Name        string `sql:"not null;unique;type:varchar(255);"`
	Album       string `sql:"type:varchar(255);"`
	Title       string `sql:"type:varchar(255);"`
	Description string `sql:"type:varchar(1024);"`
	Timestamp   int64
}
