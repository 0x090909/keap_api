package contacts

import (
	"errors"
)

type ContactsPostRequestBody_fax_numbers_field int

const (
	FAX1_CONTACTSPOSTREQUESTBODY_FAX_NUMBERS_FIELD ContactsPostRequestBody_fax_numbers_field = iota
	FAX2_CONTACTSPOSTREQUESTBODY_FAX_NUMBERS_FIELD
)

func (i ContactsPostRequestBody_fax_numbers_field) String() string {
	return []string{"FAX1", "FAX2"}[i]
}
func ParseContactsPostRequestBody_fax_numbers_field(v string) (any, error) {
	result := FAX1_CONTACTSPOSTREQUESTBODY_FAX_NUMBERS_FIELD
	switch v {
	case "FAX1":
		result = FAX1_CONTACTSPOSTREQUESTBODY_FAX_NUMBERS_FIELD
	case "FAX2":
		result = FAX2_CONTACTSPOSTREQUESTBODY_FAX_NUMBERS_FIELD
	default:
		return 0, errors.New("Unknown ContactsPostRequestBody_fax_numbers_field value: " + v)
	}
	return &result, nil
}
func SerializeContactsPostRequestBody_fax_numbers_field(values []ContactsPostRequestBody_fax_numbers_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPostRequestBody_fax_numbers_field) isMultiValue() bool {
	return false
}
