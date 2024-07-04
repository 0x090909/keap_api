package tags

import (
	"errors"
)

type Tags_key int

const (
	SUCCESS_TAGS_KEY Tags_key = iota
	DUPLICATE_TAGS_KEY
	CONTACT_NOT_FOUND_TAGS_KEY
	TAG_ID_NOT_FOUND_TAGS_KEY
	FAILURE_TAGS_KEY
	NO_PERMISSION_TAGS_KEY
)

func (i Tags_key) String() string {
	return []string{"SUCCESS", "DUPLICATE", "CONTACT_NOT_FOUND", "TAG_ID_NOT_FOUND", "FAILURE", "NO_PERMISSION"}[i]
}
func ParseTags_key(v string) (any, error) {
	result := SUCCESS_TAGS_KEY
	switch v {
	case "SUCCESS":
		result = SUCCESS_TAGS_KEY
	case "DUPLICATE":
		result = DUPLICATE_TAGS_KEY
	case "CONTACT_NOT_FOUND":
		result = CONTACT_NOT_FOUND_TAGS_KEY
	case "TAG_ID_NOT_FOUND":
		result = TAG_ID_NOT_FOUND_TAGS_KEY
	case "FAILURE":
		result = FAILURE_TAGS_KEY
	case "NO_PERMISSION":
		result = NO_PERMISSION_TAGS_KEY
	default:
		return 0, errors.New("Unknown Tags_key value: " + v)
	}
	return &result, nil
}
func SerializeTags_key(values []Tags_key) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i Tags_key) isMultiValue() bool {
	return false
}
