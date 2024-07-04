package item

import (
	"errors"
)

type WithNotePatchRequestBody_type int

const (
	APPOINTMENT_WITHNOTEPATCHREQUESTBODY_TYPE WithNotePatchRequestBody_type = iota
	CALL_WITHNOTEPATCHREQUESTBODY_TYPE
	EMAIL_WITHNOTEPATCHREQUESTBODY_TYPE
	FAX_WITHNOTEPATCHREQUESTBODY_TYPE
	LETTER_WITHNOTEPATCHREQUESTBODY_TYPE
	OTHER_WITHNOTEPATCHREQUESTBODY_TYPE
)

func (i WithNotePatchRequestBody_type) String() string {
	return []string{"Appointment", "Call", "Email", "Fax", "Letter", "Other"}[i]
}
func ParseWithNotePatchRequestBody_type(v string) (any, error) {
	result := APPOINTMENT_WITHNOTEPATCHREQUESTBODY_TYPE
	switch v {
	case "Appointment":
		result = APPOINTMENT_WITHNOTEPATCHREQUESTBODY_TYPE
	case "Call":
		result = CALL_WITHNOTEPATCHREQUESTBODY_TYPE
	case "Email":
		result = EMAIL_WITHNOTEPATCHREQUESTBODY_TYPE
	case "Fax":
		result = FAX_WITHNOTEPATCHREQUESTBODY_TYPE
	case "Letter":
		result = LETTER_WITHNOTEPATCHREQUESTBODY_TYPE
	case "Other":
		result = OTHER_WITHNOTEPATCHREQUESTBODY_TYPE
	default:
		return 0, errors.New("Unknown WithNotePatchRequestBody_type value: " + v)
	}
	return &result, nil
}
func SerializeWithNotePatchRequestBody_type(values []WithNotePatchRequestBody_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithNotePatchRequestBody_type) isMultiValue() bool {
	return false
}
