package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var DB = map[string]string{
	"User1Key": "User1Secret",
	"USer2Key": "USer2Secret",
}

func Server(apiKey, sign string, data []byte){
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC )
}

func main()  {
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	data := []byte("data")
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte("data"))
	sign := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign)
	fmt.Println(sign)

	Server(apiKey, sign, data)
}