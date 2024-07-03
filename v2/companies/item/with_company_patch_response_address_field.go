package item

import (
	"errors"
)

type WithCompany_PatchResponse_address_field int

const (
	ADDRESS_FIELD_UNSPECIFIED_WITHCOMPANY_PATCHRESPONSE_ADDRESS_FIELD WithCompany_PatchResponse_address_field = iota
	BILLING_WITHCOMPANY_PATCHRESPONSE_ADDRESS_FIELD
	SHIPPING_WITHCOMPANY_PATCHRESPONSE_ADDRESS_FIELD
	OTHER_WITHCOMPANY_PATCHRESPONSE_ADDRESS_FIELD
)

func (i WithCompany_PatchResponse_address_field) String() string {
	return []string{"ADDRESS_FIELD_UNSPECIFIED", "BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseWithCompany_PatchResponse_address_field(v string) (any, error) {
	result := ADDRESS_FIELD_UNSPECIFIED_WITHCOMPANY_PATCHRESPONSE_ADDRESS_FIELD
	switch v {
	case "ADDRESS_FIELD_UNSPECIFIED":
		result = ADDRESS_FIELD_UNSPECIFIED_WITHCOMPANY_PATCHRESPONSE_ADDRESS_FIELD
	case "BILLING":
		result = BILLING_WITHCOMPANY_PATCHRESPONSE_ADDRESS_FIELD
	case "SHIPPING":
		result = SHIPPING_WITHCOMPANY_PATCHRESPONSE_ADDRESS_FIELD
	case "OTHER":
		result = OTHER_WITHCOMPANY_PATCHRESPONSE_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown WithCompany_PatchResponse_address_field value: " + v)
	}
	return &result, nil
}
func SerializeWithCompany_PatchResponse_address_field(values []WithCompany_PatchResponse_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithCompany_PatchResponse_address_field) isMultiValue() bool {
	return false
}
