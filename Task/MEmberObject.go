package main
import ("fmt"
"strconv"
"strings"
"math"
)

var members []member
var result []member
var resultStartDateBelowGivenDate []member
type date struct{
	day int64
	month int64
	year int64
}

type member struct {
			name string
			enrollmentStartDate date
			enrollmentEndDate date
			code string
			year  int64
			
}
func main(){
	
	s := member{name : "santhosh",enrollmentStartDate : date{day :1,month :1,year :2015},enrollmentEndDate : date{day :1,month :1,year :2017},code : "b15cs067",year : 2019}
	p := member{name : "santhosh kumar",enrollmentStartDate : date{day :1,month :1, year:2019},enrollmentEndDate : date{day :1,month :1,year :2020},code : "b15cs067",year : 1000}
	z := member{name : "santhosh bollena",enrollmentStartDate : date{day :1,month :1, year:2022},enrollmentEndDate : date{day :1,month :1,year :2023},code : "b15cs067",year : 1000}
	members=append(members,s);
	members=append(members,p);
	members=append(members,z);
	fmt.Println("Members whose enrollment start date and end date in between given date");
	membersindate:=MembersInGivenDate(members)
	fmt.Println(membersindate)
	fmt.Println("Members whose enrollment start date is prior to the given date");
	membersStartDateGreaterThanDate,f:=MembersWithStartDate(members)
	fmt.Println("the following boolen tells us that is there any start date present below the given date")
	fmt.Println(f)
	if(f){
	fmt.Println(membersStartDateGreaterThanDate)
	}
	fmt.Println("Members whose enrollment end date is near to given date");
	membersNearEnddate:=MembersnearEndDate(members)
	fmt.Println(membersNearEnddate)

}
func abs(x int64)int64{
	if x<0{
		return -x
	}
	return x
}
func MembersnearEndDate(mems []member)member{
    fmt.Println("Enter date in day-month-year to find near end date: ")
	var inputy string
	fmt.Scanln(&inputy)
	s := strings.Split(inputy, "-")
	day , err:=strconv.ParseInt(s[0],10,64)
	month , err:=strconv.ParseInt(s[1],10,64)
	year , err:=strconv.ParseInt(s[2],10,64)
	if err == nil {	}
	var i int
	var miny,minm,mind int64=math.MaxInt64,math.MaxInt64,math.MaxInt64
	var ty,tm,td int64
	
	var near member
	for i = 0; i < len(mems);i++ {
		
		ty=abs(year-mems[i].enrollmentEndDate.year)//1000-2019
		tm=abs(month-mems[i].enrollmentEndDate.month)
		td=abs(day-mems[i].enrollmentEndDate.day)
		if(ty<miny){
			miny=ty//1019
			near=mems[i]
			if(tm<minm){
				minm=tm
				near=mems[i]
				if(td<mind){
					mind=td
					near=mems[i]
				}
			}
		}
	}

	return near
}
func MembersWithStartDate(mems []member)([]member,bool){
	var isMemberPresentBeforeDate bool 
	fmt.Println("Enter date in day-month-year format to get wheather there are start dates below this date: ")
	var inputy string
	fmt.Scanln(&inputy)
	s := strings.Split(inputy, "-")
	day , err:=strconv.ParseInt(s[0],10,64)
	month , err:=strconv.ParseInt(s[1],10,64)
	year , err:=strconv.ParseInt(s[2],10,64)
	if err == nil {	}
	var i int
	isMemberPresentBeforeDate=false
	for i = 0; i < len(mems);i++ {
		if(year>mems[i].enrollmentStartDate.year){
			resultStartDateBelowGivenDate=append(resultStartDateBelowGivenDate,mems[i]);
			// fmt.Println("1st if")
			// fmt.Println(result)
			isMemberPresentBeforeDate=true
	 }
	 if(year==mems[i].enrollmentStartDate.year){
		if(month>mems[i].enrollmentStartDate.month){
			resultStartDateBelowGivenDate=append(resultStartDateBelowGivenDate,mems[i]);
			isMemberPresentBeforeDate=true
			// fmt.Println("2st if")
			// fmt.Println(result)
		}
		if(month==mems[i].enrollmentStartDate.month){
			if(day>mems[i].enrollmentStartDate.day){
				resultStartDateBelowGivenDate=append(resultStartDateBelowGivenDate,mems[i]);
				isMemberPresentBeforeDate=true
			// 	fmt.Println("3st if")
			// fmt.Println(result)
			}
		}
	 }
	}
	// fmt.Println()
	// fmt.Println(resultStartDateBelowGivenDate)
	// fmt.Println()
	return resultStartDateBelowGivenDate,isMemberPresentBeforeDate
}
func MembersInGivenDate(mems []member)[]member{
    fmt.Println("Enter date in day-month-year format: ")
	var inputy string
	fmt.Scanln(&inputy)
	s := strings.Split(inputy, "-")
	day , err:=strconv.ParseInt(s[0],10,64)
	month , err:=strconv.ParseInt(s[1],10,64)
	year , err:=strconv.ParseInt(s[2],10,64)
	if err == nil {	}
	var i int
	
	for i = 0; i < len(mems);i++ {
		
	if(year<mems[i].enrollmentEndDate.year&&year>mems[i].enrollmentStartDate.year){
			result=append(result,mems[i]);
	 }
	 if(year==mems[i].enrollmentEndDate.year&&year>mems[i].enrollmentStartDate.year){
		if(month<mems[i].enrollmentEndDate.month){
			result=append(result,mems[i]);
		}
		if(month==mems[i].enrollmentEndDate.month){
			if(day<mems[i].enrollmentEndDate.day&&day>mems[i].enrollmentStartDate.day){
				result=append(result,mems[i]);

			}
		}
	 }
	 if(year<mems[i].enrollmentEndDate.year&&year==mems[i].enrollmentStartDate.year){
		if(month>mems[i].enrollmentStartDate.month){
			result=append(result,mems[i]);
		}
		if(month==mems[i].enrollmentEndDate.month&&month>mems[i].enrollmentStartDate.month){
			if(day>mems[i].enrollmentStartDate.day){
				result=append(result,mems[i]);

			}
		}
	 }
	}
	return result
}