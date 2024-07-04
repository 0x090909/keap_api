package item

import (
	"errors"
)

type ContactIdGetResponse_social_accounts_type int

const (
	FACEBOOK_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE ContactIdGetResponse_social_accounts_type = iota
	TWITTER_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	LINKEDIN_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	INSTAGRAM_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	SNAPCHAT_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	YOUTUBE_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	PINTEREST_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
)

func (i ContactIdGetResponse_social_accounts_type) String() string {
	return []string{"Facebook", "Twitter", "LinkedIn", "Instagram", "Snapchat", "YouTube", "Pinterest"}[i]
}
func ParseContactIdGetResponse_social_accounts_type(v string) (any, error) {
	result := FACEBOOK_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	switch v {
	case "Facebook":
		result = FACEBOOK_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Twitter":
		result = TWITTER_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "LinkedIn":
		result = LINKEDIN_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Instagram":
		result = INSTAGRAM_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Snapchat":
		result = SNAPCHAT_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "YouTube":
		result = YOUTUBE_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Pinterest":
		result = PINTEREST_CONTACTIDGETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	default:
		return 0, errors.New("Unknown ContactIdGetResponse_social_accounts_type value: " + v)
	}
	return &result, nil
}
func SerializeContactIdGetResponse_social_accounts_type(values []ContactIdGetResponse_social_accounts_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactIdGetResponse_social_accounts_type) isMultiValue() bool {
	return false
}
