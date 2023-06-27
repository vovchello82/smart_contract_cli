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

	address := common.HexToAddress("0x4e492c31c51a8b73433a9b9c30fa0242e4e96362")
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
	price, err := tokenManger.instance.CurrentPriceInWei(nil)
	if err != nil {
		log.Printf("error on getting price %s", err)
	} else {
		log.Printf("current price %d wei", price)
	}

	return nil
}

func Withdraw() error {
	log.Println("transfer all money to the owner")
	return nil
}

func UseLastToken() error {
	log.Println("use last token of mine")
	return nil
}
