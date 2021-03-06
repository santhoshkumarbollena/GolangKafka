package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var members []member
var resultDateBwStartDateEndDate []member
var resultStartDateBelowGivenDate []member

type member struct {
	name                string
	enrollmentStartDate string
	enrollmentEndDate   string
	code                string
}

func main() {

	s := member{name: "santhosh", enrollmentStartDate: "2019-1-12", enrollmentEndDate: "2019-12-28", code: "b15cs067"}
	p := member{name: "santhosh", enrollmentStartDate: "2000-4-5", enrollmentEndDate: "2017-5-9", code: "b15cs067"}
	z := member{name: "santhosh", enrollmentStartDate: "1000-4-6", enrollmentEndDate: "1008-5-4", code: "b15cs067"}
	members = append(members, s)
	members = append(members, p)
	members = append(members, z)
	fmt.Println("Members whose enrollment start date and end date in between given date")
	membersindate := MembersInGivenDate(members)
	fmt.Println(membersindate)
	//"Members whose enrollment start date is prior to the given date"
	membersStartDateGreaterThanDate, f := MembersWithStartDate(members)
	//the following boolen tells us that is there any start date present below the given date")
	fmt.Println(f)
	if f {
		fmt.Println(membersStartDateGreaterThanDate)
	}
	fmt.Println("Members whose enrollment end date is near to given date")
	membersNearEnddate := MembersnearEndDate(members)
	fmt.Println(membersNearEnddate)

}
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
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
