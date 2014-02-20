package models

type UserSession struct {
	Id        int64
	UserId    int64
	Key       string `sql:"type:varchar(255);"`
	Active    bool
	Timestamp int64
}
