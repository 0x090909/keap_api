package contacts

import (
	"errors"
)

type ContactsGetResponse_contacts_source_type int

const (
	APPOINTMENT_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE ContactsGetResponse_contacts_source_type = iota
	FORMAPIHOSTED_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	FORMAPIINTERNAL_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	WEBFORM_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	INTERNALFORM_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	LANDINGPAGE_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	IMPORT_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	MANUAL_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	API_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	OTHER_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	UNKNOWN_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
)

func (i ContactsGetResponse_contacts_source_type) String() string {
	return []string{"APPOINTMENT", "FORMAPIHOSTED", "FORMAPIINTERNAL", "WEBFORM", "INTERNALFORM", "LANDINGPAGE", "IMPORT", "MANUAL", "API", "OTHER", "UNKNOWN"}[i]
}
func ParseContactsGetResponse_contacts_source_type(v string) (any, error) {
	result := APPOINTMENT_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	switch v {
	case "APPOINTMENT":
		result = APPOINTMENT_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	case "FORMAPIHOSTED":
		result = FORMAPIHOSTED_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	case "FORMAPIINTERNAL":
		result = FORMAPIINTERNAL_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	case "WEBFORM":
		result = WEBFORM_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	case "INTERNALFORM":
		result = INTERNALFORM_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	case "LANDINGPAGE":
		result = LANDINGPAGE_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	case "IMPORT":
		result = IMPORT_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	case "MANUAL":
		result = MANUAL_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	case "API":
		result = API_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	case "OTHER":
		result = OTHER_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	case "UNKNOWN":
		result = UNKNOWN_CONTACTSGETRESPONSE_CONTACTS_SOURCE_TYPE
	default:
		return 0, errors.New("Unknown ContactsGetResponse_contacts_source_type value: " + v)
	}
	return &result, nil
}
func SerializeContactsGetResponse_contacts_source_type(values []ContactsGetResponse_contacts_source_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsGetResponse_contacts_source_type) isMultiValue() bool {
	return false
}
