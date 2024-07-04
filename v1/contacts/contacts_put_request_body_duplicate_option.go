package contacts

import (
	"errors"
)

type ContactsPutRequestBody_duplicate_option int

const (
	EMAIL_CONTACTSPUTREQUESTBODY_DUPLICATE_OPTION ContactsPutRequestBody_duplicate_option = iota
	EMAILANDNAME_CONTACTSPUTREQUESTBODY_DUPLICATE_OPTION
)

func (i ContactsPutRequestBody_duplicate_option) String() string {
	return []string{"Email", "EmailAndName"}[i]
}
func ParseContactsPutRequestBody_duplicate_option(v string) (any, error) {
	result := EMAIL_CONTACTSPUTREQUESTBODY_DUPLICATE_OPTION
	switch v {
	case "Email":
		result = EMAIL_CONTACTSPUTREQUESTBODY_DUPLICATE_OPTION
	case "EmailAndName":
		result = EMAILANDNAME_CONTACTSPUTREQUESTBODY_DUPLICATE_OPTION
	default:
		return 0, errors.New("Unknown ContactsPutRequestBody_duplicate_option value: " + v)
	}
	return &result, nil
}
func SerializeContactsPutRequestBody_duplicate_option(values []ContactsPutRequestBody_duplicate_option) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPutRequestBody_duplicate_option) isMultiValue() bool {
	return false
}
