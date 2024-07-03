package companies

import (
	"errors"
)

type CompaniesPostRequestBody_email_address_field int

const (
	EMAIL_FIELD_UNSPECIFIED_COMPANIESPOSTREQUESTBODY_EMAIL_ADDRESS_FIELD CompaniesPostRequestBody_email_address_field = iota
	EMAIL1_COMPANIESPOSTREQUESTBODY_EMAIL_ADDRESS_FIELD
	EMAIL2_COMPANIESPOSTREQUESTBODY_EMAIL_ADDRESS_FIELD
	EMAIL3_COMPANIESPOSTREQUESTBODY_EMAIL_ADDRESS_FIELD
)

func (i CompaniesPostRequestBody_email_address_field) String() string {
	return []string{"EMAIL_FIELD_UNSPECIFIED", "EMAIL1", "EMAIL2", "EMAIL3"}[i]
}
func ParseCompaniesPostRequestBody_email_address_field(v string) (any, error) {
	result := EMAIL_FIELD_UNSPECIFIED_COMPANIESPOSTREQUESTBODY_EMAIL_ADDRESS_FIELD
	switch v {
	case "EMAIL_FIELD_UNSPECIFIED":
		result = EMAIL_FIELD_UNSPECIFIED_COMPANIESPOSTREQUESTBODY_EMAIL_ADDRESS_FIELD
	case "EMAIL1":
		result = EMAIL1_COMPANIESPOSTREQUESTBODY_EMAIL_ADDRESS_FIELD
	case "EMAIL2":
		result = EMAIL2_COMPANIESPOSTREQUESTBODY_EMAIL_ADDRESS_FIELD
	case "EMAIL3":
		result = EMAIL3_COMPANIESPOSTREQUESTBODY_EMAIL_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown CompaniesPostRequestBody_email_address_field value: " + v)
	}
	return &result, nil
}
func SerializeCompaniesPostRequestBody_email_address_field(values []CompaniesPostRequestBody_email_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CompaniesPostRequestBody_email_address_field) isMultiValue() bool {
	return false
}
