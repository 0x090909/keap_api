package products

import (
	"errors"
)

type ProductsPostResponse_subscription_plans_cycle_type int

const (
	YEAR_PRODUCTSPOSTRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE ProductsPostResponse_subscription_plans_cycle_type = iota
	MONTH_PRODUCTSPOSTRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	WEEK_PRODUCTSPOSTRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	DAY_PRODUCTSPOSTRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
)

func (i ProductsPostResponse_subscription_plans_cycle_type) String() string {
	return []string{"YEAR", "MONTH", "WEEK", "DAY"}[i]
}
func ParseProductsPostResponse_subscription_plans_cycle_type(v string) (any, error) {
	result := YEAR_PRODUCTSPOSTRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	switch v {
	case "YEAR":
		result = YEAR_PRODUCTSPOSTRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	case "MONTH":
		result = MONTH_PRODUCTSPOSTRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	case "WEEK":
		result = WEEK_PRODUCTSPOSTRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	case "DAY":
		result = DAY_PRODUCTSPOSTRESPONSE_SUBSCRIPTION_PLANS_CYCLE_TYPE
	default:
		return 0, errors.New("Unknown ProductsPostResponse_subscription_plans_cycle_type value: " + v)
	}
	return &result, nil
}
func SerializeProductsPostResponse_subscription_plans_cycle_type(values []ProductsPostResponse_subscription_plans_cycle_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ProductsPostResponse_subscription_plans_cycle_type) isMultiValue() bool {
	return false
}
