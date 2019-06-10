package main

import (
	"encoding/json"
	"fmt"
)

type member struct {
	Name                string
	EnrollmentStartDate string
	EnrollmentEndDate   string
	Code                string
}

func main() {
	a := &member{"santhu", "2015", "2019", "b15cs067"}

	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	value := string(out)
	fmt.Println(value)
}
