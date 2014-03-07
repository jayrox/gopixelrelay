package models

type TagList struct {
	Id    int64 `json:"-"`
	TagId int64
	Tag   string
}
