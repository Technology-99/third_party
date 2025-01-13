package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenPassword(pwd string) (basePwd, key, salt string) {
	key = RandSalt(32)
	salt = RandSalt(16)
	_, basePwd = AesEncryptByCTR(pwd+salt, key)
	return basePwd, key, salt
}

func GenAccessAndSecret() (AccessKey, AccessSecret string, err error) {
	AccessKey = GenerateTicket("", 16)
	//创建secret
	hash := sha1.New()
	hash.Write([]byte(AccessKey))
	sum := hash.Sum(nil)
	sign := hex.EncodeToString(sum)
	AccessSecret = ToUp(sign)
	return AccessKey, AccessSecret, nil
}

func GenTenantRsa() (AccessPriKey, AccessPubKey, RefreshPriKey, RefreshPubKey, EncryptionPriKey, EncryptionPubKey []byte, err error) {
	//生成rsa
	AccessPriKey, AccessPubKey, err = GenRsaKey()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	RefreshPriKey, RefreshPubKey, err = GenRsaKey()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	EncryptionPriKey, EncryptionPubKey, err = GenRsaKey()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	return AccessPriKey, AccessPubKey, RefreshPriKey, RefreshPubKey, EncryptionPriKey, EncryptionPubKey, nil
}

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
