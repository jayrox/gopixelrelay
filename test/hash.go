// Route: /test/hash
// Route: /test/hash/:user_id/:image_id

/*
 Used to test hashid model
*/

package test

import (
	"log"
	"strconv"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"pixelrelay/encoder"
	"pixelrelay/models"
	"pixelrelay/utils"
)

func Hash(args martini.Params, su models.User, r render.Render, p *models.Page) {
	var th models.HashID

	user_id, err := strconv.ParseInt(args["user_id"], 10, 10)
	if err != nil {
		log.Println(err)
	}

	image_id, err := strconv.ParseInt(args["image_id"], 10, 10)
	if err != nil {
		log.Println(err)
	}

	time_now := time.Now().Unix()

	th.Init(utils.AppCfg.SecretKey(), 6)
	th.SetIds(int(user_id), int(image_id), int(time_now))
	enc := th.Encrypt()
	log.Println(enc)
	dec := th.Decrypt()
	log.Println(dec)

	p.SetUser(su)
	p.SetTitle("")
	p.Data = th
	encoder.Render(p.Encoding, 200, "test/hash", p, r)
}
