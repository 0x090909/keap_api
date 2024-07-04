package files

import (
	"errors"
)

type FilesPostResponse_file_descriptor_category int

const (
	ATTACHMENTS_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY FilesPostResponse_file_descriptor_category = iota
	CART_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	DOCUMENTS_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	HIDDEN_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	INVOICE_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	LOGO_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	TICKETS_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	WEBFORM_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	FUNNEL_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
)

func (i FilesPostResponse_file_descriptor_category) String() string {
	return []string{"Attachments", "Cart", "Documents", "Hidden", "Invoice", "Logo", "Tickets", "WebForm", "Funnel"}[i]
}
func ParseFilesPostResponse_file_descriptor_category(v string) (any, error) {
	result := ATTACHMENTS_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	switch v {
	case "Attachments":
		result = ATTACHMENTS_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Cart":
		result = CART_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Documents":
		result = DOCUMENTS_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Hidden":
		result = HIDDEN_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Invoice":
		result = INVOICE_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Logo":
		result = LOGO_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Tickets":
		result = TICKETS_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "WebForm":
		result = WEBFORM_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	case "Funnel":
		result = FUNNEL_FILESPOSTRESPONSE_FILE_DESCRIPTOR_CATEGORY
	default:
		return 0, errors.New("Unknown FilesPostResponse_file_descriptor_category value: " + v)
	}
	return &result, nil
}
func SerializeFilesPostResponse_file_descriptor_category(values []FilesPostResponse_file_descriptor_category) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i FilesPostResponse_file_descriptor_category) isMultiValue() bool {
	return false
}
