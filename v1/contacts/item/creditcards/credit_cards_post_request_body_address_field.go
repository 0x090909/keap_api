package creditcards

import (
	"errors"
)

type CreditCardsPostRequestBody_address_field int

const (
	BILLING_CREDITCARDSPOSTREQUESTBODY_ADDRESS_FIELD CreditCardsPostRequestBody_address_field = iota
	SHIPPING_CREDITCARDSPOSTREQUESTBODY_ADDRESS_FIELD
	OTHER_CREDITCARDSPOSTREQUESTBODY_ADDRESS_FIELD
)

func (i CreditCardsPostRequestBody_address_field) String() string {
	return []string{"BILLING", "SHIPPING", "OTHER"}[i]
}
func ParseCreditCardsPostRequestBody_address_field(v string) (any, error) {
	result := BILLING_CREDITCARDSPOSTREQUESTBODY_ADDRESS_FIELD
	switch v {
	case "BILLING":
		result = BILLING_CREDITCARDSPOSTREQUESTBODY_ADDRESS_FIELD
	case "SHIPPING":
		result = SHIPPING_CREDITCARDSPOSTREQUESTBODY_ADDRESS_FIELD
	case "OTHER":
		result = OTHER_CREDITCARDSPOSTREQUESTBODY_ADDRESS_FIELD
	default:
		return 0, errors.New("Unknown CreditCardsPostRequestBody_address_field value: " + v)
	}
	return &result, nil
}
func SerializeCreditCardsPostRequestBody_address_field(values []CreditCardsPostRequestBody_address_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CreditCardsPostRequestBody_address_field) isMultiValue() bool {
	return false
}
