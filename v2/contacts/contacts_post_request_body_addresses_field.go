package contacts

import (
	"errors"
)

type ContactsPostRequestBody_addresses_field int

const (
	ADDRESS_FIELD_UNSPECIFIED_CONTACTSPOSTREQUESTBODY_ADDRESSES_FIELD ContactsPostRequestBody_addresses_field = iota
	BILLING_CONTACTSPOSTREQUESTBODY_ADDRESSES_FIELD
	SHIPPING_CONTACTSPOSTREQUESTBODY_ADDRESSES_FIELD
	OTHER_CONTACTSPOSTREQUESTBODY_ADDRESSES_FIELD
)

func (i ContactsPostRequestBody_addresses_field) String() string {
	return []string{"ADDRESS_FIELD_UNSPECIFIED", "BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseContactsPostRequestBody_addresses_field(v string) (any, error) {
	result := ADDRESS_FIELD_UNSPECIFIED_CONTACTSPOSTREQUESTBODY_ADDRESSES_FIELD
	switch v {
	case "ADDRESS_FIELD_UNSPECIFIED":
		result = ADDRESS_FIELD_UNSPECIFIED_CONTACTSPOSTREQUESTBODY_ADDRESSES_FIELD
	case "BILLING":
		result = BILLING_CONTACTSPOSTREQUESTBODY_ADDRESSES_FIELD
	case "SHIPPING":
		result = SHIPPING_CONTACTSPOSTREQUESTBODY_ADDRESSES_FIELD
	case "OTHER":
		result = OTHER_CONTACTSPOSTREQUESTBODY_ADDRESSES_FIELD
	default:
		return 0, errors.New("Unknown ContactsPostRequestBody_addresses_field value: " + v)
	}
	return &result, nil
}
func SerializeContactsPostRequestBody_addresses_field(values []ContactsPostRequestBody_addresses_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPostRequestBody_addresses_field) isMultiValue() bool {
	return false
}
