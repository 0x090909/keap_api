package item

import (
	"errors"
)

type WithEmailPutResponse_status int

const (
	UNENGAGEDMARKETABLE_WITHEMAILPUTRESPONSE_STATUS WithEmailPutResponse_status = iota
	UNCONFIRMED_WITHEMAILPUTRESPONSE_STATUS
	CONFIRMEDLEGACY_WITHEMAILPUTRESPONSE_STATUS
	CONFIRMED_WITHEMAILPUTRESPONSE_STATUS
	UNENGAGEDNONMARKETABLE_WITHEMAILPUTRESPONSE_STATUS
	NONMARKETABLE_WITHEMAILPUTRESPONSE_STATUS
	LOCKDOWN_WITHEMAILPUTRESPONSE_STATUS
	SOFTBOUNCE_WITHEMAILPUTRESPONSE_STATUS
	HARDBOUNCE_WITHEMAILPUTRESPONSE_STATUS
	OPTOUT_WITHEMAILPUTRESPONSE_STATUS
	OPTOUTADMIN_WITHEMAILPUTRESPONSE_STATUS
	OPTOUTSYSTEM_WITHEMAILPUTRESPONSE_STATUS
	LISTUNSUBSCRIBE_WITHEMAILPUTRESPONSE_STATUS
	PROVIDEDFEEDBACK_WITHEMAILPUTRESPONSE_STATUS
	REPORTEDSPAM_WITHEMAILPUTRESPONSE_STATUS
	INVALIDEMAIL_WITHEMAILPUTRESPONSE_STATUS
	DEACTIVATEDDELINQUENTMAILBOX_WITHEMAILPUTRESPONSE_STATUS
)

func (i WithEmailPutResponse_status) String() string {
	return []string{"Unengaged Marketable", "Unconfirmed", "Confirmed (Legacy)", "Confirmed", "Unengaged NonMarketable", "Non-marketable", "Lockdown", "Soft Bounce", "Hard Bounce", "Opt-Out", "Opt-Out: Admin", "Opt-Out: System", "List Unsubscribe", "Provided Feedback", "Reported Spam", "Invalid Email", "Deactivated/Delinquent Mailbox"}[i]
}
func ParseWithEmailPutResponse_status(v string) (any, error) {
	result := UNENGAGEDMARKETABLE_WITHEMAILPUTRESPONSE_STATUS
	switch v {
	case "Unengaged Marketable":
		result = UNENGAGEDMARKETABLE_WITHEMAILPUTRESPONSE_STATUS
	case "Unconfirmed":
		result = UNCONFIRMED_WITHEMAILPUTRESPONSE_STATUS
	case "Confirmed (Legacy)":
		result = CONFIRMEDLEGACY_WITHEMAILPUTRESPONSE_STATUS
	case "Confirmed":
		result = CONFIRMED_WITHEMAILPUTRESPONSE_STATUS
	case "Unengaged NonMarketable":
		result = UNENGAGEDNONMARKETABLE_WITHEMAILPUTRESPONSE_STATUS
	case "Non-marketable":
		result = NONMARKETABLE_WITHEMAILPUTRESPONSE_STATUS
	case "Lockdown":
		result = LOCKDOWN_WITHEMAILPUTRESPONSE_STATUS
	case "Soft Bounce":
		result = SOFTBOUNCE_WITHEMAILPUTRESPONSE_STATUS
	case "Hard Bounce":
		result = HARDBOUNCE_WITHEMAILPUTRESPONSE_STATUS
	case "Opt-Out":
		result = OPTOUT_WITHEMAILPUTRESPONSE_STATUS
	case "Opt-Out: Admin":
		result = OPTOUTADMIN_WITHEMAILPUTRESPONSE_STATUS
	case "Opt-Out: System":
		result = OPTOUTSYSTEM_WITHEMAILPUTRESPONSE_STATUS
	case "List Unsubscribe":
		result = LISTUNSUBSCRIBE_WITHEMAILPUTRESPONSE_STATUS
	case "Provided Feedback":
		result = PROVIDEDFEEDBACK_WITHEMAILPUTRESPONSE_STATUS
	case "Reported Spam":
		result = REPORTEDSPAM_WITHEMAILPUTRESPONSE_STATUS
	case "Invalid Email":
		result = INVALIDEMAIL_WITHEMAILPUTRESPONSE_STATUS
	case "Deactivated/Delinquent Mailbox":
		result = DEACTIVATEDDELINQUENTMAILBOX_WITHEMAILPUTRESPONSE_STATUS
	default:
		return 0, errors.New("Unknown WithEmailPutResponse_status value: " + v)
	}
	return &result, nil
}
func SerializeWithEmailPutResponse_status(values []WithEmailPutResponse_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithEmailPutResponse_status) isMultiValue() bool {
	return false
}
