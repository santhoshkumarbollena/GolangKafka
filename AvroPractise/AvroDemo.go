package main

import (
	"fmt"

	"github.com/linkedin/goavro"
)

var (
	codec *goavro.Codec
)

func main() {
	schema := `{
		"namespace": "my.namespace.com",
		"type":	"record",
		"name": "value_TestingGolangKafkaObjects",
		"fields": [
			{ "name": "OffshoreRestrictedIndicator", "type": "string"},
			{ "name": "ProfileIdentifier", "type": "string"},
			{ "name": "SecureClassIdentifier", "type": "string" }	,
			{ "name": "EnrollmentEffectiveDate", "type": "string" }	,
			{ "name": "EnrollmentTerminationDate", "type": "string" }
		]
	}`

	//Create Schema Once
	codec, err := goavro.NewCodec(string(schema))
	if err != nil {
		panic(err)
	}
	//Sample Data
	Member := &Member{"pranay Kumar ", "9999", "Y", "2015", "2019"}

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
	OffshoreRestrictedIndicator string
	ProfileIdentifier           string
	SecureClassIdentifier       string
	EnrollmentEffectiveDate     string
	EnrollmentTerminationDate   string
}

// Address holds information about an address.

// ToStringMap returns a map representation of the User.
func (u *Member) ToStringMap() map[string]interface{} {
	datumIn := map[string]interface{}{
		"OffshoreRestrictedIndicator": string(u.OffshoreRestrictedIndicator),
		"ProfileIdentifier":           string(u.ProfileIdentifier),
		"SecureClassIdentifier":       string(u.SecureClassIdentifier),
		"EnrollmentEffectiveDate":     string(u.EnrollmentEffectiveDate),
		"EnrollmentTerminationDate":   string(u.EnrollmentTerminationDate),
	}

	return datumIn
}

//StringMapToUser returns a User from a map representation of the User.
func StringMapToMember(data map[string]interface{}) *Member {
	ind := &Member{}
	for k, v := range data {
		switch k {
		case "OffshoreRestrictedIndicator":
			if value, ok := v.(string); ok {
				ind.OffshoreRestrictedIndicator = value
			}
		case "ProfileIdentifier":
			if value, ok := v.(string); ok {
				ind.ProfileIdentifier = value
			}
		case "SecureClassIdentifier":
			if value, ok := v.(string); ok {
				ind.SecureClassIdentifier = value
			}
		case "EnrollmentEffectiveDate":
			if value, ok := v.(string); ok {
				ind.EnrollmentEffectiveDate = value
			}
		case "EnrollmentTerminationDate":
			if value, ok := v.(string); ok {
				ind.EnrollmentTerminationDate = value
			}

		}
	}
	return ind
}
