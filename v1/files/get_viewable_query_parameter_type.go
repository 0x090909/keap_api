package files

import (
	"errors"
)

type GetViewableQueryParameterType int

const (
	PUBLIC_GETVIEWABLEQUERYPARAMETERTYPE GetViewableQueryParameterType = iota
	PRIVATE_GETVIEWABLEQUERYPARAMETERTYPE
	BOTH_GETVIEWABLEQUERYPARAMETERTYPE
)

func (i GetViewableQueryParameterType) String() string {
	return []string{"PUBLIC", "PRIVATE", "BOTH"}[i]
}
func ParseGetViewableQueryParameterType(v string) (any, error) {
	result := PUBLIC_GETVIEWABLEQUERYPARAMETERTYPE
	switch v {
	case "PUBLIC":
		result = PUBLIC_GETVIEWABLEQUERYPARAMETERTYPE
	case "PRIVATE":
		result = PRIVATE_GETVIEWABLEQUERYPARAMETERTYPE
	case "BOTH":
		result = BOTH_GETVIEWABLEQUERYPARAMETERTYPE
	default:
		return 0, errors.New("Unknown GetViewableQueryParameterType value: " + v)
	}
	return &result, nil
}
func SerializeGetViewableQueryParameterType(values []GetViewableQueryParameterType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i GetViewableQueryParameterType) isMultiValue() bool {
	return false
}
