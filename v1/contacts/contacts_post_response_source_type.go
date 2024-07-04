package contacts

import (
	"errors"
)

type ContactsPostResponse_source_type int

const (
	APPOINTMENT_CONTACTSPOSTRESPONSE_SOURCE_TYPE ContactsPostResponse_source_type = iota
	FORMAPIHOSTED_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	FORMAPIINTERNAL_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	WEBFORM_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	INTERNALFORM_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	LANDINGPAGE_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	IMPORT_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	MANUAL_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	API_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	OTHER_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	UNKNOWN_CONTACTSPOSTRESPONSE_SOURCE_TYPE
)

func (i ContactsPostResponse_source_type) String() string {
	return []string{"APPOINTMENT", "FORMAPIHOSTED", "FORMAPIINTERNAL", "WEBFORM", "INTERNALFORM", "LANDINGPAGE", "IMPORT", "MANUAL", "API", "OTHER", "UNKNOWN"}[i]
}
func ParseContactsPostResponse_source_type(v string) (any, error) {
	result := APPOINTMENT_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	switch v {
	case "APPOINTMENT":
		result = APPOINTMENT_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	case "FORMAPIHOSTED":
		result = FORMAPIHOSTED_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	case "FORMAPIINTERNAL":
		result = FORMAPIINTERNAL_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	case "WEBFORM":
		result = WEBFORM_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	case "INTERNALFORM":
		result = INTERNALFORM_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	case "LANDINGPAGE":
		result = LANDINGPAGE_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	case "IMPORT":
		result = IMPORT_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	case "MANUAL":
		result = MANUAL_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	case "API":
		result = API_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	case "OTHER":
		result = OTHER_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	case "UNKNOWN":
		result = UNKNOWN_CONTACTSPOSTRESPONSE_SOURCE_TYPE
	default:
		return 0, errors.New("Unknown ContactsPostResponse_source_type value: " + v)
	}
	return &result, nil
}
func SerializeContactsPostResponse_source_type(values []ContactsPostResponse_source_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPostResponse_source_type) isMultiValue() bool {
	return false
}
