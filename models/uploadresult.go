package models

import "encoding/json"

type UploadResult struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
	Name   string `json:"name"`
}

func (ur *UploadResult) SetCode(code int) {
	ur.Code = code
}

func (ur *UploadResult) SetResult(result string) {
	ur.Result = result
}

func (ur *UploadResult) SetName(name string) {
	ur.Name = name
}

func (ur *UploadResult) GetName() string {
	return ur.Name
}

func (ur UploadResult) String() (s string) {
	b, err := json.Marshal(ur)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}
