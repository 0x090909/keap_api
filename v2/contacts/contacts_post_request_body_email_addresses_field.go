package contacts

import (
	"errors"
)

type ContactsPostRequestBody_email_addresses_field int

const (
	EMAIL_FIELD_UNSPECIFIED_CONTACTSPOSTREQUESTBODY_EMAIL_ADDRESSES_FIELD ContactsPostRequestBody_email_addresses_field = iota
	EMAIL1_CONTACTSPOSTREQUESTBODY_EMAIL_ADDRESSES_FIELD
	EMAIL2_CONTACTSPOSTREQUESTBODY_EMAIL_ADDRESSES_FIELD
	EMAIL3_CONTACTSPOSTREQUESTBODY_EMAIL_ADDRESSES_FIELD
)

func (i ContactsPostRequestBody_email_addresses_field) String() string {
	return []string{"EMAIL_FIELD_UNSPECIFIED", "EMAIL1", "EMAIL2", "EMAIL3"}[i]
}
func ParseContactsPostRequestBody_email_addresses_field(v string) (any, error) {
	result := EMAIL_FIELD_UNSPECIFIED_CONTACTSPOSTREQUESTBODY_EMAIL_ADDRESSES_FIELD
	switch v {
	case "EMAIL_FIELD_UNSPECIFIED":
		result = EMAIL_FIELD_UNSPECIFIED_CONTACTSPOSTREQUESTBODY_EMAIL_ADDRESSES_FIELD
	case "EMAIL1":
		result = EMAIL1_CONTACTSPOSTREQUESTBODY_EMAIL_ADDRESSES_FIELD
	case "EMAIL2":
		result = EMAIL2_CONTACTSPOSTREQUESTBODY_EMAIL_ADDRESSES_FIELD
	case "EMAIL3":
		result = EMAIL3_CONTACTSPOSTREQUESTBODY_EMAIL_ADDRESSES_FIELD
	default:
		return 0, errors.New("Unknown ContactsPostRequestBody_email_addresses_field value: " + v)
	}
	return &result, nil
}
func SerializeContactsPostRequestBody_email_addresses_field(values []ContactsPostRequestBody_email_addresses_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPostRequestBody_email_addresses_field) isMultiValue() bool {
	return false
}
