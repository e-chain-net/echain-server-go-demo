package test

import (
	"fmt"
	"github.com/e-chain-net/echain-go-sdk-721/sdk"
	"testing"
)

//全局变量
var owner = sdk.Account{
	Address: "0xdb222aaaefb28a8a2b1533c6b098b819a80665b6",
	Private: "500a9faf63d80563207702295b9bf9f5dc98956fd253d81b8ba062f69b5cf6b1",
}

var user1 = sdk.Account{
	Address: "0x95a1a99be965777d8b0e42fe5ec1c161f6c3a5da",
	Private: "20af5ca9552563576673abda1af0540ff6c72ea631b1de8b11296c94167a6b06",
}

var user2 = sdk.Account{
	Address: "0xf53baf7526a2c8aec2f185ed48e94316e29e9e95",
	Private: "42e548a753fe86d0937372b24ae472559966929fb3f8d0672376849d23f6a43d",
}

var contractAddress = "0x48394ca653ed914ec4747a1f13f30de8c9edc404"
var tokenId = "1234568"
var blockNumber = int64(56315)

func TestSignMint(t *testing.T){
	toAddress := user1.Address
	privateKey := owner.Private

	txHash,signed,err := sdk.SignMint(toAddress,tokenId,contractAddress,privateKey,blockNumber)
	if err != nil{
		t.Error(err)
	}
	fmt.Println("SignMint tx-hash:",txHash)
	fmt.Println("SignMint tx-signed:",signed)
}

func TestSignTransfer(t *testing.T){
	fromAddress := user1.Address
	toAddress := user2.Address
	privateKey := user1.Private

	txHash,signed,err := sdk.SignTransferFrom(fromAddress,toAddress,tokenId,contractAddress,privateKey,blockNumber)
	if err != nil{
		t.Error(err)
	}
	fmt.Println("SignTransferFrom tx-hash:",txHash)
	fmt.Println("SignTransferFrom tx-signed:",signed)
}

func TestSignBurn(t *testing.T){
	privateKey := user2.Private

	txHash,signed,err := sdk.SignBurn(tokenId,contractAddress,privateKey,blockNumber)
	if err != nil{
		t.Error(err)
	}
	fmt.Println("SignBurn tx-hash:",txHash)
	fmt.Println("SignBurn tx-signed:",signed)
}