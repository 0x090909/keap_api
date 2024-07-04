package orders

import (
	"errors"
)

type GetOrderQueryParameterType int

const (
	ORDER_DATE_GETORDERQUERYPARAMETERTYPE GetOrderQueryParameterType = iota
	UPDATE_DATE_GETORDERQUERYPARAMETERTYPE
)

func (i GetOrderQueryParameterType) String() string {
	return []string{"order_date", "update_date"}[i]
}
func ParseGetOrderQueryParameterType(v string) (any, error) {
	result := ORDER_DATE_GETORDERQUERYPARAMETERTYPE
	switch v {
	case "order_date":
		result = ORDER_DATE_GETORDERQUERYPARAMETERTYPE
	case "update_date":
		result = UPDATE_DATE_GETORDERQUERYPARAMETERTYPE
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
