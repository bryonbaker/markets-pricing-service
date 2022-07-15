package main

import (
	"fmt"
	"os"

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

	//TODO: Move this to a test case. Delete once the readers are working.
	testWriter(publisher)

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
				fmt.Printf("Quote Details:\n%s\n", m)
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

// This function is used to test the publisher. This should get moved to a test case.
func testWriter(publisher *market_data_publisher.ConsolePublisher) {
	// var price market_data_publisher.FxPrice
	// var prices []market_data_publisher.FxPrice

	// price.BaseCurrency = "USD"
	// price.Currency = "AUD"
	// price.Ask = "0.76894"
	// price.Bid = "0.76881"
	// price.Date = "2022-04-19T23:59:59+0000"
	// price.HighAsk = "0.77038"
	// price.HighBid = "0.77027"
	// price.LowAsk = "0.76688"
	// price.LowBid = "0.76675"
	// price.Midpoint = "0.76888"

	// prices = append(prices, price)

	// price.BaseCurrency = "USD"
	// price.Currency = "NZD"
	// price.Ask = "0.72894"
	// price.Bid = "0.72890"
	// price.Date = "2022-04-19T23:59:59+0000"
	// price.HighAsk = "0.76038"
	// price.HighBid = "0.75027"
	// price.LowAsk = "0.75688"
	// price.LowBid = "0.75675"
	// price.Midpoint = "0.75888"

	// prices = append(prices, price)

	publisher.PublishPricingData("Test message")
}
