package orders

import (
	"errors"
)

type OrdersPostResponse_order_items_product_product_options_type int

const (
	FIXEDLIST_ORDERSPOSTRESPONSE_ORDER_ITEMS_PRODUCT_PRODUCT_OPTIONS_TYPE OrdersPostResponse_order_items_product_product_options_type = iota
	VARIABLE_ORDERSPOSTRESPONSE_ORDER_ITEMS_PRODUCT_PRODUCT_OPTIONS_TYPE
)

func (i OrdersPostResponse_order_items_product_product_options_type) String() string {
	return []string{"FixedList", "Variable"}[i]
}
func ParseOrdersPostResponse_order_items_product_product_options_type(v string) (any, error) {
	result := FIXEDLIST_ORDERSPOSTRESPONSE_ORDER_ITEMS_PRODUCT_PRODUCT_OPTIONS_TYPE
	switch v {
	case "FixedList":
		result = FIXEDLIST_ORDERSPOSTRESPONSE_ORDER_ITEMS_PRODUCT_PRODUCT_OPTIONS_TYPE
	case "Variable":
		result = VARIABLE_ORDERSPOSTRESPONSE_ORDER_ITEMS_PRODUCT_PRODUCT_OPTIONS_TYPE
	default:
		return 0, errors.New("Unknown OrdersPostResponse_order_items_product_product_options_type value: " + v)
	}
	return &result, nil
}
func SerializeOrdersPostResponse_order_items_product_product_options_type(values []OrdersPostResponse_order_items_product_product_options_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i OrdersPostResponse_order_items_product_product_options_type) isMultiValue() bool {
	return false
}
