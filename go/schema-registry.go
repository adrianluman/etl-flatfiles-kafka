package main

import (
	// "encoding/binary"
	// "bufio"
	"fmt"
	// "os"
	// "strings"
	// "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/elodina/go-kafka-avro"
	// "github.com/datamountaineer/schema-registry"
)

func main() {

	// client, client_err := schemaregistry.NewClient(schemaregistry.DefaultUrl)
	// if client_err != nil {
	// 	fmt.Fprintf(os.Stderr, "Schema Registry client error: %s", client_err)
	// }
	// subjects, subjects_err := client.Subjects()
	// if subjects_err != nil {
	// 	fmt.Fprintf(os.Stderr, "Schema Registry get subjects error: %s", client_err)
	// }
	// fmt.Println(subjects)

	client := avro.NewCachedSchemaRegistryClient("http://localhost:8081")
	fmt.Println(client)
}
