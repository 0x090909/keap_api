package files

import (
	"errors"
)

type GetPermissionQueryParameterType int

const (
	USER_GETPERMISSIONQUERYPARAMETERTYPE GetPermissionQueryParameterType = iota
	COMPANY_GETPERMISSIONQUERYPARAMETERTYPE
	BOTH_GETPERMISSIONQUERYPARAMETERTYPE
)

func (i GetPermissionQueryParameterType) String() string {
	return []string{"USER", "COMPANY", "BOTH"}[i]
}
func ParseGetPermissionQueryParameterType(v string) (any, error) {
	result := USER_GETPERMISSIONQUERYPARAMETERTYPE
	switch v {
	case "USER":
		result = USER_GETPERMISSIONQUERYPARAMETERTYPE
	case "COMPANY":
		result = COMPANY_GETPERMISSIONQUERYPARAMETERTYPE
	case "BOTH":
		result = BOTH_GETPERMISSIONQUERYPARAMETERTYPE
	default:
		return 0, errors.New("Unknown GetPermissionQueryParameterType value: " + v)
	}
	return &result, nil
}
func SerializeGetPermissionQueryParameterType(values []GetPermissionQueryParameterType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i GetPermissionQueryParameterType) isMultiValue() bool {
	return false
}
