package main

import (
	"fmt"
	"os"
	"strings"

	"os-climate.org/market-pricing/pkg/market_data_publisher"
	"os-climate.org/market-pricing/pkg/market_data_source"
	"os-climate.org/market-pricing/pkg/market_reader"
)

func main() {
	fmt.Println("Initialising...")

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	c := make(chan string) // Channel for passing pricing information
	quit := make(chan int) // Channel for sending quit signals.
	defer close(sigchan)
	defer close(c)
	defer close(quit)

	reader := &market_reader.TimerReader{}
	dataSource := &market_data_source.MarketSimulator{}
	publisher := &market_data_publisher.ConsolePublisher{}
	reader.SetMarketProvider(dataSource)

	currencies := []string{"AUD", "NZD", "EUR", "GBP"} // TODO: Replace this with currencies from a config file.

	// Start the reader thread
	reader.Initialise(c, quit)
	go reader.GetFxPricing(currencies)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			m := <-c // Test the channel to see if the price getter has retrieved a quote
			if m != "" {
				SendToPublisher(publisher, m)
			}
		}
	}
	quit <- 0 // Send a quit signal

	// Wait for clean termination response from the thread.
	for q := <-c; q != "done"; {
		continue
	}
	fmt.Printf("Received clean termination signal from all threads.\n")
	fmt.Printf("Exiting")

	return
}

// Send the key/value to the instantiated Market Data Publisher
func SendToPublisher(publisher *market_data_publisher.ConsolePublisher, priceData string) {
	arr := strings.SplitN(priceData, ",", 2)

	// Check the data is formatted properly
	if len(arr) == 2 {
		publisher.PublishPricingData(arr[0], arr[1])
	} else {
		fmt.Printf("ERROR: Badly formatted data in SendToPublisher. No comma separater: %s", priceData)
	}

}
