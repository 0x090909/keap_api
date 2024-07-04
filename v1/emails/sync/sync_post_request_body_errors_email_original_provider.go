package sync

import (
	"errors"
)

// Provider that sent the email case insensitive, must be in list [GOOGLE, INFUSIONSOFT].  If omitted gets mapped to UNKNOWN.
type SyncPostRequestBody_errors_email_original_provider int

const (
	UNKNOWN_SYNCPOSTREQUESTBODY_ERRORS_EMAIL_ORIGINAL_PROVIDER SyncPostRequestBody_errors_email_original_provider = iota
	INFUSIONSOFT_SYNCPOSTREQUESTBODY_ERRORS_EMAIL_ORIGINAL_PROVIDER
	MICROSOFT_SYNCPOSTREQUESTBODY_ERRORS_EMAIL_ORIGINAL_PROVIDER
	GOOGLE_SYNCPOSTREQUESTBODY_ERRORS_EMAIL_ORIGINAL_PROVIDER
)

func (i SyncPostRequestBody_errors_email_original_provider) String() string {
	return []string{"UNKNOWN", "INFUSIONSOFT", "MICROSOFT", "GOOGLE"}[i]
}
func ParseSyncPostRequestBody_errors_email_original_provider(v string) (any, error) {
	result := UNKNOWN_SYNCPOSTREQUESTBODY_ERRORS_EMAIL_ORIGINAL_PROVIDER
	switch v {
	case "UNKNOWN":
		result = UNKNOWN_SYNCPOSTREQUESTBODY_ERRORS_EMAIL_ORIGINAL_PROVIDER
	case "INFUSIONSOFT":
		result = INFUSIONSOFT_SYNCPOSTREQUESTBODY_ERRORS_EMAIL_ORIGINAL_PROVIDER
	case "MICROSOFT":
		result = MICROSOFT_SYNCPOSTREQUESTBODY_ERRORS_EMAIL_ORIGINAL_PROVIDER
	case "GOOGLE":
		result = GOOGLE_SYNCPOSTREQUESTBODY_ERRORS_EMAIL_ORIGINAL_PROVIDER
	default:
		return 0, errors.New("Unknown SyncPostRequestBody_errors_email_original_provider value: " + v)
	}
	return &result, nil
}
func SerializeSyncPostRequestBody_errors_email_original_provider(values []SyncPostRequestBody_errors_email_original_provider) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i SyncPostRequestBody_errors_email_original_provider) isMultiValue() bool {
	return false
}
