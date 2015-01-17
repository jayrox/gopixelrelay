package utils

import "time"

func ExpiresHeader() string {
	return time.Now().Add(time.Hour * 168).UTC().Format("Mon, 2 Jan 2006 15:04:05 GMT")
}
