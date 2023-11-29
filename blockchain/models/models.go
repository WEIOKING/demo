package models

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Block struct {
	ParentHash      common.Hash      `json:"parentHash"`
	UncleHash       common.Hash      `json:"sha3Uncles"`
	Coinbase        common.Address   `json:"miner" `
	Root            common.Hash      `json:"stateRoot"`
	TxHash          common.Hash      `json:"transactionsRoot"`
	ReceiptHash     common.Hash      `json:"receiptsRoot"`
	Bloom           types.Bloom      `json:"logsBloom"`
	Difficulty      *big.Int         `json:"difficulty"`
	Number          *big.Int         `json:"number"`
	GasLimit        uint64           `json:"gasLimit"`
	GasUsed         uint64           `json:"gasUsed"`
	Time            uint64           `json:"timestamp"`
	Extra           []byte           `json:"extraData"`
	MixDigest       common.Hash      `json:"mixHash"`
	Nonce           types.BlockNonce `json:"nonce"`
	BaseFee         *big.Int         `json:"baseFeePerGas"`
	Hash            common.Hash      `json:"hash"`
	L1BlockNumber   *big.Int         `json:"l1BlockNumber"`
	SendCount       uint64           `json:"sendCount"`
	SendRoot        common.Hash      `json:"sendRoot"`
	Size            uint64           `json:"size"`
	TotalDifficulty *big.Int         `json:"totalDifficulty"`
	Transactions    []Transaction    `json:"transactions"`
}

// MarshalJSON marshals as JSON.
func (h Block) MarshalJSON() ([]byte, error) {
	type Header struct {
		ParentHash      common.Hash      `json:"parentHash"`
		UncleHash       common.Hash      `json:"sha3Uncles"`
		Coinbase        common.Address   `json:"miner" `
		Root            common.Hash      `json:"stateRoot"`
		TxHash          common.Hash      `json:"transactionsRoot"`
		ReceiptHash     common.Hash      `json:"receiptsRoot"`
		Bloom           types.Bloom      `json:"logsBloom"`
		Difficulty      *hexutil.Big     `json:"difficulty"`
		Number          *hexutil.Big     `json:"number"`
		GasLimit        hexutil.Uint64   `json:"gasLimit"`
		GasUsed         hexutil.Uint64   `json:"gasUsed"`
		Time            hexutil.Uint64   `json:"timestamp"`
		Extra           hexutil.Bytes    `json:"extraData"`
		MixDigest       common.Hash      `json:"mixHash"`
		Nonce           types.BlockNonce `json:"nonce"`
		BaseFee         *hexutil.Big     `json:"baseFeePerGas"`
		Hash            common.Hash      `json:"hash"`
		L1BlockNumber   *hexutil.Big     `json:"l1BlockNumber"`
		SendCount       hexutil.Uint64   `json:"sendCount"`
		SendRoot        common.Hash      `json:"sendRoot"`
		Size            hexutil.Uint64   `json:"size"`
		TotalDifficulty *hexutil.Big     `json:"totalDifficulty"`
		Transactions    []Transaction    `json:"transactions"`
	}
	var enc Header
	enc.ParentHash = h.ParentHash
	enc.UncleHash = h.UncleHash
	enc.Coinbase = h.Coinbase
	enc.Root = h.Root
	enc.TxHash = h.TxHash
	enc.ReceiptHash = h.ReceiptHash
	enc.Bloom = h.Bloom
	enc.Difficulty = (*hexutil.Big)(h.Difficulty)
	enc.Number = (*hexutil.Big)(h.Number)
	enc.GasLimit = hexutil.Uint64(h.GasLimit)
	enc.GasUsed = hexutil.Uint64(h.GasUsed)
	enc.Time = hexutil.Uint64(h.Time)
	enc.Extra = h.Extra
	enc.MixDigest = h.MixDigest
	enc.Nonce = h.Nonce
	enc.BaseFee = (*hexutil.Big)(h.BaseFee)
	enc.L1BlockNumber = (*hexutil.Big)(h.L1BlockNumber)
	enc.SendCount = hexutil.Uint64(h.SendCount)
	enc.SendRoot = h.SendRoot
	enc.Size = hexutil.Uint64(h.Size)
	enc.TotalDifficulty = (*hexutil.Big)(h.TotalDifficulty)
	enc.Transactions = h.Transactions
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (h *Block) UnmarshalJSON(input []byte) error {
	type Header struct {
		ParentHash      *common.Hash      `json:"parentHash"`
		UncleHash       *common.Hash      `json:"sha3Uncles"`
		Coinbase        *common.Address   `json:"miner"`
		Root            *common.Hash      `json:"stateRoot"`
		TxHash          *common.Hash      `json:"transactionsRoot"`
		ReceiptHash     *common.Hash      `json:"receiptsRoot"`
		Bloom           *types.Bloom      `json:"logsBloom"`
		Difficulty      *hexutil.Big      `json:"difficulty"`
		Number          *hexutil.Big      `json:"number" `
		GasLimit        *hexutil.Uint64   `json:"gasLimit"`
		GasUsed         *hexutil.Uint64   `json:"gasUsed"`
		Time            *hexutil.Uint64   `json:"timestamp"`
		Extra           *hexutil.Bytes    `json:"extraData"`
		MixDigest       *common.Hash      `json:"mixHash"`
		Nonce           *types.BlockNonce `json:"nonce"`
		BaseFee         *hexutil.Big      `json:"baseFeePerGas"`
		L1BlockNumber   *hexutil.Big      `json:"l1BlockNumber"`
		SendCount       *hexutil.Uint64   `json:"sendCount"`
		SendRoot        *common.Hash      `json:"sendRoot"`
		Size            *hexutil.Uint64   `json:"size"`
		TotalDifficulty *hexutil.Big      `json:"totalDifficulty"`
		Transactions    []Transaction     `json:"transactions"`
	}
	var dec Header
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for Header")
	}
	h.ParentHash = *dec.ParentHash
	if dec.UncleHash == nil {
		return errors.New("missing required field 'sha3Uncles' for Header")
	}
	h.UncleHash = *dec.UncleHash
	if dec.Coinbase == nil {
		return errors.New("missing required field 'miner' for Header")
	}
	h.Coinbase = *dec.Coinbase
	if dec.Root == nil {
		return errors.New("missing required field 'stateRoot' for Header")
	}
	h.Root = *dec.Root
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionsRoot' for Header")
	}
	h.TxHash = *dec.TxHash
	if dec.ReceiptHash == nil {
		return errors.New("missing required field 'receiptsRoot' for Header")
	}
	h.ReceiptHash = *dec.ReceiptHash
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Header")
	}
	h.Bloom = *dec.Bloom
	if dec.Difficulty == nil {
		return errors.New("missing required field 'difficulty' for Header")
	}
	h.Difficulty = (*big.Int)(dec.Difficulty)
	if dec.Number == nil {
		return errors.New("missing required field 'number' for Header")
	}
	h.Number = (*big.Int)(dec.Number)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Header")
	}
	h.GasLimit = uint64(*dec.GasLimit)
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Header")
	}
	h.GasUsed = uint64(*dec.GasUsed)
	if dec.Time == nil {
		return errors.New("missing required field 'timestamp' for Header")
	}
	h.Time = uint64(*dec.Time)
	if dec.Extra == nil {
		return errors.New("missing required field 'extraData' for Header")
	}
	h.Extra = *dec.Extra
	if dec.MixDigest != nil {
		h.MixDigest = *dec.MixDigest
	}
	if dec.Nonce != nil {
		h.Nonce = *dec.Nonce
	}
	if dec.BaseFee != nil {
		h.BaseFee = (*big.Int)(dec.BaseFee)
	}
	if dec.L1BlockNumber != nil {
		h.L1BlockNumber = (*big.Int)(dec.L1BlockNumber)
	}
	if dec.SendCount != nil {
		h.SendCount = uint64(*dec.SendCount)
	}
	if dec.SendRoot != nil {
		h.SendRoot = *dec.SendRoot
	}
	if dec.Size != nil {
		h.Size = uint64(*dec.Size)
	}
	if dec.TotalDifficulty != nil {
		h.TotalDifficulty = (*big.Int)(dec.TotalDifficulty)
	}
	if dec.Transactions != nil && len(dec.Transactions) > 0 {
		h.Transactions = dec.Transactions
	}
	return nil
}

type Transaction struct {
	BlockHash        common.Hash      `json:"blockHash"`
	BlockNumber      *hexutil.Big     `json:"blockNumber"`
	From             common.Address   `json:"from"`
	Gas              uint64           `json:"gas"`
	GasPrice         uint64           `json:"gasPrice"`
	Hash             common.Hash      `json:"hash"`
	Input            string           `json:"input"`
	Nonce            types.BlockNonce `json:"nonce"`
	To               common.Address   `json:"to"`
	TransactionIndex string           `json:"transactionIndex"`
	Value            *hexutil.Big     `json:"value"`
	Type             byte             `json:"type"`
	ChainId          *big.Int         `json:"chainId"`
	V                string           `json:"v"`
	R                string           `json:"r"`
	S                string           `json:"s"`
}

// MarshalJSON marshals as JSON.
func (h Transaction) MarshalJSON() ([]byte, error) {
	type Header struct {
		ParentHash      common.Hash      `json:"parentHash"`
		UncleHash       common.Hash      `json:"sha3Uncles"`
		Coinbase        common.Address   `json:"miner" `
		Root            common.Hash      `json:"stateRoot"`
		TxHash          common.Hash      `json:"transactionsRoot"`
		ReceiptHash     common.Hash      `json:"receiptsRoot"`
		Bloom           types.Bloom      `json:"logsBloom"`
		Difficulty      *hexutil.Big     `json:"difficulty"`
		Number          *hexutil.Big     `json:"number"`
		GasLimit        hexutil.Uint64   `json:"gasLimit"`
		GasUsed         hexutil.Uint64   `json:"gasUsed"`
		Time            hexutil.Uint64   `json:"timestamp"`
		Extra           hexutil.Bytes    `json:"extraData"`
		MixDigest       common.Hash      `json:"mixHash"`
		Nonce           types.BlockNonce `json:"nonce"`
		BaseFee         *hexutil.Big     `json:"baseFeePerGas"`
		Hash            common.Hash      `json:"hash"`
		L1BlockNumber   *hexutil.Big     `json:"l1BlockNumber"`
		SendCount       hexutil.Uint64   `json:"sendCount"`
		SendRoot        common.Hash      `json:"sendRoot"`
		Size            hexutil.Uint64   `json:"size"`
		TotalDifficulty *hexutil.Big     `json:"totalDifficulty"`
		Transactions    []Transaction    `json:"transactions"`
	}
	var enc Header
	enc.ParentHash = h.ParentHash
	enc.UncleHash = h.UncleHash
	enc.Coinbase = h.Coinbase
	enc.Root = h.Root
	enc.TxHash = h.TxHash
	enc.ReceiptHash = h.ReceiptHash
	enc.Bloom = h.Bloom
	enc.Difficulty = (*hexutil.Big)(h.Difficulty)
	enc.Number = (*hexutil.Big)(h.Number)
	enc.GasLimit = hexutil.Uint64(h.GasLimit)
	enc.GasUsed = hexutil.Uint64(h.GasUsed)
	enc.Time = hexutil.Uint64(h.Time)
	enc.Extra = h.Extra
	enc.MixDigest = h.MixDigest
	enc.Nonce = h.Nonce
	enc.BaseFee = (*hexutil.Big)(h.BaseFee)
	enc.L1BlockNumber = (*hexutil.Big)(h.L1BlockNumber)
	enc.SendCount = hexutil.Uint64(h.SendCount)
	enc.SendRoot = h.SendRoot
	enc.Size = hexutil.Uint64(h.Size)
	enc.TotalDifficulty = (*hexutil.Big)(h.TotalDifficulty)
	enc.Transactions = h.Transactions
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (h *Transaction) UnmarshalJSON(input []byte) error {
	type Header struct {
		ParentHash      *common.Hash      `json:"parentHash"`
		UncleHash       *common.Hash      `json:"sha3Uncles"`
		Coinbase        *common.Address   `json:"miner"`
		Root            *common.Hash      `json:"stateRoot"`
		TxHash          *common.Hash      `json:"transactionsRoot"`
		ReceiptHash     *common.Hash      `json:"receiptsRoot"`
		Bloom           *types.Bloom      `json:"logsBloom"`
		Difficulty      *hexutil.Big      `json:"difficulty"`
		Number          *hexutil.Big      `json:"number" `
		GasLimit        *hexutil.Uint64   `json:"gasLimit"`
		GasUsed         *hexutil.Uint64   `json:"gasUsed"`
		Time            *hexutil.Uint64   `json:"timestamp"`
		Extra           *hexutil.Bytes    `json:"extraData"`
		MixDigest       *common.Hash      `json:"mixHash"`
		Nonce           *types.BlockNonce `json:"nonce"`
		BaseFee         *hexutil.Big      `json:"baseFeePerGas"`
		L1BlockNumber   *hexutil.Big      `json:"l1BlockNumber"`
		SendCount       *hexutil.Uint64   `json:"sendCount"`
		SendRoot        *common.Hash      `json:"sendRoot"`
		Size            *hexutil.Uint64   `json:"size"`
		TotalDifficulty *hexutil.Big      `json:"totalDifficulty"`
		Transactions    []Transaction     `json:"transactions"`
	}
	var dec Header
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for Header")
	}
	h.ParentHash = *dec.ParentHash
	if dec.UncleHash == nil {
		return errors.New("missing required field 'sha3Uncles' for Header")
	}
	h.UncleHash = *dec.UncleHash
	if dec.Coinbase == nil {
		return errors.New("missing required field 'miner' for Header")
	}
	h.Coinbase = *dec.Coinbase
	if dec.Root == nil {
		return errors.New("missing required field 'stateRoot' for Header")
	}
	h.Root = *dec.Root
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionsRoot' for Header")
	}
	h.TxHash = *dec.TxHash
	if dec.ReceiptHash == nil {
		return errors.New("missing required field 'receiptsRoot' for Header")
	}
	h.ReceiptHash = *dec.ReceiptHash
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Header")
	}
	h.Bloom = *dec.Bloom
	if dec.Difficulty == nil {
		return errors.New("missing required field 'difficulty' for Header")
	}
	h.Difficulty = (*big.Int)(dec.Difficulty)
	if dec.Number == nil {
		return errors.New("missing required field 'number' for Header")
	}
	h.Number = (*big.Int)(dec.Number)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Header")
	}
	h.GasLimit = uint64(*dec.GasLimit)
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Header")
	}
	h.GasUsed = uint64(*dec.GasUsed)
	if dec.Time == nil {
		return errors.New("missing required field 'timestamp' for Header")
	}
	h.Time = uint64(*dec.Time)
	if dec.Extra == nil {
		return errors.New("missing required field 'extraData' for Header")
	}
	h.Extra = *dec.Extra
	if dec.MixDigest != nil {
		h.MixDigest = *dec.MixDigest
	}
	if dec.Nonce != nil {
		h.Nonce = *dec.Nonce
	}
	if dec.BaseFee != nil {
		h.BaseFee = (*big.Int)(dec.BaseFee)
	}
	if dec.L1BlockNumber != nil {
		h.L1BlockNumber = (*big.Int)(dec.L1BlockNumber)
	}
	if dec.SendCount != nil {
		h.SendCount = uint64(*dec.SendCount)
	}
	if dec.SendRoot != nil {
		h.SendRoot = *dec.SendRoot
	}
	if dec.Size != nil {
		h.Size = uint64(*dec.Size)
	}
	if dec.TotalDifficulty != nil {
		h.TotalDifficulty = (*big.Int)(dec.TotalDifficulty)
	}
	if dec.Transactions != nil && len(dec.Transactions) > 0 {
		h.Transactions = dec.Transactions
	}
	return nil
}
