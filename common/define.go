package common


//测试测试
const UrlBase string = "https://qa-api.e-chain.net.cn"
//压测环境
//const UrlBase string = "https://perf-api.e-chain.net.cn"
const UrlQuery string = UrlBase + "/chain/rpc/query"
const UrlTx string = UrlBase + "/chain/rpc/tx"
const UrlDeploy string = UrlBase + "/chain/contract/deploy"

const CertPath string = "D:/yeepay/e-chain.net.cn_server.crt"

const RsaPrivate string = "MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCY/i3wvgeqo1gSusHf1AuPYU+nRMNLkZeIPpA6YjvE5iQ26DQXdrQ0WamsvAKz3GEQMwRqmtesEQhLygJj4wJdO4PDhGNqlObsWQx4N1cyrbvGouTbO2hUhnesP3PURMRBEpJd1A/koKmh606i4oKcXKuVBwjoBMdPPCMZ4QchcJ30KFLtiXT9OBpGim7SgFBiKnWxO/4CGd1KkixRj4ID0lyYhRnnCUvFptN822V4g5yr7vSUsH3M7IxVO4SgFzFZjM9pvuo0G2vf51NJ+fK3y9rK24/Vt671sEf59s28OXuyRXErjgfIoDdM2triJ2Fq8jZQV9kVvR6gn7rmZfUnAgMBAAECggEAQ248J1BKJr5Jsi+YBaP62F4Gcm3POb5YsFcK0IC9YSIiMgUT+Id8E1q1ewl+k3F9YltqBeZrSk5TfrvxY78JKrhxcbom6zHnuaHh6hZSG2cRTRI8lhfP+vktQ8DPt237pcaetjYiLx1UxqXkicwVzv7VLSDlnwWEJvsVaXGR5/2BT8q+/2VEK4qCe8DESNpWNlDfonXAK0FDtDzWkjwLeWzJtzWQLw0ps8gSTQsUYRA2GUBtcp3MWOy+GOAIzhbTawOYTi3EjvAsRB7YuLYLOnueid0vYVRu6IHETcOBJIpGbBxV0IpbNvYNJ53A1bgyELvKIM9xUYs/3m5HIc6mmQKBgQDKvaYx1nUTjikkkn88IC+TgnMGSBSDSKcxZd0IOUPC1ohtnB0x/IcH+mEBot8GEkn7CjnyvtbaBq3I+RxGpfrLzdzt8BH9tLxyrGA872iXfB8owRMpoOs0hDMRTT2gZPsXpNdQwDP4UvLqmsQPw0QO5id7gLnc5Rm6OSMcIaXx/QKBgQDBLvl3we7ZzN+PydXBY9AKnvAl9BeDFPgynsRXn0dYNuKDWR3PXF/IOLGraa7LHZ3L6WJY3fLRr5CMV+k8RjWo6aZMHRqFzsQGRW3ta7XczTO1yq6/ks6xHje/yQqeGdbJLD07StCwslA7JDukA5u0WuPkaozjRKLrN9ShiHDq8wKBgQCOxpoo1NekSuQsjkKuTBhVMHPiw5Y2kk60GgFbzkArETwIvP1Oe4F4m9n+9f1L4EtbUGtYyQ6zgiqWsuA33KHPLw3cPsncupBPzZcEsrEcpVuoLrhZA6tAU61HDPdOYm71yq+bfY/b3EaX8yAJ3cCrIWhCsHez2V+R5rUUFZow3QKBgQC81Sr7OfE8qrt49OTh5awNRbEemEuHUS8PZAwuTj5R50xg8fJmqDfkIi7hjCtU1f1Rvi7pCQL6nm9gD+qnhUWcd8+bJPOxChyouKMsaZXaYCcEszs/fcRWc2AxMtYTFtTRzlGILKhzn8k3FkLKHtDLafDLbK+M06Gg5PEOeK1PqwKBgQCn0r+NDE09ImX9PVymwomScpPRWC/SxVgmzx3mGDG0AaRqGjfa1hskqxwx8eRfT3exwvQ9dvYaPyfATyQZ0uEkg7bJ47Jfr7YHkcKMWM1+u8iVxN4mn6Kj3aSfy71iumPzm8J/9BZHX1cTLzXDi1OiP+mQs+UXqwXGfvx/UZHSgw=="
const MerchantNo string = "3203001000583086614"

const ERC721ABIStr string = `[{"inputs":[{"internalType":"string","name":"name_","type":"string"},{"internalType":"string","name":"symbol_","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"approved","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"},{"conflictFields":[{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":2,"value":[1]},{"kind":3,"slot":4,"value":[1]}],"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"approve","outputs":[],"selector":[157198259,523061344],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":0}],"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"selector":[1889567281,3431720718],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0},{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":2,"value":[0]},{"kind":3,"slot":4,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"burn","outputs":[],"selector":[1117154408,958323734],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":3,"slot":2,"value":[0]},{"kind":3,"slot":4,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"getApproved","outputs":[{"internalType":"address","name":"","type":"address"}],"selector":[135795452,2629401170],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0}],"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"selector":[3917867461,1275072607],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0},{"kind":3,"slot":2,"value":[1]},{"kind":4,"value":[6]}],"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"mint","outputs":[],"selector":[1086394137,1331560647],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":4,"value":[0]}],"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"selector":[117300739,2971363459],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":4,"value":[6]}],"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"selector":[2376452955,1351213768],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":3,"slot":2,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"ownerOf","outputs":[{"internalType":"address","name":"","type":"address"}],"selector":[1666326814,1460761419],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":4,"value":[6]}],"inputs":[],"name":"renounceOwnership","outputs":[],"selector":[1901074598,3631098338],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":4,"value":[2]}],"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"safeTransferFrom","outputs":[],"selector":[1115958798,2965612340],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":4,"value":[2]}],"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"selector":[3096268766,1445559354],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":0},{"kind":2,"slot":5,"value":[0]}],"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"selector":[2720838757,805160440],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":5}],"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"selector":[33540519,3934173080],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":4,"value":[1]}],"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"selector":[2514000705,1543120790],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":3,"slot":2,"value":[0]}],"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"tokenURI","outputs":[{"internalType":"string","name":"","type":"string"}],"selector":[3363526365,851218509],"stateMutability":"view","type":"function"},{"conflictFields":[{"kind":0},{"kind":2,"slot":5,"value":[0]},{"kind":3,"slot":2,"value":[2]},{"kind":3,"slot":4,"value":[2]}],"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"transferFrom","outputs":[],"selector":[599290589,2911541041],"stateMutability":"nonpayable","type":"function"},{"conflictFields":[{"kind":4,"value":[6]}],"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"selector":[4076725131,382390570],"stateMutability":"nonpayable","type":"function"}]`

type AnyType interface {
}
type JsonRpc struct{
	Method string 		`json:"method"`
	Params []AnyType 	`json:"params"`
}

type TxRequest struct{
	ReqNo 	string 	`json:"reqNo"`
	JsonRpc JsonRpc `json:"jsonRpc"`
}

type QueryRequest struct{
	JsonRpc JsonRpc `json:"jsonRpc"`
}

type DeployRequest struct{
	ReqNo 	string 		`json:"reqNo"`
	ContractType string `json:"contractType"`
	Owner 	string 		`json:"owner"`
}