package item

import (
	"errors"
)

type WithContact_GetResponse_email_addresses_field int

const (
	EMAIL_FIELD_UNSPECIFIED_WITHCONTACT_GETRESPONSE_EMAIL_ADDRESSES_FIELD WithContact_GetResponse_email_addresses_field = iota
	EMAIL1_WITHCONTACT_GETRESPONSE_EMAIL_ADDRESSES_FIELD
	EMAIL2_WITHCONTACT_GETRESPONSE_EMAIL_ADDRESSES_FIELD
	EMAIL3_WITHCONTACT_GETRESPONSE_EMAIL_ADDRESSES_FIELD
)

func (i WithContact_GetResponse_email_addresses_field) String() string {
	return []string{"EMAIL_FIELD_UNSPECIFIED", "EMAIL1", "EMAIL2", "EMAIL3"}[i]
}
func ParseWithContact_GetResponse_email_addresses_field(v string) (any, error) {
	result := EMAIL_FIELD_UNSPECIFIED_WITHCONTACT_GETRESPONSE_EMAIL_ADDRESSES_FIELD
	switch v {
	case "EMAIL_FIELD_UNSPECIFIED":
		result = EMAIL_FIELD_UNSPECIFIED_WITHCONTACT_GETRESPONSE_EMAIL_ADDRESSES_FIELD
	case "EMAIL1":
		result = EMAIL1_WITHCONTACT_GETRESPONSE_EMAIL_ADDRESSES_FIELD
	case "EMAIL2":
		result = EMAIL2_WITHCONTACT_GETRESPONSE_EMAIL_ADDRESSES_FIELD
	case "EMAIL3":
		result = EMAIL3_WITHCONTACT_GETRESPONSE_EMAIL_ADDRESSES_FIELD
	default:
		return 0, errors.New("Unknown WithContact_GetResponse_email_addresses_field value: " + v)
	}
	return &result, nil
}
func SerializeWithContact_GetResponse_email_addresses_field(values []WithContact_GetResponse_email_addresses_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithContact_GetResponse_email_addresses_field) isMultiValue() bool {
	return false
}
