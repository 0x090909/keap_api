package customfields

import (
	"errors"
)

type CustomFieldsPostResponse_record_type int

const (
	CONTACT_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE CustomFieldsPostResponse_record_type = iota
	REFERRAL_PARTNER_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	OPPORTUNITY_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	TASK_NOTE_APPOINTMENT_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	COMPANY_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	ORDER_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	SUBSCRIPTION_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
)

func (i CustomFieldsPostResponse_record_type) String() string {
	return []string{"CONTACT", "REFERRAL_PARTNER", "OPPORTUNITY", "TASK_NOTE_APPOINTMENT", "COMPANY", "ORDER", "SUBSCRIPTION"}[i]
}
func ParseCustomFieldsPostResponse_record_type(v string) (any, error) {
	result := CONTACT_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	switch v {
	case "CONTACT":
		result = CONTACT_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	case "REFERRAL_PARTNER":
		result = REFERRAL_PARTNER_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	case "OPPORTUNITY":
		result = OPPORTUNITY_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	case "TASK_NOTE_APPOINTMENT":
		result = TASK_NOTE_APPOINTMENT_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	case "COMPANY":
		result = COMPANY_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	case "ORDER":
		result = ORDER_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	case "SUBSCRIPTION":
		result = SUBSCRIPTION_CUSTOMFIELDSPOSTRESPONSE_RECORD_TYPE
	default:
		return 0, errors.New("Unknown CustomFieldsPostResponse_record_type value: " + v)
	}
	return &result, nil
}
func SerializeCustomFieldsPostResponse_record_type(values []CustomFieldsPostResponse_record_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CustomFieldsPostResponse_record_type) isMultiValue() bool {
	return false
}
