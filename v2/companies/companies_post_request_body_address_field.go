package companies

import (
	"errors"
)

type CompaniesPostRequestBody_address_field int

const (
	ADDRESS_FIELD_UNSPECIFIED_COMPANIESPOSTREQUESTBODY_ADDRESS_FIELD CompaniesPostRequestBody_address_field = iota
	BILLING_COMPANIESPOSTREQUESTBODY_ADDRESS_FIELD
	SHIPPING_COMPANIESPOSTREQUESTBODY_ADDRESS_FIELD
	OTHER_COMPANIESPOSTREQUESTBODY_ADDRESS_FIELD
)

func (i CompaniesPostRequestBody_address_field) String() string {
	return []string{"ADDRESS_FIELD_UNSPECIFIED", "BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseCompaniesPostRequestBody_address_field(v string) (any, error) {
	result := ADDRESS_FIELD_UNSPECIFIED_COMPANIESPOSTREQUESTBODY_ADDRESS_FIELD
	switch v {
	case "ADDRESS_FIELD_UNSPECIFIED":
		result = ADDRESS_FIELD_UNSPECIFIED_COMPANIESPOSTREQUESTBODY_ADDRESS_FIELD
	case "BILLING":
		result = BILLING_COMPANIESPOSTREQUESTBODY_ADDRESS_FIELD
	case "SHIPPING":
		result = SHIPPING_COMPANIESPOSTREQUESTBODY_ADDRESS_FIELD
	case "OTHER":
		result = OTHER_COMPANIESPOSTREQUESTBODY_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown CompaniesPostRequestBody_address_field value: " + v)
	}
	return &result, nil
}
func SerializeCompaniesPostRequestBody_address_field(values []CompaniesPostRequestBody_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CompaniesPostRequestBody_address_field) isMultiValue() bool {
	return false
}
