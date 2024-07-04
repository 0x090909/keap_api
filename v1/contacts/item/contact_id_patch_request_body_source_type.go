package item

import (
	"errors"
)

type ContactIdPatchRequestBody_source_type int

const (
	APPOINTMENT_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE ContactIdPatchRequestBody_source_type = iota
	FORMAPIHOSTED_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	FORMAPIINTERNAL_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	WEBFORM_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	INTERNALFORM_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	LANDINGPAGE_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	IMPORT_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	MANUAL_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	API_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	OTHER_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	UNKNOWN_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
)

func (i ContactIdPatchRequestBody_source_type) String() string {
	return []string{"APPOINTMENT", "FORMAPIHOSTED", "FORMAPIINTERNAL", "WEBFORM", "INTERNALFORM", "LANDINGPAGE", "IMPORT", "MANUAL", "API", "OTHER", "UNKNOWN"}[i]
}
func ParseContactIdPatchRequestBody_source_type(v string) (any, error) {
	result := APPOINTMENT_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	switch v {
	case "APPOINTMENT":
		result = APPOINTMENT_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	case "FORMAPIHOSTED":
		result = FORMAPIHOSTED_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	case "FORMAPIINTERNAL":
		result = FORMAPIINTERNAL_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	case "WEBFORM":
		result = WEBFORM_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	case "INTERNALFORM":
		result = INTERNALFORM_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	case "LANDINGPAGE":
		result = LANDINGPAGE_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	case "IMPORT":
		result = IMPORT_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	case "MANUAL":
		result = MANUAL_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	case "API":
		result = API_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	case "OTHER":
		result = OTHER_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	case "UNKNOWN":
		result = UNKNOWN_CONTACTIDPATCHREQUESTBODY_SOURCE_TYPE
	default:
		return 0, errors.New("Unknown ContactIdPatchRequestBody_source_type value: " + v)
	}
	return &result, nil
}
func SerializeContactIdPatchRequestBody_source_type(values []ContactIdPatchRequestBody_source_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactIdPatchRequestBody_source_type) isMultiValue() bool {
	return false
}