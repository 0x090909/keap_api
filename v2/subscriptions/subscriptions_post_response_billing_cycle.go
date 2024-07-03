package subscriptions

import (
	"errors"
)

type SubscriptionsPostResponse_billing_cycle int

const (
	YEAR_SUBSCRIPTIONSPOSTRESPONSE_BILLING_CYCLE SubscriptionsPostResponse_billing_cycle = iota
	MONTH_SUBSCRIPTIONSPOSTRESPONSE_BILLING_CYCLE
	WEEK_SUBSCRIPTIONSPOSTRESPONSE_BILLING_CYCLE
	DAY_SUBSCRIPTIONSPOSTRESPONSE_BILLING_CYCLE
)

func (i SubscriptionsPostResponse_billing_cycle) String() string {
	return []string{"YEAR", "MONTH", "WEEK", "DAY"}[i]
}
func ParseSubscriptionsPostResponse_billing_cycle(v string) (any, error) {
	result := YEAR_SUBSCRIPTIONSPOSTRESPONSE_BILLING_CYCLE
	switch v {
	case "YEAR":
		result = YEAR_SUBSCRIPTIONSPOSTRESPONSE_BILLING_CYCLE
	case "MONTH":
		result = MONTH_SUBSCRIPTIONSPOSTRESPONSE_BILLING_CYCLE
	case "WEEK":
		result = WEEK_SUBSCRIPTIONSPOSTRESPONSE_BILLING_CYCLE
	case "DAY":
		result = DAY_SUBSCRIPTIONSPOSTRESPONSE_BILLING_CYCLE
	default:
		return 0, errors.New("Unknown SubscriptionsPostResponse_billing_cycle value: " + v)
	}
	return &result, nil
}
func SerializeSubscriptionsPostResponse_billing_cycle(values []SubscriptionsPostResponse_billing_cycle) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i SubscriptionsPostResponse_billing_cycle) isMultiValue() bool {
	return false
}
