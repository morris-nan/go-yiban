package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"github.com/morris-nan/go-yiban/ybe"
)

const (
	keyLength = 16
)

// Decode
/*
 * 128位AES解密方法

 * text 需要解密的字符串
 * key（对应AppSecret值）
 * iv （对应AppID值）

 * return string 解密后的JSON字符串
 * return error 解密时产生的错误
 */
func Decode(text string, key string, iv string) (string, error) {
	var keyValue []byte
	if len(iv) == keyLength {
		keyValue = []byte(key)
	} else {
		keyValue = []byte(key[:keyLength])
	}

	ivSpec := []byte(iv[:keyLength])
	block, err := aes.NewCipher(keyValue)
	if err != nil {
		return "", err
	}

	cipherText, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}

	if len(cipherText)%aes.BlockSize != 0 {
		return "", ybe.ErrBlockSize
	}

	mode := cipher.NewCBCDecrypter(block, ivSpec)
	mode.CryptBlocks(cipherText, cipherText)

	index := bytes.IndexByte(cipherText, 0)
	if index < 0 {
		return string(cipherText), nil
	}

	return string(cipherText[:index]), nil
}
