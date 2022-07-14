package market_data_source

import (
	"fmt"
)

type MarketSimulator struct {
}

func (r *MarketSimulator) GetFxPricing(currencies []string) []FxPriceDetails {
	fmt.Println("GetFxPricing() requested for MarketSimulator")

	// Initialise a price list array that is the length of the currencies list
	var priceList = make([]FxPriceDetails, len(currencies))

	var dummyPrice FxPriceDetails

	dummyPrice.BaseCurrency = "USD"
	dummyPrice.Currency = "undefined"
	dummyPrice.Ask = "0.72894"
	dummyPrice.Bid = "0.72890"
	dummyPrice.Date = "2022-04-19T23:59:59+0000"
	dummyPrice.HighAsk = "0.76038"
	dummyPrice.HighBid = "0.75027"
	dummyPrice.LowAsk = "0.75688"
	dummyPrice.LowBid = "0.75675"
	dummyPrice.Midpoint = "0.75888"

	for i, c := range currencies {
		// fmt.Printf("Replace this with populating the response:  i = %d, c = %s\n", i, c)
		dummyPrice.Currency = c
		priceList[i] = dummyPrice
	}

	return priceList
}
