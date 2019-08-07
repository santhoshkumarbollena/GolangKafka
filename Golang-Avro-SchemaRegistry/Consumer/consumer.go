package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	cluster "github.com/bsm/sarama-cluster"
	"github.com/dangkaka/go-kafka-avro"
)

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8081"}
var topic = "testMessage1"

type Student struct {
	ApplicationLogs string `json:"ApplicationLogs"`
}

func main() {
	consumerCallbacks := kafka.ConsumerCallbacks{
		OnDataReceived: func(msg kafka.Message) {
			fmt.Println(msg.Value)
			fmt.Println(msg.Key)
			fmt.Println("--------------------------------------")
			KEY := msg.Key
			fmt.Println()

			reqBody := []byte(msg.Value)

			// 	// update our global Articles array to include
			// 	// our new Article
			var Student Student

			fmt.Println(string(reqBody))
			json.Unmarshal(reqBody, &Student)

			Student.ApplicationLogs = msg.Value

			fmt.Println(Student)

			fmt.Println("demo")
			fmt.Println(string(reqBody))
			//RequestBodyFor2ndService, _ := json.Marshal(Student)
			//Response2ndService, _ := http.Put("http://localhost:9200/applicationlogs/applicationlog/"+KEY, "application/json", bytes.NewBuffer(RequestBodyFor2ndService))
			Response2ndService, _ := http.Post("http://localhost:9200/applicationlogs/applicationlog/"+KEY, "application/json", bytes.NewBuffer([]byte(reqBody)))
			fmt.Println(Response2ndService)
			//fmt.Println("here4")
			//ResponseBody, _ := ioutil.ReadAll(Response2ndService.Body)
			//fmt.Println()

			//s := string(ResponseBody)
			//fmt.Println("here5")
			//json.Unmarshal([]byte(s), &ResponseBody)

		},
		OnError: func(err error) {
			fmt.Println("Consumer error", err)
		},
		OnNotification: func(notification *cluster.Notification) {
			fmt.Println(notification)
		},
	}

	consumer, err := kafka.NewAvroConsumer(kafkaServers, schemaRegistryServers, topic, "consumer-group", consumerCallbacks)
	if err != nil {
		fmt.Println(err)
	}
	consumer.Consume()
}
