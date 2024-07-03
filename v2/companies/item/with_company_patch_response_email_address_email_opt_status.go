package item

import (
	"errors"
)

type WithCompany_PatchResponse_email_address_email_opt_status int

const (
	UNENGAGED_MARKETABLE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS WithCompany_PatchResponse_email_address_email_opt_status = iota
	SINGLE_OPT_IN_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	DOUBLE_OPT_IN_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	CONFIRMED_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	UNENGAGED_NON_MARKETABLE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	NON_MARKETABLE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	LOCKDOWN_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	BOUNCE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	HARD_BOUNCE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	MANUAL_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	ADMIN_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	SYSTEM_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	LIST_UNSUBSCRIBE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	FEEDBACK_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	SPAM_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	INVALID_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	DEACTIVATED_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
)

func (i WithCompany_PatchResponse_email_address_email_opt_status) String() string {
	return []string{"UNENGAGED_MARKETABLE", "SINGLE_OPT_IN", "DOUBLE_OPT_IN", "CONFIRMED", "UNENGAGED_NON_MARKETABLE", "NON_MARKETABLE", "LOCKDOWN", "BOUNCE", "HARD_BOUNCE", "MANUAL", "ADMIN", "SYSTEM", "LIST_UNSUBSCRIBE", "FEEDBACK", "SPAM", "INVALID", "DEACTIVATED"}[i]
}
func ParseWithCompany_PatchResponse_email_address_email_opt_status(v string) (any, error) {
	result := UNENGAGED_MARKETABLE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	switch v {
	case "UNENGAGED_MARKETABLE":
		result = UNENGAGED_MARKETABLE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "SINGLE_OPT_IN":
		result = SINGLE_OPT_IN_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "DOUBLE_OPT_IN":
		result = DOUBLE_OPT_IN_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "CONFIRMED":
		result = CONFIRMED_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "UNENGAGED_NON_MARKETABLE":
		result = UNENGAGED_NON_MARKETABLE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "NON_MARKETABLE":
		result = NON_MARKETABLE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "LOCKDOWN":
		result = LOCKDOWN_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "BOUNCE":
		result = BOUNCE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "HARD_BOUNCE":
		result = HARD_BOUNCE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "MANUAL":
		result = MANUAL_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "ADMIN":
		result = ADMIN_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "SYSTEM":
		result = SYSTEM_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "LIST_UNSUBSCRIBE":
		result = LIST_UNSUBSCRIBE_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "FEEDBACK":
		result = FEEDBACK_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "SPAM":
		result = SPAM_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "INVALID":
		result = INVALID_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	case "DEACTIVATED":
		result = DEACTIVATED_WITHCOMPANY_PATCHRESPONSE_EMAIL_ADDRESS_EMAIL_OPT_STATUS
	default:
		return 0, errors.New("Unknown WithCompany_PatchResponse_email_address_email_opt_status value: " + v)
	}
	return &result, nil
}
func SerializeWithCompany_PatchResponse_email_address_email_opt_status(values []WithCompany_PatchResponse_email_address_email_opt_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithCompany_PatchResponse_email_address_email_opt_status) isMultiValue() bool {
	return false
}