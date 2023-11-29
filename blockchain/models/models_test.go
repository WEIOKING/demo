package models

import (
	"encoding/json"
	"fmt"
	"github.com/nanmu42/etherscan-api"
	"testing"
)

var jsonStr = "{\n  \"baseFeePerGas\": \"0x5f5e100\",\n  \"difficulty\": \"0x1\",\n  \"extraData\": \"0x3a94902a5b58ed208efe68d3fc196acc779429a21c5b21cc02663adfa01ba91d\",\n  \"gasLimit\": \"0x4000000000000\",\n  \"gasUsed\": \"0x1a3ee3\",\n  \"hash\": \"0xd9fbb04369a7b4fe00a7c2a819a4f22861e4b337843884e8c2d1153946462967\",\n  \"l1BlockNumber\": \"0x11c2f23\",\n  \"logsBloom\": \"0x00000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000001000000000000800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000\",\n  \"miner\": \"0xa4b000000000000000000073657175656e636572\",\n  \"mixHash\": \"0x0000000000017eeb00000000011c2f23000000000000000a0000000000000000\",\n  \"nonce\": \"0x00000000001242d4\",\n  \"number\": \"0x91b802a\",\n  \"parentHash\": \"0xaa0b591b0e800c790696b3c26a8dc56112f491933aa034eaf6d3bfac8e10d2b3\",\n  \"receiptsRoot\": \"0xa179f493b0f3115979855eeb38e4988ebaec1088a6bbed3eab7fa2ddb101aa75\",\n  \"sendCount\": \"0x17eeb\",\n  \"sendRoot\": \"0x3a94902a5b58ed208efe68d3fc196acc779429a21c5b21cc02663adfa01ba91d\",\n  \"sha3Uncles\": \"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347\",\n  \"size\": \"0x511\",\n  \"stateRoot\": \"0xb18aa587b214cd2c14ed23442871c330b5cda949b5b4c768b2bbaa987901dada\",\n  \"timestamp\": \"0x655d667e\",\n  \"totalDifficulty\": \"0x7c8a2e2\",\n  \"transactions\": [\n    {\n      \"blockHash\": \"0xd9fbb04369a7b4fe00a7c2a819a4f22861e4b337843884e8c2d1153946462967\",\n      \"blockNumber\": \"0x91b802a\",\n      \"from\": \"0x00000000000000000000000000000000000a4b05\",\n      \"gas\": \"0x0\",\n      \"gasPrice\": \"0x0\",\n      \"hash\": \"0x826e3aa06b1c259a75ec175197c8c98d922bd86e4ce40775a62783f632d83a38\",\n      \"input\": \"0x6bf6a42d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000011c2f2300000000000000000000000000000000000000000000000000000000091b802a0000000000000000000000000000000000000000000000000000000000000001\",\n      \"nonce\": \"0x0\",\n      \"to\": \"0x00000000000000000000000000000000000a4b05\",\n      \"transactionIndex\": \"0x0\",\n      \"value\": \"0x0\",\n      \"type\": \"0x6a\",\n      \"chainId\": \"0xa4b1\",\n      \"v\": \"0x0\",\n      \"r\": \"0x0\",\n      \"s\": \"0x0\"\n    },\n    {\n      \"blockHash\": \"0xd9fbb04369a7b4fe00a7c2a819a4f22861e4b337843884e8c2d1153946462967\",\n      \"blockNumber\": \"0x91b802a\",\n      \"from\": \"0xd7173acbc07fdfa59563d16404ae2eb1e652a611\",\n      \"gas\": \"0x2625a0\",\n      \"gasPrice\": \"0xbebc200\",\n      \"maxFeePerGas\": \"0x11e1a300\",\n      \"maxPriorityFeePerGas\": \"0x5f5e100\",\n      \"hash\": \"0x7b2dd3aaef2d7123c5636705cba2988c1f2bbb32a91be2e94382d8ad92bd3a5f\",\n      \"input\": \"0x\",\n      \"nonce\": \"0x0\",\n      \"to\": \"0x0938c63109801ee4243a487ab84dffa2bba4589e\",\n      \"transactionIndex\": \"0x1\",\n      \"value\": \"0x49749122a2000\",\n      \"type\": \"0x2\",\n      \"accessList\": [\n        \n      ],\n      \"chainId\": \"0xa4b1\",\n      \"v\": \"0x0\",\n      \"r\": \"0x15b273b40426b8aec257daa89c6b0cec17951d49d43715ac39ec51d700ad3859\",\n      \"s\": \"0x7ca6a8a09760ea7172d1e542865674ff2963e690b0a5c0ff51618ba7846c23\"\n    },\n    {\n      \"blockHash\": \"0xd9fbb04369a7b4fe00a7c2a819a4f22861e4b337843884e8c2d1153946462967\",\n      \"blockNumber\": \"0x91b802a\",\n      \"from\": \"0x685ee51518320a6625fa2bb523b2eab385df68b3\",\n      \"gas\": \"0x2f6a30\",\n      \"gasPrice\": \"0x5f5e100\",\n      \"maxFeePerGas\": \"0x5f5e100\",\n      \"maxPriorityFeePerGas\": \"0x5f5e100\",\n      \"hash\": \"0xda71e776e74d5fc48c630a9124f500b4d2d887fca1f3d0ce6a98e8c12cef6bf6\",\n      \"input\": \"0x0ddedd8400000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e 0000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000002f6a3000000000000000000000000000000000000000000000000000687324ec658000000000000000000000000000000000000000000000000000 00000000000000010000000000000000000000004b59e262003331288b3dedce2752080701650f790000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000b25ef85ec8920000000000000000000000000000000000 00000000000000000000000000000013c1cf8026387a8c66681aab7327e4ed2c7f0ff8b9f6bf45667983d78d66c4020\",\n      \"nonce\": \"0x21c0f\",\n      \"to\": \"0xc0e02aa55d10e38855e13b64a8e1387a04681a00\",\n      \"transactionIndex\": \"0x2\",\n      \"value\": \"0x0\",\n      \"type\": \"0x2\",\n      \"accessList\": [\n        \n      ],\n      \"chainId \": \"0xa4b1\",\n      \"v\": \"0x1\",\n      \"r\": \"0x4b1bddbdc9997e25073ebc19d8f71c232a76ca85cb4f8e7f34c555ef43b3ffab\",\n      \"s\": \"0x29cee12311a1d7cf32e13f27a9fbb28469b4047bca4b75cf0db6471e36c1a1f4\"\n    }\n  ],\n  \"transactionsRoot\": \"0x91e0a27a8c68301acec8f1d4e8f244dd92e111dbd664a5bec2a4f7f4ed583f59\",\n  \"uncles\": [\n    \n  ]\n}"

func TestBlockDataDeal(t *testing.T) {
	//utils.HttpPostOfJson()
	var block *Block
	if err := json.Unmarshal([]byte(jsonStr), &block); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(block.Hash)

	client := etherscan.New(etherscan.Mainnet, "[your API key]")
	txs, err := client.NormalTxByAddress()
}