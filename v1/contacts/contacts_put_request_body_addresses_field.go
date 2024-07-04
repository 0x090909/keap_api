package contacts

import (
	"errors"
)

type ContactsPutRequestBody_addresses_field int

const (
	BILLING_CONTACTSPUTREQUESTBODY_ADDRESSES_FIELD ContactsPutRequestBody_addresses_field = iota
	SHIPPING_CONTACTSPUTREQUESTBODY_ADDRESSES_FIELD
	OTHER_CONTACTSPUTREQUESTBODY_ADDRESSES_FIELD
)

func (i ContactsPutRequestBody_addresses_field) String() string {
	return []string{"BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseContactsPutRequestBody_addresses_field(v string) (any, error) {
	result := BILLING_CONTACTSPUTREQUESTBODY_ADDRESSES_FIELD
	switch v {
	case "BILLING":
		result = BILLING_CONTACTSPUTREQUESTBODY_ADDRESSES_FIELD
	case "SHIPPING":
		result = SHIPPING_CONTACTSPUTREQUESTBODY_ADDRESSES_FIELD
	case "OTHER":
		result = OTHER_CONTACTSPUTREQUESTBODY_ADDRESSES_FIELD
	default:
		return 0, errors.New("Unknown ContactsPutRequestBody_addresses_field value: " + v)
	}
	return &result, nil
}
func SerializeContactsPutRequestBody_addresses_field(values []ContactsPutRequestBody_addresses_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPutRequestBody_addresses_field) isMultiValue() bool {
	return false
}
