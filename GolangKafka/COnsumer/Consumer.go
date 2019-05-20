package main

import (
	"fmt"
	"os"
	"os/signal"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"github.com/linkedin/goavro"
	"github.com/Shopify/sarama"
)
var (
	codec *goavro.Codec
)
var (
	//./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic sampleTopic --from-beginning
	brokerList        = kingpin.Flag("brokerList", "").Default("localhost:9092").Strings()
	topic             = kingpin.Flag("topic", "TestingGolangKafkaObjects").Default("TestingGolangKafkaObjects").String()
	partition         = kingpin.Flag("partition", "Partition number").Default("0").String()
	offsetType        = kingpin.Flag("offsetType", "Offset Type (OffsetNewest | OffsetOldest)").Default("-1").Int()
	messageCountStart = kingpin.Flag("messageCountStart", "Message counter start from:").Int()
)
type Member struct {
	Name string
	Code  string
	Year    string
	EnrollmentStartDate *Date
	EnrollmentEndDate *Date
}

// Address holds information about an address.
type Date struct {
	Day int64
	Month int64
	Year    int64
}

func main() {
	kingpin.Parse()
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := *brokerList
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()
	consumer, err := master.ConsumePartition(*topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				*messageCountStart++
				schema:=`{
					"namespace": "my.namespace.com",
					"type":	"record",
					"name": "indentity",
					"fields": [
						{ "name": "Name", "type": "string"},
						{ "name": "Code", "type": "string"},
						{ "name": "Year", "type": "string" },
						{ "name": "EnrollmentStartDate", "type": ["null",{
							"namespace": "my.namespace.com",
							"type":	"record",
							"name": "enrollmentStartDate",
							"fields": [
								{ "name": "Day", "type": "long" },
								{ "name": "Month", "type": "long" },
								{ "name": "Year", "type": "long" }
							]
						}],"default":null},
						{ "name": "EnrollmentEndDate", "type": ["null",{
							"namespace": "my.namespace.com",
							"type":	"record",
							"name": "enrollmentEndDate",
							"fields": [
								{ "name": "Day", "type": "long" },
								{ "name": "Month", "type": "long" },
								{ "name": "Year", "type": "long" }
							]
						}],"default":null}
					]
				}`
				codec, err := goavro.NewCodec(string(schema))
					if err != nil {
					panic(err)
				}
				//fmt.Println(msg.Value)
				native, _, err := codec.NativeFromBinary(msg.Value)
				if err != nil {
					panic(err)
				}
				userOut := StringMapToMember(native.(map[string]interface{}))
				//fmt.Printf("user out=%+v\n", userOut)
			//	fmt.Println()
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
				//fmt.Println()
				//fmt.Println(userOut)
				fmt.Println(userOut.EnrollmentEndDate)
				fmt.Println(userOut.EnrollmentStartDate)
				// Mem:=&Member{userOut}
				// fmt.Println(Mem)
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	fmt.Println("Processed", *messageCountStart, "messages")
}
func  StringMapToMember(data map[string]interface{}) *Member {
	ind := &Member{}
	for k, v := range data {
		switch k {
		case "Name":
			if value, ok := v.(string); ok {
				ind.Name = value
			}
		case "Year":
			if value, ok := v.(string); ok {
				ind.Year = value
			}
		case "Code":
			if value, ok := v.(string); ok {
				ind.Code = value
			}
		case "EnrollmentStartDate":
			if vmap, ok := v.(map[string]interface{}); ok {
				//important need namespace and record name
				if cookieSMap, ok := vmap["my.namespace.com.enrollmentStartDate"].(map[string]interface{}); ok {
					add := &Date{}
					for k, v := range cookieSMap {
						switch k {
						case "Day":
							if value, ok := v.(int64); ok {
								add.Day = value
							}
						case "Month":
							if value, ok := v.(int64); ok {
								add.Month = value
							}
						case "Year":
							if value, ok := v.(int64); ok {
								add.Year = value
							}
						}
					}
					ind.EnrollmentStartDate = add
				}
		}
	case "EnrollmentEndDate":
		if vmap, ok := v.(map[string]interface{}); ok {
			//important need namespace and record name
			if cookieSMap, ok := vmap["my.namespace.com.enrollmentEndDate"].(map[string]interface{}); ok {
				add := &Date{}
				for k, v := range cookieSMap {
					switch k {
					case "Day":
						if value, ok := v.(int64); ok {
							add.Day = value
						}
					case "Month":
						if value, ok := v.(int64); ok {
							add.Month = value
						}
					case "Year":
						if value, ok := v.(int64); ok {
							add.Year = value
						}
					}
				}
				ind.EnrollmentEndDate = add
			}
		}

	}
}
	return ind
}

