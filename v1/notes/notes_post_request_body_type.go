package notes

import (
	"errors"
)

type NotesPostRequestBody_type int

const (
	APPOINTMENT_NOTESPOSTREQUESTBODY_TYPE NotesPostRequestBody_type = iota
	CALL_NOTESPOSTREQUESTBODY_TYPE
	EMAIL_NOTESPOSTREQUESTBODY_TYPE
	FAX_NOTESPOSTREQUESTBODY_TYPE
	LETTER_NOTESPOSTREQUESTBODY_TYPE
	OTHER_NOTESPOSTREQUESTBODY_TYPE
)

func (i NotesPostRequestBody_type) String() string {
	return []string{"Appointment", "Call", "Email", "Fax", "Letter", "Other"}[i]
}
func ParseNotesPostRequestBody_type(v string) (any, error) {
	result := APPOINTMENT_NOTESPOSTREQUESTBODY_TYPE
	switch v {
	case "Appointment":
		result = APPOINTMENT_NOTESPOSTREQUESTBODY_TYPE
	case "Call":
		result = CALL_NOTESPOSTREQUESTBODY_TYPE
	case "Email":
		result = EMAIL_NOTESPOSTREQUESTBODY_TYPE
	case "Fax":
		result = FAX_NOTESPOSTREQUESTBODY_TYPE
	case "Letter":
		result = LETTER_NOTESPOSTREQUESTBODY_TYPE
	case "Other":
		result = OTHER_NOTESPOSTREQUESTBODY_TYPE
	default:
		return 0, errors.New("Unknown NotesPostRequestBody_type value: " + v)
	}
	return &result, nil
}
func SerializeNotesPostRequestBody_type(values []NotesPostRequestBody_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i NotesPostRequestBody_type) isMultiValue() bool {
	return false
}
