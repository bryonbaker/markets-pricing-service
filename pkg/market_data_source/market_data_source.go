package market_data_source

type IMarketDataSource interface {
	GetFxPricing(currencies []string) string
}
