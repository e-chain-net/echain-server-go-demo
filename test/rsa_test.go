package test

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"github.com/e-chain-net/echain-server-go-demo/common"
	"testing"

	"encoding/base64"
	"fmt"
)

func TestRsaSign(t *testing.T) {
	message := []byte("Hello, world!")
	privKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Error(err)
	}

	hash := sha256.Sum256(message)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hash[:])
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Message: ", string(message))
	fmt.Println("Signature: ", base64.StdEncoding.EncodeToString(signature))

	pubKey := privKey.PublicKey
	err = rsa.VerifyPKCS1v15(&pubKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Signature verified")
}



func TestRsaLoadAndSign(t *testing.T) {
	private64 := "MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCY/i3wvgeqo1gSusHf1AuPYU+nRMNLkZeIPpA6YjvE5iQ26DQXdrQ0WamsvAKz3GEQMwRqmtesEQhLygJj4wJdO4PDhGNqlObsWQx4N1cyrbvGouTbO2hUhnesP3PURMRBEpJd1A/koKmh606i4oKcXKuVBwjoBMdPPCMZ4QchcJ30KFLtiXT9OBpGim7SgFBiKnWxO/4CGd1KkixRj4ID0lyYhRnnCUvFptN822V4g5yr7vSUsH3M7IxVO4SgFzFZjM9pvuo0G2vf51NJ+fK3y9rK24/Vt671sEf59s28OXuyRXErjgfIoDdM2triJ2Fq8jZQV9kVvR6gn7rmZfUnAgMBAAECggEAQ248J1BKJr5Jsi+YBaP62F4Gcm3POb5YsFcK0IC9YSIiMgUT+Id8E1q1ewl+k3F9YltqBeZrSk5TfrvxY78JKrhxcbom6zHnuaHh6hZSG2cRTRI8lhfP+vktQ8DPt237pcaetjYiLx1UxqXkicwVzv7VLSDlnwWEJvsVaXGR5/2BT8q+/2VEK4qCe8DESNpWNlDfonXAK0FDtDzWkjwLeWzJtzWQLw0ps8gSTQsUYRA2GUBtcp3MWOy+GOAIzhbTawOYTi3EjvAsRB7YuLYLOnueid0vYVRu6IHETcOBJIpGbBxV0IpbNvYNJ53A1bgyELvKIM9xUYs/3m5HIc6mmQKBgQDKvaYx1nUTjikkkn88IC+TgnMGSBSDSKcxZd0IOUPC1ohtnB0x/IcH+mEBot8GEkn7CjnyvtbaBq3I+RxGpfrLzdzt8BH9tLxyrGA872iXfB8owRMpoOs0hDMRTT2gZPsXpNdQwDP4UvLqmsQPw0QO5id7gLnc5Rm6OSMcIaXx/QKBgQDBLvl3we7ZzN+PydXBY9AKnvAl9BeDFPgynsRXn0dYNuKDWR3PXF/IOLGraa7LHZ3L6WJY3fLRr5CMV+k8RjWo6aZMHRqFzsQGRW3ta7XczTO1yq6/ks6xHje/yQqeGdbJLD07StCwslA7JDukA5u0WuPkaozjRKLrN9ShiHDq8wKBgQCOxpoo1NekSuQsjkKuTBhVMHPiw5Y2kk60GgFbzkArETwIvP1Oe4F4m9n+9f1L4EtbUGtYyQ6zgiqWsuA33KHPLw3cPsncupBPzZcEsrEcpVuoLrhZA6tAU61HDPdOYm71yq+bfY/b3EaX8yAJ3cCrIWhCsHez2V+R5rUUFZow3QKBgQC81Sr7OfE8qrt49OTh5awNRbEemEuHUS8PZAwuTj5R50xg8fJmqDfkIi7hjCtU1f1Rvi7pCQL6nm9gD+qnhUWcd8+bJPOxChyouKMsaZXaYCcEszs/fcRWc2AxMtYTFtTRzlGILKhzn8k3FkLKHtDLafDLbK+M06Gg5PEOeK1PqwKBgQCn0r+NDE09ImX9PVymwomScpPRWC/SxVgmzx3mGDG0AaRqGjfa1hskqxwx8eRfT3exwvQ9dvYaPyfATyQZ0uEkg7bJ47Jfr7YHkcKMWM1+u8iVxN4mn6Kj3aSfy71iumPzm8J/9BZHX1cTLzXDi1OiP+mQs+UXqwXGfvx/UZHSgw=="
	public64 := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmP4t8L4HqqNYErrB39QLj2FPp0TDS5GXiD6QOmI7xOYkNug0F3a0NFmprLwCs9xhEDMEaprXrBEIS8oCY+MCXTuDw4RjapTm7FkMeDdXMq27xqLk2ztoVIZ3rD9z1ETEQRKSXdQP5KCpoetOouKCnFyrlQcI6ATHTzwjGeEHIXCd9ChS7Yl0/TgaRopu0oBQYip1sTv+AhndSpIsUY+CA9JcmIUZ5wlLxabTfNtleIOcq+70lLB9zOyMVTuEoBcxWYzPab7qNBtr3+dTSfnyt8vaytuP1beu9bBH+fbNvDl7skVxK44HyKA3TNra4idhavI2UFfZFb0eoJ+65mX1JwIDAQAB"

	privateKey,err := common.LoadRsaPrivateKey(private64)
	if err != nil{
		t.Error(err)
	}

	publicKey,err := common.LoadRsaPublicKey(public64)
	if err != nil{
		t.Error(err)
	}

	message := []byte("1020000189025321" + "-" + "1681178471210")
	encoded,err := common.RsaSign(privateKey,message)
	if err != nil {
		t.Error(err)
	}
	// 使用公钥对签名进行验证
	err = common.RsaVerify(encoded,publicKey,message)
	if err != nil {
		t.Error(err)
	}


	fmt.Println("Signature:" + encoded)
	fmt.Println("Signature verified")
}
