package creditcards

import (
	"errors"
)

type CreditCardsPostResponse_address_field int

const (
	BILLING_CREDITCARDSPOSTRESPONSE_ADDRESS_FIELD CreditCardsPostResponse_address_field = iota
	SHIPPING_CREDITCARDSPOSTRESPONSE_ADDRESS_FIELD
	OTHER_CREDITCARDSPOSTRESPONSE_ADDRESS_FIELD
)

func (i CreditCardsPostResponse_address_field) String() string {
	return []string{"BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseCreditCardsPostResponse_address_field(v string) (any, error) {
	result := BILLING_CREDITCARDSPOSTRESPONSE_ADDRESS_FIELD
	switch v {
	case "BILLING":
		result = BILLING_CREDITCARDSPOSTRESPONSE_ADDRESS_FIELD
	case "SHIPPING":
		result = SHIPPING_CREDITCARDSPOSTRESPONSE_ADDRESS_FIELD
	case "OTHER":
		result = OTHER_CREDITCARDSPOSTRESPONSE_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown CreditCardsPostResponse_address_field value: " + v)
	}
	return &result, nil
}
func SerializeCreditCardsPostResponse_address_field(values []CreditCardsPostResponse_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CreditCardsPostResponse_address_field) isMultiValue() bool {
	return false
}
