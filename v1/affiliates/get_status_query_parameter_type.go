package affiliates

import (
	"errors"
)

type GetStatusQueryParameterType int

const (
	ACTIVE_GETSTATUSQUERYPARAMETERTYPE GetStatusQueryParameterType = iota
	INACTIVE_GETSTATUSQUERYPARAMETERTYPE
)

func (i GetStatusQueryParameterType) String() string {
	return []string{"active", "inactive"}[i]
}
func ParseGetStatusQueryParameterType(v string) (any, error) {
	result := ACTIVE_GETSTATUSQUERYPARAMETERTYPE
	switch v {
	case "active":
		result = ACTIVE_GETSTATUSQUERYPARAMETERTYPE
	case "inactive":
		result = INACTIVE_GETSTATUSQUERYPARAMETERTYPE
	default:
		return 0, errors.New("Unknown GetStatusQueryParameterType value: " + v)
	}
	return &result, nil
}
func SerializeGetStatusQueryParameterType(values []GetStatusQueryParameterType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i GetStatusQueryParameterType) isMultiValue() bool {
	return false
}
