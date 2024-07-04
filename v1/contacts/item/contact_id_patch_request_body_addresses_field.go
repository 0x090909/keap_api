package item

import (
	"errors"
)

type ContactIdPatchRequestBody_addresses_field int

const (
	BILLING_CONTACTIDPATCHREQUESTBODY_ADDRESSES_FIELD ContactIdPatchRequestBody_addresses_field = iota
	SHIPPING_CONTACTIDPATCHREQUESTBODY_ADDRESSES_FIELD
	OTHER_CONTACTIDPATCHREQUESTBODY_ADDRESSES_FIELD
)

func (i ContactIdPatchRequestBody_addresses_field) String() string {
	return []string{"BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseContactIdPatchRequestBody_addresses_field(v string) (any, error) {
	result := BILLING_CONTACTIDPATCHREQUESTBODY_ADDRESSES_FIELD
	switch v {
	case "BILLING":
		result = BILLING_CONTACTIDPATCHREQUESTBODY_ADDRESSES_FIELD
	case "SHIPPING":
		result = SHIPPING_CONTACTIDPATCHREQUESTBODY_ADDRESSES_FIELD
	case "OTHER":
		result = OTHER_CONTACTIDPATCHREQUESTBODY_ADDRESSES_FIELD
	default:
		return 0, errors.New("Unknown ContactIdPatchRequestBody_addresses_field value: " + v)
	}
	return &result, nil
}
func SerializeContactIdPatchRequestBody_addresses_field(values []ContactIdPatchRequestBody_addresses_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactIdPatchRequestBody_addresses_field) isMultiValue() bool {
	return false
}
