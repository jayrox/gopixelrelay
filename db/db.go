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
	db.AutoMigrate(models.Album{})
	db.AutoMigrate(models.Image{})
	db.AutoMigrate(models.ImageTag{})
	db.AutoMigrate(models.Tag{})
	db.AutoMigrate(models.Uploader{})
	db.AutoMigrate(models.User{})
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

// Tag Image
func TagImage(db *gorm.DB, tag string, name string) models.ImageTag {
	// Get tag id or create tag
	var mTag models.Tag
	db.Where(models.Tag{Name: tag}).FirstOrCreate(&mTag)
	//db.FirstOrCreate(&mTag, models.Tag{Name: tag})

	// Get Image Id
	var image models.Image
	db.First(&image, "name = ?", name)

	// Add tag to image
	imagetag := models.ImageTag{ImgId: image.Id, TagId: mTag.Id}

	// Save tag
	db.NewRecord(&imagetag)
	db.Save(&imagetag)
	db.NewRecord(&imagetag)
	
	return imagetag
}

// Get Images with Tag
func GetImagesWithTag(db *gorm.DB, tag string) []models.TaggedImage {
	var images []models.TaggedImage
	//query := db.Exec("SELECT images.id as image_id, images.name as name, tags.name as tag FROM images LEFT JOIN image_tags ON (image_tags.img_id = images.id) LEFT JOIN tags ON (image_tags.tag_id = tags.id AND image_tags.img_id = images.id) WHERE tags.name = ? ORDER BY images.id ASC", tag).Scan(&images)
	db.Table("images").Select("images.id as image_id, images.name as name, tags.name as tag").Joins("LEFT JOIN image_tags ON (image_tags.img_id = images.id) LEFT JOIN tags ON (image_tags.tag_id = tags.id AND image_tags.img_id = images.id)").Where("tags.name = ?", tag).Order("images.id ASC").Scan(&images)
	fmt.Println(images)
	return images
}

// Get all tags
func GetAllTags(db *gorm.DB) []models.Tag {
	var tags []models.Tag
	db.Where("name != ''").Find(&tags)
	return tags
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
	db.DropTable(models.ImageTag{})
	db.DropTable(models.Tag{})
	db.DropTable(models.Uploader{})
	db.DropTable(models.User{})
}

func AddTables(db *gorm.DB) {
	fmt.Println("adding tables")
	db.CreateTable(models.Album{})
	db.CreateTable(models.Image{})
	db.CreateTable(models.ImageTag{})
	db.CreateTable(models.Tag{})
	db.CreateTable(models.Uploader{})
	db.CreateTable(models.User{})
}

// Get first image in album for album thumbnail
func FirstImage(db *gorm.DB, album string) []models.Image {
	var image []models.Image
	db.First(&image, "album = ?", album)
	return image
}
