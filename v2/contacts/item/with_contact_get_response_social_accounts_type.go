package item

import (
	"errors"
)

type WithContact_GetResponse_social_accounts_type int

const (
	SOCIAL_ACCOUNT_TYPE_UNSPECIFIED_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE WithContact_GetResponse_social_accounts_type = iota
	FACEBOOK_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	LINKED_IN_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	TWITTER_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	INSTAGRAM_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	SNAPCHAT_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	YOUTUBE_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	PINTEREST_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
)

func (i WithContact_GetResponse_social_accounts_type) String() string {
	return []string{"SOCIAL_ACCOUNT_TYPE_UNSPECIFIED", "FACEBOOK", "LINKED_IN", "TWITTER", "INSTAGRAM", "SNAPCHAT", "YOUTUBE", "PINTEREST"}[i]
}
func ParseWithContact_GetResponse_social_accounts_type(v string) (any, error) {
	result := SOCIAL_ACCOUNT_TYPE_UNSPECIFIED_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	switch v {
	case "SOCIAL_ACCOUNT_TYPE_UNSPECIFIED":
		result = SOCIAL_ACCOUNT_TYPE_UNSPECIFIED_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "FACEBOOK":
		result = FACEBOOK_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "LINKED_IN":
		result = LINKED_IN_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "TWITTER":
		result = TWITTER_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "INSTAGRAM":
		result = INSTAGRAM_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "SNAPCHAT":
		result = SNAPCHAT_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "YOUTUBE":
		result = YOUTUBE_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	case "PINTEREST":
		result = PINTEREST_WITHCONTACT_GETRESPONSE_SOCIAL_ACCOUNTS_TYPE
	default:
		return 0, errors.New("Unknown WithContact_GetResponse_social_accounts_type value: " + v)
	}
	return &result, nil
}
func SerializeWithContact_GetResponse_social_accounts_type(values []WithContact_GetResponse_social_accounts_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i WithContact_GetResponse_social_accounts_type) isMultiValue() bool {
	return false
}
