package notes

import (
	"errors"
)

type NotesPostResponse_type int

const (
	APPOINTMENT_NOTESPOSTRESPONSE_TYPE NotesPostResponse_type = iota
	CALL_NOTESPOSTRESPONSE_TYPE
	EMAIL_NOTESPOSTRESPONSE_TYPE
	FAX_NOTESPOSTRESPONSE_TYPE
	LETTER_NOTESPOSTRESPONSE_TYPE
	OTHER_NOTESPOSTRESPONSE_TYPE
)

func (i NotesPostResponse_type) String() string {
	return []string{"Appointment", "Call", "Email", "Fax", "Letter", "Other"}[i]
}
func ParseNotesPostResponse_type(v string) (any, error) {
	result := APPOINTMENT_NOTESPOSTRESPONSE_TYPE
	switch v {
	case "Appointment":
		result = APPOINTMENT_NOTESPOSTRESPONSE_TYPE
	case "Call":
		result = CALL_NOTESPOSTRESPONSE_TYPE
	case "Email":
		result = EMAIL_NOTESPOSTRESPONSE_TYPE
	case "Fax":
		result = FAX_NOTESPOSTRESPONSE_TYPE
	case "Letter":
		result = LETTER_NOTESPOSTRESPONSE_TYPE
	case "Other":
		result = OTHER_NOTESPOSTRESPONSE_TYPE
	default:
		return 0, errors.New("Unknown NotesPostResponse_type value: " + v)
	}
	return &result, nil
}
func SerializeNotesPostResponse_type(values []NotesPostResponse_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i NotesPostResponse_type) isMultiValue() bool {
	return false
}
