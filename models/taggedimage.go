package models

type TaggedImage struct {
	Id      int64
	Name    string
	Tag     string
	Trashed bool
}
