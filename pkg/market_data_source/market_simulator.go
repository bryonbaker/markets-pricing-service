package market_data_source

import (
	"fmt"
)

type MarketSimulator struct {
}

func (r *MarketSimulator) GetFxPricing(currencies []string) string {
	fmt.Println("GetFxPricing() requested for MarketSimulator")

	var jsonString = `{ "AUD_USD" : {
		"ask": "0.76894",
		"bid": "0.76881",
		"date": "2022-04-19T23:59:59+0000",
		"high_ask": "0.77038",
		"high_bid": "0.77027",
		"low_ask": "0.76688",
		"low_bid": "0.76675",
		"midpoint": "0.76888"
	} }`

	return jsonString
}
