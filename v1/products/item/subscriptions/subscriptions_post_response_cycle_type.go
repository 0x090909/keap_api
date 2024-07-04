package subscriptions

import (
	"errors"
)

type SubscriptionsPostResponse_cycle_type int

const (
	YEAR_SUBSCRIPTIONSPOSTRESPONSE_CYCLE_TYPE SubscriptionsPostResponse_cycle_type = iota
	MONTH_SUBSCRIPTIONSPOSTRESPONSE_CYCLE_TYPE
	WEEK_SUBSCRIPTIONSPOSTRESPONSE_CYCLE_TYPE
	DAY_SUBSCRIPTIONSPOSTRESPONSE_CYCLE_TYPE
)

func (i SubscriptionsPostResponse_cycle_type) String() string {
	return []string{"YEAR", "MONTH", "WEEK", "DAY"}[i]
}
func ParseSubscriptionsPostResponse_cycle_type(v string) (any, error) {
	result := YEAR_SUBSCRIPTIONSPOSTRESPONSE_CYCLE_TYPE
	switch v {
	case "YEAR":
		result = YEAR_SUBSCRIPTIONSPOSTRESPONSE_CYCLE_TYPE
	case "MONTH":
		result = MONTH_SUBSCRIPTIONSPOSTRESPONSE_CYCLE_TYPE
	case "WEEK":
		result = WEEK_SUBSCRIPTIONSPOSTRESPONSE_CYCLE_TYPE
	case "DAY":
		result = DAY_SUBSCRIPTIONSPOSTRESPONSE_CYCLE_TYPE
	default:
		return 0, errors.New("Unknown SubscriptionsPostResponse_cycle_type value: " + v)
	}
	return &result, nil
}
func SerializeSubscriptionsPostResponse_cycle_type(values []SubscriptionsPostResponse_cycle_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i SubscriptionsPostResponse_cycle_type) isMultiValue() bool {
	return false
}
