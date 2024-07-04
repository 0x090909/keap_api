package item

import (
	"errors"
)

type WithCompanyGetResponse_email_status int

const (
	UNENGAGEDMARKETABLE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS WithCompanyGetResponse_email_status = iota
	SINGLEOPTIN_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	DOUBLEOPTIN_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	CONFIRMED_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	UNENGAGEDNONMARKETABLE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	NONMARKETABLE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	LOCKDOWN_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	BOUNCE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	HARDBOUNCE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	MANUAL_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	ADMIN_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	SYSTEM_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	LISTUNSUBSCRIBE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	FEEDBACK_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	SPAM_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	INVALID_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	DEACTIVATED_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
)

func (i WithCompanyGetResponse_email_status) String() string {
	return []string{"UnengagedMarketable", "SingleOptIn", "DoubleOptin", "Confirmed", "UnengagedNonMarketable", "NonMarketable", "Lockdown", "Bounce", "HardBounce", "Manual", "Admin", "System", "ListUnsubscribe", "Feedback", "Spam", "Invalid", "Deactivated"}[i]
}
func ParseWithCompanyGetResponse_email_status(v string) (any, error) {
	result := UNENGAGEDMARKETABLE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	switch v {
	case "UnengagedMarketable":
		result = UNENGAGEDMARKETABLE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "SingleOptIn":
		result = SINGLEOPTIN_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "DoubleOptin":
		result = DOUBLEOPTIN_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "Confirmed":
		result = CONFIRMED_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "UnengagedNonMarketable":
		result = UNENGAGEDNONMARKETABLE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "NonMarketable":
		result = NONMARKETABLE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "Lockdown":
		result = LOCKDOWN_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "Bounce":
		result = BOUNCE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "HardBounce":
		result = HARDBOUNCE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "Manual":
		result = MANUAL_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "Admin":
		result = ADMIN_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "System":
		result = SYSTEM_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "ListUnsubscribe":
		result = LISTUNSUBSCRIBE_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "Feedback":
		result = FEEDBACK_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "Spam":
		result = SPAM_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "Invalid":
		result = INVALID_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	case "Deactivated":
		result = DEACTIVATED_WITHCOMPANYGETRESPONSE_EMAIL_STATUS
	default:
		return 0, errors.New("Unknown WithCompanyGetResponse_email_status value: " + v)
	}
	return &result, nil
}
func SerializeWithCompanyGetResponse_email_status(values []WithCompanyGetResponse_email_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithCompanyGetResponse_email_status) isMultiValue() bool {
	return false
}