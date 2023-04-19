package common

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"strings"
)

func chunkSplit(s string, chunkSize int, endLine string) string {
	var chunks []string

	for len(s) > 0 {
		if len(s) < chunkSize {
			chunkSize = len(s)
		}
		chunks = append(chunks, s[:chunkSize])
		s = s[chunkSize:]
	}

	return strings.Join(chunks, endLine)
}

func LoadRsaPrivateKey(private64 string) (*rsa.PrivateKey,error){
	pemSplit := chunkSplit(private64,64,"\n")
	privatePem := "-----BEGIN PRIVATE KEY-----\n" + pemSplit + "\n-----END PRIVATE KEY-----\n"
	privateKeyBlock, _ := pem.Decode([]byte(privatePem))
	privateKeyParsed, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil,err
	}
	return privateKeyParsed.(*rsa.PrivateKey),nil
}

func LoadRsaPublicKey(public64 string) (*rsa.PublicKey,error){
	pemSplit := chunkSplit(public64,64,"\n")
	publicPem := "-----BEGIN PUBLIC KEY-----\n" + pemSplit + "\n-----END PUBLIC KEY-----\n"
	publicKeyBlock, _ := pem.Decode([]byte(publicPem))
	publicKeyParsed, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return nil,err
	}
	return publicKeyParsed.(*rsa.PublicKey),nil
}

func RsaSign(privateKey *rsa.PrivateKey,data []byte) (string,error){
	// 使用私钥对消息进行签名
	hash := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "",err
	}
	return base64.StdEncoding.EncodeToString([]byte(signature)),nil
}

func RsaVerify(signature string,publicKey *rsa.PublicKey,data []byte)(error){
	hash := sha256.Sum256(data)
	decoded,err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	// 使用公钥对签名进行验证
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], decoded)
	if err != nil {
		return err
	}
	return nil
}