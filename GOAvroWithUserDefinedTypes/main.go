package main

import (
	"fmt"
	"github.com/linkedin/goavro"
)

var (
	codec *goavro.Codec
)

func main() {
	schema:=`{
		"namespace": "my.namespace.com",
		"type":	"record",
		"name": "indentity",
		"fields": [
			{ "name": "Name", "type": "string"},
			{ "name": "Code", "type": "string"},
			{ "name": "Year", "type": "string" },
			{ "name": "EnrollmentStartDate", "type": ["null",{
				"namespace": "my.namespace.com",
				"type":	"record",
				"name": "enrollmentStartDate",
				"fields": [
					{ "name": "Day", "type": "long" },
					{ "name": "Month", "type": "long" },
					{ "name": "Year", "type": "long" }
				]
			}],"default":null},
			{ "name": "EnrollmentEndDate", "type": ["null",{
				"namespace": "my.namespace.com",
				"type":	"record",
				"name": "enrollmentEndDate",
				"fields": [
					{ "name": "Day", "type": "long" },
					{ "name": "Month", "type": "long" },
					{ "name": "Year", "type": "long" }
				]
			}],"default":null}
		]
	}`

	//Create Schema Once
	codec, err := goavro.NewCodec(string(schema))
	if err != nil {
		panic(err)
	}
	//Sample Data
	Member := &Member{
		Name: "santhosh",
		Code:  "b15cs067",
		Year:"2019",
		EnrollmentStartDate: &Date{
			Day: 2,
			Month:  2,
			Year:  2016,
		},
		EnrollmentEndDate: &Date{
			Day: 2,
			Month:  2,
			Year:  2018,
		},
	}

	fmt.Printf("user in=%+v\n", Member)
	//fmt.Printf((String)codec)
	///Convert Binary From Native
	fmt.Println()
	fmt.Println(Member.ToStringMap())
	fmt.Println()
	binary, err := codec.BinaryFromNative(nil, Member.ToStringMap())
	if err != nil {
		panic(err)
	}
fmt.Println()
fmt.Println(binary)
fmt.Println()
	///Convert Native from Binary
	native, _, err := codec.NativeFromBinary(binary)
	if err != nil {
		panic(err)
	}

	//Convert it back tp Native
	userOut := StringMapToMember(native.(map[string]interface{}))
	fmt.Printf("user out=%+v\n", userOut)
	fmt.Println()
	// fmt.Println(userOut.Address)
	// if ok := reflect.DeepEqual(user, userOut); !ok {
	// 	fmt.Fprintf(os.Stderr, "struct Compare Failed ok=%t\n", ok)
	// 	os.Exit(1)
	// }
}

// User holds information about a user.
type Member struct {
	Name string
	Code  string
	Year    string
	EnrollmentStartDate *Date
	EnrollmentEndDate *Date
}

// Address holds information about an address.
type Date struct {
	Day int64
	Month int64
	Year     int64
}

// ToStringMap returns a map representation of the User.
func (u *Member) ToStringMap() map[string]interface{} {
	datumIn := map[string]interface{}{
		"Name": string(u.Name),
		"Year":  string(u.Year),
		"Code": string(u.Code),
	}

	if u.EnrollmentStartDate != nil {
		addDatum1 := map[string]interface{}{
			"Day": int64(u.EnrollmentStartDate.Day),
			"Month":     int64(u.EnrollmentStartDate.Month),
			"Year":    int64(u.EnrollmentStartDate.Year),
		}
		if u.EnrollmentEndDate != nil {
			addDatum2 := map[string]interface{}{
				"Day": int64(u.EnrollmentEndDate.Day),
				"Month":     int64(u.EnrollmentEndDate.Month),
				"Year":    int64(u.EnrollmentEndDate.Year),
		}

		//important need namespace and record name
		datumIn["EnrollmentStartDate"] = goavro.Union("my.namespace.com.enrollmentStartDate", addDatum1)
		datumIn["EnrollmentEndDate"] = goavro.Union("my.namespace.com.enrollmentEndDate", addDatum2)

	} else {
		datumIn["EnrollmentStartDate"] = goavro.Union("null", nil)
		datumIn["EnrollmentEndDate"] = goavro.Union("null", nil)
	}
}
	return datumIn
}

//StringMapToUser returns a User from a map representation of the User.
func  StringMapToMember(data map[string]interface{}) *Member {
	ind := &Member{}
	for k, v := range data {
		switch k {
		case "Name":
			if value, ok := v.(string); ok {
				ind.Name = value
			}
		case "Year":
			if value, ok := v.(string); ok {
				ind.Year = value
			}
		case "Code":
			if value, ok := v.(string); ok {
				ind.Code = value
			}
		case "EnrollmentStartDate":
			if vmap, ok := v.(map[string]interface{}); ok {
				//important need namespace and record name
				if cookieSMap, ok := vmap["my.namespace.com.enrollmentStartDate"].(map[string]interface{}); ok {
					add := &Date{}
					for k, v := range cookieSMap {
						switch k {
						case "Day":
							if value, ok := v.(int64); ok {
								add.Day = value
							}
						case "Month":
							if value, ok := v.(int64); ok {
								add.Month = value
							}
						case "Year":
							if value, ok := v.(int64); ok {
								add.Year = value
							}
						}
					}
					ind.EnrollmentStartDate = add
				}
		}
	case "EnrollmentEndDate":
		if vmap, ok := v.(map[string]interface{}); ok {
			//important need namespace and record name
			if cookieSMap, ok := vmap["my.namespace.com.enrollmentEndDate"].(map[string]interface{}); ok {
				add := &Date{}
				for k, v := range cookieSMap {
					switch k {
					case "Day":
						if value, ok := v.(int64); ok {
							add.Day = value
						}
					case "Month":
						if value, ok := v.(int64); ok {
							add.Month = value
						}
					case "Year":
						if value, ok := v.(int64); ok {
							add.Year = value
						}
					}
				}
				ind.EnrollmentEndDate = add
			}
		}

	}
}
	return ind
}
