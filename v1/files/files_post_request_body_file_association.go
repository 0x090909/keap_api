package files

import (
	"errors"
)

type FilesPostRequestBody_file_association int

const (
	CONTACT_FILESPOSTREQUESTBODY_FILE_ASSOCIATION FilesPostRequestBody_file_association = iota
	USER_FILESPOSTREQUESTBODY_FILE_ASSOCIATION
	COMPANY_FILESPOSTREQUESTBODY_FILE_ASSOCIATION
)

func (i FilesPostRequestBody_file_association) String() string {
	return []string{"CONTACT", "USER", "COMPANY"}[i]
}
func ParseFilesPostRequestBody_file_association(v string) (any, error) {
	result := CONTACT_FILESPOSTREQUESTBODY_FILE_ASSOCIATION
	switch v {
	case "CONTACT":
		result = CONTACT_FILESPOSTREQUESTBODY_FILE_ASSOCIATION
	case "USER":
		result = USER_FILESPOSTREQUESTBODY_FILE_ASSOCIATION
	case "COMPANY":
		result = COMPANY_FILESPOSTREQUESTBODY_FILE_ASSOCIATION
	default:
		return 0, errors.New("Unknown FilesPostRequestBody_file_association value: " + v)
	}
	return &result, nil
}
func SerializeFilesPostRequestBody_file_association(values []FilesPostRequestBody_file_association) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i FilesPostRequestBody_file_association) isMultiValue() bool {
	return false
}
