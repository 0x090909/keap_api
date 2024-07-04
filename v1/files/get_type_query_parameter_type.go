package files

import (
	"errors"
)

type GetTypeQueryParameterType int

const (
	APPLICATION_GETTYPEQUERYPARAMETERTYPE GetTypeQueryParameterType = iota
	IMAGE_GETTYPEQUERYPARAMETERTYPE
	FAX_GETTYPEQUERYPARAMETERTYPE
	ATTACHMENT_GETTYPEQUERYPARAMETERTYPE
	TICKET_GETTYPEQUERYPARAMETERTYPE
	CONTACT_GETTYPEQUERYPARAMETERTYPE
	DIGITALPRODUCT_GETTYPEQUERYPARAMETERTYPE
	IMPORT_GETTYPEQUERYPARAMETERTYPE
	HIDDEN_GETTYPEQUERYPARAMETERTYPE
	WEBFORM_GETTYPEQUERYPARAMETERTYPE
	STYLEDCART_GETTYPEQUERYPARAMETERTYPE
	RESAMPLEDIMAGE_GETTYPEQUERYPARAMETERTYPE
	TEMPLATETHUMBNAIL_GETTYPEQUERYPARAMETERTYPE
	FUNNEL_GETTYPEQUERYPARAMETERTYPE
	LOGOTHUMBNAIL_GETTYPEQUERYPARAMETERTYPE
	UNLAYER_GETTYPEQUERYPARAMETERTYPE
)

func (i GetTypeQueryParameterType) String() string {
	return []string{"Application", "Image", "Fax", "Attachment", "Ticket", "Contact", "DigitalProduct", "Import", "Hidden", "WebForm", "StyledCart", "ReSampledImage", "TemplateThumbnail", "Funnel", "LogoThumbnail", "Unlayer"}[i]
}
func ParseGetTypeQueryParameterType(v string) (any, error) {
	result := APPLICATION_GETTYPEQUERYPARAMETERTYPE
	switch v {
	case "Application":
		result = APPLICATION_GETTYPEQUERYPARAMETERTYPE
	case "Image":
		result = IMAGE_GETTYPEQUERYPARAMETERTYPE
	case "Fax":
		result = FAX_GETTYPEQUERYPARAMETERTYPE
	case "Attachment":
		result = ATTACHMENT_GETTYPEQUERYPARAMETERTYPE
	case "Ticket":
		result = TICKET_GETTYPEQUERYPARAMETERTYPE
	case "Contact":
		result = CONTACT_GETTYPEQUERYPARAMETERTYPE
	case "DigitalProduct":
		result = DIGITALPRODUCT_GETTYPEQUERYPARAMETERTYPE
	case "Import":
		result = IMPORT_GETTYPEQUERYPARAMETERTYPE
	case "Hidden":
		result = HIDDEN_GETTYPEQUERYPARAMETERTYPE
	case "WebForm":
		result = WEBFORM_GETTYPEQUERYPARAMETERTYPE
	case "StyledCart":
		result = STYLEDCART_GETTYPEQUERYPARAMETERTYPE
	case "ReSampledImage":
		result = RESAMPLEDIMAGE_GETTYPEQUERYPARAMETERTYPE
	case "TemplateThumbnail":
		result = TEMPLATETHUMBNAIL_GETTYPEQUERYPARAMETERTYPE
	case "Funnel":
		result = FUNNEL_GETTYPEQUERYPARAMETERTYPE
	case "LogoThumbnail":
		result = LOGOTHUMBNAIL_GETTYPEQUERYPARAMETERTYPE
	case "Unlayer":
		result = UNLAYER_GETTYPEQUERYPARAMETERTYPE
	default:
		return 0, errors.New("Unknown GetTypeQueryParameterType value: " + v)
	}
	return &result, nil
}
func SerializeGetTypeQueryParameterType(values []GetTypeQueryParameterType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i GetTypeQueryParameterType) isMultiValue() bool {
	return false
}
