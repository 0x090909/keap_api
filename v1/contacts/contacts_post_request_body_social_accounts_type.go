package contacts

import (
	"errors"
)

type ContactsPostRequestBody_social_accounts_type int

const (
	FACEBOOK_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE ContactsPostRequestBody_social_accounts_type = iota
	TWITTER_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	LINKEDIN_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	INSTAGRAM_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	SNAPCHAT_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	YOUTUBE_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	PINTEREST_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
)

func (i ContactsPostRequestBody_social_accounts_type) String() string {
	return []string{"Facebook", "Twitter", "LinkedIn", "Instagram", "Snapchat", "YouTube", "Pinterest"}[i]
}
func ParseContactsPostRequestBody_social_accounts_type(v string) (any, error) {
	result := FACEBOOK_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	switch v {
	case "Facebook":
		result = FACEBOOK_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "Twitter":
		result = TWITTER_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "LinkedIn":
		result = LINKEDIN_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "Instagram":
		result = INSTAGRAM_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "Snapchat":
		result = SNAPCHAT_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "YouTube":
		result = YOUTUBE_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "Pinterest":
		result = PINTEREST_CONTACTSPOSTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	default:
		return 0, errors.New("Unknown ContactsPostRequestBody_social_accounts_type value: " + v)
	}
	return &result, nil
}
func SerializeContactsPostRequestBody_social_accounts_type(values []ContactsPostRequestBody_social_accounts_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPostRequestBody_social_accounts_type) isMultiValue() bool {
	return false
}
