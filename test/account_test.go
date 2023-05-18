package test

import (
	"fmt"
	"github.com/e-chain-net/echain-go-sdk-721/sdk"
	"testing"
)

func TestNewAccount(t *testing.T) {
	account := sdk.NewAccount()
	fmt.Println("Address:",account.Address)
	fmt.Println("Private:",account.Private)
}
