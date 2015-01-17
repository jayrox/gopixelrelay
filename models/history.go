package models

type History struct {
	Id        int64
	User      int64
	Image     int64 `sql:"not null;unique;"`
	Uploader  int64 `sql:"not null;"`
	Timestamp int64
}
