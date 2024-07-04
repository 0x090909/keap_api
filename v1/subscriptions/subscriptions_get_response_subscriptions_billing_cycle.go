package subscriptions

import (
	"errors"
)

type SubscriptionsGetResponse_subscriptions_billing_cycle int

const (
	YEAR_SUBSCRIPTIONSGETRESPONSE_SUBSCRIPTIONS_BILLING_CYCLE SubscriptionsGetResponse_subscriptions_billing_cycle = iota
	MONTH_SUBSCRIPTIONSGETRESPONSE_SUBSCRIPTIONS_BILLING_CYCLE
	WEEK_SUBSCRIPTIONSGETRESPONSE_SUBSCRIPTIONS_BILLING_CYCLE
	DAY_SUBSCRIPTIONSGETRESPONSE_SUBSCRIPTIONS_BILLING_CYCLE
)

func (i SubscriptionsGetResponse_subscriptions_billing_cycle) String() string {
	return []string{"YEAR", "MONTH", "WEEK", "DAY"}[i]
}
func ParseSubscriptionsGetResponse_subscriptions_billing_cycle(v string) (any, error) {
	result := YEAR_SUBSCRIPTIONSGETRESPONSE_SUBSCRIPTIONS_BILLING_CYCLE
	switch v {
	case "YEAR":
		result = YEAR_SUBSCRIPTIONSGETRESPONSE_SUBSCRIPTIONS_BILLING_CYCLE
	case "MONTH":
		result = MONTH_SUBSCRIPTIONSGETRESPONSE_SUBSCRIPTIONS_BILLING_CYCLE
	case "WEEK":
		result = WEEK_SUBSCRIPTIONSGETRESPONSE_SUBSCRIPTIONS_BILLING_CYCLE
	case "DAY":
		result = DAY_SUBSCRIPTIONSGETRESPONSE_SUBSCRIPTIONS_BILLING_CYCLE
	default:
		return 0, errors.New("Unknown SubscriptionsGetResponse_subscriptions_billing_cycle value: " + v)
	}
	return &result, nil
}
func SerializeSubscriptionsGetResponse_subscriptions_billing_cycle(values []SubscriptionsGetResponse_subscriptions_billing_cycle) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i SubscriptionsGetResponse_subscriptions_billing_cycle) isMultiValue() bool {
	return false
}
