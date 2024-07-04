package item

import (
	"errors"
)

type WithTransactionGetResponse_orders_source_type int

const (
	INVOICE_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE WithTransactionGetResponse_orders_source_type = iota
	API_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE
	CHECKOUT_FORM_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE
	MANUAL_PAYMENT_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE
	UNKNOWN_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE
	QBO_SYNC_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE
)

func (i WithTransactionGetResponse_orders_source_type) String() string {
	return []string{"INVOICE", "API", "CHECKOUT_FORM", "MANUAL_PAYMENT", "UNKNOWN", "QBO_SYNC"}[i]
}
func ParseWithTransactionGetResponse_orders_source_type(v string) (any, error) {
	result := INVOICE_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE
	switch v {
	case "INVOICE":
		result = INVOICE_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE
	case "API":
		result = API_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE
	case "CHECKOUT_FORM":
		result = CHECKOUT_FORM_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE
	case "MANUAL_PAYMENT":
		result = MANUAL_PAYMENT_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE
	case "UNKNOWN":
		result = UNKNOWN_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE
	case "QBO_SYNC":
		result = QBO_SYNC_WITHTRANSACTIONGETRESPONSE_ORDERS_SOURCE_TYPE
	default:
		return 0, errors.New("Unknown WithTransactionGetResponse_orders_source_type value: " + v)
	}
	return &result, nil
}
func SerializeWithTransactionGetResponse_orders_source_type(values []WithTransactionGetResponse_orders_source_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithTransactionGetResponse_orders_source_type) isMultiValue() bool {
	return false
}
