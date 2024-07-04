package item

import (
	"errors"
)

type WithCampaignGetResponse_goals_type int

const (
	WEBFORM_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE WithCampaignGetResponse_goals_type = iota
	LANDINGPAGE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	SURVEY_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	LINKCLICK_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	EMAILOPENED_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	SCORE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	PURCHASE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	FAILEDPURCHASE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	QUOTE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	WEBSITETRIGGER_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	INTERNALFORM_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	SMARTFORM_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	APPOINTMENTEVENT_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	TASK_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	STAGEMOVE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	NOTE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	TAG_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	PIPELINESTAGEMOVE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	UNLAYERLANDINGPAGE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	GROSOCIAL_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	TWITTERFORM_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	API_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	WORDPRESS_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	NEWLANDINGPAGE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	INTEGRATIONTRIGGER_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
)

func (i WithCampaignGetResponse_goals_type) String() string {
	return []string{"WebForm", "LandingPage", "Survey", "LinkClick", "EmailOpened", "Score", "Purchase", "FailedPurchase", "Quote", "WebsiteTrigger", "InternalForm", "SmartForm", "AppointmentEvent", "Task", "StageMove", "Note", "Tag", "PipelineStageMove", "UnlayerLandingPage", "GroSocial", "TwitterForm", "API", "WordPress", "NewLandingPage", "IntegrationTrigger"}[i]
}
func ParseWithCampaignGetResponse_goals_type(v string) (any, error) {
	result := WEBFORM_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	switch v {
	case "WebForm":
		result = WEBFORM_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "LandingPage":
		result = LANDINGPAGE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "Survey":
		result = SURVEY_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "LinkClick":
		result = LINKCLICK_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "EmailOpened":
		result = EMAILOPENED_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "Score":
		result = SCORE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "Purchase":
		result = PURCHASE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "FailedPurchase":
		result = FAILEDPURCHASE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "Quote":
		result = QUOTE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "WebsiteTrigger":
		result = WEBSITETRIGGER_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "InternalForm":
		result = INTERNALFORM_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "SmartForm":
		result = SMARTFORM_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "AppointmentEvent":
		result = APPOINTMENTEVENT_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "Task":
		result = TASK_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "StageMove":
		result = STAGEMOVE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "Note":
		result = NOTE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "Tag":
		result = TAG_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "PipelineStageMove":
		result = PIPELINESTAGEMOVE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "UnlayerLandingPage":
		result = UNLAYERLANDINGPAGE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "GroSocial":
		result = GROSOCIAL_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "TwitterForm":
		result = TWITTERFORM_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "API":
		result = API_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "WordPress":
		result = WORDPRESS_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "NewLandingPage":
		result = NEWLANDINGPAGE_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	case "IntegrationTrigger":
		result = INTEGRATIONTRIGGER_WITHCAMPAIGNGETRESPONSE_GOALS_TYPE
	default:
		return 0, errors.New("Unknown WithCampaignGetResponse_goals_type value: " + v)
	}
	return &result, nil
}
func SerializeWithCampaignGetResponse_goals_type(values []WithCampaignGetResponse_goals_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithCampaignGetResponse_goals_type) isMultiValue() bool {
	return false
}
