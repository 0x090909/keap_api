package item

import (
	"errors"
)

type WithNotePatchResponse_type int

const (
	APPOINTMENT_WITHNOTEPATCHRESPONSE_TYPE WithNotePatchResponse_type = iota
	CALL_WITHNOTEPATCHRESPONSE_TYPE
	EMAIL_WITHNOTEPATCHRESPONSE_TYPE
	FAX_WITHNOTEPATCHRESPONSE_TYPE
	LETTER_WITHNOTEPATCHRESPONSE_TYPE
	OTHER_WITHNOTEPATCHRESPONSE_TYPE
)

func (i WithNotePatchResponse_type) String() string {
	return []string{"Appointment", "Call", "Email", "Fax", "Letter", "Other"}[i]
}
func ParseWithNotePatchResponse_type(v string) (any, error) {
	result := APPOINTMENT_WITHNOTEPATCHRESPONSE_TYPE
	switch v {
	case "Appointment":
		result = APPOINTMENT_WITHNOTEPATCHRESPONSE_TYPE
	case "Call":
		result = CALL_WITHNOTEPATCHRESPONSE_TYPE
	case "Email":
		result = EMAIL_WITHNOTEPATCHRESPONSE_TYPE
	case "Fax":
		result = FAX_WITHNOTEPATCHRESPONSE_TYPE
	case "Letter":
		result = LETTER_WITHNOTEPATCHRESPONSE_TYPE
	case "Other":
		result = OTHER_WITHNOTEPATCHRESPONSE_TYPE
	default:
		return 0, errors.New("Unknown WithNotePatchResponse_type value: " + v)
	}
	return &result, nil
}
func SerializeWithNotePatchResponse_type(values []WithNotePatchResponse_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithNotePatchResponse_type) isMultiValue() bool {
	return false
}
