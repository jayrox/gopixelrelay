package models

import (
	"log"

	hashids "github.com/speps/go-hashids"
)

type HashID struct {
	Ids []int
	Enc string
	Dec []int
	hid *hashids.HashID
}

func (h *HashID) Init(salt string, min_length int) {
	//h.hid = hashids.New()
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = min_length
	h.hid = hashids.NewWithData(hd)
}

func (h *HashID) Salt(salt string) {
	//h.hid.Salt = salt
	log.Println("fix salt")
}

func (h *HashID) MinLen(min int) {
	//h.hid.MinLength = min
	log.Println("fix minlen")
}

func (h *HashID) SetIds(ids ...int) {
	h.Ids = ids
}

func (h *HashID) AddId(id int) {
	h.Ids = append(h.Ids, id)
}

func (h *HashID) Encrypt() string {
	e, err := h.hid.Encode(h.Ids)
	if err != nil {
		log.Println(err)
	}
	h.Enc = e
	return e
}

func (h *HashID) Encrypted(e string) {
	h.Enc = e
}

func (h *HashID) Decrypt() (dec []int) {
	dec = h.hid.Decode(h.Enc)
	h.Dec = dec
	return
}
