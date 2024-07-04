package item

import (
	"errors"
)

type ContactIdGetResponse_email_status int

const (
	UNENGAGEDMARKETABLE_CONTACTIDGETRESPONSE_EMAIL_STATUS ContactIdGetResponse_email_status = iota
	SINGLEOPTIN_CONTACTIDGETRESPONSE_EMAIL_STATUS
	DOUBLEOPTIN_CONTACTIDGETRESPONSE_EMAIL_STATUS
	CONFIRMED_CONTACTIDGETRESPONSE_EMAIL_STATUS
	UNENGAGEDNONMARKETABLE_CONTACTIDGETRESPONSE_EMAIL_STATUS
	NONMARKETABLE_CONTACTIDGETRESPONSE_EMAIL_STATUS
	LOCKDOWN_CONTACTIDGETRESPONSE_EMAIL_STATUS
	BOUNCE_CONTACTIDGETRESPONSE_EMAIL_STATUS
	HARDBOUNCE_CONTACTIDGETRESPONSE_EMAIL_STATUS
	MANUAL_CONTACTIDGETRESPONSE_EMAIL_STATUS
	ADMIN_CONTACTIDGETRESPONSE_EMAIL_STATUS
	SYSTEM_CONTACTIDGETRESPONSE_EMAIL_STATUS
	LISTUNSUBSCRIBE_CONTACTIDGETRESPONSE_EMAIL_STATUS
	FEEDBACK_CONTACTIDGETRESPONSE_EMAIL_STATUS
	SPAM_CONTACTIDGETRESPONSE_EMAIL_STATUS
	INVALID_CONTACTIDGETRESPONSE_EMAIL_STATUS
	DEACTIVATED_CONTACTIDGETRESPONSE_EMAIL_STATUS
)

func (i ContactIdGetResponse_email_status) String() string {
	return []string{"UnengagedMarketable", "SingleOptIn", "DoubleOptin", "Confirmed", "UnengagedNonMarketable", "NonMarketable", "Lockdown", "Bounce", "HardBounce", "Manual", "Admin", "System", "ListUnsubscribe", "Feedback", "Spam", "Invalid", "Deactivated"}[i]
}
func ParseContactIdGetResponse_email_status(v string) (any, error) {
	result := UNENGAGEDMARKETABLE_CONTACTIDGETRESPONSE_EMAIL_STATUS
	switch v {
	case "UnengagedMarketable":
		result = UNENGAGEDMARKETABLE_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "SingleOptIn":
		result = SINGLEOPTIN_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "DoubleOptin":
		result = DOUBLEOPTIN_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "Confirmed":
		result = CONFIRMED_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "UnengagedNonMarketable":
		result = UNENGAGEDNONMARKETABLE_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "NonMarketable":
		result = NONMARKETABLE_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "Lockdown":
		result = LOCKDOWN_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "Bounce":
		result = BOUNCE_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "HardBounce":
		result = HARDBOUNCE_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "Manual":
		result = MANUAL_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "Admin":
		result = ADMIN_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "System":
		result = SYSTEM_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "ListUnsubscribe":
		result = LISTUNSUBSCRIBE_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "Feedback":
		result = FEEDBACK_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "Spam":
		result = SPAM_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "Invalid":
		result = INVALID_CONTACTIDGETRESPONSE_EMAIL_STATUS
	case "Deactivated":
		result = DEACTIVATED_CONTACTIDGETRESPONSE_EMAIL_STATUS
	default:
		return 0, errors.New("Unknown ContactIdGetResponse_email_status value: " + v)
	}
	return &result, nil
}
func SerializeContactIdGetResponse_email_status(values []ContactIdGetResponse_email_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactIdGetResponse_email_status) isMultiValue() bool {
	return false
}
