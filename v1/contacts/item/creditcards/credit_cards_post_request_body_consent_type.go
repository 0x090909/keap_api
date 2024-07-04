package creditcards

import (
	"errors"
)

type CreditCardsPostRequestBody_consent_type int

const (
	RECURRING_CONSENT_CREDITCARDSPOSTREQUESTBODY_CONSENT_TYPE CreditCardsPostRequestBody_consent_type = iota
	IMPLICIT_CONSENT_CREDITCARDSPOSTREQUESTBODY_CONSENT_TYPE
	EXPLICIT_CONSENT_CREDITCARDSPOSTREQUESTBODY_CONSENT_TYPE
	NO_CONSENT_CREDITCARDSPOSTREQUESTBODY_CONSENT_TYPE
)

func (i CreditCardsPostRequestBody_consent_type) String() string {
	return []string{"RECURRING_CONSENT", "IMPLICIT_CONSENT", "EXPLICIT_CONSENT", "NO_CONSENT"}[i]
}
func ParseCreditCardsPostRequestBody_consent_type(v string) (any, error) {
	result := RECURRING_CONSENT_CREDITCARDSPOSTREQUESTBODY_CONSENT_TYPE
	switch v {
	case "RECURRING_CONSENT":
		result = RECURRING_CONSENT_CREDITCARDSPOSTREQUESTBODY_CONSENT_TYPE
	case "IMPLICIT_CONSENT":
		result = IMPLICIT_CONSENT_CREDITCARDSPOSTREQUESTBODY_CONSENT_TYPE
	case "EXPLICIT_CONSENT":
		result = EXPLICIT_CONSENT_CREDITCARDSPOSTREQUESTBODY_CONSENT_TYPE
	case "NO_CONSENT":
		result = NO_CONSENT_CREDITCARDSPOSTREQUESTBODY_CONSENT_TYPE
	default:
		return 0, errors.New("Unknown CreditCardsPostRequestBody_consent_type value: " + v)
	}
	return &result, nil
}
func SerializeCreditCardsPostRequestBody_consent_type(values []CreditCardsPostRequestBody_consent_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CreditCardsPostRequestBody_consent_type) isMultiValue() bool {
	return false
}
