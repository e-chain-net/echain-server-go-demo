package test

import (
	"fmt"
	"github.com/e-chain-net/echain-server-go-demo/common"
	"github.com/e-chain-net/echain-server-go-demo/net"
	"testing"
)

func TestDeployContract(t *testing.T) {
	owner := "0x2ebca12753f7e9526ef76f2698b7124e37e5ce87"
	reqNo := "123"
	addr,err := net.DeployContract(common.UrlDeploy,reqNo,owner)
	if err != nil{
		t.Error(err)
	}
	fmt.Println("合约地址：",addr)
}