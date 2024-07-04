package users

import (
	"errors"
)

type UsersGetResponse_users_fax_numbers_field int

const (
	FAX1_USERSGETRESPONSE_USERS_FAX_NUMBERS_FIELD UsersGetResponse_users_fax_numbers_field = iota
	FAX2_USERSGETRESPONSE_USERS_FAX_NUMBERS_FIELD
)

func (i UsersGetResponse_users_fax_numbers_field) String() string {
	return []string{"FAX1", "FAX2"}[i]
}
func ParseUsersGetResponse_users_fax_numbers_field(v string) (any, error) {
	result := FAX1_USERSGETRESPONSE_USERS_FAX_NUMBERS_FIELD
	switch v {
	case "FAX1":
		result = FAX1_USERSGETRESPONSE_USERS_FAX_NUMBERS_FIELD
	case "FAX2":
		result = FAX2_USERSGETRESPONSE_USERS_FAX_NUMBERS_FIELD
	default:
		return 0, errors.New("Unknown UsersGetResponse_users_fax_numbers_field value: " + v)
	}
	return &result, nil
}
func SerializeUsersGetResponse_users_fax_numbers_field(values []UsersGetResponse_users_fax_numbers_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i UsersGetResponse_users_fax_numbers_field) isMultiValue() bool {
	return false
}
