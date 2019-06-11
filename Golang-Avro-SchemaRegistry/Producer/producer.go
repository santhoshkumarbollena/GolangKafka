package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/dangkaka/go-kafka-avro"
	"github.com/linkedin/goavro"
)

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8081"}
var topic = "testMember"

type Member struct {
	OffshoreRestrictedIndicator string
	ProfileIdentifier           string
	SecureClassIdentifier       string
	EnrollmentEffectiveDate     string
	EnrollmentTerminationDate   string
}

func main() {
	var n int
	//schema for message
	schema := `{
		"namespace": "my.namespace.com",
		"type":	"record",
		"name": "value_TestingGolangKafkaObjects",
		"fields": [
			{ "name": "OffshoreRestrictedIndicator", "type": "string"},
			{ "name": "ProfileIdentifier", "type": "string"},
			{ "name": "SecureClassIdentifier", "type": "string" }	,
			{ "name": "EnrollmentEffectiveDate", "type": "string" }	,
			{ "name": "EnrollmentTerminationDate", "type": "string" }
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

var (
	codec *goavro.Codec
)

type Key struct {
	KeyFeild string
}

func addMsg(producer *kafka.AvroProducer, schema string) {

	// value := `{
	// 	"Name": "santhosh",
	// 	"enrollmentStartDate":"2019",
	// 	"enrollmentEndDate":"2019",
	// 	"Code":"b15"
	// }`

	//schema for key
	schemakey := `{
		"namespace": "my.namespace.com",
		"type":	"record",
		"name": "value_TestingGolangKafkaObjectsKey",
		"fields": [
			{ "name": "KeyFeild", "type": "string"}
		]
	}`
	//Assigning schema to Codec
	codec, err := goavro.NewCodec(string(schemakey))
	if err != nil {
		panic(err)
	}
	//Sample Data
	a := &Member{"pranay Kumar bollena ", "9999", "Y", "2015", "2019"}
	//concat ProfIden and SecureClassIden to generate key
	Key := &Key{a.ProfileIdentifier + a.SecureClassIdentifier}

	//fmt.Printf("user in=%+v\n", Key)
	//fmt.Printf((String)codec)
	///Convert Binary From Native
	//fmt.Println()
	//fmt.Println(Key.ToStringMap())
	//fmt.Println()

	//COverting key to binary format
	binary, err := codec.BinaryFromNative(nil, Key.ToStringMap())
	if err != nil {
		panic(err)
	}
	//fmt.Println()
	//fmt.Println(binary)
	//s := string(binary)

	//Converting member type to string because string to byte can be converted easyly
	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	value := string(out)
	//sending message from producer
	err = producer.Add(topic, schema, []byte(binary), []byte(value))
	//fmt.Println(key)
	//fmt.Println(value)
	if err != nil {
		fmt.Printf("Could not add a msg: %s", err)
	}
}

//Mapping KeyFeild to Avro
func (u *Key) ToStringMap() map[string]interface{} {
	datumIn := map[string]interface{}{
		"KeyFeild": string(u.KeyFeild),
	}

	return datumIn
}
