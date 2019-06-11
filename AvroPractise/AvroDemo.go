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
		"name": "value_TestingGolangKafkaObjectsKey",
		"fields": [
			{ "name": "KeyFeild", "type": "string"}
		]
	}`

	//Create Schema Once
	codec, err := goavro.NewCodec(string(schema))
	if err != nil {
		panic(err)
	}
	//Sample Data
	Member := &Key{"pranay Kumar "}

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
type Key struct {
	KeyFeild string
}

// Address holds information about an address.

// ToStringMap returns a map representation of the User.
func (u *Key) ToStringMap() map[string]interface{} {
	datumIn := map[string]interface{}{
		"KeyFeild": string(u.KeyFeild),
	}

	return datumIn
}

//StringMapToUser returns a User from a map representation of the User.
func StringMapToMember(data map[string]interface{}) *Key {
	ind := &Key{}
	for k, v := range data {
		switch k {
		case "KeyFeild":
			if value, ok := v.(string); ok {
				ind.KeyFeild = value
			}

		}
	}
	return ind
}
