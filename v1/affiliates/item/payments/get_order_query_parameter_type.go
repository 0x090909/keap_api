package payments

import (
	"errors"
)

type GetOrderQueryParameterType int

const (
	ID_GETORDERQUERYPARAMETERTYPE GetOrderQueryParameterType = iota
	DATE_GETORDERQUERYPARAMETERTYPE
	AMOUNT_GETORDERQUERYPARAMETERTYPE
)

func (i GetOrderQueryParameterType) String() string {
	return []string{"ID", "DATE", "AMOUNT"}[i]
}
func ParseGetOrderQueryParameterType(v string) (any, error) {
	result := ID_GETORDERQUERYPARAMETERTYPE
	switch v {
	case "ID":
		result = ID_GETORDERQUERYPARAMETERTYPE
	case "DATE":
		result = DATE_GETORDERQUERYPARAMETERTYPE
	case "AMOUNT":
		result = AMOUNT_GETORDERQUERYPARAMETERTYPE
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
