package qxCrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/chacha20poly1305"
	"io"
)

// **AES-GCM 加密**
func AESEncryptByGCM(plainText, key []byte) (string, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", nil, err
	}

	// 生成随机 Nonce（12 字节）
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", nil, err
	}

	// 加密 + 认证
	cipherText := aesGCM.Seal(nil, nonce, plainText, nil)

	return base64.StdEncoding.EncodeToString(cipherText), nonce, nil
}

// **AES-GCM 解密**
func AESDecryptByGCM(cipherTextBase64 string, key, nonce []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return "", err
	}

	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

// **AES-CCM 加密**
func AESEncryptByCCM(plainText, key []byte) (string, []byte, error) {
	aesCCM, err := chacha20poly1305.NewX(key)
	if err != nil {
		return "", nil, err
	}

	// 生成随机 Nonce（13 字节）
	nonce := make([]byte, 13)
	_, err = rand.Read(nonce)
	if err != nil {
		return "", nil, err
	}

	cipherText := aesCCM.Seal(nil, nonce, plainText, nil)

	return base64.StdEncoding.EncodeToString(cipherText), nonce, nil
}

// **AES-CCM 解密**
func AESDecryptByCCM(cipherTextBase64 string, key, nonce []byte) (string, error) {
	aesCCM, err := chacha20poly1305.NewX(key)
	if err != nil {
		return "", err
	}

	cipherText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return "", err
	}

	plainText, err := aesCCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

// **PKCS7 填充**
func CBCPkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// **PKCS7 去填充**
func CBCPkcs7Unpad(data []byte) []byte {
	length := len(data)
	unpad := int(data[length-1])
	return data[:(length - unpad)]
}

// **AES-CBC 加密**
func AESEncryptByCBC(plainText, key []byte) (string, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", nil, err
	}

	// 生成随机 IV（16 字节）
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", nil, err
	}

	// PKCS7 填充
	plainText = CBCPkcs7Pad(plainText, aes.BlockSize)

	cipherText := make([]byte, len(plainText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText), iv, nil
}

// **AES-CBC 解密**
func AESDecryptByCBC(cipherTextBase64 string, key, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Base64 解码密文
	cipherText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return "", err
	}

	// CBC 解密
	plainText := make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherText)

	// PKCS7 去填充
	plainText = CBCPkcs7Unpad(plainText)

	return string(plainText), nil
}

// **AES-CTR 加密**
func AESEncryptByCTR(plainText, key []byte) (string, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", nil, err
	}
	// 生成随机 Nonce（IV，16 字节）
	nonce := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", nil, err
	}

	stream := cipher.NewCTR(block, nonce)
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText), nonce, nil
}

// **AES-CTR 解密**
func AESDecryptByCTR(cipherTextBase64 string, key, nonce []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCTR(block, nonce)
	cipherText, _ := base64.StdEncoding.DecodeString(cipherTextBase64)
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}
