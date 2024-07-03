package companies

import (
	"errors"
)

type CompaniesGetResponse_companies_address_field int

const (
	ADDRESS_FIELD_UNSPECIFIED_COMPANIESGETRESPONSE_COMPANIES_ADDRESS_FIELD CompaniesGetResponse_companies_address_field = iota
	BILLING_COMPANIESGETRESPONSE_COMPANIES_ADDRESS_FIELD
	SHIPPING_COMPANIESGETRESPONSE_COMPANIES_ADDRESS_FIELD
	OTHER_COMPANIESGETRESPONSE_COMPANIES_ADDRESS_FIELD
)

func (i CompaniesGetResponse_companies_address_field) String() string {
	return []string{"ADDRESS_FIELD_UNSPECIFIED", "BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseCompaniesGetResponse_companies_address_field(v string) (any, error) {
	result := ADDRESS_FIELD_UNSPECIFIED_COMPANIESGETRESPONSE_COMPANIES_ADDRESS_FIELD
	switch v {
	case "ADDRESS_FIELD_UNSPECIFIED":
		result = ADDRESS_FIELD_UNSPECIFIED_COMPANIESGETRESPONSE_COMPANIES_ADDRESS_FIELD
	case "BILLING":
		result = BILLING_COMPANIESGETRESPONSE_COMPANIES_ADDRESS_FIELD
	case "SHIPPING":
		result = SHIPPING_COMPANIESGETRESPONSE_COMPANIES_ADDRESS_FIELD
	case "OTHER":
		result = OTHER_COMPANIESGETRESPONSE_COMPANIES_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown CompaniesGetResponse_companies_address_field value: " + v)
	}
	return &result, nil
}
func SerializeCompaniesGetResponse_companies_address_field(values []CompaniesGetResponse_companies_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CompaniesGetResponse_companies_address_field) isMultiValue() bool {
	return false
}
