package market_reader

import (
	"fmt"
	"time"

	"os-climate.org/market-pricing/pkg/market_data_source"
)

// Refined abstraction of the base class (Abstraction from GoF)
type TimerReader struct {
	marketProvider market_data_source.IMarketDataSource
	commsChannel   chan string
	quitChannel    chan int
}

// Used to initialise the inter-process communication channels.
// This may not be required if the esign calls for all GetFxPricing to be used as a go routine.
// In which case the channel initialisers move into the base class.
func (r *TimerReader) Initialise(c chan string, quit chan int) {
	r.commsChannel = c
	r.quitChannel = quit
}

// Initialise the concrete implementation
func (r *TimerReader) SetMarketProvider(mds market_data_source.IMarketDataSource) {
	r.marketProvider = mds
}

// A Go routine that retrieves pricing on scheduled intervals and puts the
// result on a channel for the main thread to pick up.
func (r *TimerReader) GetFxPricing(currencies []string) {
	fmt.Println("GetFxPricing() request for TimerReader")

	if r.commsChannel == nil || r.quitChannel == nil {
		fmt.Println("ERROR: TimeReader::GetFxPricing(): Channels not initialised.")
	} else if r.marketProvider == nil {
		fmt.Println("ERROR: TimeReader::GetFxPricing(): MarketProvider not initialised.")
	} else {
		ticker := time.NewTicker(5 * time.Second)
		for _ = range ticker.C {
			resp := r.marketProvider.GetFxPricing(currencies)

			select {
			case r.commsChannel <- resp: // Send the pricing info to the main loop via the pricing channel.
				continue
			case <-r.quitChannel: // Check if a quit signal has been received. If so, tell the main loop that all thread-termination steps are done..
				fmt.Printf("Received QUIT signal.\n")
				r.commsChannel <- "done"
				return
			}
		}
		fmt.Printf("ERROR: quoteGetter() exiting the thread incorrectly")
	}
}
