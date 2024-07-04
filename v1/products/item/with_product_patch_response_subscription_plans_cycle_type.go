package item

import (
	"errors"
)

type WithProductPatchResponse_subscription_plans_cycle_type int

const (
	YEAR_WITHPRODUCTPATCHRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE WithProductPatchResponse_subscription_plans_cycle_type = iota
	MONTH_WITHPRODUCTPATCHRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	WEEK_WITHPRODUCTPATCHRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	DAY_WITHPRODUCTPATCHRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
)

func (i WithProductPatchResponse_subscription_plans_cycle_type) String() string {
	return []string{"YEAR", "MONTH", "WEEK", "DAY"}[i]
}
func ParseWithProductPatchResponse_subscription_plans_cycle_type(v string) (any, error) {
	result := YEAR_WITHPRODUCTPATCHRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	switch v {
	case "YEAR":
		result = YEAR_WITHPRODUCTPATCHRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	case "MONTH":
		result = MONTH_WITHPRODUCTPATCHRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	case "WEEK":
		result = WEEK_WITHPRODUCTPATCHRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	case "DAY":
		result = DAY_WITHPRODUCTPATCHRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	default:
		return 0, errors.New("Unknown WithProductPatchResponse_subscription_plans_cycle_type value: " + v)
	}
	return &result, nil
}
func SerializeWithProductPatchResponse_subscription_plans_cycle_type(values []WithProductPatchResponse_subscription_plans_cycle_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithProductPatchResponse_subscription_plans_cycle_type) isMultiValue() bool {
	return false
}
