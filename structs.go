package btce

import (
	"encoding/json"
	"time"
)

type Ticker map[string]TickerPair

type TickerPair struct {
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Avg     float64 `json:"avg"`
	Vol     float64 `json:"vol"`
	VolCur  float64 `json:"vol_cur"`
	Last    float64 `json:"last"`
	Buy     float64 `json:"buy"`
	Sell    float64 `json:"sell"`
	Updated int64   `json:"updated"`
}

type Info struct {
	ServerTime int64               `json:"server_time"`
	Pairs      map[string]InfoPair `json:"pairs"`
}

type InfoPair struct {
	DecimalPlaces int     `json:"decimal_places"`
	MinPrice      float64 `json:"min_price"`
	MaxPrice      float64 `json:"max_price"`
	MinAmount     float64 `json:"min_amount"`
	Hidden        int     `json:"hidden"`
	Fee           float64 `json:"fee"`
}

type Depth map[string]DepthPair

type DepthPair struct {
	Asks []DepthItem `json:"asks"`
	Bids []DepthItem `json:"bids"`
}

type DepthItem []float64

type Trade map[string]TradePair
type TradePair []TradeItem
type TradeItem struct {
	Type      string  `json:"type"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	TID       string  `json:"tid"`
	Timestamp int64   `json:"timestamp"`
}

type Response struct {
	Success int             `json:"success"`
	Return  json.RawMessage `json:"return"`
	Error   string          `json:"error"`
}

type AccountInfo struct {
	Funds            map[string]float64 `json:"funds"`
	Rights           Rights             `json:"rights"`
	TransactionCount int64              `json:"transaction_count"`
	OpenOrders       int64              `json:"open_orders"`
	ServerTime       float64            `json:"server_time"`
}

type Rights struct {
	Info     int `json:"info"`
	Trade    int `json:"trade"`
	Withdraw int `json:"withdraw"`
}

type ActiveOrder struct {
	Pair             string  `json:"pair"`
	Type             string  `json:"type"`
	Amount           float64 `json:"amount"`
	Rate             float64 `json:"rate"`
	TimestampCreated int64   `json:"timestamp_created"`
	Status           int     `json:"status"`
}

type ActiveOrders map[string]ActiveOrder

type TradeResponse struct {
	Received float64            `json:"received"`
	Remains  float64            `json:"remains"`
	OrderID  string             `json:"order_id"`
	Funds    map[string]float64 `json:"funds"`
}

type OrderInfoItem struct {
	Pair             string  `json:"pair"`
	Type             string  `json:"type"`
	StartAmount      float64 `json:"start_amount"`
	Amount           float64 `json:"amount"`
	Rate             float64 `json:"rate"`
	TimestampCreated int64   `json:"timestamp_created"`
	Status           int     `json:"status"`
}

type OrderInfo map[string]OrderInfoItem

type CancelOrder struct {
	OrderID int                `json:"order_id"`
	Funds   map[string]float64 `json:"funds"`
}

type HistoryFilter struct {
	From   int
	Count  int
	FromID int
	EndID  int
	Order  string
	Since  time.Time
	End    time.Time
}

type TradeHistoryItem struct {
	Pair        string  `json:"pair"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Rate        float64 `json:"rate"`
	OrderID     int     `json:"order_id"`
	IsYourOrder int     `json:"is_your_order"`
	Timestamp   int64   `json:"timestamp"`
}

type TradeHistory map[string]TradeHistoryItem

type TransactionHistoryItem struct {
	Type        int     `json:"type"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"desc"`
	Status      int     `json:"status"`
	Timestamp   int64   `json:"timestamp"`
}

type TransactionHistory map[string]TransactionHistoryItem

type WithdrawCoin struct {
	TransactionID int                `json:"tId"`
	AmountSent    float64            `json:"amountSent"`
	Funds         map[string]float64 `json:"funds"`
}

type CreateCoupon struct {
	Coupon        string             `json:"coupon"`
	TransactionID int                `json:"transID"`
	Funds         map[string]float64 `json:"funds"`
}

type RedeemCoupon struct {
	CouponAmount   string             `json:"couponAmount"`
	CouponCurrency string             `json:"couponCurrency"`
	TransactionID  int                `json:"transID"`
	Funds          map[string]float64 `json:"funds"`
}
