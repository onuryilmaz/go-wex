// Package wex provides native Go client for interacting with WEX (former BTC-E) Public API v3 and Trading API.
//
// Example usage:
//
// 	package main
//
// 	import (
//		"fmt"
//		wex "github.com/KopfKrieg/go-wex"
// 	)
//
// 	func main() {
//
//		api := wex.API{}
//
//		ticker, err := api.Public.Ticker([]string{"btc_usd"})
//		if err == nil {
//			fmt.Printf("BTC buy price: %.3f \n", ticker["btc_usd"].Buy)
//			fmt.Printf("BTC sell price: %.3f \n", ticker["btc_usd"].Sell)
//		}
//
//		info, err := api.Trade.GetInfoAuth("API_KEY", "API_SECRET")
//		if err == nil {
//			fmt.Printf("BTC amount: %.3f \n", info.Funds["btc"])
//		}
// 	}
package wex

// API allows to use public and trade APIs of BTC-E
type API struct {
	Public PublicAPI
	Trade  TradeAPI
}
