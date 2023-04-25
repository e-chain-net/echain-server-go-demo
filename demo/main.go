package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/dinghongchao/bcos-go-sdk/sdk"
	"github.com/e-chain-net/echain-server-go-demo/common"
	"github.com/e-chain-net/echain-server-go-demo/net"
	"os"
	"strconv"
	"sync"
)

type TxParam struct{
	ContractAddress string
	CocurrentCount int
	TotalTxCount int
	PrivateKey string
	BlockNumber int64
	ToAddress string
}



func signAndSend(index int,txParam TxParam,wg *sync.WaitGroup){
	defer wg.Done()
	//每个线程发送交易数
	txCount := txParam.TotalTxCount/txParam.CocurrentCount
	//起始tokenId
	startId := txCount * index

	for i := startId; i<(index+1)*txCount; i++{
		tokenId := strconv.Itoa(startId + i)
		var blockLimit int64 = int64(txParam.BlockNumber + 900)
		input,err := common.EncodeMint(txParam.ToAddress,tokenId)
		if err != nil{
			fmt.Println("encode mint error:",err)
			continue
		}
		txHash,txSigned,err := sdk.CreateSignedTransaction(txParam.PrivateKey,"group0","chain0",txParam.ContractAddress,input,"",blockLimit)
		if err != nil{
			fmt.Println("CreateSignedTransaction failed:",err)
			continue
		}
		jsonRpc := common.JsonRpc{}
		jsonRpc.Method = "sendTransaction"
		jsonRpc.Params = []common.AnyType{txSigned,false}
		request := common.TxRequest{}
		request.ReqNo = txHash
		request.JsonRpc = jsonRpc

		payload,err := json.Marshal(request)
		if err != nil{
			fmt.Println("JsonMarshal request failed:",err)
			continue
		}
		//fmt.Println("Request payload:",string(payload))
		fmt.Println("TxHash:",txHash)
		resp,err := net.HttpPost(common.UrlTx,payload)
		if err != nil{
			fmt.Println("Http post failed:",err)
			continue
		}

		code,err := jsonparser.GetString(resp,"code")
		if err != nil{
			fmt.Println("jsonParser error:",err)
			continue
		}
		if code != "EC000001" {
			fmt.Println("sendTransaction error,code=",code)
		}
	}
}

func requestBlockNumber()(int64,error){
	payload := []byte(`{"jsonRpc": {"method":"getBlockNumber","params":[]}}`)
	resp,err := net.HttpPost(common.UrlQuery,payload)
	if err != nil{
		return 0,err
	}
	//fmt.Println(string(resp))
	code,err := jsonparser.GetString(resp,"code")
	if err != nil{
		return 0,err
	}
	if code != "EC000000"{
		return 0,errors.New(string(resp))
	}
	blockNumber,err := jsonparser.GetInt(resp,"data","blockNumber")
	if err != nil{
		return 0,err
	}
	return blockNumber,nil
}

//测试账户，其它测试配置见 commmon/define.go
//address:0x1d35a69766f36d958f8927aa137d1497d7321705
//private:cc531db4906f0ee7441909274d47f72cff547b7383c3a862a34bcd15650fc08d

//说明：需要提前部署一个合约，并把管理员权限转移给测试账户
//参数1：用户私钥（合约管理员）
//参数2：合约地址
//参数3：并发数
//参数4：交易总数
func main() {
	//for index, arg := range os.Args[0:] {
	//	fmt.Println(index, arg)
	//}
	if len(os.Args) < 5{
		fmt.Println("参数个数不能少于4")
		return
	}

	txParam := TxParam{}
	txParam.ToAddress = "0x7f7cd0133ba23aca1140dd41180e07b9873c5664"

	var err error = nil
	txParam.PrivateKey = os.Args[1]
	txParam.ContractAddress = os.Args[2]
	txParam.CocurrentCount,err = strconv.Atoi(os.Args[3])
	if err !=nil{
		fmt.Println("Param 3 parse int error")
		return
	}
	txParam.TotalTxCount,err = strconv.Atoi(os.Args[4])
	if err !=nil{
		fmt.Println("Param 4 parse int error")
		return
	}
	txParam.BlockNumber,err = requestBlockNumber()
	if err !=nil{
		fmt.Println("Request blockNumber error",err)
		return
	}

	timeStart := common.GetTimeMilli()
	var wg sync.WaitGroup
	for i := 0; i < txParam.CocurrentCount; i++ {
		wg.Add(1)
		go signAndSend(i,txParam,&wg)
	}
	wg.Wait()

	timeEnd := common.GetTimeMilli()
	timeUsed := timeEnd-timeStart
	tps := int64(txParam.TotalTxCount) * 1000/(timeEnd-timeStart)
	fmt.Printf("Time used:%dms,total tx:%d,send tps:%d\n",timeUsed,txParam.TotalTxCount,tps)
}


