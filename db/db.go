package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"log"
	"pixelrelay/models"
	"pixelrelay/utils"
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
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Album{})
	db.AutoMigrate(models.Image{})
	fmt.Println("completed updating tables")
}

// GetAllAlbumImages returns all Images in the album database
func GetAllAlbumImages(db *gorm.DB, album string) []models.Image {
    var images []models.Image
    db.Where("album = ?", album).Find(&images)
    return images
}

// GetAllAlbums returns all albums
func GetAllAlbums(db *gorm.DB) []models.Album {
    var albums []models.Album
    db.Where("name != ''").Find(&albums)
    return albums
}

// Add new album image
func AddImage(db *gorm.DB, image models.Image) {
	db.NewRecord(&image)
	db.Save(&image)
	db.NewRecord(&image)
	fmt.Println("add image: ", &image)
}

// Add new album
func AddAlbum(db *gorm.DB, album models.Album) {
	db.NewRecord(&album)
	db.Save(&album)
	db.NewRecord(&album)
	fmt.Println("add album: ", &album)
}

func DropTables(db *gorm.DB) {
	fmt.Println("dropping tables")
	db.DropTable(models.Album{})
	db.DropTable(models.Image{})
	db.DropTable(models.User{})
}

func AddTables(db *gorm.DB) {
	fmt.Println("adding tables")
	db.CreateTable(models.Album{})
	db.CreateTable(models.Image{})
	db.CreateTable(models.User{})
}

func FirstImage(db *gorm.DB, album string) []models.Image {
	fmt.Println("getting first album image")
	
	var image []models.Image
	db.First(&image, "album = ?", album)
	return image
}