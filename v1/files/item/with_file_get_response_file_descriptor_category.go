package item

import (
	"errors"
)

type WithFileGetResponse_file_descriptor_category int

const (
	ATTACHMENTS_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY WithFileGetResponse_file_descriptor_category = iota
	CART_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	DOCUMENTS_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	HIDDEN_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	INVOICE_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	LOGO_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	TICKETS_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	WEBFORM_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	FUNNEL_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
)

func (i WithFileGetResponse_file_descriptor_category) String() string {
	return []string{"Attachments", "Cart", "Documents", "Hidden", "Invoice", "Logo", "Tickets", "WebForm", "Funnel"}[i]
}
func ParseWithFileGetResponse_file_descriptor_category(v string) (any, error) {
	result := ATTACHMENTS_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	switch v {
	case "Attachments":
		result = ATTACHMENTS_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Cart":
		result = CART_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Documents":
		result = DOCUMENTS_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Hidden":
		result = HIDDEN_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Invoice":
		result = INVOICE_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Logo":
		result = LOGO_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Tickets":
		result = TICKETS_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "WebForm":
		result = WEBFORM_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Funnel":
		result = FUNNEL_WITHFILEGETRESPONSE_FILE_DESCRIPTOR_CATEGORY
	default:
		return 0, errors.New("Unknown WithFileGetResponse_file_descriptor_category value: " + v)
	}
	return &result, nil
}
func SerializeWithFileGetResponse_file_descriptor_category(values []WithFileGetResponse_file_descriptor_category) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithFileGetResponse_file_descriptor_category) isMultiValue() bool {
	return false
}
