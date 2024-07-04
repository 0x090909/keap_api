package item

import (
	"errors"
)

type WithNotePutRequestBody_type int

const (
	APPOINTMENT_WITHNOTEPUTREQUESTBODY_TYPE WithNotePutRequestBody_type = iota
	CALL_WITHNOTEPUTREQUESTBODY_TYPE
	EMAIL_WITHNOTEPUTREQUESTBODY_TYPE
	FAX_WITHNOTEPUTREQUESTBODY_TYPE
	LETTER_WITHNOTEPUTREQUESTBODY_TYPE
	OTHER_WITHNOTEPUTREQUESTBODY_TYPE
)

func (i WithNotePutRequestBody_type) String() string {
	return []string{"Appointment", "Call", "Email", "Fax", "Letter", "Other"}[i]
}
func ParseWithNotePutRequestBody_type(v string) (any, error) {
	result := APPOINTMENT_WITHNOTEPUTREQUESTBODY_TYPE
	switch v {
	case "Appointment":
		result = APPOINTMENT_WITHNOTEPUTREQUESTBODY_TYPE
	case "Call":
		result = CALL_WITHNOTEPUTREQUESTBODY_TYPE
	case "Email":
		result = EMAIL_WITHNOTEPUTREQUESTBODY_TYPE
	case "Fax":
		result = FAX_WITHNOTEPUTREQUESTBODY_TYPE
	case "Letter":
		result = LETTER_WITHNOTEPUTREQUESTBODY_TYPE
	case "Other":
		result = OTHER_WITHNOTEPUTREQUESTBODY_TYPE
	default:
		return 0, errors.New("Unknown WithNotePutRequestBody_type value: " + v)
	}
	return &result, nil
}
func SerializeWithNotePutRequestBody_type(values []WithNotePutRequestBody_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithNotePutRequestBody_type) isMultiValue() bool {
	return false
}
