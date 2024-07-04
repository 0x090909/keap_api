package campaigns

import (
	"errors"
)

type GetOrderQueryParameterType int

const (
	ID_GETORDERQUERYPARAMETERTYPE GetOrderQueryParameterType = iota
	NAME_GETORDERQUERYPARAMETERTYPE
	PUBLISHED_DATE_GETORDERQUERYPARAMETERTYPE
	COMPLETED_CONTACT_COUNT_GETORDERQUERYPARAMETERTYPE
	ACTIVE_CONTACT_COUNT_GETORDERQUERYPARAMETERTYPE
	DATE_CREATED_GETORDERQUERYPARAMETERTYPE
	LAST_UPDATED_GETORDERQUERYPARAMETERTYPE
	CATEGORY_GETORDERQUERYPARAMETERTYPE
)

func (i GetOrderQueryParameterType) String() string {
	return []string{"id", "name", "published_date", "completed_contact_count", "active_contact_count", "date_created", "last_updated", "category"}[i]
}
func ParseGetOrderQueryParameterType(v string) (any, error) {
	result := ID_GETORDERQUERYPARAMETERTYPE
	switch v {
	case "id":
		result = ID_GETORDERQUERYPARAMETERTYPE
	case "name":
		result = NAME_GETORDERQUERYPARAMETERTYPE
	case "published_date":
		result = PUBLISHED_DATE_GETORDERQUERYPARAMETERTYPE
	case "completed_contact_count":
		result = COMPLETED_CONTACT_COUNT_GETORDERQUERYPARAMETERTYPE
	case "active_contact_count":
		result = ACTIVE_CONTACT_COUNT_GETORDERQUERYPARAMETERTYPE
	case "date_created":
		result = DATE_CREATED_GETORDERQUERYPARAMETERTYPE
	case "last_updated":
		result = LAST_UPDATED_GETORDERQUERYPARAMETERTYPE
	case "category":
		result = CATEGORY_GETORDERQUERYPARAMETERTYPE
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
