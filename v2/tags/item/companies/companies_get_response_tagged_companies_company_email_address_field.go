package companies

import (
	"errors"
)

type CompaniesGetResponse_tagged_companies_company_email_address_field int

const (
	EMAIL_FIELD_UNSPECIFIED_COMPANIESGETRESPONSE_TAGGED_COMPANIES_COMPANY_EMAIL_ADDRESS_FIELD CompaniesGetResponse_tagged_companies_company_email_address_field = iota
	EMAIL1_COMPANIESGETRESPONSE_TAGGED_COMPANIES_COMPANY_EMAIL_ADDRESS_FIELD
	EMAIL2_COMPANIESGETRESPONSE_TAGGED_COMPANIES_COMPANY_EMAIL_ADDRESS_FIELD
	EMAIL3_COMPANIESGETRESPONSE_TAGGED_COMPANIES_COMPANY_EMAIL_ADDRESS_FIELD
)

func (i CompaniesGetResponse_tagged_companies_company_email_address_field) String() string {
	return []string{"EMAIL_FIELD_UNSPECIFIED", "EMAIL1", "EMAIL2", "EMAIL3"}[i]
}
func ParseCompaniesGetResponse_tagged_companies_company_email_address_field(v string) (any, error) {
	result := EMAIL_FIELD_UNSPECIFIED_COMPANIESGETRESPONSE_TAGGED_COMPANIES_COMPANY_EMAIL_ADDRESS_FIELD
	switch v {
	case "EMAIL_FIELD_UNSPECIFIED":
		result = EMAIL_FIELD_UNSPECIFIED_COMPANIESGETRESPONSE_TAGGED_COMPANIES_COMPANY_EMAIL_ADDRESS_FIELD
	case "EMAIL1":
		result = EMAIL1_COMPANIESGETRESPONSE_TAGGED_COMPANIES_COMPANY_EMAIL_ADDRESS_FIELD
	case "EMAIL2":
		result = EMAIL2_COMPANIESGETRESPONSE_TAGGED_COMPANIES_COMPANY_EMAIL_ADDRESS_FIELD
	case "EMAIL3":
		result = EMAIL3_COMPANIESGETRESPONSE_TAGGED_COMPANIES_COMPANY_EMAIL_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown CompaniesGetResponse_tagged_companies_company_email_address_field value: " + v)
	}
	return &result, nil
}
func SerializeCompaniesGetResponse_tagged_companies_company_email_address_field(values []CompaniesGetResponse_tagged_companies_company_email_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CompaniesGetResponse_tagged_companies_company_email_address_field) isMultiValue() bool {
	return false
}
