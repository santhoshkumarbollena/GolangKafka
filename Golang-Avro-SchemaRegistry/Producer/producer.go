package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/dangkaka/go-kafka-avro"
)

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8081"}
var topic = "testObjectsnew"

type member struct {
	Name                string
	EnrollmentStartDate string
	EnrollmentEndDate   string
	Code                string
}

func main() {
	var n int
	schema := `{
		"namespace": "my.namespace.com",
		"type":	"record",
		"name": "value_TestingGolangKafkaObjects",
		"fields": [
			{ "name": "Name", "type": "string"},
			{ "name": "Code", "type": "string"},
			{ "name": "EnrollmentStartDate", "type": "string" }	,
			{ "name": "EnrollmentEndDate", "type": "string" }	
		]
	}`
	producer, err := kafka.NewAvroProducer(kafkaServers, schemaRegistryServers)
	if err != nil {
		fmt.Printf("Could not create avro producer: %s", err)
	}
	flag.IntVar(&n, "n", 1, "number")
	flag.Parse()
	for i := 0; i < n; i++ {
		//fmt.Println(i)
		addMsg(producer, schema)
	}
}

func addMsg(producer *kafka.AvroProducer, schema string) {

	// value := `{
	// 	"Name": "santhosh",
	// 	"enrollmentStartDate":"2019",
	// 	"enrollmentEndDate":"2019",
	// 	"Code":"b15"
	// }`
	a := &member{"santhu bollena its working", "2015", "2019", "b15cs067"}

	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	value := string(out)
	key := a.Code
	err = producer.Add(topic, schema, []byte(key), []byte(value))
	fmt.Println(key)
	fmt.Println(value)
	if err != nil {
		fmt.Printf("Could not add a msg: %s", err)
	}
}
