package market_data_publisher

import (
	"fmt"
)

type ConsolePublisher struct {
}

func (p *ConsolePublisher) PublishPricingData(key string, data string) {
	fmt.Printf("ConsolePublisher::PublishPricingData()\n")

	fmt.Printf("Key: %s\nData: %s\n", key, data)
}

func (p *ConsolePublisher) Initialise() {
	// Do nothing
}

// Clean up any resources on exit
func (p *ConsolePublisher) Cleanup() {
	// Do nothing
}
