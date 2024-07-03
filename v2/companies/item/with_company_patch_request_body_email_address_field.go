package item

import (
	"errors"
)

type WithCompany_PatchRequestBody_email_address_field int

const (
	EMAIL_FIELD_UNSPECIFIED_WITHCOMPANY_PATCHREQUESTBODY_EMAIL_ADDRESS_FIELD WithCompany_PatchRequestBody_email_address_field = iota
	EMAIL1_WITHCOMPANY_PATCHREQUESTBODY_EMAIL_ADDRESS_FIELD
	EMAIL2_WITHCOMPANY_PATCHREQUESTBODY_EMAIL_ADDRESS_FIELD
	EMAIL3_WITHCOMPANY_PATCHREQUESTBODY_EMAIL_ADDRESS_FIELD
)

func (i WithCompany_PatchRequestBody_email_address_field) String() string {
	return []string{"EMAIL_FIELD_UNSPECIFIED", "EMAIL1", "EMAIL2", "EMAIL3"}[i]
}
func ParseWithCompany_PatchRequestBody_email_address_field(v string) (any, error) {
	result := EMAIL_FIELD_UNSPECIFIED_WITHCOMPANY_PATCHREQUESTBODY_EMAIL_ADDRESS_FIELD
	switch v {
	case "EMAIL_FIELD_UNSPECIFIED":
		result = EMAIL_FIELD_UNSPECIFIED_WITHCOMPANY_PATCHREQUESTBODY_EMAIL_ADDRESS_FIELD
	case "EMAIL1":
		result = EMAIL1_WITHCOMPANY_PATCHREQUESTBODY_EMAIL_ADDRESS_FIELD
	case "EMAIL2":
		result = EMAIL2_WITHCOMPANY_PATCHREQUESTBODY_EMAIL_ADDRESS_FIELD
	case "EMAIL3":
		result = EMAIL3_WITHCOMPANY_PATCHREQUESTBODY_EMAIL_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown WithCompany_PatchRequestBody_email_address_field value: " + v)
	}
	return &result, nil
}
func SerializeWithCompany_PatchRequestBody_email_address_field(values []WithCompany_PatchRequestBody_email_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithCompany_PatchRequestBody_email_address_field) isMultiValue() bool {
	return false
}
