package companies

import (
	"errors"
)

type CompaniesGetResponse_companies_fax_number_field int

const (
	FAX_NUMBER_FIELD_UNSPECIFIED_COMPANIESGETRESPONSE_COMPANIES_FAX_NUMBER_FIELD CompaniesGetResponse_companies_fax_number_field = iota
	FAX1_COMPANIESGETRESPONSE_COMPANIES_FAX_NUMBER_FIELD
	FAX2_COMPANIESGETRESPONSE_COMPANIES_FAX_NUMBER_FIELD
)

func (i CompaniesGetResponse_companies_fax_number_field) String() string {
	return []string{"FAX_NUMBER_FIELD_UNSPECIFIED", "FAX1", "FAX2"}[i]
}
func ParseCompaniesGetResponse_companies_fax_number_field(v string) (any, error) {
	result := FAX_NUMBER_FIELD_UNSPECIFIED_COMPANIESGETRESPONSE_COMPANIES_FAX_NUMBER_FIELD
	switch v {
	case "FAX_NUMBER_FIELD_UNSPECIFIED":
		result = FAX_NUMBER_FIELD_UNSPECIFIED_COMPANIESGETRESPONSE_COMPANIES_FAX_NUMBER_FIELD
	case "FAX1":
		result = FAX1_COMPANIESGETRESPONSE_COMPANIES_FAX_NUMBER_FIELD
	case "FAX2":
		result = FAX2_COMPANIESGETRESPONSE_COMPANIES_FAX_NUMBER_FIELD
	default:
		return 0, errors.New("Unknown CompaniesGetResponse_companies_fax_number_field value: " + v)
	}
	return &result, nil
}
func SerializeCompaniesGetResponse_companies_fax_number_field(values []CompaniesGetResponse_companies_fax_number_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CompaniesGetResponse_companies_fax_number_field) isMultiValue() bool {
	return false
}