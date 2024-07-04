package notes

import (
	"errors"
)

type NotesGetResponse_notes_type int

const (
	APPOINTMENT_NOTESGETRESPONSE_NOTES_TYPE NotesGetResponse_notes_type = iota
	CALL_NOTESGETRESPONSE_NOTES_TYPE
	EMAIL_NOTESGETRESPONSE_NOTES_TYPE
	FAX_NOTESGETRESPONSE_NOTES_TYPE
	LETTER_NOTESGETRESPONSE_NOTES_TYPE
	OTHER_NOTESGETRESPONSE_NOTES_TYPE
)

func (i NotesGetResponse_notes_type) String() string {
	return []string{"Appointment", "Call", "Email", "Fax", "Letter", "Other"}[i]
}
func ParseNotesGetResponse_notes_type(v string) (any, error) {
	result := APPOINTMENT_NOTESGETRESPONSE_NOTES_TYPE
	switch v {
	case "Appointment":
		result = APPOINTMENT_NOTESGETRESPONSE_NOTES_TYPE
	case "Call":
		result = CALL_NOTESGETRESPONSE_NOTES_TYPE
	case "Email":
		result = EMAIL_NOTESGETRESPONSE_NOTES_TYPE
	case "Fax":
		result = FAX_NOTESGETRESPONSE_NOTES_TYPE
	case "Letter":
		result = LETTER_NOTESGETRESPONSE_NOTES_TYPE
	case "Other":
		result = OTHER_NOTESGETRESPONSE_NOTES_TYPE
	default:
		return 0, errors.New("Unknown NotesGetResponse_notes_type value: " + v)
	}
	return &result, nil
}
func SerializeNotesGetResponse_notes_type(values []NotesGetResponse_notes_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i NotesGetResponse_notes_type) isMultiValue() bool {
	return false
}
