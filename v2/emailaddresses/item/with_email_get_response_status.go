package item

import (
	"errors"
)

type WithEmailGetResponse_status int

const (
	UNENGAGEDMARKETABLE_WITHEMAILGETRESPONSE_STATUS WithEmailGetResponse_status = iota
	UNCONFIRMED_WITHEMAILGETRESPONSE_STATUS
	CONFIRMEDLEGACY_WITHEMAILGETRESPONSE_STATUS
	CONFIRMED_WITHEMAILGETRESPONSE_STATUS
	UNENGAGEDNONMARKETABLE_WITHEMAILGETRESPONSE_STATUS
	NONMARKETABLE_WITHEMAILGETRESPONSE_STATUS
	LOCKDOWN_WITHEMAILGETRESPONSE_STATUS
	SOFTBOUNCE_WITHEMAILGETRESPONSE_STATUS
	HARDBOUNCE_WITHEMAILGETRESPONSE_STATUS
	OPTOUT_WITHEMAILGETRESPONSE_STATUS
	OPTOUTADMIN_WITHEMAILGETRESPONSE_STATUS
	OPTOUTSYSTEM_WITHEMAILGETRESPONSE_STATUS
	LISTUNSUBSCRIBE_WITHEMAILGETRESPONSE_STATUS
	PROVIDEDFEEDBACK_WITHEMAILGETRESPONSE_STATUS
	REPORTEDSPAM_WITHEMAILGETRESPONSE_STATUS
	INVALIDEMAIL_WITHEMAILGETRESPONSE_STATUS
	DEACTIVATEDDELINQUENTMAILBOX_WITHEMAILGETRESPONSE_STATUS
)

func (i WithEmailGetResponse_status) String() string {
	return []string{"Unengaged Marketable", "Unconfirmed", "Confirmed (Legacy)", "Confirmed", "Unengaged NonMarketable", "Non-marketable", "Lockdown", "Soft Bounce", "Hard Bounce", "Opt-Out", "Opt-Out: Admin", "Opt-Out: System", "List Unsubscribe", "Provided Feedback", "Reported Spam", "Invalid Email", "Deactivated/Delinquent Mailbox"}[i]
}
func ParseWithEmailGetResponse_status(v string) (any, error) {
	result := UNENGAGEDMARKETABLE_WITHEMAILGETRESPONSE_STATUS
	switch v {
	case "Unengaged Marketable":
		result = UNENGAGEDMARKETABLE_WITHEMAILGETRESPONSE_STATUS
	case "Unconfirmed":
		result = UNCONFIRMED_WITHEMAILGETRESPONSE_STATUS
	case "Confirmed (Legacy)":
		result = CONFIRMEDLEGACY_WITHEMAILGETRESPONSE_STATUS
	case "Confirmed":
		result = CONFIRMED_WITHEMAILGETRESPONSE_STATUS
	case "Unengaged NonMarketable":
		result = UNENGAGEDNONMARKETABLE_WITHEMAILGETRESPONSE_STATUS
	case "Non-marketable":
		result = NONMARKETABLE_WITHEMAILGETRESPONSE_STATUS
	case "Lockdown":
		result = LOCKDOWN_WITHEMAILGETRESPONSE_STATUS
	case "Soft Bounce":
		result = SOFTBOUNCE_WITHEMAILGETRESPONSE_STATUS
	case "Hard Bounce":
		result = HARDBOUNCE_WITHEMAILGETRESPONSE_STATUS
	case "Opt-Out":
		result = OPTOUT_WITHEMAILGETRESPONSE_STATUS
	case "Opt-Out: Admin":
		result = OPTOUTADMIN_WITHEMAILGETRESPONSE_STATUS
	case "Opt-Out: System":
		result = OPTOUTSYSTEM_WITHEMAILGETRESPONSE_STATUS
	case "List Unsubscribe":
		result = LISTUNSUBSCRIBE_WITHEMAILGETRESPONSE_STATUS
	case "Provided Feedback":
		result = PROVIDEDFEEDBACK_WITHEMAILGETRESPONSE_STATUS
	case "Reported Spam":
		result = REPORTEDSPAM_WITHEMAILGETRESPONSE_STATUS
	case "Invalid Email":
		result = INVALIDEMAIL_WITHEMAILGETRESPONSE_STATUS
	case "Deactivated/Delinquent Mailbox":
		result = DEACTIVATEDDELINQUENTMAILBOX_WITHEMAILGETRESPONSE_STATUS
	default:
		return 0, errors.New("Unknown WithEmailGetResponse_status value: " + v)
	}
	return &result, nil
}
func SerializeWithEmailGetResponse_status(values []WithEmailGetResponse_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithEmailGetResponse_status) isMultiValue() bool {
	return false
}
