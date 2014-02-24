package db

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

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
		Logger(&DB, utils.DbCfg.Debug())
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
	db.AutoMigrate(models.UserSession{})
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

// GetAllAlbumsByUserId returns all albums owned by Id
func GetAllAlbumsByUserId(db *gorm.DB, uid int64) []models.Album {
	var albums []models.Album
	db.Where("name != '' and user = ?", uid).Find(&albums)
	return albums
}

// GetAlbum returns album
func GetAlbum(db *gorm.DB, albumname string) models.Album {
	var album models.Album
	db.Where("name = ?", albumname).Find(&album)
	return album
}

// GetAlbum returns album
func GetAlbumByUserId(db *gorm.DB, albumname string, uid int64) models.Album {
	var album models.Album
	db.Where("name = ? and user = ?", albumname, uid).Find(&album)
	return album
}

func SetAlbumPrivacy(db *gorm.DB, uid int64, albumname string, state bool) {
	fmt.Println("id: ", uid)
	fmt.Println("albumname: ", albumname)
	fmt.Println("state: ", state)

	var udb models.Album
	album := GetAlbumByUserId(db, albumname, uid)
	fmt.Println("album: ", album)
	album.Private = state
	db.Model(&udb).Where("id = ? and user = ?", album.Id, uid).Limit(1).Updates(&album)
	//return user

	return
}

// Add new album image
func AddImage(db *gorm.DB, image models.Image) models.Image {
	db.NewRecord(&image)
	db.Save(&image)
	db.NewRecord(&image)
	return image
}

// Tag Image
func TagImage(db *gorm.DB, tag, name string) models.ImageTag {
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
	db.DropTable(models.UserSession{})
}

func AddTables(db *gorm.DB) {
	fmt.Println("adding tables")
	db.CreateTable(models.Album{})
	db.CreateTable(models.Image{})
	db.CreateTable(models.ImageTag{})
	db.CreateTable(models.Tag{})
	db.CreateTable(models.Uploader{})
	db.CreateTable(models.User{})
	db.CreateTable(models.UserSession{})
}

// Get first image in album for album thumbnail
func FirstImage(db *gorm.DB, album string) []models.Image {
	var image []models.Image
	db.First(&image, "album = ?", album)
	return image
}

func GetUserByEmail(db *gorm.DB, email string) models.User {
	var user models.User
	db.Where("email = ?", email).Find(&user)
	return user
}

func GetUserById(db *gorm.DB, id int64) models.User {
	var user models.User
	db.Where("id = ?", id).Find(&user)
	return user
}

func GetUserByIdSessionKey(db *gorm.DB, uid int64, sessionkey string) models.UserSession {
	var usersession models.UserSession
	
	db.Where("user_id = ? and session_key = ?", uid, sessionkey).Find(&usersession)
	return usersession
}

func GetUserIdByUserName(db *gorm.DB, auser string) int64 {
	var user models.User
	
	db.Where("user_name = ?", auser).Find(&user)
	return user.Id
}

func GetUserByUserName(db *gorm.DB, auser string) models.User {
	var user models.User
	
	db.Where("user_name = ?", auser).Find(&user)
	return user
}

func InsertUser(db *gorm.DB, user models.User) models.User {
	db.NewRecord(&user)
	db.Save(&user)
	db.NewRecord(&user)
	return user
}

func UpdateUser(db *gorm.DB, user models.User) models.User {
	var udb models.User
	db.Model(&udb).Where("id = ?", user.Id).Limit(1).Updates(&user)
	return user
}

func CreateSession(db *gorm.DB, session models.UserSession) {
	db.Save(&session)
	return
}

func DestroySession(db *gorm.DB, uid int64, sessionkey string) {
	var usersession models.UserSession
	
	db.Model(&usersession).Where("user_id = ? and session_key = ?", uid, sessionkey).Limit(1).Updates(models.UserSession{UserId: uid, Active: false, Timestamp: time.Now().Unix()})
	return
}