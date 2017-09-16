[![Go Report Card](https://goreportcard.com/badge/github.com/KopfKrieg/go-wex)](https://goreportcard.com/report/github.com/KopfKrieg/go-wex)
[![GoDoc](https://godoc.org/github.com/KopfKrieg/go-wex?status.svg)](https://godoc.org/github.com/KopfKrieg/go-wex)

## WEX (former BTC-E) API Go Client
Native Go client for interacting with [WEX](https://wex.nz/) [Public API v3](https://wex.nz/api/3/docs) and [Trading API](https://wex.nz/tapi/docs).

### Usage

```go
package main

import (
	"fmt"
	wex "github.com/KopfKrieg/go-wex"
)

func main() {

	api := wex.API{}

	ticker, err := api.Public.Ticker([]string{"btc_usd"})
	if err == nil {
		fmt.Printf("BTC buy price: %.3f \n", ticker["btc_usd"].Buy)
		fmt.Printf("BTC sell price: %.3f \n", ticker["btc_usd"].Sell)
	}

	info, err := api.Trade.GetInfoAuth("API_KEY", "API_SECRET")
	if err == nil {
		fmt.Printf("BTC amount: %.3f \n", info.Funds["btc"])
	}
}
```
