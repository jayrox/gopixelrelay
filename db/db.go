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

type Dbh struct {
	DB gorm.DB
}

func Init(db *Dbh) *Dbh {
	// Set up DB connection
	var err error

	sqlConnection := fmt.Sprintf("%s:%s@%s/%s?clientFoundRows=true&charset=UTF8",
		utils.DbCfg.User(), utils.DbCfg.Pass(), utils.DbCfg.Host(), utils.DbCfg.Name())
	db.DB, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	if utils.AppCfg.Debug() {
		db.Logger(utils.DbCfg.Debug())
	}

	db.DB.DB().SetMaxIdleConns(10)
	db.DB.DB().SetMaxOpenConns(100)
	db.DB.DB().Ping()

	return db
}

func (db *Dbh) Logger(enable bool) {
	db.DB.LogMode(enable)
}

/****************
*
*  Albums
*
 */

// Add new album
func (db *Dbh) AddAlbum(album models.Album) {
	db.DB.NewRecord(&album)
	db.DB.Save(&album)
	db.DB.NewRecord(&album)
}

// GetAlbum returns album
func (db *Dbh) GetAlbum(name string) (album models.Album) {
	db.DB.Where("name = ?", name).Find(&album)
	return
}

// GetAllAlbums returns all albums
func (db *Dbh) GetAllAlbums() (albums []models.Album) {
	db.DB.Where("name != ''").Find(&albums)
	return
}

// GetAlbum returns album
func (db *Dbh) GetAlbumByUserId(name string, uid int64) (album models.Album) {
	db.DB.Where("name = ? and user = ?", name, uid).Find(&album)
	return
}

// GetAllAlbumsByUserId returns all albums owned by Id
func (db *Dbh) GetAllAlbumsByUserId(uid int64) (albums []models.Album) {
	db.DB.Where("name != '' and user = ?", uid).Find(&albums)
	return
}

// GetAllAlbumImages returns all Images in the album database
func (db *Dbh) GetAllAlbumImages(album string) (images []models.Image) {
	db.DB.Where("album = ?", album).Find(&images)
	return
}

// SetAlbumPrivacy changes private state
func (db *Dbh) SetAlbumPrivacy(uid int64, name string, state bool) {
	var udb models.Album
	album := db.GetAlbumByUserId(name, uid)

	album.Private = state
	db.DB.Model(&udb).Where("id = ? and user = ?", album.Id, uid).Limit(1).Updates(&album)

	return
}

/****************
*
*  Images
*
 */

// Add new album image
func (db *Dbh) AddImage(image models.Image) models.Image {
	db.DB.NewRecord(&image)
	db.DB.Save(&image)
	db.DB.NewRecord(&image)
	return image
}

// Get first image in album for album thumbnail
func (db *Dbh) FirstImageByAlbum(album string) (image []models.Image) {
	db.DB.First(&image, "album = ?", album)
	return
}

// Get first image by image name
func (db *Dbh) FirstImageByName(name string) (image models.Image) {
	db.DB.First(&image, "name = ?", name)
	return
}

/****************
*
*  Tags
*
 */

// Tag Image
func (db *Dbh) TagImage(tag, name string) (imagetag models.ImageTag) {
	// Get tag id or create tag
	var mTag models.Tag
	db.DB.Where(models.Tag{Name: tag}).FirstOrCreate(&mTag)
	//db.FirstOrCreate(&mTag, models.Tag{Name: tag})

	// Get Image Id
	var image models.Image
	db.DB.First(&image, "name = ?", name)

	// Add tag to image
	imagetag = models.ImageTag{ImgId: image.Id, TagId: mTag.Id}

	// Save tag
	db.DB.NewRecord(&imagetag)
	db.DB.Save(&imagetag)
	db.DB.NewRecord(&imagetag)

	return
}

// Get Images with Tag
func (db *Dbh) GetImagesWithTag(tag string) (images []models.TaggedImage) {
	db.DB.Table("images").Select("images.id as image_id, images.name as name, tags.name as tag").Joins("LEFT JOIN image_tags ON (image_tags.img_id = images.id) LEFT JOIN tags ON (image_tags.tag_id = tags.id AND image_tags.img_id = images.id)").Where("tags.name = ?", tag).Order("images.id ASC").Scan(&images)
	log.Println(images)
	return
}

// Get all tags
func (db *Dbh) GetAllTags() (tags []models.Tag) {
	db.DB.Where("name != ''").Find(&tags)
	return
}

// Get all tags
func (db *Dbh) GetAllTagsByImageId(id int64) (tags []models.TagList) {
	db.DB.Table("image_tags").Select("image_tags.img_id as id, tags.id as tag_id, tags.name as tag").Joins("LEFT JOIN tags ON (image_tags.tag_id = tags.id)").Where("image_tags.img_id = ?", id).Scan(&tags)
	//db.DB.Where("img_id = ?", id).Find(&tags)
	return
}

/****************
*
*  Tables
*
 */

func (db *Dbh) AddTables() {
	log.Println("adding tables")
	db.DB.CreateTable(models.Album{})
	db.DB.CreateTable(models.Image{})
	db.DB.CreateTable(models.ImageTag{})
	db.DB.CreateTable(models.Tag{})
	db.DB.CreateTable(models.Uploader{})
	db.DB.CreateTable(models.User{})
	db.DB.CreateTable(models.UserSession{})
}

func (db *Dbh) DropTables() {
	log.Println("dropping tables")
	db.DB.DropTable(models.Album{})
	db.DB.DropTable(models.Image{})
	db.DB.DropTable(models.ImageTag{})
	db.DB.DropTable(models.Tag{})
	db.DB.DropTable(models.Uploader{})
	db.DB.DropTable(models.User{})
	db.DB.DropTable(models.UserSession{})
}

func (db *Dbh) MigrateDB() {
	log.Println("updating tables")
	db.DB.AutoMigrate(models.Album{})
	db.DB.AutoMigrate(models.Image{})
	db.DB.AutoMigrate(models.ImageTag{})
	db.DB.AutoMigrate(models.Tag{})
	db.DB.AutoMigrate(models.Uploader{})
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.UserSession{})
	log.Println("completed updating tables")
}

/****************
*
*  Users
*
 */

func (db *Dbh) GetUserByEmail(email string) (user models.User) {
	db.DB.Where("email = ?", email).Find(&user)
	return
}

func (db *Dbh) GetUserById(id int64) (user models.User) {
	db.DB.Where("id = ?", id).Find(&user)
	return
}

func (db *Dbh) GetUserByIdSessionKey(uid int64, sessionkey string) (usersession models.UserSession) {
	db.DB.Where("user_id = ? and session_key = ?", uid, sessionkey).Find(&usersession)
	return
}

func (db *Dbh) GetUserIdByUserName(auser string) int64 {
	var user models.User

	db.DB.Where("user_name = ?", auser).Find(&user)
	return user.Id
}

func (db *Dbh) GetUserByUserName(auser string) (user models.User) {
	db.DB.Where("user_name = ?", auser).Find(&user)
	return
}

func (db *Dbh) InsertUser(user models.User) models.User {
	db.DB.NewRecord(&user)
	db.DB.Save(&user)
	db.DB.NewRecord(&user)
	return user
}

func (db *Dbh) UpdateUser(user models.User) models.User {
	var udb models.User
	db.DB.Model(&udb).Where("id = ?", user.Id).Limit(1).Updates(&user)
	return user
}

/****************
*
*  Sessions
*
 */

func (db *Dbh) CreateSession(session models.UserSession) {
	db.DB.Save(&session)
	return
}

func (db *Dbh) DestroySession(uid int64, sessionkey string) {
	var usersession models.UserSession
	db.DB.Model(&usersession).Where("user_id = ? and session_key = ?", uid, sessionkey).Limit(1).Updates(models.UserSession{UserId: uid, Active: false, Timestamp: time.Now().Unix()})
	return
}

/****************
*
*  History
*
 */

// Add upload to uploader list
func (db *Dbh) AddUploader(upload models.Uploader) {
	db.DB.NewRecord(&upload)
	db.DB.Save(&upload)
	db.DB.NewRecord(&upload)
}

//
func (db *Dbh) GetUploaderByEmail(email string) (user models.User) {
	db.DB.Where("email = ?", email).Find(&user)
	return
}
