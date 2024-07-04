package item

import (
	"errors"
)

type ContactIdPatchResponse_addresses_field int

const (
	BILLING_CONTACTIDPATCHRESPONSE_ADDRESSES_FIELD ContactIdPatchResponse_addresses_field = iota
	SHIPPING_CONTACTIDPATCHRESPONSE_ADDRESSES_FIELD
	OTHER_CONTACTIDPATCHRESPONSE_ADDRESSES_FIELD
)

func (i ContactIdPatchResponse_addresses_field) String() string {
	return []string{"BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseContactIdPatchResponse_addresses_field(v string) (any, error) {
	result := BILLING_CONTACTIDPATCHRESPONSE_ADDRESSES_FIELD
	switch v {
	case "BILLING":
		result = BILLING_CONTACTIDPATCHRESPONSE_ADDRESSES_FIELD
	case "SHIPPING":
		result = SHIPPING_CONTACTIDPATCHRESPONSE_ADDRESSES_FIELD
	case "OTHER":
		result = OTHER_CONTACTIDPATCHRESPONSE_ADDRESSES_FIELD
	default:
		return 0, errors.New("Unknown ContactIdPatchResponse_addresses_field value: " + v)
	}
	return &result, nil
}
func SerializeContactIdPatchResponse_addresses_field(values []ContactIdPatchResponse_addresses_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactIdPatchResponse_addresses_field) isMultiValue() bool {
	return false
}
