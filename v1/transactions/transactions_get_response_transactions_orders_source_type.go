package transactions

import (
	"errors"
)

type TransactionsGetResponse_transactions_orders_source_type int

const (
	INVOICE_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE TransactionsGetResponse_transactions_orders_source_type = iota
	API_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE
	CHECKOUT_FORM_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE
	MANUAL_PAYMENT_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE
	UNKNOWN_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE
	QBO_SYNC_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE
)

func (i TransactionsGetResponse_transactions_orders_source_type) String() string {
	return []string{"INVOICE", "API", "CHECKOUT_FORM", "MANUAL_PAYMENT", "UNKNOWN", "QBO_SYNC"}[i]
}
func ParseTransactionsGetResponse_transactions_orders_source_type(v string) (any, error) {
	result := INVOICE_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE
	switch v {
	case "INVOICE":
		result = INVOICE_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE
	case "API":
		result = API_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE
	case "CHECKOUT_FORM":
		result = CHECKOUT_FORM_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE
	case "MANUAL_PAYMENT":
		result = MANUAL_PAYMENT_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE
	case "UNKNOWN":
		result = UNKNOWN_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE
	case "QBO_SYNC":
		result = QBO_SYNC_TRANSACTIONSGETRESPONSE_TRANSACTIONS_ORDERS_SOURCE_TYPE
	default:
		return 0, errors.New("Unknown TransactionsGetResponse_transactions_orders_source_type value: " + v)
	}
	return &result, nil
}
func SerializeTransactionsGetResponse_transactions_orders_source_type(values []TransactionsGetResponse_transactions_orders_source_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i TransactionsGetResponse_transactions_orders_source_type) isMultiValue() bool {
	return false
}
