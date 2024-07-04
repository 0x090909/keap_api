package orders

import (
	"errors"
)

type OrdersGetResponse_orders_source_type int

const (
	INVOICE_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE OrdersGetResponse_orders_source_type = iota
	API_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE
	CHECKOUT_FORM_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE
	MANUAL_PAYMENT_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE
	UNKNOWN_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE
	QBO_SYNC_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE
)

func (i OrdersGetResponse_orders_source_type) String() string {
	return []string{"INVOICE", "API", "CHECKOUT_FORM", "MANUAL_PAYMENT", "UNKNOWN", "QBO_SYNC"}[i]
}
func ParseOrdersGetResponse_orders_source_type(v string) (any, error) {
	result := INVOICE_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE
	switch v {
	case "INVOICE":
		result = INVOICE_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE
	case "API":
		result = API_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE
	case "CHECKOUT_FORM":
		result = CHECKOUT_FORM_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE
	case "MANUAL_PAYMENT":
		result = MANUAL_PAYMENT_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE
	case "UNKNOWN":
		result = UNKNOWN_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE
	case "QBO_SYNC":
		result = QBO_SYNC_ORDERSGETRESPONSE_ORDERS_SOURCE_TYPE
	default:
		return 0, errors.New("Unknown OrdersGetResponse_orders_source_type value: " + v)
	}
	return &result, nil
}
func SerializeOrdersGetResponse_orders_source_type(values []OrdersGetResponse_orders_source_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i OrdersGetResponse_orders_source_type) isMultiValue() bool {
	return false
}
