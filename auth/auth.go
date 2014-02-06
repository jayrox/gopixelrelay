/*Utility functions and methods for authoriztion*/
package auth

// package originally from
// https://github.com/adnaan/hamster/blob/master/auth.go
/*******************************************************
* The MIT License (MIT)
*
* Copyright (c) 2013 Adnaan Badr
*
* Permission is hereby granted, free of charge, to any person obtaining a copy of
* this software and associated documentation files (the "Software"), to deal in
* the Software without restriction, including without limitation the rights to
* use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
* the Software, and to permit persons to whom the Software is furnished to do so,
* subject to the following conditions:

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
* FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
* COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
* IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
* CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

import (
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	//"errors"
	//"fmt"
	"hash"
	"net/http"
	"strings"
	"sync"
	//"time"
)

//Get Basic user password
func getUserPassword(r *http.Request) (string, string) {

	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 || s[0] != "Basic" {
		return "", ""
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return "", ""
	}
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return "", ""
	}

	return pair[0], pair[1]

}

type Hmac struct {
	hashFunc hash.Hash
	m        sync.Mutex
}

/*Generates hash*/
func (h *Hmac) generateHash(data []byte) []byte {
	h.m.Lock()
	defer h.m.Unlock()

	h.hashFunc.Reset()
	h.hashFunc.Write(data)
	return h.hashFunc.Sum(nil)

}

/*Encrypts the password and returns password hash*/
func (h *Hmac) encrypt(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(h.generateHash(password), cost)
}

func (h *Hmac) compare(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, h.generateHash(password))
}

/*Returns a hmac type*/
func New(hash func() hash.Hash, salt []byte) *Hmac {
	hm := &Hmac{
		hashFunc: hmac.New(hash, salt),
		m:        sync.Mutex{},
	}
	return hm
}

/*Encrypts password. Returns hash+salt*/
func EncryptPassword(password string) (string, string, error) {

	salt, err0 := genUUID(16)
	if err0 != nil {
		return "", "", err0
	}

	hm := New(sha512.New, []byte(salt))
	pass := []byte(password)
	encrypted, err := hm.encrypt(pass, bcrypt.DefaultCost)

	if err != nil {

		return "", "", err
	}

	return string(encrypted), salt, nil
}

/*Match encrypted string*/
func MatchPassword(password string, hash string, salt string) bool {
	p := []byte(password)
	h := []byte(hash)
	s := []byte(salt)
	hm := New(sha512.New, s)
	err := hm.compare(h, p)
	if err != nil {
		return false
	} else {
		//matched!
		return true
	}

}

/*Encode to base64*/
func encodeBase64Token(hexVal string) string {

	token := base64.URLEncoding.EncodeToString([]byte(hexVal))

	return token

}

/*Decode from base64*/
func decodeToken(token string) string {

	hexVal, err := base64.URLEncoding.DecodeString(token)
	if err != nil {

		return ""

	}

	return string(hexVal)

}

/*Generate uuid*/
func genUUID(size int) (string, error) {
	uuid := make([]byte, size)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	uuid[8] = 0x80
	uuid[4] = 0x40

	return hex.EncodeToString(uuid), nil
}

/*
 Return a random 16-byte base64 alphabet string
*/
func randomKey() string {
	k := make([]byte, 12)
	for bytes := 0; bytes < len(k); {
		n, err := rand.Read(k[bytes:])
		if err != nil {
			panic("rand.Read() failed")
		}
		bytes += n
	}
	return base64.StdEncoding.EncodeToString(k)
}