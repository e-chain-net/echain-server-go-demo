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
		abi, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"name_","type":"string"},{"internalType":"string","name":"symbol_","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"approved","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"},{"conflictFields":[{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":2,"value":[1]},{"kind":3,"slot":4,"value":[1]}],"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"approve","outputs":[],"selector":[157198259,523061344],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":0}],"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"selector":[1889567281,3431720718],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0},{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":2,"value":[0]},{"kind":3,"slot":4,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"burn","outputs":[],"selector":[1117154408,958323734],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":3,"slot":2,"value":[0]},{"kind":3,"slot":4,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"getApproved","outputs":[{"internalType":"address","name":"","type":"address"}],"selector":[135795452,2629401170],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0}],"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"selector":[3917867461,1275072607],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0},{"kind":3,"slot":2,"value":[1]},{"kind":4,"value":[6]}],"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"mint","outputs":[],"selector":[1086394137,1331560647],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":4,"value":[0]}],"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"selector":[117300739,2971363459],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":4,"value":[6]}],"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"selector":[2376452955,1351213768],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":3,"slot":2,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"ownerOf","outputs":[{"internalType":"address","name":"","type":"address"}],"selector":[1666326814,1460761419],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":4,"value":[6]}],"inputs":[],"name":"renounceOwnership","outputs":[],"selector":[1901074598,3631098338],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":4,"value":[2]}],"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"safeTransferFrom","outputs":[],"selector":[1115958798,2965612340],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":4,"value":[2]}],"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"selector":[3096268766,1445559354],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":0},{"kind":2,"slot":5,"value":[0]}],"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"selector":[2720838757,805160440],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":5}],"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"selector":[33540519,3934173080],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":4,"value":[1]}],"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"selector":[2514000705,1543120790],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":3,"slot":2,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"tokenURI","outputs":[{"internalType":"string","name":"","type":"string"}],"selector":[3363526365,851218509],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0},{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":2,"value":[2]},{"kind":3,"slot":4,"value":[2]}],"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"transferFrom","outputs":[],"selector":[599290589,2911541041],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":4,"value":[6]}],"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"selector":[4076725131,382390570],"stateMutability":"nonpayable","type":"function"}]`))
		if err != nil {
			fmt.Println("Abi json read error:" ,err)
		}
		erc721Abi = abi
	})
}
func EncodeMint(to string,tokenId string) (string,error){
	checkAbi()
	toAddress := common.HexToAddress(to)
	tokenBig := new(big.Int)
	tokenBig.SetString(tokenId,0)

	data, err := erc721Abi.Pack("mint", toAddress,tokenBig)
	if err != nil {
		return "",err
	}
	return hexutil.Encode(data),nil
}

func EncodeTransferFrom(from string,to string,tokenId string)(string,error){
	checkAbi()
	fromAddress := common.HexToAddress(from)
	toAddress := common.HexToAddress(to)
	tokenBig := new(big.Int)
	tokenBig.SetString(tokenId,0)

	data, err := erc721Abi.Pack("transferFrom", fromAddress,toAddress,tokenBig)
	if err != nil {
		return "",err
	}
	return hexutil.Encode(data),nil
}

func EncodeBurn(tokenId string) (string,error){
	checkAbi()
	tokenBig := new(big.Int)
	tokenBig.SetString(tokenId,0)

	data, err := erc721Abi.Pack("burn", tokenBig)
	if err != nil {
		return "",err
	}
	return hexutil.Encode(data),nil
}

func GetTimeMilli() int64{
	return time.Now().UnixMilli()
}

