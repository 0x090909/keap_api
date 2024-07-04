package item

import (
	"errors"
)

type WithKeyGetResponse_status int

const (
	UNVERIFIED_WITHKEYGETRESPONSE_STATUS WithKeyGetResponse_status = iota
	VERIFIED_WITHKEYGETRESPONSE_STATUS
	INACTIVE_WITHKEYGETRESPONSE_STATUS
)

func (i WithKeyGetResponse_status) String() string {
	return []string{"Unverified", "Verified", "Inactive"}[i]
}
func ParseWithKeyGetResponse_status(v string) (any, error) {
	result := UNVERIFIED_WITHKEYGETRESPONSE_STATUS
	switch v {
	case "Unverified":
		result = UNVERIFIED_WITHKEYGETRESPONSE_STATUS
	case "Verified":
		result = VERIFIED_WITHKEYGETRESPONSE_STATUS
	case "Inactive":
		result = INACTIVE_WITHKEYGETRESPONSE_STATUS
	default:
		return 0, errors.New("Unknown WithKeyGetResponse_status value: " + v)
	}
	return &result, nil
}
func SerializeWithKeyGetResponse_status(values []WithKeyGetResponse_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithKeyGetResponse_status) isMultiValue() bool {
	return false
}
