package emails

import (
	"errors"
)

type EmailsGetResponse_emails_original_provider int

const (
	UNKNOWN_EMAILSGETRESPONSE_EMAILS_ORIGINAL_PROVIDER EmailsGetResponse_emails_original_provider = iota
	INFUSIONSOFT_EMAILSGETRESPONSE_EMAILS_ORIGINAL_PROVIDER
	MICROSOFT_EMAILSGETRESPONSE_EMAILS_ORIGINAL_PROVIDER
	GOOGLE_EMAILSGETRESPONSE_EMAILS_ORIGINAL_PROVIDER
)

func (i EmailsGetResponse_emails_original_provider) String() string {
	return []string{"UNKNOWN", "INFUSIONSOFT", "MICROSOFT", "GOOGLE"}[i]
}
func ParseEmailsGetResponse_emails_original_provider(v string) (any, error) {
	result := UNKNOWN_EMAILSGETRESPONSE_EMAILS_ORIGINAL_PROVIDER
	switch v {
	case "UNKNOWN":
		result = UNKNOWN_EMAILSGETRESPONSE_EMAILS_ORIGINAL_PROVIDER
	case "INFUSIONSOFT":
		result = INFUSIONSOFT_EMAILSGETRESPONSE_EMAILS_ORIGINAL_PROVIDER
	case "MICROSOFT":
		result = MICROSOFT_EMAILSGETRESPONSE_EMAILS_ORIGINAL_PROVIDER
	case "GOOGLE":
		result = GOOGLE_EMAILSGETRESPONSE_EMAILS_ORIGINAL_PROVIDER
	default:
		return 0, errors.New("Unknown EmailsGetResponse_emails_original_provider value: " + v)
	}
	return &result, nil
}
func SerializeEmailsGetResponse_emails_original_provider(values []EmailsGetResponse_emails_original_provider) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i EmailsGetResponse_emails_original_provider) isMultiValue() bool {
	return false
}
