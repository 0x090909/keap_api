package companies

import (
	"errors"
)

type CompaniesPostResponse_phone_number_field int

const (
	PHONE_NUMBER_FIELD_UNSPECIFIED_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD CompaniesPostResponse_phone_number_field = iota
	PHONE1_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD
	PHONE2_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD
	PHONE3_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD
	PHONE4_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD
	PHONE5_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD
)

func (i CompaniesPostResponse_phone_number_field) String() string {
	return []string{"PHONE_NUMBER_FIELD_UNSPECIFIED", "PHONE1", "PHONE2", "PHONE3", "PHONE4", "PHONE5"}[i]
}
func ParseCompaniesPostResponse_phone_number_field(v string) (any, error) {
	result := PHONE_NUMBER_FIELD_UNSPECIFIED_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD
	switch v {
	case "PHONE_NUMBER_FIELD_UNSPECIFIED":
		result = PHONE_NUMBER_FIELD_UNSPECIFIED_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD
	case "PHONE1":
		result = PHONE1_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD
	case "PHONE2":
		result = PHONE2_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD
	case "PHONE3":
		result = PHONE3_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD
	case "PHONE4":
		result = PHONE4_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD
	case "PHONE5":
		result = PHONE5_COMPANIESPOSTRESPONSE_PHONE_NUMBER_FIELD
	default:
		return 0, errors.New("Unknown CompaniesPostResponse_phone_number_field value: " + v)
	}
	return &result, nil
}
func SerializeCompaniesPostResponse_phone_number_field(values []CompaniesPostResponse_phone_number_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CompaniesPostResponse_phone_number_field) isMultiValue() bool {
	return false
}
