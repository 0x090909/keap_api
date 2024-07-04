package contacts

import (
	"errors"
)

type ContactsPutResponse_phone_numbers_field int

const (
	PHONE1_CONTACTSPUTRESPONSE_PHONE_NUMBERS_FIELD ContactsPutResponse_phone_numbers_field = iota
	PHONE2_CONTACTSPUTRESPONSE_PHONE_NUMBERS_FIELD
	PHONE3_CONTACTSPUTRESPONSE_PHONE_NUMBERS_FIELD
	PHONE4_CONTACTSPUTRESPONSE_PHONE_NUMBERS_FIELD
	PHONE5_CONTACTSPUTRESPONSE_PHONE_NUMBERS_FIELD
)

func (i ContactsPutResponse_phone_numbers_field) String() string {
	return []string{"PHONE1", "PHONE2", "PHONE3", "PHONE4", "PHONE5"}[i]
}
func ParseContactsPutResponse_phone_numbers_field(v string) (any, error) {
	result := PHONE1_CONTACTSPUTRESPONSE_PHONE_NUMBERS_FIELD
	switch v {
	case "PHONE1":
		result = PHONE1_CONTACTSPUTRESPONSE_PHONE_NUMBERS_FIELD
	case "PHONE2":
		result = PHONE2_CONTACTSPUTRESPONSE_PHONE_NUMBERS_FIELD
	case "PHONE3":
		result = PHONE3_CONTACTSPUTRESPONSE_PHONE_NUMBERS_FIELD
	case "PHONE4":
		result = PHONE4_CONTACTSPUTRESPONSE_PHONE_NUMBERS_FIELD
	case "PHONE5":
		result = PHONE5_CONTACTSPUTRESPONSE_PHONE_NUMBERS_FIELD
	default:
		return 0, errors.New("Unknown ContactsPutResponse_phone_numbers_field value: " + v)
	}
	return &result, nil
}
func SerializeContactsPutResponse_phone_numbers_field(values []ContactsPutResponse_phone_numbers_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPutResponse_phone_numbers_field) isMultiValue() bool {
	return false
}
