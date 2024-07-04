package item

import (
	"errors"
)

type ContactIdPatchResponse_social_accounts_type int

const (
	FACEBOOK_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE ContactIdPatchResponse_social_accounts_type = iota
	TWITTER_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	LINKEDIN_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	INSTAGRAM_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	SNAPCHAT_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	YOUTUBE_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	PINTEREST_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
)

func (i ContactIdPatchResponse_social_accounts_type) String() string {
	return []string{"Facebook", "Twitter", "LinkedIn", "Instagram", "Snapchat", "YouTube", "Pinterest"}[i]
}
func ParseContactIdPatchResponse_social_accounts_type(v string) (any, error) {
	result := FACEBOOK_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	switch v {
	case "Facebook":
		result = FACEBOOK_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Twitter":
		result = TWITTER_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "LinkedIn":
		result = LINKEDIN_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Instagram":
		result = INSTAGRAM_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Snapchat":
		result = SNAPCHAT_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "YouTube":
		result = YOUTUBE_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "Pinterest":
		result = PINTEREST_CONTACTIDPATCHRESPONSE_SOCIAL_ACCOUNTS_TYPE
	default:
		return 0, errors.New("Unknown ContactIdPatchResponse_social_accounts_type value: " + v)
	}
	return &result, nil
}
func SerializeContactIdPatchResponse_social_accounts_type(values []ContactIdPatchResponse_social_accounts_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactIdPatchResponse_social_accounts_type) isMultiValue() bool {
	return false
}
