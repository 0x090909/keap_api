package contacts

import (
	"errors"
)

type ContactsGetResponse_contacts_addresses_field int

const (
	BILLING_CONTACTSGETRESPONSE_CONTACTS_ADDRESSES_FIELD ContactsGetResponse_contacts_addresses_field = iota
	SHIPPING_CONTACTSGETRESPONSE_CONTACTS_ADDRESSES_FIELD
	OTHER_CONTACTSGETRESPONSE_CONTACTS_ADDRESSES_FIELD
)

func (i ContactsGetResponse_contacts_addresses_field) String() string {
	return []string{"BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseContactsGetResponse_contacts_addresses_field(v string) (any, error) {
	result := BILLING_CONTACTSGETRESPONSE_CONTACTS_ADDRESSES_FIELD
	switch v {
	case "BILLING":
		result = BILLING_CONTACTSGETRESPONSE_CONTACTS_ADDRESSES_FIELD
	case "SHIPPING":
		result = SHIPPING_CONTACTSGETRESPONSE_CONTACTS_ADDRESSES_FIELD
	case "OTHER":
		result = OTHER_CONTACTSGETRESPONSE_CONTACTS_ADDRESSES_FIELD
	default:
		return 0, errors.New("Unknown ContactsGetResponse_contacts_addresses_field value: " + v)
	}
	return &result, nil
}
func SerializeContactsGetResponse_contacts_addresses_field(values []ContactsGetResponse_contacts_addresses_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsGetResponse_contacts_addresses_field) isMultiValue() bool {
	return false
}
