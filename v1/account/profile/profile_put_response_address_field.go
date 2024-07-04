package profile

import (
	"errors"
)

type ProfilePutResponse_address_field int

const (
	BILLING_PROFILEPUTRESPONSE_ADDRESS_FIELD ProfilePutResponse_address_field = iota
	SHIPPING_PROFILEPUTRESPONSE_ADDRESS_FIELD
	OTHER_PROFILEPUTRESPONSE_ADDRESS_FIELD
)

func (i ProfilePutResponse_address_field) String() string {
	return []string{"BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseProfilePutResponse_address_field(v string) (any, error) {
	result := BILLING_PROFILEPUTRESPONSE_ADDRESS_FIELD
	switch v {
	case "BILLING":
		result = BILLING_PROFILEPUTRESPONSE_ADDRESS_FIELD
	case "SHIPPING":
		result = SHIPPING_PROFILEPUTRESPONSE_ADDRESS_FIELD
	case "OTHER":
		result = OTHER_PROFILEPUTRESPONSE_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown ProfilePutResponse_address_field value: " + v)
	}
	return &result, nil
}
func SerializeProfilePutResponse_address_field(values []ProfilePutResponse_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ProfilePutResponse_address_field) isMultiValue() bool {
	return false
}
