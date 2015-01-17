package models

/*
import (
	_ "github.com/go-sql-driver/mysql"
)
*/

type Uploader struct {
	Id        int64
	Email     string `sql:"not null;uniquetype:varchar(255);"`
	Timestamp int64
	//	db        *Dbh
}

/*
func (u *Uploader) init(db *Dbh) {
	u.db = db
}

func (u *Uploader) empty() {
	//alter table tablename AUTO_INCREMENT = 1;
	db.DB.Exec("TRUNCATE TABLE uploaders;")
	db.DB.Exec("ALTER TABLE uploaders AUTO_INCREMENT = 1;")

	return
}

func (u Uploader) GetAll() (uploaders []Uploader) {
	db.DB.Where("id > 0").Find(&uploaders)
	return
}
*/
