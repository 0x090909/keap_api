package model

import (
	"errors"
)

type ModelGetResponse_custom_fields_record_type int

const (
	CONTACT_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE ModelGetResponse_custom_fields_record_type = iota
	REFERRAL_PARTNER_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	OPPORTUNITY_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	TASK_NOTE_APPOINTMENT_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	COMPANY_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	ORDER_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	SUBSCRIPTION_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
)

func (i ModelGetResponse_custom_fields_record_type) String() string {
	return []string{"CONTACT", "REFERRAL_PARTNER", "OPPORTUNITY", "TASK_NOTE_APPOINTMENT", "COMPANY", "ORDER", "SUBSCRIPTION"}[i]
}
func ParseModelGetResponse_custom_fields_record_type(v string) (any, error) {
	result := CONTACT_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	switch v {
	case "CONTACT":
		result = CONTACT_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	case "REFERRAL_PARTNER":
		result = REFERRAL_PARTNER_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	case "OPPORTUNITY":
		result = OPPORTUNITY_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	case "TASK_NOTE_APPOINTMENT":
		result = TASK_NOTE_APPOINTMENT_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	case "COMPANY":
		result = COMPANY_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	case "ORDER":
		result = ORDER_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	case "SUBSCRIPTION":
		result = SUBSCRIPTION_MODELGETRESPONSE_CUSTOM_FIELDS_RECORD_TYPE
	default:
		return 0, errors.New("Unknown ModelGetResponse_custom_fields_record_type value: " + v)
	}
	return &result, nil
}
func SerializeModelGetResponse_custom_fields_record_type(values []ModelGetResponse_custom_fields_record_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ModelGetResponse_custom_fields_record_type) isMultiValue() bool {
	return false
}
