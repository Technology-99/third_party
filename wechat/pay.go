package wechat

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	wechatUtils "github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
)

func WechatCallBackBodyDecode(APIv3Secret, ResourceNonce, Ciphertext, AssociatedData string) ([]byte, error) {
	// 创建一个 AES cipher.Block，需要确保 key 是正确的长度
	block, err := aes.NewCipher([]byte(APIv3Secret))
	if err != nil {
		log.Printf("NewCipher failed, error=%s", err.Error())
		return nil, err
	}

	// 创建一个 AES-GCM 实例
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("NewGCM failed, error=%s", err.Error())
		return nil, err
	}

	decodedCiphertext, err := base64.StdEncoding.DecodeString(Ciphertext)
	if err != nil {
		log.Printf("base64 decode failed, error=%s", err.Error())
		return nil, err
	}

	nonce := []byte(ResourceNonce)

	// 解密数据
	decrypted, err := aesGCM.Open(nil, nonce, decodedCiphertext, []byte(AssociatedData))
	if err != nil {
		log.Printf("Open failed, error=%s", err.Error())
		return nil, err
	}
	return decrypted, nil
}

func WechatVerifySignature(pubKeyData string, signatureFile string, verifyData string) error {
	certificate, err := wechatUtils.LoadCertificate(pubKeyData)
	if err != nil {
		return err
	}
	rsaPubKey, ok := certificate.PublicKey.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("failed to parse RSA public key")
	}
	// 解码签名
	signature, err := base64.StdEncoding.DecodeString(signatureFile)
	if err != nil {
		return fmt.Errorf("failed to decode base64 signature: %w", err)
	}
	// 计算数据的SHA-256散列值
	hashed := sha256.Sum256([]byte(verifyData))
	// 验证签名
	err = rsa.VerifyPKCS1v15(rsaPubKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return fmt.Errorf("signature verification failed: %w", err)
	}
	fmt.Println("Signature verification succeeded.")
	return nil
}
