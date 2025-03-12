package qxCrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// RSA + SHA-256 进行加密
func RSAEncrypt(plainText []byte, publicKey *rsa.PublicKey) (string, error) {
	encryptedData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, plainText, nil)
	if err != nil {
		return "", fmt.Errorf("rsa encrypt failded: %v", err)
	}

	// Base64 编码
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// RSA + SHA-256 进行解密
func RSADecrypt(encryptedBase64 string, privateKey *rsa.PrivateKey) ([]byte, error) {
	// Base64 解码
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return nil, fmt.Errorf("Base64 decrypt: %v", err)
	}

	// 使用私钥解密
	decryptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedData, nil)
	if err != nil {
		return nil, fmt.Errorf("rsa decrypt failed: %v", err)
	}

	return decryptedData, nil
}

// note: 生成 RSA 公私钥并解析到字符串
func GenRsaKey2Str(btLen int) (privateKeyStr, publicKeyStr *string, err error) {
	if btLen <= 2048 {
		return nil, nil, errors.New("btLen must >= 2048")
	}
	privateKey, publicKey, err := RSAGenerateKeys(btLen)
	if err != nil {
		return nil, nil, err
	}
	privateKeyBt, err := RSAParsePrivateKey2Bt(privateKey)
	if err != nil {
		return nil, nil, err
	}
	publicKeyBt, err := RSAParsePublicKey2Bt(publicKey)
	if err != nil {
		return nil, nil, err
	}
	privateKeyStrTmp := string(privateKeyBt)
	publicKeyStrTmp := string(publicKeyBt)
	return &privateKeyStrTmp, &publicKeyStrTmp, nil
}

// note: 生成 RSA 公私钥并解析到bt
func GenRsaKey2Bt(btLen int) (privateKeyBt, publicKeyBt []byte, err error) {
	if btLen <= 2048 {
		return nil, nil, errors.New("btLen must >= 2048")
	}
	privateKey, publicKey, err := RSAGenerateKeys(btLen)
	if err != nil {
		return nil, nil, err
	}
	privateKeyBt, err = RSAParsePrivateKey2Bt(privateKey)
	if err != nil {
		return nil, nil, err
	}
	publicKeyBt, err = RSAParsePublicKey2Bt(publicKey)
	if err != nil {
		return nil, nil, err
	}
	return privateKeyBt, publicKeyBt, nil
}

// note: 生成 RSA 公私钥
func RSAGenerateKeys(bits int) (privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, err error) {
	privateKey, err = rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// note: 解析私钥到字节
func RSAParsePrivateKey2Bt(privateKey *rsa.PrivateKey) (prvKeyBt []byte, err error) {
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	return privateKeyPEM, nil
}

// note: 解析公钥到字节
func RSAParsePublicKey2Bt(publicKey *rsa.PublicKey) (pubKeyBt []byte, err error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	return publicKeyPEM, nil
}
