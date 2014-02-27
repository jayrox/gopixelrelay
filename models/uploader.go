package models

type Uploader struct {
	Id        int64
	Email     string `sql:"not null;type:varchar(255);"`
	Timestamp int64
}
