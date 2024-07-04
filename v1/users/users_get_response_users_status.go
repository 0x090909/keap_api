package users

import (
	"errors"
)

type UsersGetResponse_users_status int

const (
	ACTIVE_USERSGETRESPONSE_USERS_STATUS UsersGetResponse_users_status = iota
	INVITED_USERSGETRESPONSE_USERS_STATUS
	INACTIVE_USERSGETRESPONSE_USERS_STATUS
)

func (i UsersGetResponse_users_status) String() string {
	return []string{"Active", "Invited", "Inactive"}[i]
}
func ParseUsersGetResponse_users_status(v string) (any, error) {
	result := ACTIVE_USERSGETRESPONSE_USERS_STATUS
	switch v {
	case "Active":
		result = ACTIVE_USERSGETRESPONSE_USERS_STATUS
	case "Invited":
		result = INVITED_USERSGETRESPONSE_USERS_STATUS
	case "Inactive":
		result = INACTIVE_USERSGETRESPONSE_USERS_STATUS
	default:
		return 0, errors.New("Unknown UsersGetResponse_users_status value: " + v)
	}
	return &result, nil
}
func SerializeUsersGetResponse_users_status(values []UsersGetResponse_users_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i UsersGetResponse_users_status) isMultiValue() bool {
	return false
}
