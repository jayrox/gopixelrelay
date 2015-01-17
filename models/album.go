package models

import (
	"log"
	"time"
)

type Album struct {
	Id          int64
	Name        string `sql:"not null;unique;type:varchar(255);"`
	Description string
	User        int64
	Poster      string
	Privatekey  string `sql:"type:varchar(255);"`
	Private     bool
	Timestamp   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func (a *Album) AfterCreate() (err error) {
	log.Printf("New Album: %# v\n", a)
	return
}

func (a *Album) DefaultDescription() {
	const layout = "Auto-created 2 January 2006"
	t := time.Now()
	a.Description = t.Format(layout)
}
