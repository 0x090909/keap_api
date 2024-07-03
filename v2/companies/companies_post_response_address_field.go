package companies

import (
	"errors"
)

type CompaniesPostResponse_address_field int

const (
	ADDRESS_FIELD_UNSPECIFIED_COMPANIESPOSTRESPONSE_ADDRESS_FIELD CompaniesPostResponse_address_field = iota
	BILLING_COMPANIESPOSTRESPONSE_ADDRESS_FIELD
	SHIPPING_COMPANIESPOSTRESPONSE_ADDRESS_FIELD
	OTHER_COMPANIESPOSTRESPONSE_ADDRESS_FIELD
)

func (i CompaniesPostResponse_address_field) String() string {
	return []string{"ADDRESS_FIELD_UNSPECIFIED", "BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseCompaniesPostResponse_address_field(v string) (any, error) {
	result := ADDRESS_FIELD_UNSPECIFIED_COMPANIESPOSTRESPONSE_ADDRESS_FIELD
	switch v {
	case "ADDRESS_FIELD_UNSPECIFIED":
		result = ADDRESS_FIELD_UNSPECIFIED_COMPANIESPOSTRESPONSE_ADDRESS_FIELD
	case "BILLING":
		result = BILLING_COMPANIESPOSTRESPONSE_ADDRESS_FIELD
	case "SHIPPING":
		result = SHIPPING_COMPANIESPOSTRESPONSE_ADDRESS_FIELD
	case "OTHER":
		result = OTHER_COMPANIESPOSTRESPONSE_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown CompaniesPostResponse_address_field value: " + v)
	}
	return &result, nil
}
func SerializeCompaniesPostResponse_address_field(values []CompaniesPostResponse_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CompaniesPostResponse_address_field) isMultiValue() bool {
	return false
}
