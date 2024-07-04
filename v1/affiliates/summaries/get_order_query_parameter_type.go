package summaries

import (
	"errors"
)

type GetOrderQueryParameterType int

const (
	AFFILIATE_ID_GETORDERQUERYPARAMETERTYPE GetOrderQueryParameterType = iota
	AMOUNT_EARNED_GETORDERQUERYPARAMETERTYPE
	BALANCE_GETORDERQUERYPARAMETERTYPE
	CLAWBACKS_GETORDERQUERYPARAMETERTYPE
)

func (i GetOrderQueryParameterType) String() string {
	return []string{"affiliate_id", "amount_earned", "balance", "clawbacks"}[i]
}
func ParseGetOrderQueryParameterType(v string) (any, error) {
	result := AFFILIATE_ID_GETORDERQUERYPARAMETERTYPE
	switch v {
	case "affiliate_id":
		result = AFFILIATE_ID_GETORDERQUERYPARAMETERTYPE
	case "amount_earned":
		result = AMOUNT_EARNED_GETORDERQUERYPARAMETERTYPE
	case "balance":
		result = BALANCE_GETORDERQUERYPARAMETERTYPE
	case "clawbacks":
		result = CLAWBACKS_GETORDERQUERYPARAMETERTYPE
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
