package customfields

import (
	"errors"
)

type CustomFieldsPostRequestBody_field_type int

const (
	CURRENCY_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE CustomFieldsPostRequestBody_field_type = iota
	DATE_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	DATETIME_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	DAYOFWEEK_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	DRILLDOWN_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	EMAIL_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	MONTH_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	LISTBOX_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	NAME_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	WHOLENUMBER_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	DECIMALNUMBER_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	PERCENT_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	PHONENUMBER_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	RADIO_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	DROPDOWN_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	SOCIALSECURITYNUMBER_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	STATE_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	TEXT_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	TEXTAREA_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	USER_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	USERLISTBOX_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	WEBSITE_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	YEAR_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	YESNO_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
)

func (i CustomFieldsPostRequestBody_field_type) String() string {
	return []string{"Currency", "Date", "DateTime", "DayOfWeek", "Drilldown", "Email", "Month", "ListBox", "Name", "WholeNumber", "DecimalNumber", "Percent", "PhoneNumber", "Radio", "Dropdown", "SocialSecurityNumber", "State", "Text", "TextArea", "User", "UserListBox", "Website", "Year", "YesNo"}[i]
}
func ParseCustomFieldsPostRequestBody_field_type(v string) (any, error) {
	result := CURRENCY_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	switch v {
	case "Currency":
		result = CURRENCY_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "Date":
		result = DATE_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "DateTime":
		result = DATETIME_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "DayOfWeek":
		result = DAYOFWEEK_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "Drilldown":
		result = DRILLDOWN_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "Email":
		result = EMAIL_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "Month":
		result = MONTH_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "ListBox":
		result = LISTBOX_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "Name":
		result = NAME_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "WholeNumber":
		result = WHOLENUMBER_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "DecimalNumber":
		result = DECIMALNUMBER_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "Percent":
		result = PERCENT_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "PhoneNumber":
		result = PHONENUMBER_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "Radio":
		result = RADIO_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "Dropdown":
		result = DROPDOWN_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "SocialSecurityNumber":
		result = SOCIALSECURITYNUMBER_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "State":
		result = STATE_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "Text":
		result = TEXT_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "TextArea":
		result = TEXTAREA_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "User":
		result = USER_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "UserListBox":
		result = USERLISTBOX_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "Website":
		result = WEBSITE_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "Year":
		result = YEAR_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	case "YesNo":
		result = YESNO_CUSTOMFIELDSPOSTREQUESTBODY_FIELD_TYPE
	default:
		return 0, errors.New("Unknown CustomFieldsPostRequestBody_field_type value: " + v)
	}
	return &result, nil
}
func SerializeCustomFieldsPostRequestBody_field_type(values []CustomFieldsPostRequestBody_field_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i CustomFieldsPostRequestBody_field_type) isMultiValue() bool {
	return false
}