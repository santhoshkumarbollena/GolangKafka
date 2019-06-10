package CommaSeperatedValuesToObjects

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	ModelMember "../../GoLangProjectLayout/Model"
)

type member = ModelMember.Member

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
	ModelMember.Demo()
	data = "1_enrollmentEffectiveDate:2008-03-13,1_enrollmentTerminationDate:2008-04-21,1_offshoreRestrictedIndicator:Y,1_profileIdentifier:1306,1_secureClassIdentifier:9999,2_enrollmentEffectiveDate:2009-08-13,2_enrollmentTerminationDate:2010-12-31,2_offshoreRestrictedIndicator:Y,2_profileIdentifier:1306,2_secureClassIdentifier:9999,3_enrollmentEffectiveDate:2008-04-22,3_enrollmentTerminationDate:2009-08-12,3_offshoreRestrictedIndicator:Y,3_profileIdentifier:1306,3_secureClassIdentifier:9999,4_enrollmentEffectiveDate:2008-01-01,4_enrollmentTerminationDate:2008-03-12,4_offshoreRestrictedIndicator:Y,4_profileIdentifier:1306,4_secureClassIdentifier:9999"
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
		if strings.Compare(objectKeyAndValue[0], "enrollmentEffectiveDate") == 0 {
			mem[i1] = member{EnrollmentEffectiveDate: objectKeyAndValue[1]}
		}
		if strings.Compare(objectKeyAndValue[0], "offshoreRestrictedIndicator") == 0 {
			mem[i1] = member{OffshoreRestrictedIndicator: objectKeyAndValue[1]}
		}
		if strings.Compare(objectKeyAndValue[0], "enrollmentTerminationDate") == 0 {
			mem[i1] = member{EnrollmentTerminationDate: objectKeyAndValue[1]}
		}
		if strings.Compare(objectKeyAndValue[0], "profileIdentifier") == 0 {
			mem[i1] = member{ProfileIdentifier: objectKeyAndValue[1]}
		}
		if strings.Compare(objectKeyAndValue[0], "secureClassIdentifier") == 0 {
			mem[i1] = member{SecureClassIdentifier: objectKeyAndValue[1]}
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
		var resultOffshoreRestrictedIndicator string
		var resultEnrollmentEffectiveDate string
		var resultEnrollmentTerminationDate string
		var resultProfileIdentifier string
		var resultSecureClassIdentifier string
		for i = 0; i < len(members); i++ {
			if members[i][key] != (member{}) {
				if strings.Compare(members[i][key].OffshoreRestrictedIndicator, "") != 0 {
					resultOffshoreRestrictedIndicator = members[i][key].OffshoreRestrictedIndicator
				}
				if strings.Compare(members[i][key].EnrollmentEffectiveDate, "") != 0 {
					resultEnrollmentEffectiveDate = members[i][key].EnrollmentEffectiveDate
				}
				if strings.Compare(members[i][key].EnrollmentTerminationDate, "") != 0 {
					resultEnrollmentTerminationDate = members[i][key].EnrollmentTerminationDate
				}
				if strings.Compare(members[i][key].ProfileIdentifier, "") != 0 {
					resultProfileIdentifier = members[i][key].ProfileIdentifier
				}
				if strings.Compare(members[i][key].SecureClassIdentifier, "") != 0 {
					resultSecureClassIdentifier = members[i][key].SecureClassIdentifier
				}

			}

		}

		ResuletMember := member{OffshoreRestrictedIndicator: resultOffshoreRestrictedIndicator, EnrollmentEffectiveDate: resultEnrollmentEffectiveDate, EnrollmentTerminationDate: resultEnrollmentTerminationDate, ProfileIdentifier: resultProfileIdentifier, SecureClassIdentifier: resultSecureClassIdentifier}

		ResuletMembers = append(ResuletMembers, ResuletMember)

	}
	//fmt.Println(ResuletMembers)
	m := make(map[int]string)
	m[1] = "OffshoreRestrictedIndicator"
	m[2] = "ProfileIdentifier"
	m[3] = "SecureClassIdentifier"
	m[4] = "EnrollmentEffectiveDate"
	m[5] = "EnrollmentTerminationDate"
	for i = 0; i < len(ResuletMembers); i++ {
		fmt.Println(listOfUniqueKeys[i])
		//fmt.Println(ResuletMembers[i])

		s := reflect.ValueOf(&ResuletMembers[i]).Elem()
		typeOfT := s.Type()

		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Printf("%d:  %s = %v\n", i,
				typeOfT.Field(i).Name, f.Interface())
		}

	}
	return ResuletMembers
}
