package main

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/Shopify/sarama"
	"github.com/linkedin/goavro"
)

var schemaRegistryServers = []string{"http://localhost:8081"}
var (
	codec *goavro.Codec
)
var (
	//./kafka-console-producer.sh --broker-list localhost:9092 --topic sampleTopic
	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
	topic      = kingpin.Flag("topic", "TestingGolangKafkaObjects").Default("TestingGolangKafkaObjects").String()
	maxRetry   = kingpin.Flag("maxRetry", "Retry limit").Default("5").Int()
)

type Member struct {
	Name                string
	Code                string
	Year                string
	EnrollmentStartDate *Date
	EnrollmentEndDate   *Date
}

// Address holds information about an address.
type Date struct {
	Day   int64
	Month int64
	Year  int64
}

func main() {
	kingpin.Parse()
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = *maxRetry
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(*brokerList, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	schema := `{
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

	//Create Schema Once
	codec, err := goavro.NewCodec(string(schema))
	if err != nil {
		panic(err)
	}
	//Sample Data
	Member := &Member{
		Name: "bollena santhosh kumar ",
		Code: "b15cs067",
		Year: "2019",
		EnrollmentStartDate: &Date{
			Day:   2,
			Month: 2,
			Year:  2016,
		},
		EnrollmentEndDate: &Date{
			Day:   2,
			Month: 2,
			Year:  2018,
		},
	}

	//fmt.Printf("user in=%+v\n", Member)
	//fmt.Printf((String)codec)
	///Convert Binary From Native
	fmt.Println(Member)
	//fmt.Println(Member.ToStringMap())
	fmt.Println()
	binary, err := codec.BinaryFromNative(nil, Member.ToStringMap())
	if err != nil {
		panic(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: *topic,
		Value: sarama.StringEncoder(binary),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", *topic, partition, offset)
}

func (u *Member) ToStringMap() map[string]interface{} {
	datumIn := map[string]interface{}{
		"Name": string(u.Name),
		"Year": string(u.Year),
		"Code": string(u.Code),
	}

	if u.EnrollmentStartDate != nil {
		addDatum1 := map[string]interface{}{
			"Day":   int64(u.EnrollmentStartDate.Day),
			"Month": int64(u.EnrollmentStartDate.Month),
			"Year":  int64(u.EnrollmentStartDate.Year),
		}
		if u.EnrollmentEndDate != nil {
			addDatum2 := map[string]interface{}{
				"Day":   int64(u.EnrollmentEndDate.Day),
				"Month": int64(u.EnrollmentEndDate.Month),
				"Year":  int64(u.EnrollmentEndDate.Year),
			}

			//important need namespace and record name
			datumIn["EnrollmentStartDate"] = goavro.Union("my.namespace.com.enrollmentStartDate", addDatum1)
			datumIn["EnrollmentEndDate"] = goavro.Union("my.namespace.com.enrollmentEndDate", addDatum2)

		} else {
			datumIn["EnrollmentStartDate"] = goavro.Union("null", nil)
			datumIn["EnrollmentEndDate"] = goavro.Union("null", nil)
		}
	}
	return datumIn
}
