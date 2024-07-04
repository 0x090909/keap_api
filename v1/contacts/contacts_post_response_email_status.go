package contacts

import (
	"errors"
)

type ContactsPostResponse_email_status int

const (
	UNENGAGEDMARKETABLE_CONTACTSPOSTRESPONSE_EMAIL_STATUS ContactsPostResponse_email_status = iota
	SINGLEOPTIN_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	DOUBLEOPTIN_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	CONFIRMED_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	UNENGAGEDNONMARKETABLE_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	NONMARKETABLE_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	LOCKDOWN_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	BOUNCE_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	HARDBOUNCE_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	MANUAL_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	ADMIN_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	SYSTEM_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	LISTUNSUBSCRIBE_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	FEEDBACK_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	SPAM_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	INVALID_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	DEACTIVATED_CONTACTSPOSTRESPONSE_EMAIL_STATUS
)

func (i ContactsPostResponse_email_status) String() string {
	return []string{"UnengagedMarketable", "SingleOptIn", "DoubleOptin", "Confirmed", "UnengagedNonMarketable", "NonMarketable", "Lockdown", "Bounce", "HardBounce", "Manual", "Admin", "System", "ListUnsubscribe", "Feedback", "Spam", "Invalid", "Deactivated"}[i]
}
func ParseContactsPostResponse_email_status(v string) (any, error) {
	result := UNENGAGEDMARKETABLE_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	switch v {
	case "UnengagedMarketable":
		result = UNENGAGEDMARKETABLE_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "SingleOptIn":
		result = SINGLEOPTIN_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "DoubleOptin":
		result = DOUBLEOPTIN_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "Confirmed":
		result = CONFIRMED_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "UnengagedNonMarketable":
		result = UNENGAGEDNONMARKETABLE_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "NonMarketable":
		result = NONMARKETABLE_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "Lockdown":
		result = LOCKDOWN_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "Bounce":
		result = BOUNCE_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "HardBounce":
		result = HARDBOUNCE_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "Manual":
		result = MANUAL_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "Admin":
		result = ADMIN_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "System":
		result = SYSTEM_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "ListUnsubscribe":
		result = LISTUNSUBSCRIBE_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "Feedback":
		result = FEEDBACK_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "Spam":
		result = SPAM_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "Invalid":
		result = INVALID_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	case "Deactivated":
		result = DEACTIVATED_CONTACTSPOSTRESPONSE_EMAIL_STATUS
	default:
		return 0, errors.New("Unknown ContactsPostResponse_email_status value: " + v)
	}
	return &result, nil
}
func SerializeContactsPostResponse_email_status(values []ContactsPostResponse_email_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPostResponse_email_status) isMultiValue() bool {
	return false
}
