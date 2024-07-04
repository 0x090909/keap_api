package payments

import (
	"errors"
)

type PaymentsPostRequestBody_payment_method_type int

const (
	CREDIT_CARD_PAYMENTSPOSTREQUESTBODY_PAYMENT_METHOD_TYPE PaymentsPostRequestBody_payment_method_type = iota
	CASH_PAYMENTSPOSTREQUESTBODY_PAYMENT_METHOD_TYPE
	CHECK_PAYMENTSPOSTREQUESTBODY_PAYMENT_METHOD_TYPE
	TOKEN_PAYMENTSPOSTREQUESTBODY_PAYMENT_METHOD_TYPE
)

func (i PaymentsPostRequestBody_payment_method_type) String() string {
	return []string{"CREDIT_CARD", "CASH", "CHECK", "TOKEN"}[i]
}
func ParsePaymentsPostRequestBody_payment_method_type(v string) (any, error) {
	result := CREDIT_CARD_PAYMENTSPOSTREQUESTBODY_PAYMENT_METHOD_TYPE
	switch v {
	case "CREDIT_CARD":
		result = CREDIT_CARD_PAYMENTSPOSTREQUESTBODY_PAYMENT_METHOD_TYPE
	case "CASH":
		result = CASH_PAYMENTSPOSTREQUESTBODY_PAYMENT_METHOD_TYPE
	case "CHECK":
		result = CHECK_PAYMENTSPOSTREQUESTBODY_PAYMENT_METHOD_TYPE
	case "TOKEN":
		result = TOKEN_PAYMENTSPOSTREQUESTBODY_PAYMENT_METHOD_TYPE
	default:
		return 0, errors.New("Unknown PaymentsPostRequestBody_payment_method_type value: " + v)
	}
	return &result, nil
}
func SerializePaymentsPostRequestBody_payment_method_type(values []PaymentsPostRequestBody_payment_method_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i PaymentsPostRequestBody_payment_method_type) isMultiValue() bool {
	return false
}
