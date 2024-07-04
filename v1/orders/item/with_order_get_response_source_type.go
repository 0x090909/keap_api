package item

import (
	"errors"
)

type WithOrderGetResponse_source_type int

const (
	INVOICE_WITHORDERGETRESPONSE_SOURCE_TYPE WithOrderGetResponse_source_type = iota
	API_WITHORDERGETRESPONSE_SOURCE_TYPE
	CHECKOUT_FORM_WITHORDERGETRESPONSE_SOURCE_TYPE
	MANUAL_PAYMENT_WITHORDERGETRESPONSE_SOURCE_TYPE
	UNKNOWN_WITHORDERGETRESPONSE_SOURCE_TYPE
	QBO_SYNC_WITHORDERGETRESPONSE_SOURCE_TYPE
)

func (i WithOrderGetResponse_source_type) String() string {
	return []string{"INVOICE", "API", "CHECKOUT_FORM", "MANUAL_PAYMENT", "UNKNOWN", "QBO_SYNC"}[i]
}
func ParseWithOrderGetResponse_source_type(v string) (any, error) {
	result := INVOICE_WITHORDERGETRESPONSE_SOURCE_TYPE
	switch v {
	case "INVOICE":
		result = INVOICE_WITHORDERGETRESPONSE_SOURCE_TYPE
	case "API":
		result = API_WITHORDERGETRESPONSE_SOURCE_TYPE
	case "CHECKOUT_FORM":
		result = CHECKOUT_FORM_WITHORDERGETRESPONSE_SOURCE_TYPE
	case "MANUAL_PAYMENT":
		result = MANUAL_PAYMENT_WITHORDERGETRESPONSE_SOURCE_TYPE
	case "UNKNOWN":
		result = UNKNOWN_WITHORDERGETRESPONSE_SOURCE_TYPE
	case "QBO_SYNC":
		result = QBO_SYNC_WITHORDERGETRESPONSE_SOURCE_TYPE
	default:
		return 0, errors.New("Unknown WithOrderGetResponse_source_type value: " + v)
	}
	return &result, nil
}
func SerializeWithOrderGetResponse_source_type(values []WithOrderGetResponse_source_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithOrderGetResponse_source_type) isMultiValue() bool {
	return false
}
