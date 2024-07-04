package contacts

import (
	"errors"
)

type ContactsPutResponse_email_status int

const (
	UNENGAGEDMARKETABLE_CONTACTSPUTRESPONSE_EMAIL_STATUS ContactsPutResponse_email_status = iota
	SINGLEOPTIN_CONTACTSPUTRESPONSE_EMAIL_STATUS
	DOUBLEOPTIN_CONTACTSPUTRESPONSE_EMAIL_STATUS
	CONFIRMED_CONTACTSPUTRESPONSE_EMAIL_STATUS
	UNENGAGEDNONMARKETABLE_CONTACTSPUTRESPONSE_EMAIL_STATUS
	NONMARKETABLE_CONTACTSPUTRESPONSE_EMAIL_STATUS
	LOCKDOWN_CONTACTSPUTRESPONSE_EMAIL_STATUS
	BOUNCE_CONTACTSPUTRESPONSE_EMAIL_STATUS
	HARDBOUNCE_CONTACTSPUTRESPONSE_EMAIL_STATUS
	MANUAL_CONTACTSPUTRESPONSE_EMAIL_STATUS
	ADMIN_CONTACTSPUTRESPONSE_EMAIL_STATUS
	SYSTEM_CONTACTSPUTRESPONSE_EMAIL_STATUS
	LISTUNSUBSCRIBE_CONTACTSPUTRESPONSE_EMAIL_STATUS
	FEEDBACK_CONTACTSPUTRESPONSE_EMAIL_STATUS
	SPAM_CONTACTSPUTRESPONSE_EMAIL_STATUS
	INVALID_CONTACTSPUTRESPONSE_EMAIL_STATUS
	DEACTIVATED_CONTACTSPUTRESPONSE_EMAIL_STATUS
)

func (i ContactsPutResponse_email_status) String() string {
	return []string{"UnengagedMarketable", "SingleOptIn", "DoubleOptin", "Confirmed", "UnengagedNonMarketable", "NonMarketable", "Lockdown", "Bounce", "HardBounce", "Manual", "Admin", "System", "ListUnsubscribe", "Feedback", "Spam", "Invalid", "Deactivated"}[i]
}
func ParseContactsPutResponse_email_status(v string) (any, error) {
	result := UNENGAGEDMARKETABLE_CONTACTSPUTRESPONSE_EMAIL_STATUS
	switch v {
	case "UnengagedMarketable":
		result = UNENGAGEDMARKETABLE_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "SingleOptIn":
		result = SINGLEOPTIN_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "DoubleOptin":
		result = DOUBLEOPTIN_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "Confirmed":
		result = CONFIRMED_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "UnengagedNonMarketable":
		result = UNENGAGEDNONMARKETABLE_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "NonMarketable":
		result = NONMARKETABLE_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "Lockdown":
		result = LOCKDOWN_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "Bounce":
		result = BOUNCE_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "HardBounce":
		result = HARDBOUNCE_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "Manual":
		result = MANUAL_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "Admin":
		result = ADMIN_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "System":
		result = SYSTEM_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "ListUnsubscribe":
		result = LISTUNSUBSCRIBE_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "Feedback":
		result = FEEDBACK_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "Spam":
		result = SPAM_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "Invalid":
		result = INVALID_CONTACTSPUTRESPONSE_EMAIL_STATUS
	case "Deactivated":
		result = DEACTIVATED_CONTACTSPUTRESPONSE_EMAIL_STATUS
	default:
		return 0, errors.New("Unknown ContactsPutResponse_email_status value: " + v)
	}
	return &result, nil
}
func SerializeContactsPutResponse_email_status(values []ContactsPutResponse_email_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPutResponse_email_status) isMultiValue() bool {
	return false
}
