package market_reader

import (
	"os-climate.org/market-pricing/pkg/market_data_source"
)

type IMarketReader interface {
	GetFxPricing(c chan []market_data_source.FxPriceDetails, quit chan int, currencies []string) []market_data_source.FxPriceDetails
	SetMarketProvider(market_data_source.IMarketDataSource)
}
