package btce

import (
	"encoding/json"
	"net/http"
)

type PublicAPI struct{}

const API_URL = "https://btc-e.com/api/3/"

func (api *PublicAPI) Ticker(currency []string) (Ticker, error) {

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

func (api *PublicAPI) Info() (Info, error) {

	url := API_URL + "info"
	r, err := http.Get(url)

	if err == nil {
		data := Info{}
		json.NewDecoder(r.Body).Decode(&data)
		return data, nil
	}

	return Info{}, err

}

func (api *PublicAPI) Depth(currency []string) (Depth, error) {

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

func (api *PublicAPI) Trade(currency []string) (Trade, error) {

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
