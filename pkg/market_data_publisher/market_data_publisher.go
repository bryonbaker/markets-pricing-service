package market_data_publisher

type IMarketDataPublisher interface {
	Initialise()
	PublishPricingData(string, string)
	Cleanup()
}
