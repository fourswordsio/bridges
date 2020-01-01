package main

import (
	"github.com/linkpoolio/bridges"
	"net/http"
)

// CryptoCompare is the most basic Bridge implementation, as it only calls the api:
// https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD,JPY,EUR
type CryptoCompare struct{}

// Run is the bridge.Bridge Run implementation that returns the price response
func (cc *CryptoCompare) Run(h *bridges.Helper) (interface{}, error) {
	r := make(map[string]interface{})
	err := h.HTTPCall(
		http.MethodGet,
		"https://fourswords.io/wp-json/wc/v3/orders/?consumer_key=ck_ff6333b7ecd26852dd731e93c26cef6bebacba3f&consumer_secret=cs_e63182db4325259be0b5e22de81b5a42a4227482",
		&r,
	)
	return r, err
}

// Opts is the bridge.Bridge implementation
func (cc *CryptoCompare) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   "CryptoCompare",
		Lambda: true,
	}
}

func main() {
	bridges.NewServer(&CryptoCompare{}).Start(8080)
}
