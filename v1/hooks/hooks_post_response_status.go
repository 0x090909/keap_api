package hooks

import (
	"errors"
)

type HooksPostResponse_status int

const (
	UNVERIFIED_HOOKSPOSTRESPONSE_STATUS HooksPostResponse_status = iota
	VERIFIED_HOOKSPOSTRESPONSE_STATUS
	INACTIVE_HOOKSPOSTRESPONSE_STATUS
)

func (i HooksPostResponse_status) String() string {
	return []string{"Unverified", "Verified", "Inactive"}[i]
}
func ParseHooksPostResponse_status(v string) (any, error) {
	result := UNVERIFIED_HOOKSPOSTRESPONSE_STATUS
	switch v {
	case "Unverified":
		result = UNVERIFIED_HOOKSPOSTRESPONSE_STATUS
	case "Verified":
		result = VERIFIED_HOOKSPOSTRESPONSE_STATUS
	case "Inactive":
		result = INACTIVE_HOOKSPOSTRESPONSE_STATUS
	default:
		return 0, errors.New("Unknown HooksPostResponse_status value: " + v)
	}
	return &result, nil
}
func SerializeHooksPostResponse_status(values []HooksPostResponse_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i HooksPostResponse_status) isMultiValue() bool {
	return false
}
