package campaigns

import (
	"errors"
)

type CampaignsGetResponse_campaigns_goals_type int

const (
	WEBFORM_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE CampaignsGetResponse_campaigns_goals_type = iota
	LANDINGPAGE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	SURVEY_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	LINKCLICK_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	EMAILOPENED_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	SCORE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	PURCHASE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	FAILEDPURCHASE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	QUOTE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	WEBSITETRIGGER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	INTERNALFORM_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	SMARTFORM_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	APPOINTMENTEVENT_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	TASK_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	STAGEMOVE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	NOTE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	TAG_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	PIPELINESTAGEMOVE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	UNLAYERLANDINGPAGE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	GROSOCIAL_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	TWITTERFORM_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	API_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	WORDPRESS_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	NEWLANDINGPAGE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	INTEGRATIONTRIGGER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
)

func (i CampaignsGetResponse_campaigns_goals_type) String() string {
	return []string{"WebForm", "LandingPage", "Survey", "LinkClick", "EmailOpened", "Score", "Purchase", "FailedPurchase", "Quote", "WebsiteTrigger", "InternalForm", "SmartForm", "AppointmentEvent", "Task", "StageMove", "Note", "Tag", "PipelineStageMove", "UnlayerLandingPage", "GroSocial", "TwitterForm", "API", "WordPress", "NewLandingPage", "IntegrationTrigger"}[i]
}
func ParseCampaignsGetResponse_campaigns_goals_type(v string) (any, error) {
	result := WEBFORM_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	switch v {
	case "WebForm":
		result = WEBFORM_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "LandingPage":
		result = LANDINGPAGE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "Survey":
		result = SURVEY_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "LinkClick":
		result = LINKCLICK_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "EmailOpened":
		result = EMAILOPENED_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "Score":
		result = SCORE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "Purchase":
		result = PURCHASE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "FailedPurchase":
		result = FAILEDPURCHASE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "Quote":
		result = QUOTE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "WebsiteTrigger":
		result = WEBSITETRIGGER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "InternalForm":
		result = INTERNALFORM_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "SmartForm":
		result = SMARTFORM_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "AppointmentEvent":
		result = APPOINTMENTEVENT_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "Task":
		result = TASK_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "StageMove":
		result = STAGEMOVE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "Note":
		result = NOTE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "Tag":
		result = TAG_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "PipelineStageMove":
		result = PIPELINESTAGEMOVE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "UnlayerLandingPage":
		result = UNLAYERLANDINGPAGE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "GroSocial":
		result = GROSOCIAL_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "TwitterForm":
		result = TWITTERFORM_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "API":
		result = API_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "WordPress":
		result = WORDPRESS_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "NewLandingPage":
		result = NEWLANDINGPAGE_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	case "IntegrationTrigger":
		result = INTEGRATIONTRIGGER_CAMPAIGNSGETRESPONSE_CAMPAIGNS_GOALS_TYPE
	default:
		return 0, errors.New("Unknown CampaignsGetResponse_campaigns_goals_type value: " + v)
	}
	return &result, nil
}
func SerializeCampaignsGetResponse_campaigns_goals_type(values []CampaignsGetResponse_campaigns_goals_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CampaignsGetResponse_campaigns_goals_type) isMultiValue() bool {
	return false
}