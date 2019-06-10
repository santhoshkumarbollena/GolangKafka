package Model

import "fmt"

type Member struct {
	OffshoreRestrictedIndicator string
	ProfileIdentifier           string
	SecureClassIdentifier       string
	EnrollmentEffectiveDate     string
	EnrollmentTerminationDate   string
}

func Demo() {
	fmt.Println("demo calle")
}
