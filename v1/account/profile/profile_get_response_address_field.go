package profile

import (
	"errors"
)

type ProfileGetResponse_address_field int

const (
	BILLING_PROFILEGETRESPONSE_ADDRESS_FIELD ProfileGetResponse_address_field = iota
	SHIPPING_PROFILEGETRESPONSE_ADDRESS_FIELD
	OTHER_PROFILEGETRESPONSE_ADDRESS_FIELD
)

func (i ProfileGetResponse_address_field) String() string {
	return []string{"BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseProfileGetResponse_address_field(v string) (any, error) {
	result := BILLING_PROFILEGETRESPONSE_ADDRESS_FIELD
	switch v {
	case "BILLING":
		result = BILLING_PROFILEGETRESPONSE_ADDRESS_FIELD
	case "SHIPPING":
		result = SHIPPING_PROFILEGETRESPONSE_ADDRESS_FIELD
	case "OTHER":
		result = OTHER_PROFILEGETRESPONSE_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown ProfileGetResponse_address_field value: " + v)
	}
	return &result, nil
}
func SerializeProfileGetResponse_address_field(values []ProfileGetResponse_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ProfileGetResponse_address_field) isMultiValue() bool {
	return false
}
