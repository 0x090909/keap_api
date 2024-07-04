package contacts

import (
	"errors"
)

type ContactsPostRequestBody_source_type int

const (
	APPOINTMENT_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE ContactsPostRequestBody_source_type = iota
	FORMAPIHOSTED_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	FORMAPIINTERNAL_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	WEBFORM_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	INTERNALFORM_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	LANDINGPAGE_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	IMPORT_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	MANUAL_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	API_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	OTHER_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	UNKNOWN_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
)

func (i ContactsPostRequestBody_source_type) String() string {
	return []string{"APPOINTMENT", "FORMAPIHOSTED", "FORMAPIINTERNAL", "WEBFORM", "INTERNALFORM", "LANDINGPAGE", "IMPORT", "MANUAL", "API", "OTHER", "UNKNOWN"}[i]
}
func ParseContactsPostRequestBody_source_type(v string) (any, error) {
	result := APPOINTMENT_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	switch v {
	case "APPOINTMENT":
		result = APPOINTMENT_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "FORMAPIHOSTED":
		result = FORMAPIHOSTED_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "FORMAPIINTERNAL":
		result = FORMAPIINTERNAL_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "WEBFORM":
		result = WEBFORM_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "INTERNALFORM":
		result = INTERNALFORM_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "LANDINGPAGE":
		result = LANDINGPAGE_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "IMPORT":
		result = IMPORT_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "MANUAL":
		result = MANUAL_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "API":
		result = API_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "OTHER":
		result = OTHER_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
	case "UNKNOWN":
		result = UNKNOWN_CONTACTSPOSTREQUESTBODY_SOURCE_TYPE
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