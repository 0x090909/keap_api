package files

import (
	"errors"
)

type FilesGetResponse_files_category int

const (
	ATTACHMENTS_FILESGETRESPONSE_FILES_CATEGORY FilesGetResponse_files_category = iota
	CART_FILESGETRESPONSE_FILES_CATEGORY
	DOCUMENTS_FILESGETRESPONSE_FILES_CATEGORY
	HIDDEN_FILESGETRESPONSE_FILES_CATEGORY
	INVOICE_FILESGETRESPONSE_FILES_CATEGORY
	LOGO_FILESGETRESPONSE_FILES_CATEGORY
	TICKETS_FILESGETRESPONSE_FILES_CATEGORY
	WEBFORM_FILESGETRESPONSE_FILES_CATEGORY
	FUNNEL_FILESGETRESPONSE_FILES_CATEGORY
)

func (i FilesGetResponse_files_category) String() string {
	return []string{"Attachments", "Cart", "Documents", "Hidden", "Invoice", "Logo", "Tickets", "WebForm", "Funnel"}[i]
}
func ParseFilesGetResponse_files_category(v string) (any, error) {
	result := ATTACHMENTS_FILESGETRESPONSE_FILES_CATEGORY
	switch v {
	case "Attachments":
		result = ATTACHMENTS_FILESGETRESPONSE_FILES_CATEGORY
	case "Cart":
		result = CART_FILESGETRESPONSE_FILES_CATEGORY
	case "Documents":
		result = DOCUMENTS_FILESGETRESPONSE_FILES_CATEGORY
	case "Hidden":
		result = HIDDEN_FILESGETRESPONSE_FILES_CATEGORY
	case "Invoice":
		result = INVOICE_FILESGETRESPONSE_FILES_CATEGORY
	case "Logo":
		result = LOGO_FILESGETRESPONSE_FILES_CATEGORY
	case "Tickets":
		result = TICKETS_FILESGETRESPONSE_FILES_CATEGORY
	case "WebForm":
		result = WEBFORM_FILESGETRESPONSE_FILES_CATEGORY
	case "Funnel":
		result = FUNNEL_FILESGETRESPONSE_FILES_CATEGORY
	default:
		return 0, errors.New("Unknown FilesGetResponse_files_category value: " + v)
	}
	return &result, nil
}
func SerializeFilesGetResponse_files_category(values []FilesGetResponse_files_category) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i FilesGetResponse_files_category) isMultiValue() bool {
	return false
}
