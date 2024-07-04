package contacts

import (
	"errors"
)

type ContactsPutResponse_social_accounts_type int

const (
	FACEBOOK_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE ContactsPutResponse_social_accounts_type = iota
	TWITTER_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	LINKEDIN_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	INSTAGRAM_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	SNAPCHAT_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	YOUTUBE_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	PINTEREST_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
)

func (i ContactsPutResponse_social_accounts_type) String() string {
	return []string{"Facebook", "Twitter", "LinkedIn", "Instagram", "Snapchat", "YouTube", "Pinterest"}[i]
}
func ParseContactsPutResponse_social_accounts_type(v string) (any, error) {
	result := FACEBOOK_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	switch v {
	case "Facebook":
		result = FACEBOOK_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Twitter":
		result = TWITTER_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "LinkedIn":
		result = LINKEDIN_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Instagram":
		result = INSTAGRAM_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Snapchat":
		result = SNAPCHAT_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "YouTube":
		result = YOUTUBE_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Pinterest":
		result = PINTEREST_CONTACTSPUTRESPONSE_SOCIAL_ACCOUNTS_TYPE
	default:
		return 0, errors.New("Unknown ContactsPutResponse_social_accounts_type value: " + v)
	}
	return &result, nil
}
func SerializeContactsPutResponse_social_accounts_type(values []ContactsPutResponse_social_accounts_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsPutResponse_social_accounts_type) isMultiValue() bool {
	return false
}
