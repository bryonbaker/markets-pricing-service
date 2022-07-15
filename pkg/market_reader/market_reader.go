package market_reader

import (
	"os-climate.org/market-pricing/pkg/market_data_source"
)

type IMarketReader interface {
	GetFxPricing(c chan string, quit chan int, currencies []string)
	SetMarketProvider(market_data_source.IMarketDataSource)
}
