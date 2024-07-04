package item

import (
	"errors"
)

type WithFilePutResponse_file_descriptor_category int

const (
	ATTACHMENTS_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY WithFilePutResponse_file_descriptor_category = iota
	CART_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	DOCUMENTS_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	HIDDEN_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	INVOICE_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	LOGO_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	TICKETS_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	WEBFORM_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	FUNNEL_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
)

func (i WithFilePutResponse_file_descriptor_category) String() string {
	return []string{"Attachments", "Cart", "Documents", "Hidden", "Invoice", "Logo", "Tickets", "WebForm", "Funnel"}[i]
}
func ParseWithFilePutResponse_file_descriptor_category(v string) (any, error) {
	result := ATTACHMENTS_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	switch v {
	case "Attachments":
		result = ATTACHMENTS_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Cart":
		result = CART_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Documents":
		result = DOCUMENTS_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Hidden":
		result = HIDDEN_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Invoice":
		result = INVOICE_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Logo":
		result = LOGO_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Tickets":
		result = TICKETS_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "WebForm":
		result = WEBFORM_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Funnel":
		result = FUNNEL_WITHFILEPUTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	default:
		return 0, errors.New("Unknown WithFilePutResponse_file_descriptor_category value: " + v)
	}
	return &result, nil
}
func SerializeWithFilePutResponse_file_descriptor_category(values []WithFilePutResponse_file_descriptor_category) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithFilePutResponse_file_descriptor_category) isMultiValue() bool {
	return false
}
