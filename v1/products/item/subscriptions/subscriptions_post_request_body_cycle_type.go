package subscriptions

import (
	"errors"
)

type SubscriptionsPostRequestBody_cycle_type int

const (
	YEAR_SUBSCRIPTIONSPOSTREQUESTBODY_CYCLE_TYPE SubscriptionsPostRequestBody_cycle_type = iota
	MONTH_SUBSCRIPTIONSPOSTREQUESTBODY_CYCLE_TYPE
	WEEK_SUBSCRIPTIONSPOSTREQUESTBODY_CYCLE_TYPE
	DAY_SUBSCRIPTIONSPOSTREQUESTBODY_CYCLE_TYPE
)

func (i SubscriptionsPostRequestBody_cycle_type) String() string {
	return []string{"YEAR", "MONTH", "WEEK", "DAY"}[i]
}
func ParseSubscriptionsPostRequestBody_cycle_type(v string) (any, error) {
	result := YEAR_SUBSCRIPTIONSPOSTREQUESTBODY_CYCLE_TYPE
	switch v {
	case "YEAR":
		result = YEAR_SUBSCRIPTIONSPOSTREQUESTBODY_CYCLE_TYPE
	case "MONTH":
		result = MONTH_SUBSCRIPTIONSPOSTREQUESTBODY_CYCLE_TYPE
	case "WEEK":
		result = WEEK_SUBSCRIPTIONSPOSTREQUESTBODY_CYCLE_TYPE
	case "DAY":
		result = DAY_SUBSCRIPTIONSPOSTREQUESTBODY_CYCLE_TYPE
	default:
		return 0, errors.New("Unknown SubscriptionsPostRequestBody_cycle_type value: " + v)
	}
	return &result, nil
}
func SerializeSubscriptionsPostRequestBody_cycle_type(values []SubscriptionsPostRequestBody_cycle_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i SubscriptionsPostRequestBody_cycle_type) isMultiValue() bool {
	return false
}
