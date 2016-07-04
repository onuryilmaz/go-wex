package btce

import (
	"encoding/json"
	"net/http"
)

const API_URL = "https://btc-e.com/api/3/"

func GetTicker(currency []string) (Ticker, error) {

	url := API_URL + "ticker/"
	for _, c := range currency {
		url = url + c + "-"
	}
	r, err := http.Get(url)

	if err == nil {
		data := make(Ticker, len(currency))
		json.NewDecoder(r.Body).Decode(&data)
		return data, nil
	}

	return nil, err

}

func GetInfo() (Info, error) {

	url := API_URL + "info"
	r, err := http.Get(url)

	if err == nil {
		data := Info{}
		json.NewDecoder(r.Body).Decode(&data)
		return data, nil
	}

	return Info{}, err

}

func GetDepth(currency []string) (Depth, error) {

	url := API_URL + "depth/"
	for _, c := range currency {
		url = url + c + "-"
	}
	r, err := http.Get(url)

	if err == nil {
		data := make(Depth, len(currency))
		json.NewDecoder(r.Body).Decode(&data)
		return data, nil
	}

	return nil, err

}

func GetTrade(currency []string) (Trade, error) {

	url := API_URL + "trades/"
	for _, c := range currency {
		url = url + c + "-"
	}
	r, err := http.Get(url)

	if err == nil {
		data := make(Trade, len(currency))
		json.NewDecoder(r.Body).Decode(&data)
		return data, nil
	}

	return nil, err

}
