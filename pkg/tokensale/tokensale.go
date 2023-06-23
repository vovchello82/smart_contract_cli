package tokensale

import "log"

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

func Withdraw() error {
	log.Println("transfer all mmoney to the owner")
	return nil
}

func UseLastToken() error {
	log.Println("use last token of mine")
	return nil
}
