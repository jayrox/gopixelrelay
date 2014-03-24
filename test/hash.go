package test

import (
	"log"
	"strconv"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	hashids "github.com/speps/go-hashids"

	"pixelrelay/encoder"
	"pixelrelay/models"
	"pixelrelay/utils"
)

type testHash struct {
	Ids     []int
	Encrypt []string
	Decrypt []int
}

func Hash(args martini.Params, su models.User, r render.Render, p *models.Page) {
	var th testHash

	id, err := strconv.ParseInt(args["id"], 10, 10)
	if err != nil {
		log.Println(err)
		th.Ids = []int{0, 1, 2, 3, 5, 8, 13, 21, 34}
	} else {
		th.Ids = []int{int(id)}
	}
	log.Println("id: ", id)

	h := hashids.New()
	h.Salt = utils.AppCfg.SecretKey()
	h.MinLength = 6

	for _, i := range th.Ids {
		e, err := h.Encrypt([]int{i})
		if err != nil {
			log.Println(err)
		}
		th.Encrypt = append(th.Encrypt, e)

		d := h.Decrypt(e)
		log.Printf("i: %d, e: %s, d: %d\n", i, e, d[0])
		th.Decrypt = append(th.Decrypt, d[0])
	}

	p.SetUser(su)
	p.SetTitle("")
	p.Data = th
	encoder.Render(p.Encoding, 200, "test/hash", p, r)
}
