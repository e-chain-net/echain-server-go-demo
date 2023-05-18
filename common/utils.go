package common

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
	"sync"
	"time"
)


var (
	once   		sync.Once
	erc721Abi 	abi.ABI
)

func checkAbi(){
	once.Do(func(){
		abi, err := abi.JSON(strings.NewReader(ERC721ABIStr))
		if err != nil {
			fmt.Println("Abi json read error:" ,err)
		}
		erc721Abi = abi
	})
}

func EncodeOwnerOf(tokenId string)(string,error){
	checkAbi()
	tokenBig := new(big.Int)
	tokenBig.SetString(tokenId,0)

	data, err := erc721Abi.Pack("ownerOf", tokenBig)
	if err != nil {
		return "",err
	}
	return hexutil.Encode(data),nil
}

func EncodeBalanceOf(owner string)(string,error){
	checkAbi()
	ownerAddress := common.HexToAddress(owner)
	data,err := erc721Abi.Pack("balanceOf",ownerAddress)
	if err != nil {
		return "",err
	}
	return hexutil.Encode(data),nil
}

func GetTimeMilli() int64{
	return time.Now().UnixMilli()
}

