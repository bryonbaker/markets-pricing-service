package market_data_publisher

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaPublisher struct {
	initialised bool
}

var topic = "FX"
var kafkaProducer *kafka.Producer

// Load the configuration file and establish the connection to the broker.
func (p *KafkaPublisher) Initialise() {
	configFile := "./config/kafka.properties"
	fmt.Printf("Reading config file from: %s\n", configFile)
	conf := ReadConfig(configFile)

	var err error
	kafkaProducer, err = kafka.NewProducer(&conf)

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range kafkaProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	p.initialised = true
}

func (p *KafkaPublisher) PublishPricingData(key string, data string) {
	fmt.Printf("KafkaPublisher::PublishPricingData()\n")

	if !p.initialised {
		fmt.Printf("ERROR: KafkaPublisher in not initialised.")
		os.Exit(-1)
	}

	// fmt.Printf("Key: %s\nData: %s\n", key, data)

	kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(data),
	}, nil)

	// Wait for all messages to be delivered
	kafkaProducer.Flush(15 * 1000)
}

// Close the Kafka handle
func (p *KafkaPublisher) Cleanup() {
	kafkaProducer.Close()
}
