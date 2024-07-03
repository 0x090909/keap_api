package types

import (
	"errors"
)

type GetFilterQueryParameterType int

const (
	NAME_GETFILTERQUERYPARAMETERTYPE GetFilterQueryParameterType = iota
)

func (i GetFilterQueryParameterType) String() string {
	return []string{"name"}[i]
}
func ParseGetFilterQueryParameterType(v string) (any, error) {
	result := NAME_GETFILTERQUERYPARAMETERTYPE
	switch v {
	case "name":
		result = NAME_GETFILTERQUERYPARAMETERTYPE
	default:
		return 0, errors.New("Unknown GetFilterQueryParameterType value: " + v)
	}
	return &result, nil
}
func SerializeGetFilterQueryParameterType(values []GetFilterQueryParameterType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i GetFilterQueryParameterType) isMultiValue() bool {
	return false
}
