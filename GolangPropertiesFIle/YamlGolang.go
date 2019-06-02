package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var data string

func readData() {
	b, err := ioutil.ReadFile("Properties.yml") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	data = string(b) // convert content to a 'string'

}

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B string
	C string
}

func main() {
	readData()
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//Printing data in Object or Stuctured Format
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//printing data as it present in Properties.yaml file
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//Assigning Properties data into a map
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//Printing data by MArshelling the unmarshelled data that is same as present in input file
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
