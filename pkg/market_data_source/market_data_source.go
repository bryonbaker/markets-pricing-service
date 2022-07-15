package market_data_source

// The standard structure that market data should be returned in.
type FxPriceDetails struct {
	Fx_key        string
	Provider_resp string
}

type IMarketDataSource interface {
	GetFxPricing(currencies []string) []FxPriceDetails
}
