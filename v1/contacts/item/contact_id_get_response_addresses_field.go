package item

import (
	"errors"
)

type ContactIdGetResponse_addresses_field int

const (
	BILLING_CONTACTIDGETRESPONSE_ADDRESSES_FIELD ContactIdGetResponse_addresses_field = iota
	SHIPPING_CONTACTIDGETRESPONSE_ADDRESSES_FIELD
	OTHER_CONTACTIDGETRESPONSE_ADDRESSES_FIELD
)

func (i ContactIdGetResponse_addresses_field) String() string {
	return []string{"BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseContactIdGetResponse_addresses_field(v string) (any, error) {
	result := BILLING_CONTACTIDGETRESPONSE_ADDRESSES_FIELD
	switch v {
	case "BILLING":
		result = BILLING_CONTACTIDGETRESPONSE_ADDRESSES_FIELD
	case "SHIPPING":
		result = SHIPPING_CONTACTIDGETRESPONSE_ADDRESSES_FIELD
	case "OTHER":
		result = OTHER_CONTACTIDGETRESPONSE_ADDRESSES_FIELD
	default:
		return 0, errors.New("Unknown ContactIdGetResponse_addresses_field value: " + v)
	}
	return &result, nil
}
func SerializeContactIdGetResponse_addresses_field(values []ContactIdGetResponse_addresses_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactIdGetResponse_addresses_field) isMultiValue() bool {
	return false
}
