package forms

type Image struct {
	Version          string `form:"version" json:"version" required`
	User_Email       string `form:"user_email" json:"user_email" required`
	User_Private_Key string `form:"user_private_key" json:"user_private_key" required`
	File_Host        string `form:"file_host" json:"file_host" required`
	File_Album       string `form:"file_album" json:"file_album" required`
	File_Name        string `form:"file_name" json:"file_name" required`
	File_Mime        string `form:"file_mime" json:"file_mime" required`
}
