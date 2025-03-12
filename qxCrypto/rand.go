package qxCrypto

import "math/rand"

func GenRandByte(btLen int) ([]byte, error) {
	tmp := make([]byte, btLen) // 32 字节
	_, err := rand.Read(tmp)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}
