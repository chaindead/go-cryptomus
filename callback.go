package gocryptomus

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	CallbackStatusPaid         = "paid"
	CallbackStatusPaidOver     = "paid_over"
	CallbackStatusCancel       = "cancel"
	CallbackStatusConfirmCheck = "confirm_check"
)

type Callback struct {
	Type              string      `json:"type"`
	Uuid              string      `json:"uuid"`
	OrderId           string      `json:"order_id"`
	Amount            string      `json:"amount"`
	PaymentAmount     string      `json:"payment_amount"`
	PaymentAmountUsd  string      `json:"payment_amount_usd"`
	MerchantAmount    string      `json:"merchant_amount"`
	Commission        string      `json:"commission"`
	IsFinal           bool        `json:"is_final"`
	Status            string      `json:"status"`
	From              string      `json:"from"`
	WalletAddressUuid interface{} `json:"wallet_address_uuid"`
	Network           string      `json:"network"`
	Currency          string      `json:"currency"`
	PayerCurrency     string      `json:"payer_currency"`
	AdditionalData    interface{} `json:"additional_data"`
	Convert           struct {
		ToCurrency string      `json:"to_currency"`
		Commission interface{} `json:"commission"`
		Rate       string      `json:"rate"`
		Amount     string      `json:"amount"`
	} `json:"convert"`
	Txid string `json:"txid"`
	Sign string `json:"sign"`
}

func (c Callback) String() string {
	return c.OrderId
}

func (c *Cryptomus) VerifySignature(body []byte) bool {
	bodyStr := string(body)
	sign := gjson.Get(bodyStr, "sign").String()
	updatedJsonData, _ := sjson.Delete(bodyStr, "sign")

	// Base64 encode the JSON string
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(updatedJsonData))

	// Concatenate the Base64 encoded string with the API key
	concatenatedString := base64Encoded + c.paymentApiKey

	// Compute the MD5 hash of the concatenated string
	hash := md5.Sum([]byte(concatenatedString))

	// Convert the hash to a hexadecimal string
	hashString := fmt.Sprintf("%x", hash)

	return hashString == sign
}
