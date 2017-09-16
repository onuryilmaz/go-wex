package wex

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// TradeAPI allows to trade on the exchange and receive information about the account.
//
// To use this API, you need to create an API key. An API key can be created in your Profile in the API Keys section. After creating an API key you’ll receive a key and a secret.
// Note that the Secret can be received only during the first hour after the creation of the Key.
// API key information is used for authentication.
type TradeAPI struct {
	API_KEY    string
	API_SECRET string
	lastNonce  int64
}

const tradeURL = "https://wex.nz/tapi"

// Auth provides API key and secret setting for Trade API
func (tapi *TradeAPI) Auth(key string, secret string) {
	tapi.API_KEY = key
	tapi.API_SECRET = secret
}

// GetInfo returns information about the user’s current balance, API-key privileges, the number of open orders and Server Time.
// To use this method you need a privilege of the key info.
func (tapi *TradeAPI) GetInfo() (AccountInfo, error) {
	info := AccountInfo{}
	err := tapi.call("getInfo", &info, nil)
	if err == nil {
		return info, nil
	}
	return info, err
}

// GetInfoAuth provides GetInfo capability with authorization
func (tapi *TradeAPI) GetInfoAuth(key string, secret string) (AccountInfo, error) {
	tapi.Auth(key, secret)
	return tapi.GetInfo()

}

// Trade provide method that can be used for creating orders and trading on the exchange.
// To use this method you need an API key privilege to trade.
//
// You can only create limit orders using this method, but you can emulate market orders using rate parameters. E.g. using rate=0.1 you can sell at the best market price.
//
// Each pair has a different limit on the minimum / maximum amounts, the minimum amount and the number of digits after the decimal point. All limitations can be obtained using the info method in PublicAPI.
func (tapi *TradeAPI) Trade(pair string, orderType string, rate float64, amount float64) (TradeResponse, error) {

	tradeResponse := TradeResponse{}

	orderParams := make(map[string]string, 4)
	orderParams["pair"] = pair
	orderParams["type"] = orderType
	orderParams["rate"] = strconv.FormatFloat(rate, 'f', -1, 64)
	orderParams["amount"] = strconv.FormatFloat(amount, 'f', -1, 64)

	err := tapi.call("Trade", &tradeResponse, orderParams)

	if err == nil {
		return tradeResponse, nil
	}
	return tradeResponse, err
}

// TradeAuth provides Trade capability with authorization
func (tapi *TradeAPI) TradeAuth(key string, secret string, pair string, orderType string, rate float64, amount float64) (TradeResponse, error) {
	tapi.Auth(key, secret)
	return tapi.Trade(pair, orderType, rate, amount)

}

// ActiveOrders returns the list of your active orders.  To use this method you need a privilege of the info key.
// If the order disappears from the list, it was either executed or canceled.
func (tapi *TradeAPI) ActiveOrders(pair string) (ActiveOrders, error) {

	orderParams := make(map[string]string, 4)
	orderParams["pair"] = pair

	activeOrders := make(ActiveOrders, 0)
	err := tapi.call("ActiveOrders", &activeOrders, orderParams)
	if err == nil {
		return activeOrders, nil
	}
	return activeOrders, err
}

// ActiveOrdersAuth provides ActiveOrders capability with authorization
func (tapi *TradeAPI) ActiveOrdersAuth(key string, secret string, pair string) (ActiveOrders, error) {
	tapi.Auth(key, secret)
	return tapi.ActiveOrders(pair)
}

// OrderInfo provides the information on particular order. To use this method you need a privilege of the info key.
func (tapi *TradeAPI) OrderInfo(orderID string) (OrderInfo, error) {

	orderInfo := OrderInfo{}

	orderParams := make(map[string]string, 1)
	orderParams["order_id"] = orderID

	err := tapi.call("OrderInfo", &orderInfo, orderParams)
	if err == nil {
		return orderInfo, nil
	}
	return orderInfo, err
}

// OrderInfoAuth provides OrderInfo capability with authorization
func (tapi *TradeAPI) OrderInfoAuth(key string, secret string, orderID string) (OrderInfo, error) {
	tapi.Auth(key, secret)
	return tapi.OrderInfo(orderID)

}

// CancelOrder provides method used for order cancellation. To use this method you need a privilege of the trade key.
func (tapi *TradeAPI) CancelOrder(orderID string) (CancelOrder, error) {

	cancelReponse := CancelOrder{}

	orderParams := make(map[string]string, 1)
	orderParams["order_id"] = orderID

	err := tapi.call("CancelOrder", &cancelReponse, orderParams)

	if err == nil {
		return cancelReponse, nil
	}
	return cancelReponse, err
}

// CancelOrderAuth provides CancelOrder capability with authorization
func (tapi *TradeAPI) CancelOrderAuth(key string, secret string, orderID string) (CancelOrder, error) {
	tapi.Auth(key, secret)
	return tapi.CancelOrder(orderID)

}

// TradeHistory returns trade history. To use this method you need a privilege of the info key.
func (tapi *TradeAPI) TradeHistory(filter HistoryFilter, pair string) (TradeHistory, error) {

	tradeHistory := TradeHistory{}

	historyParams := historyFilterParams(filter)
	if pair != "" {
		historyParams["pair"] = pair
	}

	err := tapi.call("TradeHistory", &tradeHistory, historyParams)

	if err == nil {
		return tradeHistory, nil
	}
	return tradeHistory, err
}

// TradeHistoryAuth provides TradeHistory capability with authorization
func (tapi *TradeAPI) TradeHistoryAuth(key string, secret string, filter HistoryFilter, pair string) (TradeHistory, error) {
	tapi.Auth(key, secret)
	return tapi.TradeHistory(filter, pair)

}

// TransactionHistory returns the history of transactions. To use this method you need a privilege of the info key.
func (tapi *TradeAPI) TransactionHistory(filter HistoryFilter) (TransactionHistory, error) {

	transactionHistory := TransactionHistory{}

	historyParams := historyFilterParams(filter)

	err := tapi.call("TransHistory", &transactionHistory, historyParams)

	if err == nil {
		return transactionHistory, nil
	}
	return transactionHistory, err
}

// TransactionHistoryAuth provides TransactionHistory capability with authorization
func (tapi *TradeAPI) TransactionHistoryAuth(key string, secret string, filter HistoryFilter) (TransactionHistory, error) {
	tapi.Auth(key, secret)
	return tapi.TransactionHistory(filter)

}

// WithdrawCoin provides cryptocurrency withdrawals. You need to have the privilege of the Withdraw key to be able to use this method.
func (tapi *TradeAPI) WithdrawCoin(coinName string, amount float64, address string) (WithdrawCoin, error) {

	response := WithdrawCoin{}

	orderParams := make(map[string]string, 3)
	orderParams["coinName"] = coinName
	orderParams["amount"] = strconv.FormatFloat(amount, 'f', -1, 64)
	orderParams["address"] = address

	err := tapi.call("WithdrawCoin", &response, orderParams)

	if err == nil {
		return response, nil
	}
	return response, err
}

// WithdrawCoinAuth provides WithdrawCoin capability with authorization
func (tapi *TradeAPI) WithdrawCoinAuth(key string, secret string, coinName string, amount float64, address string) (WithdrawCoin, error) {
	tapi.Auth(key, secret)
	return tapi.WithdrawCoin(coinName, amount, address)

}

// CreateCoupon allows you to create Coupons. In order to use this method, you need the Coupon key privilege.
func (tapi *TradeAPI) CreateCoupon(currency string, amount float64) (CreateCoupon, error) {

	response := CreateCoupon{}

	params := make(map[string]string, 2)
	params["currency"] = currency
	params["amount"] = strconv.FormatFloat(amount, 'f', -1, 64)

	err := tapi.call("CreateCoupon", &response, params)

	if err == nil {
		return response, nil
	}
	return response, err
}

// CreateCouponAuth provides CreateCoupon capability with authorization
func (tapi *TradeAPI) CreateCouponAuth(key string, secret string, currency string, amount float64) (CreateCoupon, error) {
	tapi.Auth(key, secret)
	return tapi.CreateCoupon(currency, amount)
}

// RedeemCoupon method is used to redeem coupons. In order to use this method, you need the Coupon key privilege.
func (tapi *TradeAPI) RedeemCoupon(coupon string) (RedeemCoupon, error) {

	response := RedeemCoupon{}

	params := make(map[string]string, 1)
	params["coupon"] = coupon

	err := tapi.call("RedeemCoupon", &response, params)

	if err == nil {
		return response, nil
	}
	return response, err
}

// RedeemCouponAuth provides RedeemCoupon capability with authorization
func (tapi *TradeAPI) RedeemCouponAuth(key string, secret string, coupon string) (RedeemCoupon, error) {
	tapi.Auth(key, secret)
	return tapi.RedeemCoupon(coupon)
}

func (tapi *TradeAPI) encodePostData(method string, params map[string]string) string {
	nonce := time.Now().Unix()
	if nonce <= tapi.lastNonce {
		nonce = tapi.lastNonce + 1
	}
	tapi.lastNonce = nonce

	result := fmt.Sprintf("method=%s&nonce=%d", method, nonce)

	if len(params) > 0 {
		v := url.Values{}
		for key := range params {
			v.Add(key, params[key])
		}
		result = result + "&" + v.Encode()
	}
	return result
}

func sign(secret string, payload string) string {
	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}

func (tapi *TradeAPI) call(method string, v interface{}, params map[string]string) error {

	postData := tapi.encodePostData(method, params)

	req, err := http.NewRequest("POST", tradeURL, bytes.NewBufferString(postData))

	if err != nil {
		return err
	}

	req.Header.Add("Key", tapi.API_KEY)
	req.Header.Add("Sign", sign(tapi.API_SECRET, postData))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(postData)))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	return marshalResponse(resp, v)
}

func marshalResponse(resp *http.Response, v interface{}) error {
	// read the response
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	data := Response{}

	if err = json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	if data.Success == 1 {
		if err = json.Unmarshal(data.Return, &v); err != nil {
			return err
		}
	} else {
		return TradeError{data.Error}
	}

	return nil
}

// custom error type for server/trading errors
type TradeError struct {
	msg string
}

func (e TradeError) Error() string {
	return fmt.Sprintf("trading error: %v", e.msg)
}

// historyFilterParams creates map[string]string mapping of HistoryFilter
func historyFilterParams(filter HistoryFilter) map[string]string {
	params := make(map[string]string, 0)

	if filter.From > 0 {
		params["from"] = strconv.Itoa(filter.From)
	}
	if filter.Count > 0 {
		params["count"] = strconv.Itoa(filter.Count)
	}
	if filter.FromID > 0 {
		params["from_id"] = strconv.Itoa(filter.FromID)
	}
	if filter.EndID > 0 {
		params["end_id"] = strconv.Itoa(filter.EndID)
	}
	if filter.Order == "ASC" || filter.Order == "DESC" {
		params["order"] = filter.Order
	}
	if filter.Since.Unix() > 0 {
		params["since"] = strconv.FormatInt(filter.Since.Unix(), 10)
	}
	if filter.End.Unix() > 0 {
		params["end"] = strconv.FormatInt(filter.End.Unix(), 10)
	}
	return params
}
