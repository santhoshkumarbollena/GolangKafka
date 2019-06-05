package main

import (
	"fmt"
	"math"
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
var ResuletMembers []member

var members []map[int]member
var result map[int]member
var resultDateBwStartDateEndDate []member
var resultStartDateBelowGivenDate []member

func main() {
	data := "1_enrollmentStartDate:2019-2-23,1_enrollmentEndDate:2019-8-25,1_name:santhosh,2_name:pranay,1_name:demo,2_enrollmentStartDate:2019-2-23,2_enrollmentEndDate:2019-8-25"
	SettingCsvDataToRespectiveObject(data)
	ResuletMembersRes := GetUnique()
	fmt.Println(ResuletMembersRes)
	//Integrating Two Tasks
	fmt.Println("Members whose enrollment start date and end date in between given date")
	membersindate := MembersInGivenDate(ResuletMembersRes)
	fmt.Println(membersindate)
	//"Members whose enrollment start date is prior to the given date"
	membersStartDateGreaterThanDate, f := MembersWithStartDate(ResuletMembersRes)
	//the following boolen tells us that is there any start date present below the given date")
	fmt.Println(f)
	if f {
		fmt.Println(membersStartDateGreaterThanDate)
	}
	fmt.Println("Members whose enrollment end date is near to given date")
	membersNearEnddate := MembersnearEndDate(ResuletMembersRes)
	fmt.Println(membersNearEnddate)

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

		ResuletMembers = append(ResuletMembers, ResuletMember)

	}
	return ResuletMembers
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

//MembersnearEndDate to get members near End Date
func MembersnearEndDate(mems []member) member {
	fmt.Println("Enter date in year-month-day to find near end date: ")
	var inputy string
	fmt.Scanln(&inputy)
	s := strings.Split(inputy, "-")
	day, err := strconv.ParseInt(s[2], 10, 64)
	month, err := strconv.ParseInt(s[1], 10, 64)
	year, err := strconv.ParseInt(s[0], 10, 64)
	if err == nil {
	}
	var i int
	var miny, minm, mind int64 = math.MaxInt64, math.MaxInt64, math.MaxInt64
	var ty, tm, td int64

	var near member
	for i = 0; i < len(mems); i++ {
		if strings.Compare(mems[i].enrollmentEndDate, "") == 0 {
			break
		}
		enrollmentDateSplit := strings.Split(mems[i].enrollmentEndDate, "-")
		y, _ := strconv.ParseInt(enrollmentDateSplit[0], 10, 64)
		m, _ := strconv.ParseInt(enrollmentDateSplit[1], 10, 64)
		d, _ := strconv.ParseInt(enrollmentDateSplit[2], 10, 64)
		ty = abs(year - y) //1000-2019
		tm = abs(month - m)
		td = abs(day - d)
		if ty < miny {
			miny = ty //1019
			near = mems[i]
			if tm < minm {
				minm = tm
				near = mems[i]
				if td < mind {
					mind = td
					near = mems[i]
				}
			}
		}
	}

	return near
}

//MembersWithStartDate get wheather there are start dates below this date
func MembersWithStartDate(mems []member) ([]member, bool) {
	var isMemberPresentBeforeDate bool
	fmt.Println("Enter date in day-month-year format to get wheather there are start dates below this date: ")
	var inputy string
	fmt.Scanln(&inputy)
	s := strings.Split(inputy, "-")
	year, err := strconv.ParseInt(s[0], 10, 64)
	month, err := strconv.ParseInt(s[1], 10, 64)
	day, err := strconv.ParseInt(s[2], 10, 64)
	if err == nil {
	}
	var i int
	isMemberPresentBeforeDate = false
	for i = 0; i < len(mems); i++ {
		if strings.Compare(mems[i].enrollmentStartDate, "") == 0 {
			break
		}
		enrollmentStartDateSplit := strings.Split(mems[i].enrollmentStartDate, "-")
		esy, _ := strconv.ParseInt(enrollmentStartDateSplit[0], 10, 64)
		esm, _ := strconv.ParseInt(enrollmentStartDateSplit[1], 10, 64)
		esd, _ := strconv.ParseInt(enrollmentStartDateSplit[2], 10, 64)

		if year > esy {
			resultStartDateBelowGivenDate = append(resultStartDateBelowGivenDate, mems[i])

			isMemberPresentBeforeDate = true
		}
		if year == esy {
			if month > esm {
				resultStartDateBelowGivenDate = append(resultStartDateBelowGivenDate, mems[i])
				isMemberPresentBeforeDate = true

			}
			if month == esm {
				if day > esd {
					resultStartDateBelowGivenDate = append(resultStartDateBelowGivenDate, mems[i])
					isMemberPresentBeforeDate = true

				}
			}
		}
	}

	return resultStartDateBelowGivenDate, isMemberPresentBeforeDate
}

//MembersInGivenDate to get all the members who are in between the given date;Date lies between start date and end date
func MembersInGivenDate(mems []member) []member {
	fmt.Println(mems)
	fmt.Println("Enter date in year-month-day format: ")
	var inputy string
	fmt.Scanln(&inputy)
	s := strings.Split(inputy, "-")
	year, err := strconv.ParseInt(s[0], 10, 64)
	month, err := strconv.ParseInt(s[1], 10, 64)
	day, err := strconv.ParseInt(s[2], 10, 64)
	if err == nil {
	}
	var i int

	for i = 0; i < len(mems); i++ {
		if strings.Compare(mems[i].enrollmentStartDate, "") == 0 {
			break
		}
		if strings.Compare(mems[i].enrollmentEndDate, "") == 0 {
			break
		}
		enrollmentStartDateSplit := strings.Split(mems[i].enrollmentStartDate, "-")
		esy, _ := strconv.ParseInt(enrollmentStartDateSplit[0], 10, 64)
		esm, _ := strconv.ParseInt(enrollmentStartDateSplit[1], 10, 64)
		esd, _ := strconv.ParseInt(enrollmentStartDateSplit[2], 10, 64)
		enrollmentEndDateSplit := strings.Split(mems[i].enrollmentEndDate, "-")
		eey, _ := strconv.ParseInt(enrollmentEndDateSplit[0], 10, 64)
		eem, _ := strconv.ParseInt(enrollmentEndDateSplit[1], 10, 64)
		eed, _ := strconv.ParseInt(enrollmentEndDateSplit[2], 10, 64)

		if year < eey && year > esy {
			resultDateBwStartDateEndDate = append(resultDateBwStartDateEndDate, mems[i])
		}
		if year == eey && year > esy {
			if month < eem {
				resultDateBwStartDateEndDate = append(resultDateBwStartDateEndDate, mems[i])
			}
			if month == eem {
				if day < eed && day > esd {
					resultDateBwStartDateEndDate = append(resultDateBwStartDateEndDate, mems[i])

				}
			}
		}
		if year < eey && year == esy {
			if month > esm {
				resultDateBwStartDateEndDate = append(resultDateBwStartDateEndDate, mems[i])
			}
			if month == eem && month > esm {
				if day > esd {
					resultDateBwStartDateEndDate = append(resultDateBwStartDateEndDate, mems[i])

				}
			}
		}
		if year == eey && year == esy {
			if month <= eem && month >= esm {
				resultDateBwStartDateEndDate = append(resultDateBwStartDateEndDate, mems[i])
			}
			if month == eem && month == esm {
				if day < esd && day > eed {
					resultDateBwStartDateEndDate = append(resultDateBwStartDateEndDate, mems[i])
				}
			}
		}
	}
	return resultDateBwStartDateEndDate
}
