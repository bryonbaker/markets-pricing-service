package market_data_publisher

import (
	"fmt"
)

type ConsolePublisher struct {
}

func (p *ConsolePublisher) PublishPricingData(data string) {
	fmt.Printf("ConsolePublisher::PublishPricingData()\n")

	fmt.Printf("%s\n", data)
	// for _, v := range prices {
	// 	// Marshall the structure into JSON
	// 	var key string = v.BaseCurrency + "_" + v.Currency
	// 	value, _ := json.Marshal(v)

	// 	fmt.Printf("Key: %s\n", key)
	// 	fmt.Printf("Value: %s\n", value)
	// }
}
