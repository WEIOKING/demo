package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type EthRpcParam struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      string          `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	pending := big.NewInt(-1)
	if number.Cmp(pending) == 0 {
		return "pending"
	}
	return hexutil.EncodeBig(number)
}

func BlockByNumber(rpcUrl string, number *big.Int) (json.RawMessage, error) {
	data, err := RequestEth(rpcUrl, "eth_getBlockByNumber", toBlockNumArg(number), true)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	iTransactions := m["transactions"]
	transJson, err := json.Marshal(iTransactions)
	if err != nil {
		return nil, err
	}
	var iList []map[string]interface{}
	err = json.Unmarshal(transJson, &iList)
	if err != nil {
		return nil, err
	}
	var tempList []map[string]interface{}
	for _, tran := range iList {
		s := tran["type"]
		tranType := hexutil.MustDecodeUint64(s.(string))
		switch tranType {
		case types.LegacyTxType:
			tempList = append(tempList, tran)
		case types.AccessListTxType:
			tempList = append(tempList, tran)
		case types.DynamicFeeTxType:
			tempList = append(tempList, tran)
		default:
			fmt.Println("filter:", tran["hash"])
			continue
		}
	}
	m["transactions"] = tempList
	marshal, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func RequestEth(rpcUrl, method string, paramsIn ...interface{}) (json.RawMessage, error) {
	msg := &EthRpcParam{Version: "2.0", ID: "1", Method: method}
	if paramsIn != nil {
		var err error
		if msg.Params, err = json.Marshal(paramsIn); err != nil {
			return nil, err
		}
	}
	header := make(map[string]string)
	header["Accept"] = "application/json"
	header["Content-Type"] = "application/json"
	httpResponse, err := HttpPostOfJson(rpcUrl, msg, header)
	if err != nil {
		return nil, err
	}
	var resp = &EthRpcParam{}
	err = json.Unmarshal(httpResponse.Body(), resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

type rpcBlock struct {
	Hash         common.Hash      `json:"hash"`
	Transactions []rpcTransaction `json:"transactions"`
	UncleHashes  []common.Hash    `json:"uncles"`
}

type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}
