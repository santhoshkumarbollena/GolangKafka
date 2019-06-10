package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var data string

func readData() {
	b, err := ioutil.ReadFile("../GoLangProjectLayout/Properties.yml") // just pass the file name
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

	m := make(map[string]string)

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//Assigning Properties data into a map
	fmt.Printf("--- m:\n%v\n\n", m)
	fmt.Println(m["c"])

}
