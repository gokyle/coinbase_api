package coinbase_api

// This file contains the data types and constants used in the package.

import (
	"fmt"
	"strconv"
)

// BaseURL for the API
const api_base = "https://coinbase.com/api/v1/"

// ApiKey should contain the API key. It is global because the user needs
// to set it.
var ApiKey string

// Set to true to dump request JSON files.
var Debug = false

// The minimum number of bitcoins that may be purchased.
const MinimumPurchase = 0.10

// A collection of errors returned by the API.
var (
	ErrNotAuthenticated        = fmt.Errorf("Invalid API key.")
	ErrMinimumSubtotal         = fmt.Errorf("The minimum subtotal allowed is 0.10 BTC.  Please enter a larger BTC value.")
	ErrFirstPurchaseIncomplete = fmt.Errorf("Please wait until your first bitcoin purchase completes before making additional purchases")
	ErrSiteMaxed               = fmt.Errorf("Sorry, the maximum number of purchases on Coinbase has been reached for today.  Please try again in 24 hours.")
)

// The Exchange type represents an exchange ratio between BTC and various currencies.
type ExchangeRate map[string]string

// GetAuthenticated represents a data type that may be used as the body for an
// authenticated GET request.
type GetAuthenticated struct {
	Key string `json:"api_key"`
}

// Set the API key on an authenticated GET request.
func (get *GetAuthenticated) SetApiKey() {
	get.Key = ApiKey
}

// A TranasactionRequest is a request for purchasing some amount of bitcoins.
type TransactionRequest struct {
	Qty float64 `json:"qty"`
	Key string  `json:"api_key"`
}

func (tr *TransactionRequest) SetApiKey() {
	tr.Key = ApiKey
}

// Fee represents a specific fee incurred durring a transaction.
type Fee struct {
	Cents    int    `json:"cents"`
	Currency string `json:"currency_iso"`
}

// Balance represents some amount of money paired to a currency.
type Balance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// Numeric retrieves the numeric value of a balance.
func (b *Balance) Numeric() (float64, error) {
	amount, err := strconv.ParseFloat(b.Amount, 64)
	return amount, err
}

// Transaction represents an attempted transaction, whether a purchase
// or sale.
type TransactionResult struct {
	Success  bool     `json:"success"`
	Errors   []string `json:"errors"`
	Transfer struct {
		Type     string         `json:"_type"`
		Code     string         `json:"code"`
		Created  string         `json:"created_at"`
		Fees     map[string]Fee `json:"fees"`
		Status   string         `json:"created"`
		Payout   string         `json:"payout_date"`
		BTC      Balance        `json:"btc"`
		Subtotal Balance        `json:"subtotal"`
		Total    Balance        `json:"total"`
	} `json:"transfer"`
}

// ReceiveAddress contains the bitcoin address the user can send bitcoins to
// for sale.
type ReceiveAddress struct {
	Address     string `json:"address"`
	CallbackURL string `json:"callback_url"`
}

// Type User represents a Coinbase user.
type UserOverview struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Type Transaction represents a successful transaction. Due to the way that
// the Coinbase operates, the transaction data is stored in the T field.
type Transaction struct {
	Id        string       `json:"id"`
	CreatedAt string       `json:"created_at"`
	Amount    Balance      `json:"amount"`
	Request   bool         `json:"request"`
	Status    string       `json:"status"`
	Sender    UserOverview `json:"string"`
	Recipient UserOverview `json:"string"`
}

type TransactionContainer struct {
	T       Transaction `json:"transaction"`
	Success bool        `json:"success"`
	Error   string      `json:"error"`
}

func (tc *TransactionContainer) GetError() (err error) {
	if tc.Error != "" {
		err = fmt.Errorf(tc.Error)
	}
	return
}

// Type TransactionList stores a list of transactions for a user.
type TransactionList struct {
	CurrentUser    UserOverview  `json:"current_user"`
	CurrentBalance Balance       `json:"balance"`
	TotalCount     int           `json:"total_count"`
	NumPages       int           `json:"num_pages"`
	CurrentPage    int           `json:"current_pages"`
	Transactions   []Transaction `json:"transactions"`
}

// A TransactionListRequest is used to request the list of transactions for
// the currently-logged in user.
type TransactionListRequest struct {
	Page int    `json:"page"`
	Key  string `json:"key"`
}

func (tlr *TransactionListRequest) SetApiKey() {
	tlr.Key = ApiKey
}

// UserContainer is required due to the way the API returns user results.
type UserContainer struct {
	U     User   `json:"user"`
	Error string `json:"string"`
}

func (uc *UserContainer) GetError() (err error) {
	if uc.Error != "" {
		err = fmt.Errorf(uc.Error)
	}
	return
}

// Type User represents the currently authenticated user.
type User struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	TimeZone       string  `json:"time_zone"`
	NativeCurrency string  `json:"native_currency"`
	CurrentBalance Balance `json:"balance"`
	BuyLevel       int     `json:"buy_level"`
	SellLevel      int     `json:"sell_level"`
	BuyLimit       Balance `json:"buy_limit"`
	SellLimit      Balance `json:"sell_limit"`
}
