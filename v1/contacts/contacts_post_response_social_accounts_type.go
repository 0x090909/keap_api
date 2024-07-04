package contacts

import (
	"errors"
)

type ContactsPostResponse_social_accounts_type int

const (
	FACEBOOK_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE ContactsPostResponse_social_accounts_type = iota
	TWITTER_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	LINKEDIN_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	INSTAGRAM_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	SNAPCHAT_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	YOUTUBE_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	PINTEREST_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
)

func (i ContactsPostResponse_social_accounts_type) String() string {
	return []string{"Facebook", "Twitter", "LinkedIn", "Instagram", "Snapchat", "YouTube", "Pinterest"}[i]
}
func ParseContactsPostResponse_social_accounts_type(v string) (any, error) {
	result := FACEBOOK_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	switch v {
	case "Facebook":
		result = FACEBOOK_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Twitter":
		result = TWITTER_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "LinkedIn":
		result = LINKEDIN_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Instagram":
		result = INSTAGRAM_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Snapchat":
		result = SNAPCHAT_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "YouTube":
		result = YOUTUBE_CONTACTSPOSTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Pinterest":
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
