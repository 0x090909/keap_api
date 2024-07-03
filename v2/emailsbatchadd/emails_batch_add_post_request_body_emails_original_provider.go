package emailsbatchadd

import (
	"errors"
)

// Provider that sent the email, defaults to UNKNOWN
type EmailsBatchAddPostRequestBody_emails_original_provider int

const (
	UNKNOWN_EMAILSBATCHADDPOSTREQUESTBODY_EMAILS_ORIGINAL_PROVIDER EmailsBatchAddPostRequestBody_emails_original_provider = iota
	INFUSIONSOFT_EMAILSBATCHADDPOSTREQUESTBODY_EMAILS_ORIGINAL_PROVIDER
	MICROSOFT_EMAILSBATCHADDPOSTREQUESTBODY_EMAILS_ORIGINAL_PROVIDER
	GOOGLE_EMAILSBATCHADDPOSTREQUESTBODY_EMAILS_ORIGINAL_PROVIDER
)

func (i EmailsBatchAddPostRequestBody_emails_original_provider) String() string {
	return []string{"UNKNOWN", "INFUSIONSOFT", "MICROSOFT", "GOOGLE"}[i]
}
func ParseEmailsBatchAddPostRequestBody_emails_original_provider(v string) (any, error) {
	result := UNKNOWN_EMAILSBATCHADDPOSTREQUESTBODY_EMAILS_ORIGINAL_PROVIDER
	switch v {
	case "UNKNOWN":
		result = UNKNOWN_EMAILSBATCHADDPOSTREQUESTBODY_EMAILS_ORIGINAL_PROVIDER
	case "INFUSIONSOFT":
		result = INFUSIONSOFT_EMAILSBATCHADDPOSTREQUESTBODY_EMAILS_ORIGINAL_PROVIDER
	case "MICROSOFT":
		result = MICROSOFT_EMAILSBATCHADDPOSTREQUESTBODY_EMAILS_ORIGINAL_PROVIDER
	case "GOOGLE":
		result = GOOGLE_EMAILSBATCHADDPOSTREQUESTBODY_EMAILS_ORIGINAL_PROVIDER
	default:
		return 0, errors.New("Unknown EmailsBatchAddPostRequestBody_emails_original_provider value: " + v)
	}
	return &result, nil
}
func SerializeEmailsBatchAddPostRequestBody_emails_original_provider(values []EmailsBatchAddPostRequestBody_emails_original_provider) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i EmailsBatchAddPostRequestBody_emails_original_provider) isMultiValue() bool {
	return false
}
