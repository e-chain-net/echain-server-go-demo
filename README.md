# echain-server-go-demo

# 示例说明
测试用例包名：``test``

|  示例名称   | 描述  |
|  ----  | ----  |
| account_test  | 生成随机账户地址、私钥 |
| sign_test  | 对标准`Erc721`合约进行铸造、转移、销毁进行签名，得到交易哈希、签名后的交易体 |
| query_test  | 查询区块号、交易收据、NFT拥有者地址、NFT余额的请求示例 |
| sendtx_test  | 发送交易的请求示例 |
| contractdeploy_test  | 部署标准``ERC721``合约的请求示例 |

**注意：** 发送请求前需要先设置服务端证书路径（证书需要向亿链开发人员要一下），当前示例配置在：`common/define.go` 文件 ``CertPath`` 变量

