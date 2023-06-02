package net

import (
	"encoding/json"
	"errors"
	"github.com/buger/jsonparser"
	"github.com/e-chain-net/echain-server-go-demo/common"
	"math/big"
	"strconv"
)

func GetBlockNumber(url string) (int64,error){
	// Create request payload
	payload := []byte(`{"jsonRpc": {"method":"getBlockNumber","params":[]}}`)

	responseBody,err := HttpPost(url,payload)
	if err != nil{
		return -1,err
	}
	code,err := jsonparser.GetString(responseBody,"code")
	if err != nil{
		return -1,err
	}
	if code != "EC000000" {
		message,_ := jsonparser.GetString(responseBody,"message")
		return -1, errors.New(message)
	}
	number,_ := jsonparser.GetInt(responseBody,"data","blockNumber")
	return number,nil
}


func GetTransactionReceipt(url string,txHash string)(string,error){
	jsonRpc := common.JsonRpc{}
	jsonRpc.Method = "getTransactionReceipt"
	jsonRpc.Params = []common.AnyType{txHash,false}
	request := common.QueryRequest{}
	request.JsonRpc = jsonRpc

	payload,err := json.Marshal(request)
	if err != nil{
		return "",err
	}
	responseBody,err := HttpPost(url,payload)
	if err != nil{
		return "",err
	}
	code,err := jsonparser.GetString(responseBody,"code")
	if err != nil{
		return "",err
	}
	if code != "EC000000" {
		message,_ := jsonparser.GetString(responseBody,"message")
		return "",errors.New(message)
	}
	txReceipt,_,_,err := jsonparser.Get(responseBody,"data")
	if err != nil{
		return "",err
	}
	return string(txReceipt),nil
}

func GetOwnerOf(url string,contractAddress string,tokenId string) (string,error){
	ownerOfData,err := common.EncodeOwnerOf(tokenId)
	if err != nil {
		return "",err
	}

	jsonRpc := common.JsonRpc{}
	jsonRpc.Method = "call"
	jsonRpc.Params = []common.AnyType{contractAddress,ownerOfData}
	request := common.QueryRequest{}
	request.JsonRpc = jsonRpc

	payload,err := json.Marshal(request)
	if err != nil{
		return "",err
	}

	responseBody,err := HttpPost(url,payload)
	if err != nil{
		return "",err
	}
	code,err := jsonparser.GetString(responseBody,"code")
	if err != nil{
		return "",err
	}
	if code != "EC000000" {
		message,_ := jsonparser.GetString(responseBody,"message")
		return "",errors.New(message)
	}

	status,err := jsonparser.GetInt(responseBody,"data","result","status")
	if err != nil{
		return "",err
	}
	if status != 0{
		return "",errors.New("Get token owner failed,status=" + strconv.Itoa(int(status)))
	}
	output,err := jsonparser.GetString(responseBody,"data","result","output")
	if err != nil{
		return "",err
	}
	return "0x" + output[26:] , nil
}

func GetBalanceOf(url string,contractAddress string,owner string)(int64,error){
	balanceOfData,err := common.EncodeBalanceOf(owner)
	if err != nil {
		return -1,err
	}

	jsonRpc := common.JsonRpc{}
	jsonRpc.Method = "call"
	jsonRpc.Params = []common.AnyType{contractAddress,balanceOfData}
	request := common.QueryRequest{}
	request.JsonRpc = jsonRpc

	payload,err := json.Marshal(request)
	if err != nil{
		return -1,err
	}

	responseBody,err := HttpPost(url,payload)
	if err != nil{
		return -1,err
	}
	code,err := jsonparser.GetString(responseBody,"code")
	if err != nil{
		return -1,err
	}
	if code != "EC000000" {
		message,_ := jsonparser.GetString(responseBody,"message")
		return -1,errors.New(message)
	}

	status,err := jsonparser.GetInt(responseBody,"data","result","status")
	if err != nil{
		return -1,err
	}
	if status != 0{
		return -1,errors.New("Get token owner failed,status=" + strconv.Itoa(int(status)))
	}
	output,err := jsonparser.GetString(responseBody,"data","result","output")
	if err != nil{
		return -1,err
	}
	balanceBig := new(big.Int)
	balanceBig.SetString(output[2:],16)
	return balanceBig.Int64(),nil
}

func SendTx(url string,txHash string,txSigned string,callbackUrl string) error{
	jsonRpc := common.JsonRpc{}
	jsonRpc.Method = "sendTransaction"
	jsonRpc.Params = []common.AnyType{txSigned,false}
	request := common.TxRequest{}
	request.ReqNo = txHash
	request.JsonRpc = jsonRpc
	request.CallbackUrl = callbackUrl

	payload,err := json.Marshal(request)
	if err != nil{
		//fmt.Println("JsonMarshal request failed:",err)
		return err
	}
	//fmt.Println("Request payload:",string(payload))
	//fmt.Println("TxHash:",txHash)
	resp,err := HttpPost(url,payload)
	if err != nil{
		//fmt.Println("Http post failed:",err)
		return err
	}

	code,err := jsonparser.GetString(resp,"code")
	if err != nil{
		//fmt.Println("jsonParser error:",err)
		return err
	}
	if code != "EC000001" {
		return errors.New("sendTransaction error,code=" + code)
	}
	return nil
}

func DeployContract(url string,reqNo string,owner string) (string,error){
	req := common.DeployRequest{}
	req.ReqNo = reqNo
	req.ContractType = "ERC721"
	req.Owner = owner

	payload,err := json.Marshal(req)
	if err != nil{
		//fmt.Println("JsonMarshal request failed:",err)
		return "",err
	}
	//fmt.Println("Request payload:",string(payload))
	//fmt.Println("TxHash:",txHash)
	resp,err := HttpPost(url,payload)
	if err != nil{
		//fmt.Println("Http post failed:",err)
		return "",err
	}
	code,err := jsonparser.GetString(resp,"code")
	if err != nil{
		//fmt.Println("jsonParser error:",err)
		return "",err
	}
	if code != "EC000000" {
		return "",errors.New("DeployContract error,code=" + code)
	}
	contractAddr,err := jsonparser.GetString(resp,"data","contractAddress")
	if err != nil{
		//fmt.Println("jsonParser error:",err)
		return "",err
	}
	return contractAddr,nil
}