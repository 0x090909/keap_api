package users

import (
	"errors"
)

type UsersGetResponse_users_phone_numbers_field int

const (
	PHONE1_USERSGETRESPONSE_USERS_PHONE_NUMBERS_FIELD UsersGetResponse_users_phone_numbers_field = iota
	PHONE2_USERSGETRESPONSE_USERS_PHONE_NUMBERS_FIELD
	PHONE3_USERSGETRESPONSE_USERS_PHONE_NUMBERS_FIELD
	PHONE4_USERSGETRESPONSE_USERS_PHONE_NUMBERS_FIELD
	PHONE5_USERSGETRESPONSE_USERS_PHONE_NUMBERS_FIELD
)

func (i UsersGetResponse_users_phone_numbers_field) String() string {
	return []string{"PHONE1", "PHONE2", "PHONE3", "PHONE4", "PHONE5"}[i]
}
func ParseUsersGetResponse_users_phone_numbers_field(v string) (any, error) {
	result := PHONE1_USERSGETRESPONSE_USERS_PHONE_NUMBERS_FIELD
	switch v {
	case "PHONE1":
		result = PHONE1_USERSGETRESPONSE_USERS_PHONE_NUMBERS_FIELD
	case "PHONE2":
		result = PHONE2_USERSGETRESPONSE_USERS_PHONE_NUMBERS_FIELD
	case "PHONE3":
		result = PHONE3_USERSGETRESPONSE_USERS_PHONE_NUMBERS_FIELD
	case "PHONE4":
		result = PHONE4_USERSGETRESPONSE_USERS_PHONE_NUMBERS_FIELD
	case "PHONE5":
		result = PHONE5_USERSGETRESPONSE_USERS_PHONE_NUMBERS_FIELD
	default:
		return 0, errors.New("Unknown UsersGetResponse_users_phone_numbers_field value: " + v)
	}
	return &result, nil
}
func SerializeUsersGetResponse_users_phone_numbers_field(values []UsersGetResponse_users_phone_numbers_field) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i UsersGetResponse_users_phone_numbers_field) isMultiValue() bool {
	return false
}
