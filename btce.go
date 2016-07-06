// Package btce provides native Go client for interacting with BTC-E Public API v3 and Trading API.
//
// Example usage:
//
// 	package main
//
// 	import (
//		"fmt"
//		btc "github.com/onuryilmaz/go-btce"
// 	)
//
// 	func main() {
//
//		api := btc.API{}
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
package btce

// API allows to use public and trade APIs of BTC-E
type API struct {
	Public PublicAPI
	Trade  TradeAPI
}
