package test

import (
	"fmt"
	util "github.com/e-chain-net/echain-server-go-demo/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
	"testing"
)

func TestABI(t *testing.T) {
	// 定义 Solidity 合约中的函数参数类型
	inputABI, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"name_","type":"string"},{"internalType":"string","name":"symbol_","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"approved","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"},{"conflictFields":[{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":2,"value":[1]},{"kind":3,"slot":4,"value":[1]}],"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"approve","outputs":[],"selector":[157198259,523061344],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":0}],"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"selector":[1889567281,3431720718],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0},{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":2,"value":[0]},{"kind":3,"slot":4,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"burn","outputs":[],"selector":[1117154408,958323734],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":3,"slot":2,"value":[0]},{"kind":3,"slot":4,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"getApproved","outputs":[{"internalType":"address","name":"","type":"address"}],"selector":[135795452,2629401170],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0}],"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"selector":[3917867461,1275072607],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0},{"kind":3,"slot":2,"value":[1]},{"kind":4,"value":[6]}],"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"mint","outputs":[],"selector":[1086394137,1331560647],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":4,"value":[0]}],"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"selector":[117300739,2971363459],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":4,"value":[6]}],"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"selector":[2376452955,1351213768],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":3,"slot":2,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"ownerOf","outputs":[{"internalType":"address","name":"","type":"address"}],"selector":[1666326814,1460761419],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":4,"value":[6]}],"inputs":[],"name":"renounceOwnership","outputs":[],"selector":[1901074598,3631098338],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":4,"value":[2]}],"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"safeTransferFrom","outputs":[],"selector":[1115958798,2965612340],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":4,"value":[2]}],"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"selector":[3096268766,1445559354],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":0},{"kind":2,"slot":5,"value":[0]}],"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"selector":[2720838757,805160440],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":5}],"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"selector":[33540519,3934173080],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":4,"value":[1]}],"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"selector":[2514000705,1543120790],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":3,"slot":2,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"tokenURI","outputs":[{"internalType":"string","name":"","type":"string"}],"selector":[3363526365,851218509],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0},{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":2,"value":[2]},{"kind":3,"slot":4,"value":[2]}],"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"transferFrom","outputs":[],"selector":[599290589,2911541041],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":4,"value":[6]}],"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"selector":[4076725131,382390570],"stateMutability":"nonpayable","type":"function"}]`))
	if err != nil {
		t.Error(err)
	}
	// 定义函数输入参数
	to := common.HexToAddress("0x9847b8f7bf06fa6687f38475ab621c188689d11e")
	tokenId := big.NewInt(123)
	data, err := inputABI.Pack("mint", to,tokenId)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Encoded mint:", hexutil.Encode(data))

	from := common.HexToAddress("0x2242eaaedb3ecb4d02c43aef87dd25e4ef559c29")
	data, err = inputABI.Pack("transferFrom", from, to,tokenId)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Encoded transfer:", hexutil.Encode(data))
}

func TestUtil(t *testing.T){
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