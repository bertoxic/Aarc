package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"
)

func GenerateVerificationCode() (string, time.Time) {
	b := make([]byte, 8)
	n, err := rand.Read(b)
	log.Println(n)
	if err != nil {
		log.Println(err)
	}
	expiry := time.Now().Add(30 * time.Minute)
//	log.Println(base64.URLEncoding.EncodeToString(b))

	return base64.URLEncoding.EncodeToString(b), expiry
}