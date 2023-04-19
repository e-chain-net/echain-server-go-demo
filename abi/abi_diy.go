package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

//想自己定义一套AbiEncode的，时间原因没跑通，依然用go-ethereum

// BytesToAddress returns Address with value b.
// If b is larger than len(h), b will be cropped from the left.
func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}
// has0xPrefix validates str begins with '0x' or '0X'.
func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

// Hex2Bytes returns the bytes represented by the hexadecimal string str.
func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}
// FromHex returns the bytes represented by the hexadecimal string s.
// s may be prefixed with "0x".
func FromHex(s string) []byte {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return Hex2Bytes(s)
}

// HexToAddress returns Address with byte values of s.
// If s is larger than len(h), s will be cropped from the left.
func HexToAddress(s string) Address { return BytesToAddress(FromHex(s)) }


type input struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

type output struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

type abi struct {
	Inputs          []input  `json:"inputs"`
	Name            string   `json:"name"`
	Outputs         []output `json:"outputs"`
	StateMutability string   `json:"stateMutability"`
	Type            string   `json:"type"`
}

func encodeABI(contractABI string, functionName string, args ...interface{}) (string, error) {
	var contract []abi
	err := json.Unmarshal([]byte(contractABI), &contract)
	if err != nil {
		return "", err
	}
	for _, f := range contract {
		if f.Name == functionName {
			var b strings.Builder
			b.WriteString(f.Name)
			b.WriteRune('(')
			for i, input := range f.Inputs {
				if i > 0 {
					b.WriteRune(',')
				}
				b.WriteString(input.Type)
			}
			b.WriteRune(')')
			signature := b.String()

			encoded, err := encode(f, args...)
			if err != nil {
				return "", err
			}
			return signature + hex.EncodeToString([]byte(encoded)), nil
		}
	}
	return "", fmt.Errorf("Function name not found in ABI")
}

func encode(f abi, args ...interface{}) ([]byte, error) {
	if len(args) != len(f.Inputs) {
		return nil, fmt.Errorf("Wrong number of arguments")
	}
	encoded := []byte{}
	for i := range args {
		arg, err := encodeValue(args[i], f.Inputs[i].Type)
		if err != nil {
			return nil, err
		}
		encoded = append(encoded, arg...)
	}
	return encoded, nil
}

func encodeValue(arg interface{}, argType string) ([]byte, error) {
	switch argType {
	case "uint256":
		i, ok := arg.(uint64)
		if !ok {
			return nil, fmt.Errorf("Argument is not uint64")
		}
		buf := make([]byte, 32)
		for pos, val := range fmt.Sprintf("%064b", i) {
			bytePos := 31 - (pos / 8)
			bitPos := pos % 8
			if val == '1' {
				buf[bytePos] |= (1 << bitPos)
			}
		}
		return buf, nil

	case "address":
		addr, ok := arg.(string)
		if !ok {
			return nil, fmt.Errorf("Argument is not an address string")
		}
		parsedAddr := HexToAddress(addr)
		return parsedAddr.Bytes(), nil

	case "bytes":
		b, ok := arg.([]byte)
		if !ok {
			return nil, fmt.Errorf("Argument is not a byte array")
		}
		buf := make([]byte, 32)
		copy(buf, b)
		return buf, nil

	// Handle other argument types here...

	default:
		return nil, fmt.Errorf("Unsupported argument type: %s", argType)
	}
}

func main() {
	// Define the contract ABI
	myABI := `[{"inputs":[{"internalType":"string","name":"name_","type":"string"},{"internalType":"string","name":"symbol_","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"approved","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"},{"conflictFields":[{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":2,"value":[1]},{"kind":3,"slot":4,"value":[1]}],"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"approve","outputs":[],"selector":[157198259,523061344],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":0}],"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"selector":[1889567281,3431720718],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0},{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":2,"value":[0]},{"kind":3,"slot":4,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"burn","outputs":[],"selector":[1117154408,958323734],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":3,"slot":2,"value":[0]},{"kind":3,"slot":4,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"getApproved","outputs":[{"internalType":"address","name":"","type":"address"}],"selector":[135795452,2629401170],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0}],"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"selector":[3917867461,1275072607],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0},{"kind":3,"slot":2,"value":[1]},{"kind":4,"value":[6]}],"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"mint","outputs":[],"selector":[1086394137,1331560647],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":4,"value":[0]}],"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"selector":[117300739,2971363459],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":4,"value":[6]}],"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"selector":[2376452955,1351213768],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":3,"slot":2,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"ownerOf","outputs":[{"internalType":"address","name":"","type":"address"}],"selector":[1666326814,1460761419],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":4,"value":[6]}],"inputs":[],"name":"renounceOwnership","outputs":[],"selector":[1901074598,3631098338],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":4,"value":[2]}],"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"safeTransferFrom","outputs":[],"selector":[1115958798,2965612340],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":4,"value":[2]}],"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"selector":[3096268766,1445559354],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":0},{"kind":2,"slot":5,"value":[0]}],"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"selector":[2720838757,805160440],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":5}],"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"selector":[33540519,3934173080],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":4,"value":[1]}],"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"selector":[2514000705,1543120790],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":3,"slot":2,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"tokenURI","outputs":[{"internalType":"string","name":"","type":"string"}],"selector":[3363526365,851218509],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0},{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":2,"value":[2]},{"kind":3,"slot":4,"value":[2]}],"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"transferFrom","outputs":[],"selector":[599290589,2911541041],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":4,"value":[6]}],"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"selector":[4076725131,382390570],"stateMutability":"nonpayable","type":"function"}]`

	// Define function name and arguments
	functionName := "mint"
	arg1 := "0xae1bd933fe320d6f926513d159f39b3861497366"
	arg2 := uint64(123)
	//arg2 := []byte{0x01, 0x02, 0x03}
	// Encode the parameters for the setValue function
	encoded, err := encodeABI(myABI, functionName, arg1, arg2)
	if err != nil {
		fmt.Println("Error encoding parameters:", err)
		return
	}

	fmt.Printf("Encoded ABI: %s	", encoded)
}