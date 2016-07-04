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

type TradeAPI struct {
	API_KEY    string
	API_SECRET string
	lastNonce  int64
}

const TRADE_URL = "https://btc-e.com/tapi"

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

	req, err := http.NewRequest("POST", TRADE_URL, bytes.NewBufferString(postData))
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

func (tapi *TradeAPI) Auth(key string, secret string) {
	tapi.API_KEY = key
	tapi.API_SECRET = secret
}

func (tapi *TradeAPI) AccountInfo() (AccountInfo, error) {
	info := AccountInfo{}
	err := tapi.call("getInfo", &info, nil)
	if err == nil {
		return info, nil
	} else {
		return info, err
	}
}

func (tapi *TradeAPI) AccountInfoAuth(key string, secret string) (AccountInfo, error) {
	tapi.Auth(key, secret)
	return tapi.AccountInfo()

}

func (tapi *TradeAPI) ActiveOrders(pair string) (ActiveOrders, error) {

	orderParams := make(map[string]string, 4)
	orderParams["pair"] = pair

	activeOrders := make(ActiveOrders, 0)
	err := tapi.call("ActiveOrders", &activeOrders, orderParams)
	if err == nil {
		return activeOrders, nil
	} else {
		return activeOrders, err
	}
}

func (tapi *TradeAPI) ActiveOrdersAuth(key string, secret string, pair string) (ActiveOrders, error) {
	tapi.Auth(key, secret)
	return tapi.ActiveOrders(pair)
}

func (tapi *TradeAPI) Trade(pair string, orderType string, rate float64, amount float64) (Trade, error) {

	tradeResponse := Trade{}

	orderParams := make(map[string]string, 4)
	orderParams["pair"] = pair
	orderParams["type"] = orderType
	orderParams["rate"] = strconv.FormatFloat(rate, 'f', -1, 64)
	orderParams["amount"] = strconv.FormatFloat(amount, 'f', -1, 64)

	err := tapi.call("Trade", &tradeResponse, orderParams)

	if err == nil {
		return tradeResponse, nil
	} else {
		return tradeResponse, err
	}

}

func (tapi *TradeAPI) TradeAuth(key string, secret string, pair string, orderType string, rate float64, amount float64) (Trade, error) {
	tapi.Auth(key, secret)
	return tapi.Trade(pair, orderType, rate, amount)

}

func (tapi *TradeAPI) OrderInfo(orderID string) (OrderInfo, error) {

	orderInfo := OrderInfo{}

	orderParams := make(map[string]string, 1)
	orderParams["order_id"] = orderID

	err := tapi.call("OrderInfo", &orderInfo, orderParams)
	if err == nil {
		fmt.Println(err)
		return orderInfo, nil
	} else {
		return orderInfo, err
	}

}

func (tapi *TradeAPI) OrderInfoAuth(key string, secret string, orderID string) (OrderInfo, error) {
	tapi.Auth(key, secret)
	return tapi.OrderInfo(orderID)

}

func (tapi *TradeAPI) CancelOrder(orderID string) (CancelOrder, error) {

	cancelReponse := CancelOrder{}

	orderParams := make(map[string]string, 1)
	orderParams["order_id"] = orderID

	err := tapi.call("CancelOrder", &cancelReponse, orderParams)

	if err == nil {
		return cancelReponse, nil
	} else {
		return cancelReponse, err
	}

}

func (tapi *TradeAPI) CancelOrderAuth(key string, secret string, orderID string) (CancelOrder, error) {
	tapi.Auth(key, secret)
	return tapi.CancelOrder(orderID)

}

func (tapi *TradeAPI) TradeHistory(filter HistoryFilter, pair string) (TradeHistory, error) {

	tradeHistory := TradeHistory{}

	historyParams := historyFilterParams(filter)
	if pair != "" {
		historyParams["pair"] = pair
	}

	err := tapi.call("TradeHistory", &tradeHistory, historyParams)

	if err == nil {
		return tradeHistory, nil
	} else {
		return tradeHistory, err
	}

}

func (tapi *TradeAPI) TradeHistoryAuth(key string, secret string, filter HistoryFilter, pair string) (TradeHistory, error) {
	tapi.Auth(key, secret)
	return tapi.TradeHistory(filter, pair)

}

func (tapi *TradeAPI) TransactionHistory(filter HistoryFilter) (TransactionHistory, error) {

	transactionHistory := TransactionHistory{}

	historyParams := historyFilterParams(filter)

	err := tapi.call("TransHistory", &transactionHistory, historyParams)

	if err == nil {
		return transactionHistory, nil
	} else {
		return transactionHistory, err
	}

}

func (tapi *TradeAPI) TransactionHistoryAuth(key string, secret string, filter HistoryFilter) (TransactionHistory, error) {
	tapi.Auth(key, secret)
	return tapi.TransactionHistory(filter)

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

func (tapi *TradeAPI) WithdrawCoin(coinName string, amount float64, address string) (WithdrawCoin, error) {

	response := WithdrawCoin{}

	orderParams := make(map[string]string, 3)
	orderParams["coinName"] = coinName
	orderParams["amount"] = strconv.FormatFloat(amount, 'f', -1, 64)
	orderParams["address"] = address

	err := tapi.call("WithdrawCoin", &response, orderParams)

	if err == nil {
		return response, nil
	} else {
		return response, err
	}

}

func (tapi *TradeAPI) WithdrawCoinAuth(key string, secret string, coinName string, amount float64, address string) (WithdrawCoin, error) {
	tapi.Auth(key, secret)
	return tapi.WithdrawCoin(coinName, amount, address)

}

func (tapi *TradeAPI) CreateCoupon(currency string, amount float64) (CreateCoupon, error) {

	response := CreateCoupon{}

	params := make(map[string]string, 2)
	params["currency"] = currency
	params["amount"] = strconv.FormatFloat(amount, 'f', -1, 64)

	err := tapi.call("CreateCoupon", &response, params)

	if err == nil {
		return response, nil
	} else {
		return response, err
	}

}

func (tapi *TradeAPI) CreateCouponAuth(key string, secret string, currency string, amount float64) (CreateCoupon, error) {
	tapi.Auth(key, secret)
	return tapi.CreateCoupon(currency, amount)
}

func (tapi *TradeAPI) RedeemCoupon(coupon string) (RedeemCoupon, error) {

	response := RedeemCoupon{}

	params := make(map[string]string, 1)
	params["coupon"] = coupon

	err := tapi.call("RedeemCoupon", &response, params)

	if err == nil {
		return response, nil
	} else {
		return response, err
	}

}

func (tapi *TradeAPI) RedeemCouponAuth(key string, secret string, coupon string) (RedeemCoupon, error) {
	tapi.Auth(key, secret)
	return tapi.RedeemCoupon(coupon)
}
