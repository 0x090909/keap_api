package programs

import (
	"errors"
)

type GetOrder_directionQueryParameterType int

const (
	ASCENDING_GETORDER_DIRECTIONQUERYPARAMETERTYPE GetOrder_directionQueryParameterType = iota
	DESCENDING_GETORDER_DIRECTIONQUERYPARAMETERTYPE
)

func (i GetOrder_directionQueryParameterType) String() string {
	return []string{"ASCENDING", "DESCENDING"}[i]
}
func ParseGetOrder_directionQueryParameterType(v string) (any, error) {
	result := ASCENDING_GETORDER_DIRECTIONQUERYPARAMETERTYPE
	switch v {
	case "ASCENDING":
		result = ASCENDING_GETORDER_DIRECTIONQUERYPARAMETERTYPE
	case "DESCENDING":
		result = DESCENDING_GETORDER_DIRECTIONQUERYPARAMETERTYPE
	default:
		return 0, errors.New("Unknown GetOrder_directionQueryParameterType value: " + v)
	}
	return &result, nil
}
func SerializeGetOrder_directionQueryParameterType(values []GetOrder_directionQueryParameterType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i GetOrder_directionQueryParameterType) isMultiValue() bool {
	return false
}
