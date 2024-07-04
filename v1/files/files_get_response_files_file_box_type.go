package files

import (
	"errors"
)

type FilesGetResponse_files_file_box_type int

const (
	APPLICATION_FILESGETRESPONSE_FILES_FILE_BOX_TYPE FilesGetResponse_files_file_box_type = iota
	IMAGE_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	FAX_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	ATTACHMENT_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	TICKET_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	CONTACT_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	DIGITALPRODUCT_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	IMPORT_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	HIDDEN_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	WEBFORM_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	STYLEDCART_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	RESAMPLEDIMAGE_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	TEMPLATETHUMBNAIL_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	FUNNEL_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	LOGOTHUMBNAIL_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	UNLAYER_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
)

func (i FilesGetResponse_files_file_box_type) String() string {
	return []string{"Application", "Image", "Fax", "Attachment", "Ticket", "Contact", "DigitalProduct", "Import", "Hidden", "WebForm", "StyledCart", "ReSampledImage", "TemplateThumbnail", "Funnel", "LogoThumbnail", "Unlayer"}[i]
}
func ParseFilesGetResponse_files_file_box_type(v string) (any, error) {
	result := APPLICATION_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	switch v {
	case "Application":
		result = APPLICATION_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "Image":
		result = IMAGE_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "Fax":
		result = FAX_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "Attachment":
		result = ATTACHMENT_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "Ticket":
		result = TICKET_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "Contact":
		result = CONTACT_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "DigitalProduct":
		result = DIGITALPRODUCT_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "Import":
		result = IMPORT_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "Hidden":
		result = HIDDEN_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "WebForm":
		result = WEBFORM_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "StyledCart":
		result = STYLEDCART_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "ReSampledImage":
		result = RESAMPLEDIMAGE_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "TemplateThumbnail":
		result = TEMPLATETHUMBNAIL_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "Funnel":
		result = FUNNEL_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "LogoThumbnail":
		result = LOGOTHUMBNAIL_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	case "Unlayer":
		result = UNLAYER_FILESGETRESPONSE_FILES_FILE_BOX_TYPE
	default:
		return 0, errors.New("Unknown FilesGetResponse_files_file_box_type value: " + v)
	}
	return &result, nil
}
func SerializeFilesGetResponse_files_file_box_type(values []FilesGetResponse_files_file_box_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i FilesGetResponse_files_file_box_type) isMultiValue() bool {
	return false
}
