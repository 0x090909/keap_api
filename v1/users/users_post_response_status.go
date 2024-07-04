package users

import (
	"errors"
)

type UsersPostResponse_status int

const (
	ACTIVE_USERSPOSTRESPONSE_STATUS UsersPostResponse_status = iota
	INVITED_USERSPOSTRESPONSE_STATUS
	INACTIVE_USERSPOSTRESPONSE_STATUS
)

func (i UsersPostResponse_status) String() string {
	return []string{"Active", "Invited", "Inactive"}[i]
}
func ParseUsersPostResponse_status(v string) (any, error) {
	result := ACTIVE_USERSPOSTRESPONSE_STATUS
	switch v {
	case "Active":
		result = ACTIVE_USERSPOSTRESPONSE_STATUS
	case "Invited":
		result = INVITED_USERSPOSTRESPONSE_STATUS
	case "Inactive":
		result = INACTIVE_USERSPOSTRESPONSE_STATUS
	default:
		return 0, errors.New("Unknown UsersPostResponse_status value: " + v)
	}
	return &result, nil
}
func SerializeUsersPostResponse_status(values []UsersPostResponse_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i UsersPostResponse_status) isMultiValue() bool {
	return false
}
