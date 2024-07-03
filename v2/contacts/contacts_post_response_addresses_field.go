package contacts

import (
	"errors"
)

type ContactsPostResponse_addresses_field int

const (
	ADDRESS_FIELD_UNSPECIFIED_CONTACTSPOSTRESPONSE_ADDRESSES_FIELD ContactsPostResponse_addresses_field = iota
	BILLING_CONTACTSPOSTRESPONSE_ADDRESSES_FIELD
	SHIPPING_CONTACTSPOSTRESPONSE_ADDRESSES_FIELD
	OTHER_CONTACTSPOSTRESPONSE_ADDRESSES_FIELD
)

func (i ContactsPostResponse_addresses_field) String() string {
	return []string{"ADDRESS_FIELD_UNSPECIFIED", "BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseContactsPostResponse_addresses_field(v string) (any, error) {
	result := ADDRESS_FIELD_UNSPECIFIED_CONTACTSPOSTRESPONSE_ADDRESSES_FIELD
	switch v {
	case "ADDRESS_FIELD_UNSPECIFIED":
		result = ADDRESS_FIELD_UNSPECIFIED_CONTACTSPOSTRESPONSE_ADDRESSES_FIELD
	case "BILLING":
		result = BILLING_CONTACTSPOSTRESPONSE_ADDRESSES_FIELD
	case "SHIPPING":
		result = SHIPPING_CONTACTSPOSTRESPONSE_ADDRESSES_FIELD
	case "OTHER":
		result = OTHER_CONTACTSPOSTRESPONSE_ADDRESSES_FIELD
	default:
		return 0, errors.New("Unknown ContactsPostResponse_addresses_field value: " + v)
	}
	return &result, nil
}
func SerializeContactsPostResponse_addresses_field(values []ContactsPostResponse_addresses_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPostResponse_addresses_field) isMultiValue() bool {
	return false
}
