package item

import (
	"errors"
)

type WithProductGetResponse_subscription_plans_cycle_type int

const (
	YEAR_WITHPRODUCTGETRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE WithProductGetResponse_subscription_plans_cycle_type = iota
	MONTH_WITHPRODUCTGETRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	WEEK_WITHPRODUCTGETRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	DAY_WITHPRODUCTGETRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
)

func (i WithProductGetResponse_subscription_plans_cycle_type) String() string {
	return []string{"YEAR", "MONTH", "WEEK", "DAY"}[i]
}
func ParseWithProductGetResponse_subscription_plans_cycle_type(v string) (any, error) {
	result := YEAR_WITHPRODUCTGETRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	switch v {
	case "YEAR":
		result = YEAR_WITHPRODUCTGETRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	case "MONTH":
		result = MONTH_WITHPRODUCTGETRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	case "WEEK":
		result = WEEK_WITHPRODUCTGETRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	case "DAY":
		result = DAY_WITHPRODUCTGETRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	default:
		return 0, errors.New("Unknown WithProductGetResponse_subscription_plans_cycle_type value: " + v)
	}
	return &result, nil
}
func SerializeWithProductGetResponse_subscription_plans_cycle_type(values []WithProductGetResponse_subscription_plans_cycle_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithProductGetResponse_subscription_plans_cycle_type) isMultiValue() bool {
	return false
}
