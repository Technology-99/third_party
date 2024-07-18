package wechat

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	wechatUtils "github.com/wechatpay-apiv3/wechatpay-go/utils"
)

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
