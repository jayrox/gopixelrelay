package models

type UserSession struct {
	Id         int64
	UserId     int64
	SessionKey string `sql:"type:varchar(255);"`
	Active     bool
	Timestamp  int64
}
