package verify

import (
	"errors"
)

type VerifyPostResponse_status int

const (
	UNVERIFIED_VERIFYPOSTRESPONSE_STATUS VerifyPostResponse_status = iota
	VERIFIED_VERIFYPOSTRESPONSE_STATUS
	INACTIVE_VERIFYPOSTRESPONSE_STATUS
)

func (i VerifyPostResponse_status) String() string {
	return []string{"Unverified", "Verified", "Inactive"}[i]
}
func ParseVerifyPostResponse_status(v string) (any, error) {
	result := UNVERIFIED_VERIFYPOSTRESPONSE_STATUS
	switch v {
	case "Unverified":
		result = UNVERIFIED_VERIFYPOSTRESPONSE_STATUS
	case "Verified":
		result = VERIFIED_VERIFYPOSTRESPONSE_STATUS
	case "Inactive":
		result = INACTIVE_VERIFYPOSTRESPONSE_STATUS
	default:
		return 0, errors.New("Unknown VerifyPostResponse_status value: " + v)
	}
	return &result, nil
}
func SerializeVerifyPostResponse_status(values []VerifyPostResponse_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i VerifyPostResponse_status) isMultiValue() bool {
	return false
}
