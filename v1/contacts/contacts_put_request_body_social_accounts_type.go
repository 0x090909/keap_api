package contacts

import (
	"errors"
)

type ContactsPutRequestBody_social_accounts_type int

const (
	FACEBOOK_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE ContactsPutRequestBody_social_accounts_type = iota
	TWITTER_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	LINKEDIN_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	INSTAGRAM_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	SNAPCHAT_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	YOUTUBE_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	PINTEREST_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
)

func (i ContactsPutRequestBody_social_accounts_type) String() string {
	return []string{"Facebook", "Twitter", "LinkedIn", "Instagram", "Snapchat", "YouTube", "Pinterest"}[i]
}
func ParseContactsPutRequestBody_social_accounts_type(v string) (any, error) {
	result := FACEBOOK_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	switch v {
	case "Facebook":
		result = FACEBOOK_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "Twitter":
		result = TWITTER_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "LinkedIn":
		result = LINKEDIN_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "Instagram":
		result = INSTAGRAM_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "Snapchat":
		result = SNAPCHAT_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "YouTube":
		result = YOUTUBE_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "Pinterest":
		result = PINTEREST_CONTACTSPUTREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	default:
		return 0, errors.New("Unknown ContactsPutRequestBody_social_accounts_type value: " + v)
	}
	return &result, nil
}
func SerializeContactsPutRequestBody_social_accounts_type(values []ContactsPutRequestBody_social_accounts_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPutRequestBody_social_accounts_type) isMultiValue() bool {
	return false
}
