package main

import (
	"fmt"
	"strconv"
	"strings"
)

type member struct {
	name                string
	enrollmentStartDate string
	enrollmentEndDate   string
	code                string
}

var keys []int

var members []map[int]member

func main() {
	result := make(map[int]member)
	//fmt.Println("demo")
	data := "1_enrollmentStartDate:2019-2-23,1_enrollmentEndDate:2019-8-25,1_name:santhosh,2_name:pranay"
	//san_ := member{name: "", enrollmentStartDate: date{day: 0, month: 0, year: 0}, enrollmentEndDate: date{day: 0, month: 0, year: 0}, code: "", year: 0}
	s := strings.Split(data, ",")
	//fmt.Println(s)
	mem := make(map[int]member)
	for _, v := range s {
		object := strings.Split(v, "_")
		ObjectAndValue := map[string]string{object[0]: object[1]} //0-key of the object //1-value of the object

		//fmt.Println(ObjectAndValue)
		objectKeyAndValue := strings.Split(ObjectAndValue[object[0]], ":") //object[0] to get the key of the object
		//fmt.Print(object[0])
		i1, _ := strconv.Atoi(object[0])
		keys = append(keys, i1) //Storing all keys in a slice for furthur reference
		//fmt.Print(" ")
		//fmt.Println(objectKeyAndValue)
		mem = map[int]member{i1: member{}} //assigning empty member for respective key
		members = append(members, mem)     //each member is added to members slice
		// fmt.Print(objectKeyAndValue[0])
		// fmt.Print("    ")
		// fmt.Println("1_enrollmentStartDate")
		//Assigning values to respctve feilds
		if strings.Compare(objectKeyAndValue[0], "enrollmentStartDate") == 0 {
			mem[i1] = member{enrollmentStartDate: objectKeyAndValue[1]}
		}
		if strings.Compare(objectKeyAndValue[0], "name") == 0 {
			mem[i1] = member{name: objectKeyAndValue[1]}
		}
		if strings.Compare(objectKeyAndValue[0], "enrollmentEndDate") == 0 {
			mem[i1] = member{enrollmentEndDate: objectKeyAndValue[1]}
		}
		if strings.Compare(objectKeyAndValue[0], "code") == 0 {
			mem[i1] = member{code: objectKeyAndValue[1]}
		}
		//fmt.Println(mem)
	}
	//fmt.Println(members)
	//fmt.Println(keys)
	var i int
	var j int
	var ResuletMembers []map[int]member

	keysm := make(map[int]bool)
	list := []int{} //removing duplicates in keys
	for _, entry := range keys {
		if _, value := keysm[entry]; !value {
			keysm[entry] = true
			list = append(list, entry)
		}
	}
	//list consists of non repeted keys
	fmt.Println(list)
	for j = 0; j < len(list); j++ {
		key := list[j]
		var resultname string
		var resultEnrolmentStartDate string
		var resultEnrolmentEndDate string
		var resultCode string
		for i = 0; i < len(members); i++ {
			if members[i][key] != (member{}) {
				fmt.Println(members[i][key])
				if strings.Compare(members[i][key].name, "") != 0 {
					resultname = members[i][key].name
				}
				if strings.Compare(members[i][key].enrollmentStartDate, "") != 0 {
					resultEnrolmentStartDate = members[i][key].enrollmentStartDate
				}
				if strings.Compare(members[i][key].enrollmentEndDate, "") != 0 {
					resultEnrolmentEndDate = members[i][key].enrollmentEndDate
				}
				if strings.Compare(members[i][key].code, "") != 0 {
					resultCode = members[i][key].code
				}

			}

		}

		ResuletMember := member{name: resultname, enrollmentStartDate: resultEnrolmentStartDate, enrollmentEndDate: resultEnrolmentEndDate, code: resultCode}
		result = map[int]member{key: ResuletMember}
		ResuletMembers = append(ResuletMembers, result)

	}
	fmt.Println(ResuletMembers)

	// }
	// result := map[int]member{i1: member{name: members[i][i1].name, enrollmentStartDate: members[i][i1].enrollmentStartDate, enrollmentEndDate: members[i][i1].enrollmentEndDate, code: members[i][i1].code}}
	// fmt.Println(result)

}
