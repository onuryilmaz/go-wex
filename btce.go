package btce

// API allows to use public and trade APIs of BTC-E
type API struct {
	public PublicAPI
	trade  TradeAPI
}
