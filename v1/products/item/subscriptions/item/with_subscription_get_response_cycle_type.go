package item

import (
	"errors"
)

type WithSubscriptionGetResponse_cycle_type int

const (
	YEAR_WITHSUBSCRIPTIONGETRESPONSE_CYCLE_TYPE WithSubscriptionGetResponse_cycle_type = iota
	MONTH_WITHSUBSCRIPTIONGETRESPONSE_CYCLE_TYPE
	WEEK_WITHSUBSCRIPTIONGETRESPONSE_CYCLE_TYPE
	DAY_WITHSUBSCRIPTIONGETRESPONSE_CYCLE_TYPE
)

func (i WithSubscriptionGetResponse_cycle_type) String() string {
	return []string{"YEAR", "MONTH", "WEEK", "DAY"}[i]
}
func ParseWithSubscriptionGetResponse_cycle_type(v string) (any, error) {
	result := YEAR_WITHSUBSCRIPTIONGETRESPONSE_CYCLE_TYPE
	switch v {
	case "YEAR":
		result = YEAR_WITHSUBSCRIPTIONGETRESPONSE_CYCLE_TYPE
	case "MONTH":
		result = MONTH_WITHSUBSCRIPTIONGETRESPONSE_CYCLE_TYPE
	case "WEEK":
		result = WEEK_WITHSUBSCRIPTIONGETRESPONSE_CYCLE_TYPE
	case "DAY":
		result = DAY_WITHSUBSCRIPTIONGETRESPONSE_CYCLE_TYPE
	default:
		return 0, errors.New("Unknown WithSubscriptionGetResponse_cycle_type value: " + v)
	}
	return &result, nil
}
func SerializeWithSubscriptionGetResponse_cycle_type(values []WithSubscriptionGetResponse_cycle_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithSubscriptionGetResponse_cycle_type) isMultiValue() bool {
	return false
}
