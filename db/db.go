package db

import (
	"fmt"
	"log"
	"pixelrelay/models"
	"pixelrelay/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() gorm.DB {
	// Set up DB connection
	var err error
	var db gorm.DB
	sqlConnection := fmt.Sprintf("%s:%s@%s/%s?clientFoundRows=true&charset=UTF8", 
					 utils.DbCfg.User(), utils.DbCfg.Pass(), utils.DbCfg.Host(), utils.DbCfg.Name())
	db, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
			log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	
	db.LogMode(true)

	return db
}

func MigrateDB(db *gorm.DB) {
	fmt.Println("updating tables")
	db.AutoMigrate(models.Users{})
	db.AutoMigrate(models.Albums{})
	db.AutoMigrate(models.Images{})
	fmt.Println("completed updating tables")
}

// GetAllAlbumImages returns all Images in the album database
func GetAllAlbumImages(db *gorm.DB, album string) []models.Images {
    var images []models.Images
    db.Where("album = ?", album).Find(&images)
    return images
}

// GetAllAlbums returns all albums
func GetAllAlbums(db *gorm.DB) []models.Albums {
    var albums []models.Albums
    db.Where("name != ''").Find(&albums)
    return albums
}

// Add new album image
func AddImage(db *gorm.DB, image models.Images) {
	db.NewRecord(&image)
	db.Save(&image)
	db.NewRecord(&image)
	fmt.Println("add image: ", &image)
}

// Add new album
func AddAlbum(db *gorm.DB, album models.Albums) {
	db.NewRecord(&album)
	db.Save(&album)
	db.NewRecord(&album)
	fmt.Println("add album: ", &album)
}