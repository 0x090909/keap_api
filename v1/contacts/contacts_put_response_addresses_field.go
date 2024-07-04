package contacts

import (
	"errors"
)

type ContactsPutResponse_addresses_field int

const (
	BILLING_CONTACTSPUTRESPONSE_ADDRESSES_FIELD ContactsPutResponse_addresses_field = iota
	SHIPPING_CONTACTSPUTRESPONSE_ADDRESSES_FIELD
	OTHER_CONTACTSPUTRESPONSE_ADDRESSES_FIELD
)

func (i ContactsPutResponse_addresses_field) String() string {
	return []string{"BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseContactsPutResponse_addresses_field(v string) (any, error) {
	result := BILLING_CONTACTSPUTRESPONSE_ADDRESSES_FIELD
	switch v {
	case "BILLING":
		result = BILLING_CONTACTSPUTRESPONSE_ADDRESSES_FIELD
	case "SHIPPING":
		result = SHIPPING_CONTACTSPUTRESPONSE_ADDRESSES_FIELD
	case "OTHER":
		result = OTHER_CONTACTSPUTRESPONSE_ADDRESSES_FIELD
	default:
		return 0, errors.New("Unknown ContactsPutResponse_addresses_field value: " + v)
	}
	return &result, nil
}
func SerializeContactsPutResponse_addresses_field(values []ContactsPutResponse_addresses_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPutResponse_addresses_field) isMultiValue() bool {
	return false
}
