[![Build Status](https://travis-ci.org/onuryilmaz/go-btce.svg?branch=master)](https://travis-ci.org/onuryilmaz/go-btce)
[![Go Report Card](https://goreportcard.com/badge/github.com/onuryilmaz/go-btce)](https://goreportcard.com/report/github.com/onuryilmaz/go-btce)
[![GoDoc](https://godoc.org/github.com/onuryilmaz/go-btce?status.svg)](https://godoc.org/github.com/onuryilmaz/go-btce)
[![Coverage Status](https://coveralls.io/repos/github/onuryilmaz/go-btce/badge.svg?branch=master)](https://coveralls.io/github/onuryilmaz/go-btce?branch=master)

## BTC-E API Go Client
Native Go client for interacting with [BTC-E](https://btc-e.com/) [Public API v3](https://btc-e.com/api/3/docs) and [Trading API](https://btc-e.com/tapi/docs).

### Usage

```go
package main

import (
	"fmt"
	btc "github.com/onuryilmaz/go-btce"
)

func main() {

	api := btc.API{}

	ticker, err := api.Public.Ticker([]string{"btc_usd"})
	if err == nil {
		fmt.Printf("BTC buy price: %.3f \n", ticker["btc_usd"].Buy)
		fmt.Printf("BTC sell price: %.3f \n", ticker["btc_usd"].Sell)
	}

	info, err := api.Trade.GetInfoAuth("API_KEY", "API_SECRET")
	if err == nil {
		fmt.Printf("BTC amount: %.3f \n", info.Funds["btc_usd"])
	}
}
```
