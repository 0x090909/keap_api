package item

import (
	"errors"
)

type WithFilePutRequestBody_file_association int

const (
	CONTACT_WITHFILEPUTREQUESTBODY_FILE_ASSOCIATION WithFilePutRequestBody_file_association = iota
	USER_WITHFILEPUTREQUESTBODY_FILE_ASSOCIATION
	COMPANY_WITHFILEPUTREQUESTBODY_FILE_ASSOCIATION
)

func (i WithFilePutRequestBody_file_association) String() string {
	return []string{"CONTACT", "USER", "COMPANY"}[i]
}
func ParseWithFilePutRequestBody_file_association(v string) (any, error) {
	result := CONTACT_WITHFILEPUTREQUESTBODY_FILE_ASSOCIATION
	switch v {
	case "CONTACT":
		result = CONTACT_WITHFILEPUTREQUESTBODY_FILE_ASSOCIATION
	case "USER":
		result = USER_WITHFILEPUTREQUESTBODY_FILE_ASSOCIATION
	case "COMPANY":
		result = COMPANY_WITHFILEPUTREQUESTBODY_FILE_ASSOCIATION
	default:
		return 0, errors.New("Unknown WithFilePutRequestBody_file_association value: " + v)
	}
	return &result, nil
}
func SerializeWithFilePutRequestBody_file_association(values []WithFilePutRequestBody_file_association) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithFilePutRequestBody_file_association) isMultiValue() bool {
	return false
}
