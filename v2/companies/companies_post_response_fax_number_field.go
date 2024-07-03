package companies

import (
	"errors"
)

type CompaniesPostResponse_fax_number_field int

const (
	FAX_NUMBER_FIELD_UNSPECIFIED_COMPANIESPOSTRESPONSE_FAX_NUMBER_FIELD CompaniesPostResponse_fax_number_field = iota
	FAX1_COMPANIESPOSTRESPONSE_FAX_NUMBER_FIELD
	FAX2_COMPANIESPOSTRESPONSE_FAX_NUMBER_FIELD
)

func (i CompaniesPostResponse_fax_number_field) String() string {
	return []string{"FAX_NUMBER_FIELD_UNSPECIFIED", "FAX1", "FAX2"}[i]
}
func ParseCompaniesPostResponse_fax_number_field(v string) (any, error) {
	result := FAX_NUMBER_FIELD_UNSPECIFIED_COMPANIESPOSTRESPONSE_FAX_NUMBER_FIELD
	switch v {
	case "FAX_NUMBER_FIELD_UNSPECIFIED":
		result = FAX_NUMBER_FIELD_UNSPECIFIED_COMPANIESPOSTRESPONSE_FAX_NUMBER_FIELD
	case "FAX1":
		result = FAX1_COMPANIESPOSTRESPONSE_FAX_NUMBER_FIELD
	case "FAX2":
		result = FAX2_COMPANIESPOSTRESPONSE_FAX_NUMBER_FIELD
	default:
		return 0, errors.New("Unknown CompaniesPostResponse_fax_number_field value: " + v)
	}
	return &result, nil
}
func SerializeCompaniesPostResponse_fax_number_field(values []CompaniesPostResponse_fax_number_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CompaniesPostResponse_fax_number_field) isMultiValue() bool {
	return false
}
