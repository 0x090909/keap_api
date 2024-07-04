package delayedverify

import (
	"errors"
)

type DelayedVerifyPostResponse_status int

const (
	UNVERIFIED_DELAYEDVERIFYPOSTRESPONSE_STATUS DelayedVerifyPostResponse_status = iota
	VERIFIED_DELAYEDVERIFYPOSTRESPONSE_STATUS
	INACTIVE_DELAYEDVERIFYPOSTRESPONSE_STATUS
)

func (i DelayedVerifyPostResponse_status) String() string {
	return []string{"Unverified", "Verified", "Inactive"}[i]
}
func ParseDelayedVerifyPostResponse_status(v string) (any, error) {
	result := UNVERIFIED_DELAYEDVERIFYPOSTRESPONSE_STATUS
	switch v {
	case "Unverified":
		result = UNVERIFIED_DELAYEDVERIFYPOSTRESPONSE_STATUS
	case "Verified":
		result = VERIFIED_DELAYEDVERIFYPOSTRESPONSE_STATUS
	case "Inactive":
		result = INACTIVE_DELAYEDVERIFYPOSTRESPONSE_STATUS
	default:
		return 0, errors.New("Unknown DelayedVerifyPostResponse_status value: " + v)
	}
	return &result, nil
}
func SerializeDelayedVerifyPostResponse_status(values []DelayedVerifyPostResponse_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i DelayedVerifyPostResponse_status) isMultiValue() bool {
	return false
}
