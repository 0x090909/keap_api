package hooks

import (
	"errors"
)

type Hooks_status int

const (
	UNVERIFIED_HOOKS_STATUS Hooks_status = iota
	VERIFIED_HOOKS_STATUS
	INACTIVE_HOOKS_STATUS
)

func (i Hooks_status) String() string {
	return []string{"Unverified", "Verified", "Inactive"}[i]
}
func ParseHooks_status(v string) (any, error) {
	result := UNVERIFIED_HOOKS_STATUS
	switch v {
	case "Unverified":
		result = UNVERIFIED_HOOKS_STATUS
	case "Verified":
		result = VERIFIED_HOOKS_STATUS
	case "Inactive":
		result = INACTIVE_HOOKS_STATUS
	default:
		return 0, errors.New("Unknown Hooks_status value: " + v)
	}
	return &result, nil
}
func SerializeHooks_status(values []Hooks_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i Hooks_status) isMultiValue() bool {
	return false
}
