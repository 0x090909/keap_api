package emails

import (
	"errors"
)

type EmailsPostResponse_original_provider int

const (
	UNKNOWN_EMAILSPOSTRESPONSE_ORIGINAL_PROVIDER EmailsPostResponse_original_provider = iota
	INFUSIONSOFT_EMAILSPOSTRESPONSE_ORIGINAL_PROVIDER
	MICROSOFT_EMAILSPOSTRESPONSE_ORIGINAL_PROVIDER
	GOOGLE_EMAILSPOSTRESPONSE_ORIGINAL_PROVIDER
)

func (i EmailsPostResponse_original_provider) String() string {
	return []string{"UNKNOWN", "INFUSIONSOFT", "MICROSOFT", "GOOGLE"}[i]
}
func ParseEmailsPostResponse_original_provider(v string) (any, error) {
	result := UNKNOWN_EMAILSPOSTRESPONSE_ORIGINAL_PROVIDER
	switch v {
	case "UNKNOWN":
		result = UNKNOWN_EMAILSPOSTRESPONSE_ORIGINAL_PROVIDER
	case "INFUSIONSOFT":
		result = INFUSIONSOFT_EMAILSPOSTRESPONSE_ORIGINAL_PROVIDER
	case "MICROSOFT":
		result = MICROSOFT_EMAILSPOSTRESPONSE_ORIGINAL_PROVIDER
	case "GOOGLE":
		result = GOOGLE_EMAILSPOSTRESPONSE_ORIGINAL_PROVIDER
	default:
		return 0, errors.New("Unknown EmailsPostResponse_original_provider value: " + v)
	}
	return &result, nil
}
func SerializeEmailsPostResponse_original_provider(values []EmailsPostResponse_original_provider) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i EmailsPostResponse_original_provider) isMultiValue() bool {
	return false
}
