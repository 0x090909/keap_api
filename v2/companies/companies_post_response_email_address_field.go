package companies

import (
	"errors"
)

type CompaniesPostResponse_email_address_field int

const (
	EMAIL_FIELD_UNSPECIFIED_COMPANIESPOSTRESPONSE_EMAIL_ADDRESS_FIELD CompaniesPostResponse_email_address_field = iota
	EMAIL1_COMPANIESPOSTRESPONSE_EMAIL_ADDRESS_FIELD
	EMAIL2_COMPANIESPOSTRESPONSE_EMAIL_ADDRESS_FIELD
	EMAIL3_COMPANIESPOSTRESPONSE_EMAIL_ADDRESS_FIELD
)

func (i CompaniesPostResponse_email_address_field) String() string {
	return []string{"EMAIL_FIELD_UNSPECIFIED", "EMAIL1", "EMAIL2", "EMAIL3"}[i]
}
func ParseCompaniesPostResponse_email_address_field(v string) (any, error) {
	result := EMAIL_FIELD_UNSPECIFIED_COMPANIESPOSTRESPONSE_EMAIL_ADDRESS_FIELD
	switch v {
	case "EMAIL_FIELD_UNSPECIFIED":
		result = EMAIL_FIELD_UNSPECIFIED_COMPANIESPOSTRESPONSE_EMAIL_ADDRESS_FIELD
	case "EMAIL1":
		result = EMAIL1_COMPANIESPOSTRESPONSE_EMAIL_ADDRESS_FIELD
	case "EMAIL2":
		result = EMAIL2_COMPANIESPOSTRESPONSE_EMAIL_ADDRESS_FIELD
	case "EMAIL3":
		result = EMAIL3_COMPANIESPOSTRESPONSE_EMAIL_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown CompaniesPostResponse_email_address_field value: " + v)
	}
	return &result, nil
}
func SerializeCompaniesPostResponse_email_address_field(values []CompaniesPostResponse_email_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CompaniesPostResponse_email_address_field) isMultiValue() bool {
	return false
}
