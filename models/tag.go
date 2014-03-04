package models

type Tag struct {
	Id   int64  `form:"-"`
	Name string `form:"tag" sql:"not null;unique;type:varchar(255);" attr:"type:text;placeholder:Tag;label:Tag;value:input;required" required`
}
