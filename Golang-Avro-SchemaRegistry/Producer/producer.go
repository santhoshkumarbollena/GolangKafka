package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/dangkaka/go-kafka-avro"
	"github.com/gorilla/mux"
	"github.com/linkedin/goavro"
)

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8081"}
var topic = "testMessage1"
var ApplicationLogs12 []string
var RequestIdKey string

type Message struct {
	ApplicationLogs string
}

type Input struct {
	RequestId                     string   `json:"requestId"`
	MemberId                      string   `json:"memberId"`
	MemberIdType                  string   `json:"memberIdType"`
	ReferedToSpecialtyCategory    string   `json:"referedToSpecialtyCategory"`
	ReferedToSpecialityCode       []string `json:"referedToSpecialityCode"`
	ReferedToSpecialityAreaOfBody string   `json:"referedToSpecialityAreaOfBody"`
	ProviderIds                   []string `json:"providerIds"`
	SearchFilterCriteria          string   `json:"searchFilterCriteria"`
	CallingApp                    string   `json:"callingApp"`
	CallingAppType                string   `json:"callingAppType"`
}
type Providers struct {
	NPI     string
	Ranking string
}
type Object struct {
	StatusCode    string
	StatusMessage string
}
type Output struct {
	ResponseId string      `json:"responseId"`
	Providers  []Providers `json:"providers"`

	ResponseStatus Object `json:"responseStatus"`
}
type Logger struct {
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	prefix string     // prefix to write at beginning of each line
	flag   int        // properties
	out    io.Writer  // destination for output
	buf    []byte     // for accumulating text to write
}

func FirstService(w http.ResponseWriter, r *http.Request) {
	// resp, _ := http.Get("http://localhost:10000/GetAllEmployes")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	DateTimeInMilliseconds := time.Now()
	reqBody, _ := ioutil.ReadAll(r.Body)
	OutputFromSecondService := Output{}
	var Input Input
	var Output Output
	json.Unmarshal(reqBody, &Input)

	// f, err := os.OpenFile("logs.txt",  os.O_CREATE | os.O_RDWR, 0666)
	//     if err != nil {
	//         fmt.Printf("error opening file: %v", err)
	//     }

	//     // don't forget to close it
	//     defer f.Close()

	//     // assign it to the standard logger
	//     log.SetOutput(f)

	// 	log.Output(1, "this is an event")

	Log1 := DateTimeInMilliseconds.Format("2006-01-02 15:04:05.0000") + " 1st Service"
	ApplicationLogs12 = append(ApplicationLogs12, Log1)
	//fmt.Println(t.Format("2006-01-02 15:04:05.0000")+" 1st Service")
	//fmt.Println(demo2)
	// demo2,_:=fmt.Println("1st Service")
	// fmt.Println(demo1+demo2)
	// var l *Logger
	// l=log.Print(Input)
	//fmt.Println()
	//fmt.Println(Students)
	//fmt.Println(string(reqBody))
	//fmt.Println("here1")
	RequestIdKey = Input.RequestId
	if Input.RequestId == "" {
		Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: "Error no RequestId | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if Input.MemberId == "" {
		Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: "Error no MemberId | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if Input.MemberIdType == "" {
		Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: " Error no MemberIdType | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if Input.ReferedToSpecialtyCategory == "" {
		Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: " Error no ReferedToSpecialtyCategory | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if len(Input.ProviderIds) == 0 {
		Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: " Error no ProviderIds | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if Input.SearchFilterCriteria == "" {
		Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: " Error no SearchFilterCriteria | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if Input.CallingApp == "" {
		Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: " Error no CallingApp | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	if Input.CallingAppType == "" {
		Output.ResponseStatus = Object{StatusCode: "901", StatusMessage: " Error no CallingAppType | Passed"}
		json.NewEncoder(w).Encode((Output))
		return
	}
	//fmt.Println("here2")
	RequestBodyFor2ndService, _ := json.Marshal(Input)

	Log2 := DateTimeInMilliseconds.Format("2006-01-02 15:04:05.0000") + string(RequestBodyFor2ndService)
	ApplicationLogs12 = append(ApplicationLogs12, Log2)
	//log.Println(string(RequestBodyFor2ndService))
	Response2ndService, _ := http.Post("http://localhost:10000/SecondService", "application/json", bytes.NewBuffer(RequestBodyFor2ndService))

	defer Response2ndService.Body.Close()
	//fmt.Println("here4")
	ResponseBody, _ := ioutil.ReadAll(Response2ndService.Body)
	//fmt.Println()
	Log5 := DateTimeInMilliseconds.Format("2006-01-02 15:04:05.0000") + "Printing Response"
	ApplicationLogs12 = append(ApplicationLogs12, Log5)
	//log.Println("Printing Response")

	Log6 := DateTimeInMilliseconds.Format("2006-01-02 15:04:05.0000") + string(ResponseBody)
	ApplicationLogs12 = append(ApplicationLogs12, Log6)
	//log.Println(string(ResponseBody))
	//log.Printf(string(ResponseBody))

	fmt.Println()
	//fmt.Println(tim)
	//fmt.Println(ApplicationLogs12)
	//fmt.Println(log.Println(string(ResponseBody)))
	s := string(ResponseBody)
	//fmt.Println("here5")
	json.Unmarshal([]byte(s), &OutputFromSecondService)

	json.NewEncoder(w).Encode(OutputFromSecondService)
	//kafka Producer
	var n int
	//schema for message
	schema := `{
		"namespace": "my.namespace.com",
		"type":	"record",
		"name": "value_TestingGolangKafkaObjects",
		"fields": [
			{ "name": "ApplicationLogs", "type": "string"}
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
func SecondService(w http.ResponseWriter, r *http.Request) {
	DateTimeInMilliseconds := time.Now()
	var Output Output
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Input Input
	json.Unmarshal(reqBody, &Input)

	//fmt.Println()
	Log3 := DateTimeInMilliseconds.Format("2006-01-02 15:04:05.0000") + "2nd Service"
	ApplicationLogs12 = append(ApplicationLogs12, Log3)
	//log.Println("2nd Service")
	//fmt.Println()
	//COnverting Input member tpe to string
	out, err := json.Marshal(Input)
	if err != nil {
		panic(err)
	}
	value := string(out)
	Log4 := DateTimeInMilliseconds.Format("2006-01-02 15:04:05.0000") + string(value)
	ApplicationLogs12 = append(ApplicationLogs12, Log4)
	//log.Println(Input)
	//fmt.Println()
	//Setting output values in service 2
	Output.ResponseId = "1231"
	Provider := Providers{"key1", "1"}
	//Object := Object{"Response Code","Response Status"}
	var ProvidersList []Providers
	ProvidersList = append(ProvidersList, Provider)
	Provider2 := Providers{"key2", "2"}
	ProvidersList = append(ProvidersList, Provider2)
	Output.Providers = ProvidersList

	Output.ResponseStatus = Object{StatusCode: "200", StatusMessage: "sucess"}

	json.NewEncoder(w).Encode(Output)
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/FirstService", FirstService).Methods("POST")
	myRouter.HandleFunc("/SecondService", SecondService).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))

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
	ApplicationLogStirng := strings.Join(ApplicationLogs12, ";")
	fmt.Println("string")
	fmt.Println(ApplicationLogStirng)
	a := &Message{ApplicationLogStirng}
	//concat ProfIden and SecureClassIden to generate key
	Key := &Key{RequestIdKey}

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
