package btce

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var API_KEY = ""
var API_SECRET = ""
var lastNonce int64 = 0

const TRADE_URL = "https://btc-e.com/tapi"

func encodePostData(method string, params map[string]string) string {
	nonce := time.Now().Unix()
	if nonce <= lastNonce {
		nonce = nonce + 1
	}
	lastNonce = nonce
	result := fmt.Sprintf("method=%s&nonce=%d", method, nonce)

	// params are unordered, but after method and nonce
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

func call(method string, v interface{}, params map[string]string) error {

	postData := encodePostData(method, params)

	req, err := http.NewRequest("POST", TRADE_URL, bytes.NewBufferString(postData))
	req.Header.Add("Key", API_KEY)
	req.Header.Add("Sign", sign(API_SECRET, postData))
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
		fmt.Println("ReadAll() failed: %v", err)
		return err
	}

	data := Response{}

	if err = json.Unmarshal(bytes, &data); err != nil {
		fmt.Println(err)
		return err
	}

	if data.Success == 1 {
		if err = json.Unmarshal(data.Return, &v); err != nil {
			fmt.Println(err)
			return err
		}
	} else {
		fmt.Println(data.Error)
		return errors.New(data.Error)
	}

	return nil
}

func Auth(key string, secret string) {
	API_KEY = key
	API_SECRET = secret
}
func GetAccountInfo() (AccountInfo, error) {
	info := AccountInfo{}
	err := call("getInfo", &info, nil)
	if err == nil {
		return info, nil
	} else {
		return info, err
	}
}
func GetAccountInfoAuth(key string, secret string) (AccountInfo, error) {
	Auth(key, secret)
	return GetAccountInfo()

}

func GetActiveOrders(pair string) (ActiveOrders, error) {

	orderParams := make(map[string]string, 4)
	orderParams["pair"] = pair

	activeOrders := make(ActiveOrders, 0)
	err := call("ActiveOrders", &activeOrders, orderParams)
	if err == nil {
		return activeOrders, nil
	} else {
		return activeOrders, err
	}
}
func GetActiveOrdersAuth(key string, secret string, pair string) (ActiveOrders, error) {
	Auth(key, secret)
	return GetActiveOrders(pair)

}

func OrderTrade(pair string, orderType string, rate float64, amount float64) (TradeResponse, error) {

	tradeResponse := TradeResponse{}

	orderParams := make(map[string]string, 4)
	orderParams["pair"] = pair
	orderParams["type"] = orderType
	orderParams["rate"] = strconv.FormatFloat(rate, 'f', -1, 64)
	orderParams["amount"] = strconv.FormatFloat(amount, 'f', -1, 64)

	err := call("Trade", &tradeResponse, orderParams)

	if err == nil {
		return tradeResponse, nil
	} else {
		return tradeResponse, err
	}

}

func OrderTradeAuth(key string, secret string, pair string, orderType string, rate float64, amount float64) (TradeResponse, error) {
	Auth(key, secret)
	return OrderTrade(pair, orderType, rate, amount)

}

func GetOrderInfo(orderID string) (OrderInfos, error) {

	orderInfo := OrderInfos{}

	orderParams := make(map[string]string, 1)
	orderParams["order_id"] = orderID

	err := call("OrderInfo", &orderInfo, orderParams)
	if err == nil {
		fmt.Println(err)
		return orderInfo, nil
	} else {
		return orderInfo, err
	}

}

func GetOrderInfoAuth(key string, secret string, orderID string) (OrderInfos, error) {
	Auth(key, secret)
	return GetOrderInfo(orderID)

}

func CancelOrder(orderID string) (CancelOrderResponse, error) {

	cancelReponse := CancelOrderResponse{}

	orderParams := make(map[string]string, 1)
	orderParams["order_id"] = orderID

	err := call("CancelOrder", &cancelReponse, orderParams)

	if err == nil {
		return cancelReponse, nil
	} else {
		return cancelReponse, err
	}

}

func CancelOrderAuth(key string, secret string, orderID string) (CancelOrderResponse, error) {
	Auth(key, secret)
	return CancelOrder(orderID)

}

func GetTradeHistory(filter HistoryFilter, pair string) (TradeHistory, error) {

	tradeHistory := TradeHistory{}

	historyParams := historyFilterParams(filter)
	if pair != "" {
		historyParams["pair"] = pair
	}

	err := call("TradeHistory", &tradeHistory, historyParams)

	if err == nil {
		return tradeHistory, nil
	} else {
		return tradeHistory, err
	}

}

func GetTradeHistoryAuth(key string, secret string, filter HistoryFilter, pair string) (TradeHistory, error) {
	Auth(key, secret)
	return GetTradeHistory(filter, pair)

}

func GetTransactionHistory(filter HistoryFilter) (TransactionHistory, error) {

	transactionHistory := TransactionHistory{}

	historyParams := historyFilterParams(filter)

	err := call("TransHistory", &transactionHistory, historyParams)

	if err == nil {
		return transactionHistory, nil
	} else {
		return transactionHistory, err
	}

}

func GetTransactionHistoryAuth(key string, secret string, filter HistoryFilter) (TransactionHistory, error) {
	Auth(key, secret)
	return GetTransactionHistory(filter)

}
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

func WithdrawCoin(coinName string, amount float64, address string) (WithdrawCoinResponse, error) {

	response := WithdrawCoinResponse{}

	orderParams := make(map[string]string, 3)
	orderParams["coinName"] = coinName
	orderParams["amount"] = strconv.FormatFloat(amount, 'f', -1, 64)
	orderParams["address"] = address

	err := call("WithdrawCoin", &response, orderParams)

	if err == nil {
		return response, nil
	} else {
		return response, err
	}

}

func WithdrawCoinAuth(key string, secret string, coinName string, amount float64, address string) (WithdrawCoinResponse, error) {
	Auth(key, secret)
	return WithdrawCoin(coinName, amount, address)

}

func CreateCoupon(currency string, amount float64) (CreateCouponResponse, error) {

	response := CreateCouponResponse{}

	params := make(map[string]string, 2)
	params["currency"] = currency
	params["amount"] = strconv.FormatFloat(amount, 'f', -1, 64)

	err := call("CreateCoupon", &response, params)

	if err == nil {
		return response, nil
	} else {
		return response, err
	}

}

func CreateCouponAuth(key string, secret string, currency string, amount float64) (CreateCouponResponse, error) {
	Auth(key, secret)
	return CreateCoupon(currency, amount)
}

func RedeemCoupon(coupon string) (RedeemCouponResponse, error) {

	response := RedeemCouponResponse{}

	params := make(map[string]string, 1)
	params["coupon"] = coupon

	err := call("RedeemCoupon", &response, params)

	if err == nil {
		return response, nil
	} else {
		return response, err
	}

}

func RedeemCouponAuth(key string, secret string, coupon string) (RedeemCouponResponse, error) {
	Auth(key, secret)
	return RedeemCoupon(coupon)
}
