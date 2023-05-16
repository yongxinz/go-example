package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now().Unix()
	secret := "this is secret"

	stringToSign := fmt.Sprintf("%v", timestamp) + "@" + secret

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	signature := hex.EncodeToString(h.Sum(nil))

	fmt.Println(signature)
}
