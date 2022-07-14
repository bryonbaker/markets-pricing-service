package market_data_publisher

type IMarketDataPublisher interface {
	PublishPricingData(FxPrice)
}

// This is the datagram that will be stored in the base class' publisher.
type FxPrice struct {
	Currency     string `json:"currency"`
	BaseCurrency string `json:"base_currency"`
	Ask          string `json:"ask"`
	Bid          string `json:"bid"`
	Date         string `json:"date"`
	HighAsk      string `json:"high_ask"`
	HighBid      string `json:"high_bid"`
	LowAsk       string `json:"low_ask"`
	LowBid       string `json:"low_bid"`
	Midpoint     string `json:"midpoint"`
}
