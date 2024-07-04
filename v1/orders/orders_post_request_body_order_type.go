package orders

import (
	"errors"
)

type OrdersPostRequestBody_order_type int

const (
	OFFLINE_ORDERSPOSTREQUESTBODY_ORDER_TYPE OrdersPostRequestBody_order_type = iota
	ONLINE_ORDERSPOSTREQUESTBODY_ORDER_TYPE
)

func (i OrdersPostRequestBody_order_type) String() string {
	return []string{"Offline", "Online"}[i]
}
func ParseOrdersPostRequestBody_order_type(v string) (any, error) {
	result := OFFLINE_ORDERSPOSTREQUESTBODY_ORDER_TYPE
	switch v {
	case "Offline":
		result = OFFLINE_ORDERSPOSTREQUESTBODY_ORDER_TYPE
	case "Online":
		result = ONLINE_ORDERSPOSTREQUESTBODY_ORDER_TYPE
	default:
		return 0, errors.New("Unknown OrdersPostRequestBody_order_type value: " + v)
	}
	return &result, nil
}
func SerializeOrdersPostRequestBody_order_type(values []OrdersPostRequestBody_order_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i OrdersPostRequestBody_order_type) isMultiValue() bool {
	return false
}
