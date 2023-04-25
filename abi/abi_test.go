package main

import (
	"fmt"
	"testing"
)

func TestAbi(t *testing.T) {
	method := "transferFrom(address,address,uint256)"
	formatter := new(Formatter)
	methodSig := formatter.ToMethodFormat(method) // "a9059cbb"
	address := "0x9847b8f7bf06fa6687f38475ab621c188689d11e"
	addressFromSig := formatter.ToAddressFormat(address)
	address = "0x2242eaaedb3ecb4d02c43aef87dd25e4ef559c29"
	addressToSig := formatter.ToAddressFormat(address)
	value := "123"
	intSig := formatter.ToIntegerFormat(value, 64) // "000000000000000000000000000000000000000000000de0b6b3a7640000"

	finalSig := methodSig + addressFromSig + addressToSig + intSig
	fmt.Println(finalSig)
}
