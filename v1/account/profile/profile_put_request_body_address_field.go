package profile

import (
	"errors"
)

type ProfilePutRequestBody_address_field int

const (
	BILLING_PROFILEPUTREQUESTBODY_ADDRESS_FIELD ProfilePutRequestBody_address_field = iota
	SHIPPING_PROFILEPUTREQUESTBODY_ADDRESS_FIELD
	OTHER_PROFILEPUTREQUESTBODY_ADDRESS_FIELD
)

func (i ProfilePutRequestBody_address_field) String() string {
	return []string{"BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseProfilePutRequestBody_address_field(v string) (any, error) {
	result := BILLING_PROFILEPUTREQUESTBODY_ADDRESS_FIELD
	switch v {
	case "BILLING":
		result = BILLING_PROFILEPUTREQUESTBODY_ADDRESS_FIELD
	case "SHIPPING":
		result = SHIPPING_PROFILEPUTREQUESTBODY_ADDRESS_FIELD
	case "OTHER":
		result = OTHER_PROFILEPUTREQUESTBODY_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown ProfilePutRequestBody_address_field value: " + v)
	}
	return &result, nil
}
func SerializeProfilePutRequestBody_address_field(values []ProfilePutRequestBody_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ProfilePutRequestBody_address_field) isMultiValue() bool {
	return false
}
