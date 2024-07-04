package commissions

import (
	"errors"
)

type GetOrderQueryParameterType int

const (
	DATE_EARNED_GETORDERQUERYPARAMETERTYPE GetOrderQueryParameterType = iota
)

func (i GetOrderQueryParameterType) String() string {
	return []string{"DATE_EARNED"}[i]
}
func ParseGetOrderQueryParameterType(v string) (any, error) {
	result := DATE_EARNED_GETORDERQUERYPARAMETERTYPE
	switch v {
	case "DATE_EARNED":
		result = DATE_EARNED_GETORDERQUERYPARAMETERTYPE
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
