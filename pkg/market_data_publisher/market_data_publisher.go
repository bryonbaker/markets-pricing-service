package market_data_publisher

type IMarketDataPublisher interface {
	PublishPricingData(string, string)
}
