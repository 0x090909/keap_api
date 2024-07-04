package campaigns

import (
	"errors"
)

type CampaignsGetResponse_campaigns_sequences_paths_items_type int

const (
	START_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE CampaignsGetResponse_campaigns_sequences_paths_items_type = iota
	WAIT_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	DELAYTIMER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	CONTACTTIMER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	DATETIMER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	APPOINTMENTTIMER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	EMAIL_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	BARDEMAIL_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	UNLAYEREMAIL_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	EMAILCONFIRM_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	VOICE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	FAX_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	LETTER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	AUTOMATEDSMS_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	TAG_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	OPPORTUNITY_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	NOTE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	TASK_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	COMPLETETASKS_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	APPOINTMENT_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	ASSIGNOWNER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	FIELDVALUE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	FULFILLMENT_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	CREATEORDER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	DEAL_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	ADDTOSEQUENCE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	CANCELSUBSCRIPTION_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	MOVEOPPORTUNITY_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	HTTP_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	CUSTOMERHUB_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	HTTPREQUEST_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	ACTIONSET_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
)

func (i CampaignsGetResponse_campaigns_sequences_paths_items_type) String() string {
	return []string{"Start", "Wait", "DelayTimer", "ContactTimer", "DateTimer", "AppointmentTimer", "Email", "BardEmail", "UnlayerEmail", "EmailConfirm", "Voice", "Fax", "Letter", "AutomatedSms", "Tag", "Opportunity", "Note", "Task", "CompleteTasks", "Appointment", "AssignOwner", "FieldValue", "Fulfillment", "CreateOrder", "Deal", "AddToSequence", "CancelSubscription", "MoveOpportunity", "Http", "CustomerHub", "HttpRequest", "ActionSet"}[i]
}
func ParseCampaignsGetResponse_campaigns_sequences_paths_items_type(v string) (any, error) {
	result := START_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	switch v {
	case "Start":
		result = START_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Wait":
		result = WAIT_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "DelayTimer":
		result = DELAYTIMER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "ContactTimer":
		result = CONTACTTIMER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "DateTimer":
		result = DATETIMER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "AppointmentTimer":
		result = APPOINTMENTTIMER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Email":
		result = EMAIL_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "BardEmail":
		result = BARDEMAIL_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "UnlayerEmail":
		result = UNLAYEREMAIL_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "EmailConfirm":
		result = EMAILCONFIRM_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Voice":
		result = VOICE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Fax":
		result = FAX_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Letter":
		result = LETTER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "AutomatedSms":
		result = AUTOMATEDSMS_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Tag":
		result = TAG_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Opportunity":
		result = OPPORTUNITY_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Note":
		result = NOTE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Task":
		result = TASK_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "CompleteTasks":
		result = COMPLETETASKS_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Appointment":
		result = APPOINTMENT_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "AssignOwner":
		result = ASSIGNOWNER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "FieldValue":
		result = FIELDVALUE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Fulfillment":
		result = FULFILLMENT_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "CreateOrder":
		result = CREATEORDER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Deal":
		result = DEAL_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "AddToSequence":
		result = ADDTOSEQUENCE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "CancelSubscription":
		result = CANCELSUBSCRIPTION_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "MoveOpportunity":
		result = MOVEOPPORTUNITY_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "Http":
		result = HTTP_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "CustomerHub":
		result = CUSTOMERHUB_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "HttpRequest":
		result = HTTPREQUEST_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	case "ActionSet":
		result = ACTIONSET_CAMPAIGNSGETRESPONSE_CAMPAIGNS_SEQUENCES_PATHS_ITEMS_TYPE
	default:
		return 0, errors.New("Unknown CampaignsGetResponse_campaigns_sequences_paths_items_type value: " + v)
	}
	return &result, nil
}
func SerializeCampaignsGetResponse_campaigns_sequences_paths_items_type(values []CampaignsGetResponse_campaigns_sequences_paths_items_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CampaignsGetResponse_campaigns_sequences_paths_items_type) isMultiValue() bool {
	return false
}
