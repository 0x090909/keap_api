package item

import (
	"errors"
)

type WithNotePutResponse_type int

const (
	APPOINTMENT_WITHNOTEPUTRESPONSE_TYPE WithNotePutResponse_type = iota
	CALL_WITHNOTEPUTRESPONSE_TYPE
	EMAIL_WITHNOTEPUTRESPONSE_TYPE
	FAX_WITHNOTEPUTRESPONSE_TYPE
	LETTER_WITHNOTEPUTRESPONSE_TYPE
	OTHER_WITHNOTEPUTRESPONSE_TYPE
)

func (i WithNotePutResponse_type) String() string {
	return []string{"Appointment", "Call", "Email", "Fax", "Letter", "Other"}[i]
}
func ParseWithNotePutResponse_type(v string) (any, error) {
	result := APPOINTMENT_WITHNOTEPUTRESPONSE_TYPE
	switch v {
	case "Appointment":
		result = APPOINTMENT_WITHNOTEPUTRESPONSE_TYPE
	case "Call":
		result = CALL_WITHNOTEPUTRESPONSE_TYPE
	case "Email":
		result = EMAIL_WITHNOTEPUTRESPONSE_TYPE
	case "Fax":
		result = FAX_WITHNOTEPUTRESPONSE_TYPE
	case "Letter":
		result = LETTER_WITHNOTEPUTRESPONSE_TYPE
	case "Other":
		result = OTHER_WITHNOTEPUTRESPONSE_TYPE
	default:
		return 0, errors.New("Unknown WithNotePutResponse_type value: " + v)
	}
	return &result, nil
}
func SerializeWithNotePutResponse_type(values []WithNotePutResponse_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithNotePutResponse_type) isMultiValue() bool {
	return false
}
