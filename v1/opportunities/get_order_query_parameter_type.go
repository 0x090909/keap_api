package opportunities

import (
	"errors"
)

type GetOrderQueryParameterType int

const (
	NEXT_ACTION_GETORDERQUERYPARAMETERTYPE GetOrderQueryParameterType = iota
	OPPORTUNITY_NAME_GETORDERQUERYPARAMETERTYPE
	CONTACT_NAME_GETORDERQUERYPARAMETERTYPE
	DATE_CREATED_GETORDERQUERYPARAMETERTYPE
)

func (i GetOrderQueryParameterType) String() string {
	return []string{"next_action", "opportunity_name", "contact_name", "date_created"}[i]
}
func ParseGetOrderQueryParameterType(v string) (any, error) {
	result := NEXT_ACTION_GETORDERQUERYPARAMETERTYPE
	switch v {
	case "next_action":
		result = NEXT_ACTION_GETORDERQUERYPARAMETERTYPE
	case "opportunity_name":
		result = OPPORTUNITY_NAME_GETORDERQUERYPARAMETERTYPE
	case "contact_name":
		result = CONTACT_NAME_GETORDERQUERYPARAMETERTYPE
	case "date_created":
		result = DATE_CREATED_GETORDERQUERYPARAMETERTYPE
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
