package util

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	b64 "encoding/base64"
)

func EncString(s []byte) string {
	res := b64.StdEncoding.EncodeToString([]byte(s))
	return res
}

func DencString(s string) []byte {
	res, _ := b64.StdEncoding.DecodeString(s)
	return res
}

func RsaEncString(s string, key *rsa.PublicKey) ([]byte, error) {
	var err error
	var enc []byte
	md5_hash := md5.New()

	if encrypted, er := rsa.EncryptOAEP(md5_hash, rand.Reader, key, []byte(s), []byte("t")); er != nil {
		err = er
	} else {
		enc = encrypted
	}
	return enc, err
}
