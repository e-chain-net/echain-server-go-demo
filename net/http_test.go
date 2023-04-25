package net

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/e-chain-net/echain-server-go-demo/common"
	"testing"
)


func TestGetBlockNumber(t *testing.T) {
	// Create request payload
	payload := []byte(`{"jsonRpc": {"method":"getBlockNumber","params":[]}}`)

	responseBody,err := HttpPost(common.UrlQuery,payload)
	if err != nil{
		t.Error(err)
	}
	code,err := jsonparser.GetString(responseBody,"code")
	if err != nil{
		t.Error(err)
	}
	if code != "EC000000" {
		message,_ := jsonparser.GetString(responseBody,"message")
		t.Errorf("sendTransaction error,msg=%s",message)
	}
	number,_ := jsonparser.GetInt(responseBody,"data","blockNumber")
	fmt.Println("BlockNumber:", number)
}

func TestTransactionReceipt(t *testing.T) {
	txHash := "0x3ac02bbaca5e7e0adc05d0e36954c86ee39108d543542a49eed7420d445d2536"
	jsonRpc := common.JsonRpc{}
	jsonRpc.Method = "getTransactionReceipt"
	jsonRpc.Params = []common.AnyType{txHash,false}
	request := common.QueryRequest{}
	request.JsonRpc = jsonRpc

	payload,err := json.Marshal(request)
	if err != nil{
		t.Error(err)
	}
	responseBody,err := HttpPost(common.UrlQuery,payload)
	if err != nil{
		t.Error(err)
	}
	code,err := jsonparser.GetString(responseBody,"code")
	if err != nil{
		t.Error(err)
	}
	if code != "EC000000" {
		message,_ := jsonparser.GetString(responseBody,"message")
		t.Errorf("sendTransaction error,msg=%s",message)
	}
	txReceipt,_,_,err := jsonparser.Get(responseBody,"data")
	if err != nil{
		t.Error(err)
	}
	fmt.Println("TransactionReceipt:",string(txReceipt))
}

func TestTokenOwner(t *testing.T){
	contractAddress := "0xc0f2254a5e506d6cda5e5ccd98ced32bd0e81609";
	ownerOfData,err := common.EncodeOwnerOf("1000")
	if err != nil {
		t.Error(err)
	}

	jsonRpc := common.JsonRpc{}
	jsonRpc.Method = "call"
	jsonRpc.Params = []common.AnyType{contractAddress,ownerOfData}
	request := common.QueryRequest{}
	request.JsonRpc = jsonRpc

	payload,err := json.Marshal(request)
	if err != nil{
		t.Error(err)
	}

	responseBody,err := HttpPost(common.UrlQuery,payload)
	if err != nil{
		t.Error(err)
	}
	code,err := jsonparser.GetString(responseBody,"code")
	if err != nil{
		t.Error(err)
	}
	if code != "EC000000" {
		message,_ := jsonparser.GetString(responseBody,"message")
		t.Errorf("sendTransaction error,msg=%s",message)
	}

	status,err := jsonparser.GetInt(responseBody,"data","jsonRpcResp","result","status")
	if err != nil{
		t.Error(err)
	}
	fmt.Println("OwnerOf status:",status)
	output,err := jsonparser.GetString(responseBody,"data","jsonRpcResp","result","output")
	if err != nil{
		t.Error(err)
	}
	fmt.Println("OwnerOf output:",output)
}


