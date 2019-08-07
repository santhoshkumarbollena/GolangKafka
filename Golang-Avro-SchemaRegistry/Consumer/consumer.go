package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	cluster "github.com/bsm/sarama-cluster"
	"github.com/dangkaka/go-kafka-avro"
)

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8081"}
var topic = "testMessage1"

func main() {
	consumerCallbacks := kafka.ConsumerCallbacks{
		OnDataReceived: func(msg kafka.Message) {
			fmt.Println(msg.Value)
			fmt.Println(msg.Key)
			fmt.Println("--------------------------------------")
			// KEY := msg.Key
			// fmt.Println()

			// reqBody := []byte(msg.Value)

			// // 	// update our global Articles array to include
			// // 	// our new Article

			// fmt.Println(string(reqBody))

			// fmt.Println("demo")
			// fmt.Println(KEY)
			// fmt.Println(string(reqBody))
			// //RequestBodyFor2ndService, _ := json.Marshal(Student)
			// //Response2ndService, _ := http.Put("http://localhost:9200/applicationlogs/applicationlog/"+KEY, "application/json", bytes.NewBuffer(RequestBodyFor2ndService))
			// Response2ndService, _ := http.Post("http://localhost:9200/applicationlogs/applicationlog/"+KEY, "application/json", bytes.NewBuffer(reqBody))
			// fmt.Println(Response2ndService)
			// //fmt.Println("here4")
			// //ResponseBody, _ := ioutil.ReadAll(Response2ndService.Body)
			// //fmt.Println()

			// //s := string(ResponseBody)
			// //fmt.Println("here5")
			// //json.Unmarshal([]byte(s), &ResponseBody)

			//vars := mux.Vars(r)

			//strings.Join(kk, msg.Value)
			KEY := "pra"
			// 	fmt.Println()
			//   var Students []Student
			// 	reqBody, _ := ioutil.ReadAll(r.Body)
			// 	fmt.Println(string(reqBody))
			// 	var Student Student
			// 	json.Unmarshal(reqBody, &Student)
			// // 	// update our global Articles array to include
			// // 	// our new Article
			// 	Students = append(Students, Student)
			reqBody := []byte(msg.Value)
			//fmt.Println(Student)
			//RequestBodyFor2ndService, _ := json.Marshal(Student)
			//Response2ndService, _ := http.Put("http://localhost:9200/applicationlogs/applicationlog/"+KEY, "application/json", bytes.NewBuffer(RequestBodyFor2ndService))
			fmt.Println("demo")
			fmt.Println(KEY)
			fmt.Println(string(reqBody))
			Response2ndService, _ := http.Post("http://localhost:9200/applicationlogs/applicationlog/"+KEY, "application/json", bytes.NewBuffer(reqBody))
			fmt.Println()
			fmt.Println(Response2ndService)
			defer Response2ndService.Body.Close()
			//fmt.Println("here4")
			ResponseBody, _ := ioutil.ReadAll(Response2ndService.Body)
			//fmt.Println()

			s := string(ResponseBody)
			//fmt.Println("here5")
			json.Unmarshal([]byte(s), &ResponseBody)

			//json.NewEncoder(w).Encode(s)
			fmt.Println("after")

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
