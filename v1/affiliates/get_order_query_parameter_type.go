package affiliates

import (
	"errors"
)

type GetOrderQueryParameterType int

const (
	ID_GETORDERQUERYPARAMETERTYPE GetOrderQueryParameterType = iota
	NAME_GETORDERQUERYPARAMETERTYPE
)

func (i GetOrderQueryParameterType) String() string {
	return []string{"id", "name"}[i]
}
func ParseGetOrderQueryParameterType(v string) (any, error) {
	result := ID_GETORDERQUERYPARAMETERTYPE
	switch v {
	case "id":
		result = ID_GETORDERQUERYPARAMETERTYPE
	case "name":
		result = NAME_GETORDERQUERYPARAMETERTYPE
	default:
		return 0, errors.New("Unknown GetOrderQueryParameterType value: " + v)
	}
	return &result, nil
}
func SerializeGetOrderQueryParameterType(values []GetOrderQueryParameterType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i GetOrderQueryParameterType) isMultiValue() bool {
	return false
}
