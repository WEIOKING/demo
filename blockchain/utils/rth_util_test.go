package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"testing"
)

var arbitrumRpcUrl = "https://endpoints.omniatech.io/v1/arbitrum/goerli/public"

func TestBlockByNumber(t *testing.T) {
	blockNumber := big.NewInt(57108200)
	data, err := BlockByNumber(arbitrumRpcUrl, blockNumber)
	if err != nil {
		fmt.Println(err.Error())
	}
	var head *types.Header
	var body rpcBlock
	if err := json.Unmarshal(data, &head); err != nil {
		println(err.Error())
		return
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		println(err.Error())
	}
	if err := json.Unmarshal(data, &body); err != nil {
		println(err.Error())
		return
	}
	for _, transaction := range body.Transactions {
		tx := transaction.tx
		println(tx.Hash().Hex())
		println(tx.To().Hex())
		println(tx.GasPrice().String())
		println(tx.GasTipCap().String())
		println(tx.GasFeeCap().String())
		println(tx.Gas())
		println(string(tx.Data()))
		println(tx.Value().String())
		println()
	}
	return
}
