package applicationsgetconfiguration

import (
	"errors"
)

type GetFieldsQueryParameterType int

const (
	AFFILIATE_GETFIELDSQUERYPARAMETERTYPE GetFieldsQueryParameterType = iota
	APPOINTMENT_GETFIELDSQUERYPARAMETERTYPE
	CONTACT_GETFIELDSQUERYPARAMETERTYPE
	ECOMMERCE_GETFIELDSQUERYPARAMETERTYPE
	EMAIL_GETFIELDSQUERYPARAMETERTYPE
	FORMS_GETFIELDSQUERYPARAMETERTYPE
	FULFILLMENT_GETFIELDSQUERYPARAMETERTYPE
	INVOICE_GETFIELDSQUERYPARAMETERTYPE
	NOTE_GETFIELDSQUERYPARAMETERTYPE
	OPPORTUNITY_GETFIELDSQUERYPARAMETERTYPE
	TASK_GETFIELDSQUERYPARAMETERTYPE
	TEMPLATE_GETFIELDSQUERYPARAMETERTYPE
)

func (i GetFieldsQueryParameterType) String() string {
	return []string{"affiliate", "appointment", "contact", "ecommerce", "email", "forms", "fulfillment", "invoice", "note", "opportunity", "task", "template"}[i]
}
func ParseGetFieldsQueryParameterType(v string) (any, error) {
	result := AFFILIATE_GETFIELDSQUERYPARAMETERTYPE
	switch v {
	case "affiliate":
		result = AFFILIATE_GETFIELDSQUERYPARAMETERTYPE
	case "appointment":
		result = APPOINTMENT_GETFIELDSQUERYPARAMETERTYPE
	case "contact":
		result = CONTACT_GETFIELDSQUERYPARAMETERTYPE
	case "ecommerce":
		result = ECOMMERCE_GETFIELDSQUERYPARAMETERTYPE
	case "email":
		result = EMAIL_GETFIELDSQUERYPARAMETERTYPE
	case "forms":
		result = FORMS_GETFIELDSQUERYPARAMETERTYPE
	case "fulfillment":
		result = FULFILLMENT_GETFIELDSQUERYPARAMETERTYPE
	case "invoice":
		result = INVOICE_GETFIELDSQUERYPARAMETERTYPE
	case "note":
		result = NOTE_GETFIELDSQUERYPARAMETERTYPE
	case "opportunity":
		result = OPPORTUNITY_GETFIELDSQUERYPARAMETERTYPE
	case "task":
		result = TASK_GETFIELDSQUERYPARAMETERTYPE
	case "template":
		result = TEMPLATE_GETFIELDSQUERYPARAMETERTYPE
	default:
		return 0, errors.New("Unknown GetFieldsQueryParameterType value: " + v)
	}
	return &result, nil
}
func SerializeGetFieldsQueryParameterType(values []GetFieldsQueryParameterType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i GetFieldsQueryParameterType) isMultiValue() bool {
	return false
}
