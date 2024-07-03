package item

import (
	"errors"
)

type WithContact_GetResponse_source_type int

const (
	SOURCE_TYPE_UNSPECIFIED_WITHCONTACT_GETRESPONSE_SOURCE_TYPE WithContact_GetResponse_source_type = iota
	API_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	APPOINTMENT_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	FORM_API_HOSTED_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	FORM_API_INTERNAL_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	IMPORT_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	INTERNAL_FORM_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	LANDING_PAGE_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	MANUAL_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	OTHER_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	UNKNOWN_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	WEBFORM_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
)

func (i WithContact_GetResponse_source_type) String() string {
	return []string{"SOURCE_TYPE_UNSPECIFIED", "API", "APPOINTMENT", "FORM_API_HOSTED", "FORM_API_INTERNAL", "IMPORT", "INTERNAL_FORM", "LANDING_PAGE", "MANUAL", "OTHER", "UNKNOWN", "WEBFORM"}[i]
}
func ParseWithContact_GetResponse_source_type(v string) (any, error) {
	result := SOURCE_TYPE_UNSPECIFIED_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	switch v {
	case "SOURCE_TYPE_UNSPECIFIED":
		result = SOURCE_TYPE_UNSPECIFIED_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	case "API":
		result = API_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	case "APPOINTMENT":
		result = APPOINTMENT_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	case "FORM_API_HOSTED":
		result = FORM_API_HOSTED_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	case "FORM_API_INTERNAL":
		result = FORM_API_INTERNAL_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	case "IMPORT":
		result = IMPORT_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	case "INTERNAL_FORM":
		result = INTERNAL_FORM_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	case "LANDING_PAGE":
		result = LANDING_PAGE_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	case "MANUAL":
		result = MANUAL_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	case "OTHER":
		result = OTHER_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	case "UNKNOWN":
		result = UNKNOWN_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	case "WEBFORM":
		result = WEBFORM_WITHCONTACT_GETRESPONSE_SOURCE_TYPE
	default:
		return 0, errors.New("Unknown WithContact_GetResponse_source_type value: " + v)
	}
	return &result, nil
}
func SerializeWithContact_GetResponse_source_type(values []WithContact_GetResponse_source_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithContact_GetResponse_source_type) isMultiValue() bool {
	return false
}
