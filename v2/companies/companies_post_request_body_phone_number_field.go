package companies

import (
	"errors"
)

type CompaniesPostRequestBody_phone_number_field int

const (
	PHONE_NUMBER_FIELD_UNSPECIFIED_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD CompaniesPostRequestBody_phone_number_field = iota
	PHONE1_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD
	PHONE2_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD
	PHONE3_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD
	PHONE4_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD
	PHONE5_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD
)

func (i CompaniesPostRequestBody_phone_number_field) String() string {
	return []string{"PHONE_NUMBER_FIELD_UNSPECIFIED", "PHONE1", "PHONE2", "PHONE3", "PHONE4", "PHONE5"}[i]
}
func ParseCompaniesPostRequestBody_phone_number_field(v string) (any, error) {
	result := PHONE_NUMBER_FIELD_UNSPECIFIED_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD
	switch v {
	case "PHONE_NUMBER_FIELD_UNSPECIFIED":
		result = PHONE_NUMBER_FIELD_UNSPECIFIED_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD
	case "PHONE1":
		result = PHONE1_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD
	case "PHONE2":
		result = PHONE2_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD
	case "PHONE3":
		result = PHONE3_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD
	case "PHONE4":
		result = PHONE4_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD
	case "PHONE5":
		result = PHONE5_COMPANIESPOSTREQUESTBODY_PHONE_NUMBER_FIELD
	default:
		return 0, errors.New("Unknown CompaniesPostRequestBody_phone_number_field value: " + v)
	}
	return &result, nil
}
func SerializeCompaniesPostRequestBody_phone_number_field(values []CompaniesPostRequestBody_phone_number_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CompaniesPostRequestBody_phone_number_field) isMultiValue() bool {
	return false
}
