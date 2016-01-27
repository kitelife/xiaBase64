package main

import (
	"encoding/base64"
	"log"
	"strings"

	"xiaBase64"
)

func main() {
	rawStr := "您好，世界！"

	xiaBase64Encode := string(xiaBase64.Encode([]byte(rawStr)))
	stdBase64Encode := base64.StdEncoding.EncodeToString([]byte(rawStr))
	if strings.Compare(xiaBase64Encode, stdBase64Encode) != 0 {
		log.Fatalln("xiaBase64Encode is wrong!")
		return
	}
	log.Println("xiaBase64Encode is right!")

	xiaBase64Decode, err := xiaBase64.Decode([]byte(xiaBase64Encode))
	if err != nil {
		log.Fatalln(err)
		return
	}
	stdBase64Decode, err := base64.StdEncoding.DecodeString(xiaBase64Encode)
	if err != nil {
		log.Fatalln(err)
	}
	if strings.Compare(string(xiaBase64Decode), string(stdBase64Decode)) != 0 {
		log.Fatalln("xiaBase64Decode is wrong!")
	}
	log.Println("xiaBase64Decode is right!")
}
