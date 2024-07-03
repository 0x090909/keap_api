package companies

import (
	"errors"
)

type CompaniesPostRequestBody_fax_number_field int

const (
	FAX_NUMBER_FIELD_UNSPECIFIED_COMPANIESPOSTREQUESTBODY_FAX_NUMBER_FIELD CompaniesPostRequestBody_fax_number_field = iota
	FAX1_COMPANIESPOSTREQUESTBODY_FAX_NUMBER_FIELD
	FAX2_COMPANIESPOSTREQUESTBODY_FAX_NUMBER_FIELD
)

func (i CompaniesPostRequestBody_fax_number_field) String() string {
	return []string{"FAX_NUMBER_FIELD_UNSPECIFIED", "FAX1", "FAX2"}[i]
}
func ParseCompaniesPostRequestBody_fax_number_field(v string) (any, error) {
	result := FAX_NUMBER_FIELD_UNSPECIFIED_COMPANIESPOSTREQUESTBODY_FAX_NUMBER_FIELD
	switch v {
	case "FAX_NUMBER_FIELD_UNSPECIFIED":
		result = FAX_NUMBER_FIELD_UNSPECIFIED_COMPANIESPOSTREQUESTBODY_FAX_NUMBER_FIELD
	case "FAX1":
		result = FAX1_COMPANIESPOSTREQUESTBODY_FAX_NUMBER_FIELD
	case "FAX2":
		result = FAX2_COMPANIESPOSTREQUESTBODY_FAX_NUMBER_FIELD
	default:
		return 0, errors.New("Unknown CompaniesPostRequestBody_fax_number_field value: " + v)
	}
	return &result, nil
}
func SerializeCompaniesPostRequestBody_fax_number_field(values []CompaniesPostRequestBody_fax_number_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CompaniesPostRequestBody_fax_number_field) isMultiValue() bool {
	return false
}
