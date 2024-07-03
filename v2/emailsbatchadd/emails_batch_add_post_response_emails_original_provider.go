package emailsbatchadd

import (
	"errors"
)

type EmailsBatchAddPostResponse_emails_original_provider int

const (
	UNKNOWN_EMAILSBATCHADDPOSTRESPONSE_EMAILS_ORIGINAL_PROVIDER EmailsBatchAddPostResponse_emails_original_provider = iota
	INFUSIONSOFT_EMAILSBATCHADDPOSTRESPONSE_EMAILS_ORIGINAL_PROVIDER
	MICROSOFT_EMAILSBATCHADDPOSTRESPONSE_EMAILS_ORIGINAL_PROVIDER
	GOOGLE_EMAILSBATCHADDPOSTRESPONSE_EMAILS_ORIGINAL_PROVIDER
)

func (i EmailsBatchAddPostResponse_emails_original_provider) String() string {
	return []string{"UNKNOWN", "INFUSIONSOFT", "MICROSOFT", "GOOGLE"}[i]
}
func ParseEmailsBatchAddPostResponse_emails_original_provider(v string) (any, error) {
	result := UNKNOWN_EMAILSBATCHADDPOSTRESPONSE_EMAILS_ORIGINAL_PROVIDER
	switch v {
	case "UNKNOWN":
		result = UNKNOWN_EMAILSBATCHADDPOSTRESPONSE_EMAILS_ORIGINAL_PROVIDER
	case "INFUSIONSOFT":
		result = INFUSIONSOFT_EMAILSBATCHADDPOSTRESPONSE_EMAILS_ORIGINAL_PROVIDER
	case "MICROSOFT":
		result = MICROSOFT_EMAILSBATCHADDPOSTRESPONSE_EMAILS_ORIGINAL_PROVIDER
	case "GOOGLE":
		result = GOOGLE_EMAILSBATCHADDPOSTRESPONSE_EMAILS_ORIGINAL_PROVIDER
	default:
		return 0, errors.New("Unknown EmailsBatchAddPostResponse_emails_original_provider value: " + v)
	}
	return &result, nil
}
func SerializeEmailsBatchAddPostResponse_emails_original_provider(values []EmailsBatchAddPostResponse_emails_original_provider) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i EmailsBatchAddPostResponse_emails_original_provider) isMultiValue() bool {
	return false
}
