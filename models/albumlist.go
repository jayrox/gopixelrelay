package models

type AlbumList struct {
	Id          int64
	Name        string
	Poster      string
	Private     bool
	Privatekey  string
	Owner       int64
	Description string
	Editable    bool
}
