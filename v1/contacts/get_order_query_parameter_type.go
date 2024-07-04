package contacts

import (
	"errors"
)

type GetOrderQueryParameterType int

const (
	ID_GETORDERQUERYPARAMETERTYPE GetOrderQueryParameterType = iota
	DATE_CREATED_GETORDERQUERYPARAMETERTYPE
	LAST_UPDATED_GETORDERQUERYPARAMETERTYPE
	NAME_GETORDERQUERYPARAMETERTYPE
	FIRSTNAME_GETORDERQUERYPARAMETERTYPE
	EMAIL_GETORDERQUERYPARAMETERTYPE
)

func (i GetOrderQueryParameterType) String() string {
	return []string{"id", "date_created", "last_updated", "name", "firstName", "email"}[i]
}
func ParseGetOrderQueryParameterType(v string) (any, error) {
	result := ID_GETORDERQUERYPARAMETERTYPE
	switch v {
	case "id":
		result = ID_GETORDERQUERYPARAMETERTYPE
	case "date_created":
		result = DATE_CREATED_GETORDERQUERYPARAMETERTYPE
	case "last_updated":
		result = LAST_UPDATED_GETORDERQUERYPARAMETERTYPE
	case "name":
		result = NAME_GETORDERQUERYPARAMETERTYPE
	case "firstName":
		result = FIRSTNAME_GETORDERQUERYPARAMETERTYPE
	case "email":
		result = EMAIL_GETORDERQUERYPARAMETERTYPE
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
