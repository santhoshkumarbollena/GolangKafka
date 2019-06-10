package OperationsOnMembersObject

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	d "../../GoLangProjectLayout/CommaSeperatedValuesToObjects"
	m "../../GoLangProjectLayout/Model"
)

var ff bool = false
var resultDateBwStartDateEndDate []member
var resultStartDateBelowGivenDate []member

type member = m.Member

func OperationsOnMembersObjectFunction() {
	//fmt.Println(d.GetMembersObject())
	ResuletMembersRes := d.GetMembersObject()
	fmt.Println(ResuletMembersRes)
	fmt.Println("Members whose enrollment start date and end date in between given date")
	//res := []member(ResuletMembersRes)
	membersindate := MembersInGivenDate(ResuletMembersRes)
	//fmt.Println(membersindate)
	for i := 0; i < len(membersindate); i++ {
		s := reflect.ValueOf(&membersindate[i]).Elem()
		typeOfT := s.Type()

		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Printf("%d:  %s = %v\n", i,
				typeOfT.Field(i).Name, f.Interface())
		}
	}
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
		if strings.Compare(mems[i].EnrollmentEffectiveDate, "") == 0 {
			break
		}
		enrollmentStartDateSplit := strings.Split(mems[i].EnrollmentEffectiveDate, "-")
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
	//fmt.Println(mems)
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
		if strings.Compare(mems[i].EnrollmentEffectiveDate, "") == 0 {
			break
		}
		if strings.Compare(mems[i].EnrollmentTerminationDate, "") == 0 {
			break
		}
		enrollmentStartDateSplit := strings.Split(mems[i].EnrollmentEffectiveDate, "-")
		esy, _ := strconv.ParseInt(enrollmentStartDateSplit[0], 10, 64)
		esm, _ := strconv.ParseInt(enrollmentStartDateSplit[1], 10, 64)
		esd, _ := strconv.ParseInt(enrollmentStartDateSplit[2], 10, 64)
		enrollmentEndDateSplit := strings.Split(mems[i].EnrollmentTerminationDate, "-")
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
		if strings.Compare(mems[i].EnrollmentTerminationDate, "") == 0 {
			break
		}
		enrollmentDateSplit := strings.Split(mems[i].EnrollmentTerminationDate, "-")
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
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
