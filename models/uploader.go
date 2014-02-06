package models

type Uploader struct {
	Id        int64
	User      int64
	Image     int64  `sql:"not null;unique;"`
	Email     string `sql:"not null;type:varchar(255);"`
	Timestamp int64
}
