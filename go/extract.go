package main

import (
	"encoding/binary"
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	// Get filename from os Args
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filepath>\n",
			os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]
	fmt.Println(filename)

	// Get kafka broker host from Env variables
	kafka_broker, env_err := os.LookupEnv("KAFKA")
	if env_err == false {
		fmt.Fprintf(os.Stderr, "No $KAFKA env\n")
		os.Exit(1)
	}
	fmt.Println(kafka_broker)

	// Create Kafka Producer
	kafka_producer, kafka_err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafka_broker})
	if kafka_err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Kafka Producer: %s\n", kafka_err)
		os.Exit(1)
	}
	fmt.Println(kafka_producer)

	// Open input flat files
	input, input_err := os.Open(filename)
	if input_err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Open %s: %s\n", filename, input_err)
		os.Exit(1)
	}
	defer input.Close()

	// Create channel to buffer the message
	done_channel := make(chan bool)
	go func() {
		defer close(done_channel)
		for e := range kafka_producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
				return

			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	} ()

	// Scan input
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "|")
		if s[0] == "IDX" {
			switch s[4] {
				case "2": {
					topic := "datafeed-test-serialize"
					bs := make([]byte,4)
					binary.BigEndian.PutUint32(bs, 1)
					line := "hello go"
					hdr := append([]byte{0}, bs...)
					val := append(hdr, line...)
					// Input message to Kafka
					fmt.Println(hdr)
					fmt.Println(val)
					// fmt.Println(append(hdr))
					fmt.Println(line)
					fmt.Println(topic)
					kafka_producer.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: val}
				}
			}
		}
	}
	if scanner.Err() != nil {
		fmt.Printf("error: %s\n", scanner.Err())
	}

	// Wait for delivery report goroutine to finish
	_ = <-done_channel

	// Close Kafka Producer
	kafka_producer.Close()
}
