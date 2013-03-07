package coinbase_api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// An AuthenticatedRequest needs to have a way to load the API key.
type AuthenticatedRequest interface {
	SetApiKey()
}

// PostAuthenticatedRequest makes an authenticated JSON POST request to the API.
// The structure containing the result should be passed into res.
func PostAuthenticatedRequest(data AuthenticatedRequest, endpoint string, res interface{}) (err error) {
        if ApiKey == "" {
                return ErrNotAuthenticated
        }
	data.SetApiKey()

	request_json_body, err := json.Marshal(data)
	if err != nil {
		return
	}
	request_body := bytes.NewBuffer(request_json_body)

	client := &http.Client{}
	req, err := http.NewRequest("POST", api_base+endpoint, request_body)
	if err != nil {
		return
	}

	req.Header.Add("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	response_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.Unmarshal(response_body, &res)
	return
}

// Make an authenticated GET request to the API.
func GetAuthenticatedRequest(data AuthenticatedRequest, endpoint string, res interface{}) (err error) {
        if ApiKey == "" {
                return ErrNotAuthenticated
        }
	data.SetApiKey()
	request_json_body, err := json.Marshal(data)
	if err != nil {
		return
	}
	request_body := bytes.NewBuffer(request_json_body)

	client := &http.Client{}
	req, err := http.NewRequest("GET", api_base+endpoint, request_body)
	if err != nil {
		return
	}

	req.Header.Add("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	response_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.Unmarshal(response_body, &res)
	return
}

// Retrieve exchange rates for a list of currencies.
func GetExchangeRates(currencies []string) (ExchangeRate, error) {
	exch := make(ExchangeRate, 0)
	for _, currency := range currencies {
		exch[currency] = ""
	}
	endpoint := api_base + "currencies/exchange_rates"
	resp, err := http.Get(endpoint)
	if err != nil {
		return exch, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return exch, err
	}

	err = json.Unmarshal(body, &exch)
	return exch, err
}

// PurchaseBTC attempts to purchase the specified quantity of bitcoins.
func PurchaseBTC(qty float64) (p *Transaction, err error) {
	pr := new(TransactionRequest)
	pr.Qty = qty

	p = new(Transaction)
	endpoint := "buys"

	err = PostAuthenticatedRequest(pr, endpoint, &p)
	return
}

// SellBTC attempts to sell the specified quantity of bitcoins.
func SellBTC(qty float64) (p *Transaction, err error) {
	pr := new(TransactionRequest)
	pr.Qty = qty

	p = new(Transaction)
	endpoint := "buys"

	err = PostAuthenticatedRequest(pr, endpoint, &p)
	return
}



// GetAccountBalance retrieves the number of bitcoins in the user's account.
func GetAccountBalance() (b *Balance, err error) {
	get := new(GetAuthenticated)
	endpoint := "account/balance"

	b = new(Balance)
	err = GetAuthenticatedRequest(get, endpoint, &b)
	return
}

// GetBitcoinAddress retrieves the number of bitcoins in the user's account.
func GetReceiveAddress() (a *ReceiveAddress, err error) {
	get := new(GetAuthenticated)
	endpoint := "account/receive_address"

	a = new(ReceiveAddress)
	err = GetAuthenticatedRequest(get, endpoint, &a)
	return
}
