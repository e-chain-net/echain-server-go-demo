package test

import (
	"fmt"
	util "github.com/e-chain-net/echain-server-go-demo/common"
	"testing"
)

func TestABI(t *testing.T){
	mintData,err := util.EncodeMint("0x9847b8f7bf06fa6687f38475ab621c188689d11e","123")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Encoded mint:", mintData)

	transferData,err := util.EncodeTransferFrom("0x9847b8f7bf06fa6687f38475ab621c188689d11e","0x2242eaaedb3ecb4d02c43aef87dd25e4ef559c29","123")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Encoded transfer:", transferData)

	burnData,err := util.EncodeBurn("123")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Encoded burn:", burnData)
}