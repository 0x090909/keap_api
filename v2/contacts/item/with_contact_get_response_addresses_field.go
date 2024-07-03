package item

import (
	"errors"
)

type WithContact_GetResponse_addresses_field int

const (
	ADDRESS_FIELD_UNSPECIFIED_WITHCONTACT_GETRESPONSE_ADDRESSES_FIELD WithContact_GetResponse_addresses_field = iota
	BILLING_WITHCONTACT_GETRESPONSE_ADDRESSES_FIELD
	SHIPPING_WITHCONTACT_GETRESPONSE_ADDRESSES_FIELD
	OTHER_WITHCONTACT_GETRESPONSE_ADDRESSES_FIELD
)

func (i WithContact_GetResponse_addresses_field) String() string {
	return []string{"ADDRESS_FIELD_UNSPECIFIED", "BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseWithContact_GetResponse_addresses_field(v string) (any, error) {
	result := ADDRESS_FIELD_UNSPECIFIED_WITHCONTACT_GETRESPONSE_ADDRESSES_FIELD
	switch v {
	case "ADDRESS_FIELD_UNSPECIFIED":
		result = ADDRESS_FIELD_UNSPECIFIED_WITHCONTACT_GETRESPONSE_ADDRESSES_FIELD
	case "BILLING":
		result = BILLING_WITHCONTACT_GETRESPONSE_ADDRESSES_FIELD
	case "SHIPPING":
		result = SHIPPING_WITHCONTACT_GETRESPONSE_ADDRESSES_FIELD
	case "OTHER":
		result = OTHER_WITHCONTACT_GETRESPONSE_ADDRESSES_FIELD
	default:
		return 0, errors.New("Unknown WithContact_GetResponse_addresses_field value: " + v)
	}
	return &result, nil
}
func SerializeWithContact_GetResponse_addresses_field(values []WithContact_GetResponse_addresses_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithContact_GetResponse_addresses_field) isMultiValue() bool {
	return false
}
