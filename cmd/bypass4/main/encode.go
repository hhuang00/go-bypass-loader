package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base32"
	"encoding/base64"
	"fmt"
)

// 填充字符串（末尾）
func PaddingText1(str []byte, blockSize int) []byte {
	//需要填充的数据长度
	paddingCount := blockSize - len(str)%blockSize
	//填充数据为：paddingCount ,填充的值为：paddingCount
	paddingStr := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	newPaddingStr := append(str, paddingStr...)
	//fmt.Println(newPaddingStr)
	return newPaddingStr
}

// ---------------DES加密--------------------
func EncyptogAES(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(nil)
		return nil
	}
	src = PaddingText1(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src
}

func main() {

	shellcode := []byte{}
	str := base64.StdEncoding.EncodeToString(shellcode)

	//密钥长度16
	key := []byte("AofqwwWicshoiqQq")

	// XOR 操作
	xordMessage := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		xordMessage[i] = str[i] ^ 0xff
	}

	src := EncyptogAES(xordMessage, key)
	message := base32.HexEncoding.EncodeToString(src)
	fmt.Println(message)
}
