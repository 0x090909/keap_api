package affiliates

import (
	"errors"
)

type AffiliatesPostRequestBody_status int

const (
	ACTIVE_AFFILIATESPOSTREQUESTBODY_STATUS AffiliatesPostRequestBody_status = iota
	INACTIVE_AFFILIATESPOSTREQUESTBODY_STATUS
)

func (i AffiliatesPostRequestBody_status) String() string {
	return []string{"active", "inactive"}[i]
}
func ParseAffiliatesPostRequestBody_status(v string) (any, error) {
	result := ACTIVE_AFFILIATESPOSTREQUESTBODY_STATUS
	switch v {
	case "active":
		result = ACTIVE_AFFILIATESPOSTREQUESTBODY_STATUS
	case "inactive":
		result = INACTIVE_AFFILIATESPOSTREQUESTBODY_STATUS
	default:
		return 0, errors.New("Unknown AffiliatesPostRequestBody_status value: " + v)
	}
	return &result, nil
}
func SerializeAffiliatesPostRequestBody_status(values []AffiliatesPostRequestBody_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i AffiliatesPostRequestBody_status) isMultiValue() bool {
	return false
}
