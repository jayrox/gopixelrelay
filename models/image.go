package models

import "time"

type Image struct {
	Id          int64
	User        int64
	Name        string `sql:"not null;type:varchar(255);"`
	HashId      string `sql:"not null;unique;type:varchar(255);"`
	Album       string `sql:"type:varchar(255);"`
	AlbumId     int64
	Title       string `sql:"type:varchar(255);"`
	Description string `sql:"type:varchar(1024);"`
	Trashed     bool
	Timestamp   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
