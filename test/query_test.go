package test

import (
	"fmt"
	"github.com/e-chain-net/echain-server-go-demo/common"
	"github.com/e-chain-net/echain-server-go-demo/net"
	"testing"
)



//查询区块号
func TestGetBlockNumber(t *testing.T) {
	number,err := net.GetBlockNumber(common.UrlQuery)
	if err != nil{
		t.Error(err)
	}
	fmt.Println("BlockNumber:", number)
}

//查询交易收据
func TestTransactionReceipt(t *testing.T) {
	txHash := "0x3ac02bbaca5e7e0adc05d0e36954c86ee39108d543542a49eed7420d445d2536"
	txReceipt,err := net.GetTransactionReceipt(common.UrlQuery,txHash)
	if err != nil{
		t.Error(err)
	}
	fmt.Println("TransactionReceipt:",string(txReceipt))
}

//查询NFT的拥有者
func TestTokenOwner(t *testing.T){
	contractAddress := "0x18f8597118953b3374c2515ecf799ce4750361bb";
	owner,err := net.GetOwnerOf(common.UrlQuery,contractAddress,"1002")
	if err != nil{
		t.Error(err)
	}
	fmt.Println("OwnerOf output:",owner)
}

//查询账户余额
func TestBalanceOf(t *testing.T){
	contractAddress := "0x18f8597118953b3374c2515ecf799ce4750361bb";
	owner := "0x95a1a99be965777d8b0e42fe5ec1c161f6c3a5da"
	balance,err := net.GetBalanceOf(common.UrlQuery,contractAddress,owner)
	if err != nil{
		t.Error(err)
	}
	fmt.Println("balanceOf output:",balance)
}

