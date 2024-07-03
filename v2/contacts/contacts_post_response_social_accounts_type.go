package contacts

import (
	"errors"
)

type ContactsPostResponse_social_accounts_type int

const (
	SOCIAL_ACCOUNT_TYPE_UNSPECIFIED_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE ContactsPostResponse_social_accounts_type = iota
	FACEBOOK_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	LINKED_IN_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	TWITTER_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	INSTAGRAM_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	SNAPCHAT_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	YOUTUBE_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	PINTEREST_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
)

func (i ContactsPostResponse_social_accounts_type) String() string {
	return []string{"SOCIAL_ACCOUNT_TYPE_UNSPECIFIED", "FACEBOOK", "LINKED_IN", "TWITTER", "INSTAGRAM", "SNAPCHAT", "YOUTUBE", "PINTEREST"}[i]
}
func ParseContactsPostResponse_social_accounts_type(v string) (any, error) {
	result := SOCIAL_ACCOUNT_TYPE_UNSPECIFIED_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	switch v {
	case "SOCIAL_ACCOUNT_TYPE_UNSPECIFIED":
		result = SOCIAL_ACCOUNT_TYPE_UNSPECIFIED_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "FACEBOOK":
		result = FACEBOOK_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "LINKED_IN":
		result = LINKED_IN_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "TWITTER":
		result = TWITTER_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "INSTAGRAM":
		result = INSTAGRAM_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "SNAPCHAT":
		result = SNAPCHAT_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "YOUTUBE":
		result = YOUTUBE_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "PINTEREST":
		result = PINTEREST_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	default:
		return 0, errors.New("Unknown ContactsPostResponse_social_accounts_type value: " + v)
	}
	return &result, nil
}
func SerializeContactsPostResponse_social_accounts_type(values []ContactsPostResponse_social_accounts_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPostResponse_social_accounts_type) isMultiValue() bool {
	return false
}
