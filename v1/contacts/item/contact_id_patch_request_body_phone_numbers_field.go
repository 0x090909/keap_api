package item

import (
	"errors"
)

type ContactIdPatchRequestBody_phone_numbers_field int

const (
	PHONE1_CONTACTIDPATCHREQUESTBODY_PHONE_NUMBERS_FIELD ContactIdPatchRequestBody_phone_numbers_field = iota
	PHONE2_CONTACTIDPATCHREQUESTBODY_PHONE_NUMBERS_FIELD
	PHONE3_CONTACTIDPATCHREQUESTBODY_PHONE_NUMBERS_FIELD
	PHONE4_CONTACTIDPATCHREQUESTBODY_PHONE_NUMBERS_FIELD
	PHONE5_CONTACTIDPATCHREQUESTBODY_PHONE_NUMBERS_FIELD
)

func (i ContactIdPatchRequestBody_phone_numbers_field) String() string {
	return []string{"PHONE1", "PHONE2", "PHONE3", "PHONE4", "PHONE5"}[i]
}
func ParseContactIdPatchRequestBody_phone_numbers_field(v string) (any, error) {
	result := PHONE1_CONTACTIDPATCHREQUESTBODY_PHONE_NUMBERS_FIELD
	switch v {
	case "PHONE1":
		result = PHONE1_CONTACTIDPATCHREQUESTBODY_PHONE_NUMBERS_FIELD
	case "PHONE2":
		result = PHONE2_CONTACTIDPATCHREQUESTBODY_PHONE_NUMBERS_FIELD
	case "PHONE3":
		result = PHONE3_CONTACTIDPATCHREQUESTBODY_PHONE_NUMBERS_FIELD
	case "PHONE4":
		result = PHONE4_CONTACTIDPATCHREQUESTBODY_PHONE_NUMBERS_FIELD
	case "PHONE5":
		result = PHONE5_CONTACTIDPATCHREQUESTBODY_PHONE_NUMBERS_FIELD
	default:
		return 0, errors.New("Unknown ContactIdPatchRequestBody_phone_numbers_field value: " + v)
	}
	return &result, nil
}
func SerializeContactIdPatchRequestBody_phone_numbers_field(values []ContactIdPatchRequestBody_phone_numbers_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactIdPatchRequestBody_phone_numbers_field) isMultiValue() bool {
	return false
}
