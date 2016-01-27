package main

import (
	"encoding/base64"
	"log"
	"strings"

	"github.com/youngsterxyf/xiaBase64"
)

func testIt(rawStr string) {
	log.Printf("*** %s ***\n", rawStr)
	xiaBase64Encode := string(xiaBase64.Encode([]byte(rawStr)))
	stdBase64Encode := base64.StdEncoding.EncodeToString([]byte(rawStr))
	if strings.Compare(xiaBase64Encode, stdBase64Encode) != 0 {
		log.Println("xiaBase64Encode is wrong!")
		log.Printf("xia: %s , std: %s", xiaBase64Encode, stdBase64Encode)
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
		log.Println("xiaBase64Decode is wrong!")
		log.Printf("xia: %s , std: %s", xiaBase64Decode, stdBase64Decode)
		return
	}
	log.Println("xiaBase64Decode is right!")
}

func main() {
	rawStr := "您好，世界！"
	testIt(rawStr)
	rawStr = "Hello, world!"
	testIt(rawStr)
	rawStr = "Hello, 世界！"
	testIt(rawStr)
	rawStr = "https://github.com/youngsterxyf/xiaBase64/blob/master/example/test.go"
	testIt(rawStr)
}
