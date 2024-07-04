package item

import (
	"errors"
)

type WithNoteGetResponse_type int

const (
	APPOINTMENT_WITHNOTEGETRESPONSE_TYPE WithNoteGetResponse_type = iota
	CALL_WITHNOTEGETRESPONSE_TYPE
	EMAIL_WITHNOTEGETRESPONSE_TYPE
	FAX_WITHNOTEGETRESPONSE_TYPE
	LETTER_WITHNOTEGETRESPONSE_TYPE
	OTHER_WITHNOTEGETRESPONSE_TYPE
)

func (i WithNoteGetResponse_type) String() string {
	return []string{"Appointment", "Call", "Email", "Fax", "Letter", "Other"}[i]
}
func ParseWithNoteGetResponse_type(v string) (any, error) {
	result := APPOINTMENT_WITHNOTEGETRESPONSE_TYPE
	switch v {
	case "Appointment":
		result = APPOINTMENT_WITHNOTEGETRESPONSE_TYPE
	case "Call":
		result = CALL_WITHNOTEGETRESPONSE_TYPE
	case "Email":
		result = EMAIL_WITHNOTEGETRESPONSE_TYPE
	case "Fax":
		result = FAX_WITHNOTEGETRESPONSE_TYPE
	case "Letter":
		result = LETTER_WITHNOTEGETRESPONSE_TYPE
	case "Other":
		result = OTHER_WITHNOTEGETRESPONSE_TYPE
	default:
		return 0, errors.New("Unknown WithNoteGetResponse_type value: " + v)
	}
	return &result, nil
}
func SerializeWithNoteGetResponse_type(values []WithNoteGetResponse_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithNoteGetResponse_type) isMultiValue() bool {
	return false
}
