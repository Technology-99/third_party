package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

//AES	    分组长度(字节)	密钥长度(字节)	加密轮数
//AES-128	16	            16	10
//AES-192	16	            24	12
//AES-256	16	            32	14

// PKCS5填充方式
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

// Zero填充方式
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{0}, padding)

	return append(ciphertext, padtext...)
}

// PKCS5 反填充
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// Zero反填充
func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 加密
func AesEncrypt(encodeBytes []byte, key []byte, iv []byte) ([]byte, error) {
	//encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	//encodeBytes = ZeroPadding(encodeBytes, blockSize

	encodeBytes = PKCS5Padding(encodeBytes, blockSize) //PKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)
	return crypted, nil
	//return base64.StdEncoding.EncodeToString(crypted), nil
}

// 解密
func AesDecrypt(decodeBytes []byte, key []byte, iv []byte) ([]byte, error) {
	//decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)//先解密base64
	//decodeBytes, err := hex.DecodeString(decodeStr)
	//if err != nil {
	//	return nil, err
	//}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	//origData = ZeroUnPadding(origData) // origData = PKCS5UnPadding(origData)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func Byte2base64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

func Base642byte(str string) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func Hex2Byte(text string) ([]byte, error) {
	bytes, err := hex.DecodeString(text)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func Bytes2Hex(bytes []byte) string {
	str := hex.EncodeToString(bytes)
	return str
}
