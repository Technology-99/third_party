package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenAccessAndPwdAndRsa() (AccessKey, AccessSecret string, AccessPriKey, AccessPubKey, RefreshPriKey, RefreshPubKey []byte, err error) {
	AccessKey = GenerateTicket("", 16)

	//创建secret
	hash := sha1.New()
	hash.Write([]byte(AccessKey))
	sum := hash.Sum(nil)
	sign := hex.EncodeToString(sum)
	AccessSecret = ToUp(sign)

	//生成rsa
	AccessPriKey, AccessPubKey, err = GenRsaKey()
	if err != nil {
		return "", "", nil, nil, nil, nil, err
	}
	RefreshPriKey, RefreshPubKey, err = GenRsaKey()
	if err != nil {
		return "", "", nil, nil, nil, nil, err
	}

	return AccessKey, AccessSecret, AccessPriKey, AccessPubKey, RefreshPriKey, RefreshPubKey, nil
}
