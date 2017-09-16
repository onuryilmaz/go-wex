[![Build Status](https://travis-ci.org/onuryilmaz/go-wex.svg?branch=master)](https://travis-ci.org/onuryilmaz/go-wex)
[![Go Report Card](https://goreportcard.com/badge/github.com/onuryilmaz/go-wex)](https://goreportcard.com/report/github.com/onuryilmaz/go-wex)
[![GoDoc](https://godoc.org/github.com/onuryilmaz/go-wex?status.svg)](https://godoc.org/github.com/onuryilmaz/go-wex)
[![Coverage Status](https://coveralls.io/repos/github/onuryilmaz/go-wex/badge.svg?branch=master)](https://coveralls.io/github/onuryilmaz/go-wex?branch=master)

## WEX (former BTC-E) API Go Client
Native Go client for interacting with [WEX](https://wex.nz/) [Public API v3](https://wex.nz/api/3/docs) and [Trading API](https://wex.nz/tapi/docs).

### Usage

```go
package main

import (
	"fmt"
	wex "github.com/onuryilmaz/go-wex"
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
