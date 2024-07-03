package item

import (
	"errors"
)

type WithContact_GetResponse_phone_numbers_field int

const (
	PHONE_NUMBER_FIELD_UNSPECIFIED_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD WithContact_GetResponse_phone_numbers_field = iota
	PHONE1_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD
	PHONE2_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD
	PHONE3_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD
	PHONE4_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD
	PHONE5_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD
)

func (i WithContact_GetResponse_phone_numbers_field) String() string {
	return []string{"PHONE_NUMBER_FIELD_UNSPECIFIED", "PHONE1", "PHONE2", "PHONE3", "PHONE4", "PHONE5"}[i]
}
func ParseWithContact_GetResponse_phone_numbers_field(v string) (any, error) {
	result := PHONE_NUMBER_FIELD_UNSPECIFIED_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD
	switch v {
	case "PHONE_NUMBER_FIELD_UNSPECIFIED":
		result = PHONE_NUMBER_FIELD_UNSPECIFIED_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD
	case "PHONE1":
		result = PHONE1_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD
	case "PHONE2":
		result = PHONE2_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD
	case "PHONE3":
		result = PHONE3_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD
	case "PHONE4":
		result = PHONE4_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD
	case "PHONE5":
		result = PHONE5_WITHCONTACT_GETRESPONSE_PHONE_NUMBERS_FIELD
	default:
		return 0, errors.New("Unknown WithContact_GetResponse_phone_numbers_field value: " + v)
	}
	return &result, nil
}
func SerializeWithContact_GetResponse_phone_numbers_field(values []WithContact_GetResponse_phone_numbers_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithContact_GetResponse_phone_numbers_field) isMultiValue() bool {
	return false
}
