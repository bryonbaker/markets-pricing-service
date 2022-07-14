package market_data_source

// The standard structure that market data should be returned in.
type FxPriceDetails struct {
	Currency     string
	BaseCurrency string
	Ask          string
	Bid          string
	Date         string
	HighAsk      string
	HighBid      string
	LowAsk       string
	LowBid       string
	Midpoint     string
}

type IMarketDataSource interface {
	GetFxPricing(currencies []string) []FxPriceDetails
}
