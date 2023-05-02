package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"os"
)

var walletFilePath = "D:\\MyProject\\walletFile\\erc20"
var wallet1File = "D:\\MyProject\\walletFile\\erc20\\UTC--2023-04-30T15-39-31.888702400Z--f8057ae999615e356ecbd6a9464f83d66465eaf1"
var wallet2File = "D:\\MyProject\\walletFile\\erc20\\UTC--2023-04-30T04-27-58.45000000Z--0bbd943e04ecb39ac6c09a2982d780e7e5a5290c.json"
var ethRpcUrl = "https://goerli.infura.io/v3/89aae6dea4504bc8a4485ad6219df6b3"
var polygonRpcUrl = "https://rpc-mumbai.maticvigil.com/v1/99300cd360366c25a0222fc8b60323ba84f975a1"
var password = "123456"
var addressTo = "0x221CFd8877880EF2CA4847d8D114E77669243045"

func main() {
	//url, address := createWallet()
	//fmt.Println(url)
	//fmt.Println(address)
	_, address := getPrivateKey(wallet2File, password)
	client, err := ethclient.Dial(ethRpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	balance(client, address)
	//balance(client, addressTo)
	keyStore := importWallet(walletFilePath)
	//transferEth(client, keyStore, address, b, addressTo)
	transferWithEip1559(client, keyStore, address, 0.0001, addressTo)
	//getGapPrice(client, ethereum.CallMsg{})
}

func getGapPrice(client *ethclient.Client, msg ethereum.CallMsg) (*big.Int, *big.Int, uint64) {
	estimateGas, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Fatal(err)
	}
	maxPriorityFeePerGas, err1 := client.SuggestGasTipCap(context.Background())
	gasPrice, err1 := client.SuggestGasPrice(context.Background())
	if err1 != nil {
		log.Fatal(err1)
	}
	gasPriceDecimal := decimal.NewFromBigInt(gasPrice, 0)
	maxPriorityFeePerGasDecimal := decimal.NewFromBigInt(maxPriorityFeePerGas, 0)
	baseGasDecimal := gasPriceDecimal.Sub(maxPriorityFeePerGasDecimal)
	maxFeePerGas := baseGasDecimal.Mul(decimal.NewFromFloat(1.5)).Add(maxPriorityFeePerGasDecimal)
	fmt.Println("estimateGas: ", estimateGas)
	fmt.Println("baseGasDecimal: ", weiToGWei(baseGasDecimal.BigInt()))
	fmt.Println("maxPriorityFeePerGas: ", weiToGWei(maxPriorityFeePerGas))
	fmt.Println("maxFeePerGas: ", weiToGWei(maxFeePerGas.BigInt()))
	return maxPriorityFeePerGas, maxFeePerGas.BigInt(), estimateGas
}

func createWallet() (key, address string) {
	ks := keystore.NewKeyStore(walletFilePath, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	address = account.Address.Hex()
	return account.URL.String(), address
}

func importWallet(filePath string) *keystore.KeyStore {
	ks := keystore.NewKeyStore(filePath, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks
}

func getPrivateKey(filePath, password string) (string, string) {
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	key, err2 := keystore.DecryptKey(readFile, password)
	if err2 != nil {
		log.Fatal(err2)
	}
	privateKey := hex.EncodeToString(key.PrivateKey.D.Bytes())
	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	fmt.Println("walletAddress: ", address)
	log.Println(privateKey)
	return privateKey, address.String()
}

func balance(client *ethclient.Client, address string) *big.Float {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address, " balance: ", balance)
	etherBalance := weiToEther(balance)
	fmt.Println(address, " etherBalance: ", etherBalance)
	return etherBalance.BigFloat()
}

func transferEth(client *ethclient.Client, ks *keystore.KeyStore, addressFrom string, float float64, addressTo string) string {
	fromAccount := accounts.Account{Address: common.HexToAddress(addressFrom)}
	toAccount := accounts.Account{Address: common.HexToAddress(addressTo)}
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(addressFrom))
	if err != nil {
		log.Fatal(err)
	}
	numWei := etherToWei(float)
	gasLimit := uint64(21000)
	gasPrice, err1 := client.SuggestGasPrice(context.Background())
	if err1 != nil {
		log.Fatal(err1)
	}
	chainID, err2 := client.ChainID(context.Background())
	if err2 != nil {
		log.Fatal(err2)
	}
	transaction := types.NewTx(&types.AccessListTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAccount.Address,
		Value:    numWei.BigInt(),
		ChainID:  chainID,
		Data:     nil,
	})
	signAccount, err := ks.Find(fromAccount)
	if err != nil {
		log.Fatal(err)
	}
	err = ks.Unlock(signAccount, password)
	if err != nil {
		log.Fatal(err)
	}
	signTx, err := ks.SignTx(fromAccount, transaction, chainID)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}
	hash := signTx.Hash().Hex()
	fmt.Println(hash)
	return hash
}

func transferWithEip1559(client *ethclient.Client, ks *keystore.KeyStore, addressFrom string, float float64, addressTo string) string {
	fromAccount := accounts.Account{Address: common.HexToAddress(addressFrom)}
	toAccount := accounts.Account{Address: common.HexToAddress(addressTo)}
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(addressFrom))
	if err != nil {
		log.Fatal(err)
	}
	numWei := etherToWei(float)
	chainID, err2 := client.ChainID(context.Background())
	if err2 != nil {
		log.Fatal(err2)
	}
	maxPriorityFeePerGas, maxFeePerGas, gasLimit := getGapPrice(client, ethereum.CallMsg{})
	transaction := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		GasTipCap: maxPriorityFeePerGas,
		GasFeeCap: maxFeePerGas,
		Gas:       gasLimit,
		To:        &toAccount.Address,
		Value:     numWei.BigInt(),
		ChainID:   chainID,
		Data:      nil,
	})
	signAccount, err := ks.Find(fromAccount)
	if err != nil {
		log.Fatal(err)
	}
	err = ks.Unlock(signAccount, password)
	if err != nil {
		log.Fatal(err)
	}
	signTx, err := ks.SignTx(fromAccount, transaction, chainID)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}
	hash := signTx.Hash().Hex()
	fmt.Println(hash)
	return hash
}

func tokenTransfer(client *ethclient.Client, ks *keystore.KeyStore, contract string, addressFrom string, float float64, addressTo string) string {
	return ""
}

func weiToEther(i *big.Int) decimal.Decimal {
	bf := decimal.NewFromBigInt(i, 0)
	float := decimal.NewFromFloat(params.Ether)
	div := bf.DivRound(float, 18)
	return div
}

func etherToWei(bf float64) decimal.Decimal {
	de := decimal.NewFromFloat(bf)
	return de.Mul(decimal.NewFromFloat(params.Ether))
}

func weiToGWei(i *big.Int) decimal.Decimal {
	bf := decimal.NewFromBigInt(i, 0)
	return bf.DivRound(decimal.NewFromFloat(params.GWei), 18)
}

func gWeiToWei(bf float64) decimal.Decimal {
	de := decimal.NewFromFloat(bf)
	return de.Mul(decimal.NewFromFloat(params.GWei))
}
