package models

type AlbumList struct {
	Id      int64
	Name    string
	Poster  string
	Private bool
	Owner   int64
}
