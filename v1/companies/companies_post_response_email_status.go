package companies

import (
	"errors"
)

type CompaniesPostResponse_email_status int

const (
	UNENGAGEDMARKETABLE_COMPANIESPOSTRESPONSE_EMAIL_STATUS CompaniesPostResponse_email_status = iota
	SINGLEOPTIN_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	DOUBLEOPTIN_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	CONFIRMED_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	UNENGAGEDNONMARKETABLE_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	NONMARKETABLE_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	LOCKDOWN_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	BOUNCE_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	HARDBOUNCE_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	MANUAL_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	ADMIN_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	SYSTEM_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	LISTUNSUBSCRIBE_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	FEEDBACK_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	SPAM_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	INVALID_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	DEACTIVATED_COMPANIESPOSTRESPONSE_EMAIL_STATUS
)

func (i CompaniesPostResponse_email_status) String() string {
	return []string{"UnengagedMarketable", "SingleOptIn", "DoubleOptin", "Confirmed", "UnengagedNonMarketable", "NonMarketable", "Lockdown", "Bounce", "HardBounce", "Manual", "Admin", "System", "ListUnsubscribe", "Feedback", "Spam", "Invalid", "Deactivated"}[i]
}
func ParseCompaniesPostResponse_email_status(v string) (any, error) {
	result := UNENGAGEDMARKETABLE_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	switch v {
	case "UnengagedMarketable":
		result = UNENGAGEDMARKETABLE_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "SingleOptIn":
		result = SINGLEOPTIN_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "DoubleOptin":
		result = DOUBLEOPTIN_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "Confirmed":
		result = CONFIRMED_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "UnengagedNonMarketable":
		result = UNENGAGEDNONMARKETABLE_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "NonMarketable":
		result = NONMARKETABLE_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "Lockdown":
		result = LOCKDOWN_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "Bounce":
		result = BOUNCE_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "HardBounce":
		result = HARDBOUNCE_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "Manual":
		result = MANUAL_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "Admin":
		result = ADMIN_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "System":
		result = SYSTEM_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "ListUnsubscribe":
		result = LISTUNSUBSCRIBE_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "Feedback":
		result = FEEDBACK_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "Spam":
		result = SPAM_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "Invalid":
		result = INVALID_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	case "Deactivated":
		result = DEACTIVATED_COMPANIESPOSTRESPONSE_EMAIL_STATUS
	default:
		return 0, errors.New("Unknown CompaniesPostResponse_email_status value: " + v)
	}
	return &result, nil
}
func SerializeCompaniesPostResponse_email_status(values []CompaniesPostResponse_email_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CompaniesPostResponse_email_status) isMultiValue() bool {
	return false
}
