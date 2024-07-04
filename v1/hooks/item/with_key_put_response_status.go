package item

import (
	"errors"
)

type WithKeyPutResponse_status int

const (
	UNVERIFIED_WITHKEYPUTRESPONSE_STATUS WithKeyPutResponse_status = iota
	VERIFIED_WITHKEYPUTRESPONSE_STATUS
	INACTIVE_WITHKEYPUTRESPONSE_STATUS
)

func (i WithKeyPutResponse_status) String() string {
	return []string{"Unverified", "Verified", "Inactive"}[i]
}
func ParseWithKeyPutResponse_status(v string) (any, error) {
	result := UNVERIFIED_WITHKEYPUTRESPONSE_STATUS
	switch v {
	case "Unverified":
		result = UNVERIFIED_WITHKEYPUTRESPONSE_STATUS
	case "Verified":
		result = VERIFIED_WITHKEYPUTRESPONSE_STATUS
	case "Inactive":
		result = INACTIVE_WITHKEYPUTRESPONSE_STATUS
	default:
		return 0, errors.New("Unknown WithKeyPutResponse_status value: " + v)
	}
	return &result, nil
}
func SerializeWithKeyPutResponse_status(values []WithKeyPutResponse_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithKeyPutResponse_status) isMultiValue() bool {
	return false
}
