package tokensale

import (
	"log"
	tokensale "tokensale/third_party"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TokenSale struct {
	instance *tokensale.Tokensale
}

var tokenManger *TokenSale = NewTokenSale()

func NewTokenSale() *TokenSale {
	client, err := ethclient.Dial("https://rpc2.sepolia.org")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("")
	contractInstance, err := tokensale.NewTokensale(address, client)
	if err != nil {
		log.Fatalf("couldn't access smart contract %s", err)
	}

	return &TokenSale{
		instance: contractInstance,
	}
}

func BuyToken() error {
	log.Println("buy a token")
	return nil
}

func PushToken(token string) error {
	log.Printf("push a new token %s", token)
	return nil
}

func UpdateTokenPrice(newPrice uint64) error {
	log.Printf("set token price to %d", newPrice)
	return nil
}

func GetCurrentPriceInWei() error {
	log.Printf("get current price in wei")
	return nil
}

func Withdraw() error {
	log.Println("transfer all mmoney to the owner")
	return nil
}

func UseLastToken() error {
	log.Println("use last token of mine")
	return nil
}
