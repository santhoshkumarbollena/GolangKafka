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
var i int
var j int

//ResuletMembers consits of result members
var ResuletMembers []map[int]member

var members []map[int]member
var result map[int]member

func main() {
	SettingCsvDataToRespectiveObject()
	ResuletMembersRes := GetUnique()
	fmt.Println(ResuletMembersRes)

}

//SettingCsvDataToRespectiveObject to set data to respective Object
func SettingCsvDataToRespectiveObject() {
	data := "1_enrollmentStartDate:2019-2-23,1_enrollmentEndDate:2019-8-25,1_name:santhosh,2_name:pranay,1_name:demo,2_name:p"
	CommaSeperatedData := strings.Split(data, ",")
	mem := make(map[int]member)

	for _, v := range CommaSeperatedData {
		object := strings.Split(v, "_")
		ObjectAndValue := map[string]string{object[0]: object[1]} //0-key of the object //1-value of the object

		objectKeyAndValue := strings.Split(ObjectAndValue[object[0]], ":") //object[0] to get the key of the object

		i1, _ := strconv.Atoi(object[0])
		keys = append(keys, i1) //Storing all keys in a slice for furthur reference

		mem = map[int]member{i1: member{}} //assigning empty member for respective key
		members = append(members, mem)     //each member is added to members slice

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

	}
}

//GetUniqueKeys is a method to get unique keys for the given data
func GetUniqueKeys(keys []int) []int {
	keysm := make(map[int]bool)
	list := []int{} //removing duplicates in keys
	for _, entry := range keys {
		if _, value := keysm[entry]; !value {
			keysm[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//GetUnique to get unique
func GetUnique() []map[int]member {

	listOfUniqueKeys := GetUniqueKeys(keys)
	//listOfUniqueKeys consists of non repeted keys
	for j = 0; j < len(listOfUniqueKeys); j++ {
		key := listOfUniqueKeys[j]
		var resultname string
		var resultEnrolmentStartDate string
		var resultEnrolmentEndDate string
		var resultCode string
		for i = 0; i < len(members); i++ {
			if members[i][key] != (member{}) {
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
	return ResuletMembers
}
