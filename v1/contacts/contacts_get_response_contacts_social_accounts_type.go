package contacts

import (
	"errors"
)

type ContactsGetResponse_contacts_social_accounts_type int

const (
	FACEBOOK_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE ContactsGetResponse_contacts_social_accounts_type = iota
	TWITTER_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	LINKEDIN_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	INSTAGRAM_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	SNAPCHAT_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	YOUTUBE_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	PINTEREST_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
)

func (i ContactsGetResponse_contacts_social_accounts_type) String() string {
	return []string{"Facebook", "Twitter", "LinkedIn", "Instagram", "Snapchat", "YouTube", "Pinterest"}[i]
}
func ParseContactsGetResponse_contacts_social_accounts_type(v string) (any, error) {
	result := FACEBOOK_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	switch v {
	case "Facebook":
		result = FACEBOOK_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	case "Twitter":
		result = TWITTER_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	case "LinkedIn":
		result = LINKEDIN_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	case "Instagram":
		result = INSTAGRAM_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	case "Snapchat":
		result = SNAPCHAT_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	case "YouTube":
		result = YOUTUBE_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	case "Pinterest":
		result = PINTEREST_CONTACTSGETRESPONSE_CONTACTS_SOCIAL_ACCOUNTS_TYPE
	default:
		return 0, errors.New("Unknown ContactsGetResponse_contacts_social_accounts_type value: " + v)
	}
	return &result, nil
}
func SerializeContactsGetResponse_contacts_social_accounts_type(values []ContactsGetResponse_contacts_social_accounts_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ContactsGetResponse_contacts_social_accounts_type) isMultiValue() bool {
	return false
}
