package utils

import (
	"fmt"
	"net/http"
)

type Response map[string]interface{}

func Verify(res http.ResponseWriter, req *http.Request) {
	pk := req.FormValue("user_private_key")
	a := req.FormValue("album")
	if  pk == "" || pk != ImageCfg.SecretKey() {
		http.Error(res, "Invalid Private Key", http.StatusUnauthorized)
		res.Header().Set("Content-Type", "application/json")
        fmt.Fprint(res, Response{"error": http.StatusUnauthorized, "code": "Invalid Private Key", "name": a})
	}
	
	return 
}