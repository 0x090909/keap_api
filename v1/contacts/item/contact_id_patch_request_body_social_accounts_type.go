package item

import (
	"errors"
)

type ContactIdPatchRequestBody_social_accounts_type int

const (
	FACEBOOK_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE ContactIdPatchRequestBody_social_accounts_type = iota
	TWITTER_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	LINKEDIN_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	INSTAGRAM_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	SNAPCHAT_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	YOUTUBE_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	PINTEREST_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
)

func (i ContactIdPatchRequestBody_social_accounts_type) String() string {
	return []string{"Facebook", "Twitter", "LinkedIn", "Instagram", "Snapchat", "YouTube", "Pinterest"}[i]
}
func ParseContactIdPatchRequestBody_social_accounts_type(v string) (any, error) {
	result := FACEBOOK_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	switch v {
	case "Facebook":
		result = FACEBOOK_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "Twitter":
		result = TWITTER_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "LinkedIn":
		result = LINKEDIN_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "Instagram":
		result = INSTAGRAM_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "Snapchat":
		result = SNAPCHAT_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "YouTube":
		result = YOUTUBE_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	case "Pinterest":
		result = PINTEREST_CONTACTIDPATCHREQUESTBODY_SOCIAL_ACCOUNTS_TYPE
	default:
		return 0, errors.New("Unknown ContactIdPatchRequestBody_social_accounts_type value: " + v)
	}
	return &result, nil
}
func SerializeContactIdPatchRequestBody_social_accounts_type(values []ContactIdPatchRequestBody_social_accounts_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactIdPatchRequestBody_social_accounts_type) isMultiValue() bool {
	return false
}
