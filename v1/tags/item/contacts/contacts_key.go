package contacts

import (
	"errors"
)

type Contacts_key int

const (
	SUCCESS_CONTACTS_KEY Contacts_key = iota
	DUPLICATE_CONTACTS_KEY
	CONTACT_NOT_FOUND_CONTACTS_KEY
	TAG_ID_NOT_FOUND_CONTACTS_KEY
	FAILURE_CONTACTS_KEY
	NO_PERMISSION_CONTACTS_KEY
)

func (i Contacts_key) String() string {
	return []string{"SUCCESS", "DUPLICATE", "CONTACT_NOT_FOUND", "TAG_ID_NOT_FOUND", "FAILURE", "NO_PERMISSION"}[i]
}
func ParseContacts_key(v string) (any, error) {
	result := SUCCESS_CONTACTS_KEY
	switch v {
	case "SUCCESS":
		result = SUCCESS_CONTACTS_KEY
	case "DUPLICATE":
		result = DUPLICATE_CONTACTS_KEY
	case "CONTACT_NOT_FOUND":
		result = CONTACT_NOT_FOUND_CONTACTS_KEY
	case "TAG_ID_NOT_FOUND":
		result = TAG_ID_NOT_FOUND_CONTACTS_KEY
	case "FAILURE":
		result = FAILURE_CONTACTS_KEY
	case "NO_PERMISSION":
		result = NO_PERMISSION_CONTACTS_KEY
	default:
		return 0, errors.New("Unknown Contacts_key value: " + v)
	}
	return &result, nil
}
func SerializeContacts_key(values []Contacts_key) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i Contacts_key) isMultiValue() bool {
	return false
}
