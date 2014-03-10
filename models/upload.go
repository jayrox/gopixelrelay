package models

import "mime/multipart"

type ImageUpload struct {
	File       *multipart.FileHeader `form:"uploaded_file"`
	Version    string                `form:"version"`
	Email      string                `form:"user_email"`
	PrivateKey string                `form:"user_private_key"`
	Host       string                `form:"file_host"`
	Album      string                `form:"file_album"`
	Name       string                `form:"file_name"`
	Mime       string                `form:"file_mime"`
}
