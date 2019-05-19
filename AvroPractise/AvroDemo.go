package main
import (
    "fmt"

    "github.com/linkedin/goavro"
)
type date struct{
	day int64
	month int64
	year int64
}
type member struct {
	name string
	enrollmentStartDate date
	enrollmentEndDate date
	code string
	year  int64
	
}
func main(){
	//s := member{name : "santhosh",enrollmentStartDate : date{day :1,month :1,year :2015},enrollmentEndDate : date{day :1,month :1,year :2017},code : "b15cs067",year : 2019}
	codec, err := goavro.NewCodec(`
        {
          "type": "record",
          "name": "LongList",
          "fields" : [
            {"name": "next", "type": ["null", "LongList"], "default": null}
          ]
        }`)
    if err != nil {
        fmt.Println(err)
    }

    // NOTE: May omit fields when using default value
    textual := []byte(`{"next":{"LongList":{}}}`)

    // Convert textual Avro data (in Avro JSON format) to native Go form
    native, _, err := codec.NativeFromTextual(textual)
    if err != nil {
        fmt.Println(err)
    }

    // Convert native Go form to binary Avro data
    binary, err := codec.BinaryFromNative(nil, native)
    if err != nil {
        fmt.Println(err)
    }

    // Convert binary Avro data back to native Go form
    native, _, err = codec.NativeFromBinary(binary)
    if err != nil {
        fmt.Println(err)
    }

    // Convert native Go form to textual Avro data
    textual, err = codec.TextualFromNative(nil, native)
    if err != nil {
        fmt.Println(err)
    }

    // NOTE: Textual encoding will show all fields, even those with values that
    // match their default values
    fmt.Println(string(textual))

}