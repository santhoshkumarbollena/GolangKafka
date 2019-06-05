package CommaSeperatedValuesToObjects

import (
	"strconv"
	"strings"

	t "../../GoLangProjectLayout/HiveGoConnection"
	m "../../GoLangProjectLayout/Model"
)

type member = m.Member

var data string

var keys []int
var i int
var j int

//ResuletMembers consits of result members
var ResuletMembers []member

var members []map[int]member
var result map[int]member

//CSVDataToMemberObjects
func CSVDataToMemberObjects() {

	data = t.GetDataFromHive()
}
func GetMembersObject() []member {
	CSVDataToMemberObjects()
	//fmt.Println(data)
	SettingCsvDataToRespectiveObject(data)
	ResuletMembersRes := GetUnique()
	//fmt.Println(ResuletMembersRes)
	return ResuletMembersRes
}

//SettingCsvDataToRespectiveObject to set data to respective Object
func SettingCsvDataToRespectiveObject(data string) {

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
			mem[i1] = member{EnrollmentStartDate: objectKeyAndValue[1]}
		}
		if strings.Compare(objectKeyAndValue[0], "name") == 0 {
			mem[i1] = member{Name: objectKeyAndValue[1]}
		}
		if strings.Compare(objectKeyAndValue[0], "enrollmentEndDate") == 0 {
			mem[i1] = member{EnrollmentEndDate: objectKeyAndValue[1]}
		}
		if strings.Compare(objectKeyAndValue[0], "code") == 0 {
			mem[i1] = member{Code: objectKeyAndValue[1]}
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
func GetUnique() []member {

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
				if strings.Compare(members[i][key].Name, "") != 0 {
					resultname = members[i][key].Name
				}
				if strings.Compare(members[i][key].EnrollmentStartDate, "") != 0 {
					resultEnrolmentStartDate = members[i][key].EnrollmentStartDate
				}
				if strings.Compare(members[i][key].EnrollmentEndDate, "") != 0 {
					resultEnrolmentEndDate = members[i][key].EnrollmentEndDate
				}
				if strings.Compare(members[i][key].Code, "") != 0 {
					resultCode = members[i][key].Code
				}

			}

		}

		ResuletMember := member{Name: resultname, EnrollmentStartDate: resultEnrolmentStartDate, EnrollmentEndDate: resultEnrolmentEndDate, Code: resultCode}

		ResuletMembers = append(ResuletMembers, ResuletMember)

	}
	return ResuletMembers
}
