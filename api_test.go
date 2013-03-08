package coinbase_api

import (
	"fmt"
	"os"
	"testing"
)

var NotAuthenticated = fmt.Errorf("no API key present: can't make authenticated requests")

func init() {
	ApiKey = os.Getenv("CB_API_KEY")
}

// FailWithError is a utility for dumping errors and failing the test.
func FailWithError(t *testing.T, err error) {
	fmt.Println("failed")
	if err != nil {
		fmt.Println("[!] ", err.Error())
	}
	t.FailNow()
}

func TestGetCurrencies(t *testing.T) {
	fmt.Printf("Getting list of supported currencies: ")
	_, err := GetCurrencies()
	if err != nil {
		FailWithError(t, err)
	}
	fmt.Println("ok")
}

func TestExchangeRate(t *testing.T) {
	fmt.Printf("Getting exchange rate: ")
	currencies := []string{"btc_to_usd", "usd_to_btc"}

	_, err := GetExchangeRates(currencies)
	if err != nil {
		FailWithError(t, err)
	}

	fmt.Println("ok")
}

func TestGetAccountBalance(t *testing.T) {
	fmt.Printf("Getting account balance: ")
	if ApiKey == "" {
		FailWithError(t, NotAuthenticated)
	}
	_, err := GetAccountBalance()
	if err != nil {
		FailWithError(t, err)
	}
	fmt.Println("ok")
}

func TestGetReceiveAddress(t *testing.T) {
	fmt.Printf("Getting receive address: ")
	if ApiKey == "" {
		FailWithError(t, NotAuthenticated)
	}
	_, err := GetReceiveAddress()
	if err != nil {
		FailWithError(t, err)
	}
	fmt.Println("ok")
}

func TestSellPrice(t *testing.T) {
	fmt.Printf("Getting sale price for 1 BTC: ")
	_, err := GetSellPrice(1.00)
	if err != nil {
		FailWithError(t, err)
	}
	fmt.Println("ok")
}

func TestBuyPrice(t *testing.T) {
	fmt.Printf("Getting purchase cost for 1 BTC: ")
	_, err := GetBuyPrice(1.00)
	if err != nil {
		FailWithError(t, err)
	}
	fmt.Println("ok")
}

func TestGetTransactions(t *testing.T) {
	fmt.Printf("Getting transaction list: ")
	_, err := GetTransactions(0)
	if err != nil {
		FailWithError(t, err)
	}
	fmt.Println("ok")
}

func TestGetTransaction(t *testing.T) {
	fmt.Printf("Getting a single transaction: ")
	_, err := GetTransaction("51382d38aa5add28f8000062")
	if err != nil {
		FailWithError(t, err)
	}
	fmt.Println("ok")
}

func TestGetUser(t *testing.T) {
	fmt.Printf("Getting information for the current user: ")
	_, err := GetUser()
	fmt.Printf("u: %+v\n", u)
	if err != nil {
		FailWithError(t, err)
	}
	fmt.Println("ok")
}
