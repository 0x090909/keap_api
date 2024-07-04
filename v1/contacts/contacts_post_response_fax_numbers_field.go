package contacts

import (
	"errors"
)

type ContactsPostResponse_fax_numbers_field int

const (
	FAX1_CONTACTSPOSTRESPONSE_FAX_NUMBERS_FIELD ContactsPostResponse_fax_numbers_field = iota
	FAX2_CONTACTSPOSTRESPONSE_FAX_NUMBERS_FIELD
)

func (i ContactsPostResponse_fax_numbers_field) String() string {
	return []string{"FAX1", "FAX2"}[i]
}
func ParseContactsPostResponse_fax_numbers_field(v string) (any, error) {
	result := FAX1_CONTACTSPOSTRESPONSE_FAX_NUMBERS_FIELD
	switch v {
	case "FAX1":
		result = FAX1_CONTACTSPOSTRESPONSE_FAX_NUMBERS_FIELD
	case "FAX2":
		result = FAX2_CONTACTSPOSTRESPONSE_FAX_NUMBERS_FIELD
	default:
		return 0, errors.New("Unknown ContactsPostResponse_fax_numbers_field value: " + v)
	}
	return &result, nil
}
func SerializeContactsPostResponse_fax_numbers_field(values []ContactsPostResponse_fax_numbers_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPostResponse_fax_numbers_field) isMultiValue() bool {
	return false
}
