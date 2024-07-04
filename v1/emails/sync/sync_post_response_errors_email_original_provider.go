package sync

import (
	"errors"
)

// Provider that sent the email case insensitive, must be in list [GOOGLE, INFUSIONSOFT].  If omitted gets mapped to UNKNOWN.
type SyncPostResponse_errors_email_original_provider int

const (
	UNKNOWN_SYNCPOSTRESPONSE_ERRORS_EMAIL_ORIGINAL_PROVIDER SyncPostResponse_errors_email_original_provider = iota
	INFUSIONSOFT_SYNCPOSTRESPONSE_ERRORS_EMAIL_ORIGINAL_PROVIDER
	MICROSOFT_SYNCPOSTRESPONSE_ERRORS_EMAIL_ORIGINAL_PROVIDER
	GOOGLE_SYNCPOSTRESPONSE_ERRORS_EMAIL_ORIGINAL_PROVIDER
)

func (i SyncPostResponse_errors_email_original_provider) String() string {
	return []string{"UNKNOWN", "INFUSIONSOFT", "MICROSOFT", "GOOGLE"}[i]
}
func ParseSyncPostResponse_errors_email_original_provider(v string) (any, error) {
	result := UNKNOWN_SYNCPOSTRESPONSE_ERRORS_EMAIL_ORIGINAL_PROVIDER
	switch v {
	case "UNKNOWN":
		result = UNKNOWN_SYNCPOSTRESPONSE_ERRORS_EMAIL_ORIGINAL_PROVIDER
	case "INFUSIONSOFT":
		result = INFUSIONSOFT_SYNCPOSTRESPONSE_ERRORS_EMAIL_ORIGINAL_PROVIDER
	case "MICROSOFT":
		result = MICROSOFT_SYNCPOSTRESPONSE_ERRORS_EMAIL_ORIGINAL_PROVIDER
	case "GOOGLE":
		result = GOOGLE_SYNCPOSTRESPONSE_ERRORS_EMAIL_ORIGINAL_PROVIDER
	default:
		return 0, errors.New("Unknown SyncPostResponse_errors_email_original_provider value: " + v)
	}
	return &result, nil
}
func SerializeSyncPostResponse_errors_email_original_provider(values []SyncPostResponse_errors_email_original_provider) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i SyncPostResponse_errors_email_original_provider) isMultiValue() bool {
	return false
}
