package contacts

import (
	"errors"
)

type ContactsGetResponse_contacts_fax_numbers_field int

const (
	FAX_NUMBER_FIELD_UNSPECIFIED_CONTACTSGETRESPONSE_CONTACTS_FAX_NUMBERS_FIELD ContactsGetResponse_contacts_fax_numbers_field = iota
	FAX1_CONTACTSGETRESPONSE_CONTACTS_FAX_NUMBERS_FIELD
	FAX2_CONTACTSGETRESPONSE_CONTACTS_FAX_NUMBERS_FIELD
)

func (i ContactsGetResponse_contacts_fax_numbers_field) String() string {
	return []string{"FAX_NUMBER_FIELD_UNSPECIFIED", "FAX1", "FAX2"}[i]
}
func ParseContactsGetResponse_contacts_fax_numbers_field(v string) (any, error) {
	result := FAX_NUMBER_FIELD_UNSPECIFIED_CONTACTSGETRESPONSE_CONTACTS_FAX_NUMBERS_FIELD
	switch v {
	case "FAX_NUMBER_FIELD_UNSPECIFIED":
		result = FAX_NUMBER_FIELD_UNSPECIFIED_CONTACTSGETRESPONSE_CONTACTS_FAX_NUMBERS_FIELD
	case "FAX1":
		result = FAX1_CONTACTSGETRESPONSE_CONTACTS_FAX_NUMBERS_FIELD
	case "FAX2":
		result = FAX2_CONTACTSGETRESPONSE_CONTACTS_FAX_NUMBERS_FIELD
	default:
		return 0, errors.New("Unknown ContactsGetResponse_contacts_fax_numbers_field value: " + v)
	}
	return &result, nil
}
func SerializeContactsGetResponse_contacts_fax_numbers_field(values []ContactsGetResponse_contacts_fax_numbers_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsGetResponse_contacts_fax_numbers_field) isMultiValue() bool {
	return false
}
