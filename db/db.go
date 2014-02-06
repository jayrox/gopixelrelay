package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"pixelrelay/models"
	"pixelrelay/utils"
)

var DB gorm.DB

func InitDB() gorm.DB {
	// Set up DB connection
	var err error

	sqlConnection := fmt.Sprintf("%s:%s@%s/%s?clientFoundRows=true&charset=UTF8",
		utils.DbCfg.User(), utils.DbCfg.Pass(), utils.DbCfg.Host(), utils.DbCfg.Name())
	DB, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	if utils.AppCfg.Debug() {
		Logger(&DB, true)
	}

	return DB
}

func Logger(db *gorm.DB, enable bool) {
	db.LogMode(enable)
}

func MigrateDB(db *gorm.DB) {
	fmt.Println("updating tables")
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Album{})
	db.AutoMigrate(models.Image{})
	db.AutoMigrate(models.Uploader{})
	fmt.Println("completed updating tables")
}

// GetAllAlbumImages returns all Images in the album database
func GetAllAlbumImages(db *gorm.DB, album string) []models.Image {
	var images []models.Image
	db.Where("album = ?", album).Find(&images)
	fmt.Println(images)
	return images
}

// GetAllAlbums returns all albums
func GetAllAlbums(db *gorm.DB) []models.Album {
	var albums []models.Album
	db.Where("name != ''").Find(&albums)
	return albums
}

// GetAlbum returns album
func GetAlbum(db *gorm.DB, albumname string) models.Album {
	var album models.Album
	fmt.Println(albumname)
	db.Where("name = ?", albumname).Find(&album)
	return album
}

// Add new album image
func AddImage(db *gorm.DB, image models.Image) models.Image {
	db.NewRecord(&image)
	db.Save(&image)
	db.NewRecord(&image)
	return image
}

// Add upload to uploader history/logging
func AddUpload(db *gorm.DB, upload models.Uploader) {
	db.NewRecord(&upload)
	db.Save(&upload)
	db.NewRecord(&upload)
}

// Add new album
func AddAlbum(db *gorm.DB, album models.Album) {
	db.NewRecord(&album)
	db.Save(&album)
	db.NewRecord(&album)
}

func DropTables(db *gorm.DB) {
	fmt.Println("dropping tables")
	db.DropTable(models.Album{})
	db.DropTable(models.Image{})
	db.DropTable(models.Uploader{})
	db.DropTable(models.User{})
}

func AddTables(db *gorm.DB) {
	fmt.Println("adding tables")
	db.CreateTable(models.Album{})
	db.CreateTable(models.Image{})
	db.CreateTable(models.Uploader{})
	db.CreateTable(models.User{})
}

// Get first image in album for album thumbnail
func FirstImage(db *gorm.DB, album string) []models.Image {
	var image []models.Image
	db.First(&image, "album = ?", album)
	return image
}
