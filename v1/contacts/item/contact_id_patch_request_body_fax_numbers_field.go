package item

import (
	"errors"
)

type ContactIdPatchRequestBody_fax_numbers_field int

const (
	FAX1_CONTACTIDPATCHREQUESTBODY_FAX_NUMBERS_FIELD ContactIdPatchRequestBody_fax_numbers_field = iota
	FAX2_CONTACTIDPATCHREQUESTBODY_FAX_NUMBERS_FIELD
)

func (i ContactIdPatchRequestBody_fax_numbers_field) String() string {
	return []string{"FAX1", "FAX2"}[i]
}
func ParseContactIdPatchRequestBody_fax_numbers_field(v string) (any, error) {
	result := FAX1_CONTACTIDPATCHREQUESTBODY_FAX_NUMBERS_FIELD
	switch v {
	case "FAX1":
		result = FAX1_CONTACTIDPATCHREQUESTBODY_FAX_NUMBERS_FIELD
	case "FAX2":
		result = FAX2_CONTACTIDPATCHREQUESTBODY_FAX_NUMBERS_FIELD
	default:
		return 0, errors.New("Unknown ContactIdPatchRequestBody_fax_numbers_field value: " + v)
	}
	return &result, nil
}
func SerializeContactIdPatchRequestBody_fax_numbers_field(values []ContactIdPatchRequestBody_fax_numbers_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactIdPatchRequestBody_fax_numbers_field) isMultiValue() bool {
	return false
}