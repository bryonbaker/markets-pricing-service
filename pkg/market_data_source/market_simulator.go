package market_data_source

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type MarketSimulator struct {
}

// The  structure that dummy market data should be returned in.
type MockFxProviderResponse struct {
	Currency     string  `json:"currency"`
	BaseCurrency string  `json:"base_currency"`
	Ask          float32 `json:"ask"`
	Bid          float32 `json:"bid"`
	Date         string  `json:"date"`
	HighAsk      string  `json:"high_ask"`
	HighBid      string  `json:"high_bid"`
	LowAsk       string  `json:"low_ask"`
	LowBid       string  `json:"low_bid"`
	Midpoint     string  `json:"midpoint"`
}

type MockFxProviderResponseList []MockFxProviderResponse

var dummyPrice MockFxProviderResponse

func (r *MarketSimulator) GetFxPricing(currencies []string) []FxPriceDetails {
	fmt.Println("GetFxPricing() requested for MarketSimulator")

	var mockResponses []FxPriceDetails

	// Set up the standard dataset for the simulation
	dummyPrice.BaseCurrency = "USD"
	dummyPrice.Currency = "undefined"
	dummyPrice.Ask = 0.72894
	dummyPrice.Bid = 0.72890
	dummyPrice.Date = "2022-04-19T23:59:59+0000"
	dummyPrice.HighAsk = "0.76038"
	dummyPrice.HighBid = "0.75027"
	dummyPrice.LowAsk = "0.75688"
	dummyPrice.LowBid = "0.75675"
	dummyPrice.Midpoint = "0.75888"

	for _, c := range currencies {
		var mockPrice FxPriceDetails

		dummyPrice.Currency = c

		r.simulateMarketFluctuations()

		mockPrice.Fx_key = dummyPrice.BaseCurrency + "_" + dummyPrice.Currency

		// Convert to json
		jsonPrice, _ := json.Marshal(dummyPrice)
		mockPrice.Provider_resp = string(jsonPrice)

		mockResponses = append(mockResponses, mockPrice)
	}

	return mockResponses
}

func (r *MarketSimulator) simulateMarketFluctuations() {
	dummyPrice.Ask = dummyPrice.Ask + (rand.Float32()-0.5)/100
	dummyPrice.Bid = dummyPrice.Ask - 0.00002

	// TODO: Update date
}
