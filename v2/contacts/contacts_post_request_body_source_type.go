package contacts

import (
	"errors"
)

type ContactsPostRequestBody_source_type int

const (
	SOURCE_TYPE_UNSPECIFIED_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE ContactsPostRequestBody_source_type = iota
	API_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	APPOINTMENT_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	FORM_API_HOSTED_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	FORM_API_INTERNAL_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	IMPORT_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	INTERNAL_FORM_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	LANDING_PAGE_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	MANUAL_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	OTHER_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	UNKNOWN_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	WEBFORM_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
)

func (i ContactsPostRequestBody_source_type) String() string {
	return []string{"SOURCE_TYPE_UNSPECIFIED", "API", "APPOINTMENT", "FORM_API_HOSTED", "FORM_API_INTERNAL", "IMPORT", "INTERNAL_FORM", "LANDING_PAGE", "MANUAL", "OTHER", "UNKNOWN", "WEBFORM"}[i]
}
func ParseContactsPostRequestBody_source_type(v string) (any, error) {
	result := SOURCE_TYPE_UNSPECIFIED_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	switch v {
	case "SOURCE_TYPE_UNSPECIFIED":
		result = SOURCE_TYPE_UNSPECIFIED_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "API":
		result = API_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "APPOINTMENT":
		result = APPOINTMENT_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "FORM_API_HOSTED":
		result = FORM_API_HOSTED_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "FORM_API_INTERNAL":
		result = FORM_API_INTERNAL_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "IMPORT":
		result = IMPORT_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "INTERNAL_FORM":
		result = INTERNAL_FORM_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "LANDING_PAGE":
		result = LANDING_PAGE_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "MANUAL":
		result = MANUAL_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "OTHER":
		result = OTHER_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "UNKNOWN":
		result = UNKNOWN_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "WEBFORM":
		result = WEBFORM_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	default:
		return 0, errors.New("Unknown ContactsPostRequestBody_source_type value: " + v)
	}
	return &result, nil
}
func SerializeContactsPostRequestBody_source_type(values []ContactsPostRequestBody_source_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPostRequestBody_source_type) isMultiValue() bool {
	return false
}
