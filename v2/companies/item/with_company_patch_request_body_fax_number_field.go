package item

import (
	"errors"
)

type WithCompany_PatchRequestBody_fax_number_field int

const (
	FAX_NUMBER_FIELD_UNSPECIFIED_WITHCOMPANY_PATCHREQUESTBODY_FAX_NUMBER_FIELD WithCompany_PatchRequestBody_fax_number_field = iota
	FAX1_WITHCOMPANY_PATCHREQUESTBODY_FAX_NUMBER_FIELD
	FAX2_WITHCOMPANY_PATCHREQUESTBODY_FAX_NUMBER_FIELD
)

func (i WithCompany_PatchRequestBody_fax_number_field) String() string {
	return []string{"FAX_NUMBER_FIELD_UNSPECIFIED", "FAX1", "FAX2"}[i]
}
func ParseWithCompany_PatchRequestBody_fax_number_field(v string) (any, error) {
	result := FAX_NUMBER_FIELD_UNSPECIFIED_WITHCOMPANY_PATCHREQUESTBODY_FAX_NUMBER_FIELD
	switch v {
	case "FAX_NUMBER_FIELD_UNSPECIFIED":
		result = FAX_NUMBER_FIELD_UNSPECIFIED_WITHCOMPANY_PATCHREQUESTBODY_FAX_NUMBER_FIELD
	case "FAX1":
		result = FAX1_WITHCOMPANY_PATCHREQUESTBODY_FAX_NUMBER_FIELD
	case "FAX2":
		result = FAX2_WITHCOMPANY_PATCHREQUESTBODY_FAX_NUMBER_FIELD
	default:
		return 0, errors.New("Unknown WithCompany_PatchRequestBody_fax_number_field value: " + v)
	}
	return &result, nil
}
func SerializeWithCompany_PatchRequestBody_fax_number_field(values []WithCompany_PatchRequestBody_fax_number_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithCompany_PatchRequestBody_fax_number_field) isMultiValue() bool {
	return false
}
