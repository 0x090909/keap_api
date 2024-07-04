package item

import (
	"errors"
)

type ContactIdPatchResponse_email_addresses_field int

const (
	EMAIL1_CONTACTIDPATCHRESPONSE_EMAIL_ADDRESSES_FIELD ContactIdPatchResponse_email_addresses_field = iota
	EMAIL2_CONTACTIDPATCHRESPONSE_EMAIL_ADDRESSES_FIELD
	EMAIL3_CONTACTIDPATCHRESPONSE_EMAIL_ADDRESSES_FIELD
)

func (i ContactIdPatchResponse_email_addresses_field) String() string {
	return []string{"EMAIL1", "EMAIL2", "EMAIL3"}[i]
}
func ParseContactIdPatchResponse_email_addresses_field(v string) (any, error) {
	result := EMAIL1_CONTACTIDPATCHRESPONSE_EMAIL_ADDRESSES_FIELD
	switch v {
	case "EMAIL1":
		result = EMAIL1_CONTACTIDPATCHRESPONSE_EMAIL_ADDRESSES_FIELD
	case "EMAIL2":
		result = EMAIL2_CONTACTIDPATCHRESPONSE_EMAIL_ADDRESSES_FIELD
	case "EMAIL3":
		result = EMAIL3_CONTACTIDPATCHRESPONSE_EMAIL_ADDRESSES_FIELD
	default:
		return 0, errors.New("Unknown ContactIdPatchResponse_email_addresses_field value: " + v)
	}
	return &result, nil
}
func SerializeContactIdPatchResponse_email_addresses_field(values []ContactIdPatchResponse_email_addresses_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactIdPatchResponse_email_addresses_field) isMultiValue() bool {
	return false
}
