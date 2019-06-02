package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/dangkaka/go-kafka-avro"
)

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8081"}
var topic = "testObjects"

func main() {
	var n int
	schema := `{
		"namespace": "my.namespace.com",
		"type":	"record",
		"name": "value_TestingGolangKafkaObjects",
		"fields": [
			{ "name": "Name", "type": "string"},
			{ "name": "Code", "type": "string"},
			{ "name": "enrollmentStartDate", "type": "string" }	,
			{ "name": "enrollmentEndDate", "type": "string" }	
		]
	}`
	producer, err := kafka.NewAvroProducer(kafkaServers, schemaRegistryServers)
	if err != nil {
		fmt.Printf("Could not create avro producer: %s", err)
	}
	flag.IntVar(&n, "n", 1, "number")
	flag.Parse()
	for i := 0; i < n; i++ {
		fmt.Println(i)
		addMsg(producer, schema)
	}
}

func addMsg(producer *kafka.AvroProducer, schema string) {

	value := `{
		"Name": "santhosh",
		"enrollmentStartDate":"2019",
		"enrollmentEndDate":"2019",
		"Code":"b15"
	}`
	key := time.Now().String()
	err := producer.Add(topic, schema, []byte(key), []byte(value))
	fmt.Println(key)
	if err != nil {
		fmt.Printf("Could not add a msg: %s", err)
	}
}
