package emails

import (
	"errors"
)

// Provider that sent the email, defaults to UNKNOWN
type EmailsPostRequestBody_original_provider int

const (
	UNKNOWN_EMAILSPOSTREQUESTBODY_ORIGINAL_PROVIDER EmailsPostRequestBody_original_provider = iota
	INFUSIONSOFT_EMAILSPOSTREQUESTBODY_ORIGINAL_PROVIDER
	MICROSOFT_EMAILSPOSTREQUESTBODY_ORIGINAL_PROVIDER
	GOOGLE_EMAILSPOSTREQUESTBODY_ORIGINAL_PROVIDER
)

func (i EmailsPostRequestBody_original_provider) String() string {
	return []string{"UNKNOWN", "INFUSIONSOFT", "MICROSOFT", "GOOGLE"}[i]
}
func ParseEmailsPostRequestBody_original_provider(v string) (any, error) {
	result := UNKNOWN_EMAILSPOSTREQUESTBODY_ORIGINAL_PROVIDER
	switch v {
	case "UNKNOWN":
		result = UNKNOWN_EMAILSPOSTREQUESTBODY_ORIGINAL_PROVIDER
	case "INFUSIONSOFT":
		result = INFUSIONSOFT_EMAILSPOSTREQUESTBODY_ORIGINAL_PROVIDER
	case "MICROSOFT":
		result = MICROSOFT_EMAILSPOSTREQUESTBODY_ORIGINAL_PROVIDER
	case "GOOGLE":
		result = GOOGLE_EMAILSPOSTREQUESTBODY_ORIGINAL_PROVIDER
	default:
		return 0, errors.New("Unknown EmailsPostRequestBody_original_provider value: " + v)
	}
	return &result, nil
}
func SerializeEmailsPostRequestBody_original_provider(values []EmailsPostRequestBody_original_provider) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i EmailsPostRequestBody_original_provider) isMultiValue() bool {
	return false
}
