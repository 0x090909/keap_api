package sync

import (
	"errors"
)

type SyncGetResponse_product_statuses_product_product_options_type int

const (
	FIXEDLIST_SYNCGETRESPONSE_PRODUCT_STATUSES_PRODUCT_PRODUCT_OPTIONS_TYPE SyncGetResponse_product_statuses_product_product_options_type = iota
	VARIABLE_SYNCGETRESPONSE_PRODUCT_STATUSES_PRODUCT_PRODUCT_OPTIONS_TYPE
)

func (i SyncGetResponse_product_statuses_product_product_options_type) String() string {
	return []string{"FixedList", "Variable"}[i]
}
func ParseSyncGetResponse_product_statuses_product_product_options_type(v string) (any, error) {
	result := FIXEDLIST_SYNCGETRESPONSE_PRODUCT_STATUSES_PRODUCT_PRODUCT_OPTIONS_TYPE
	switch v {
	case "FixedList":
		result = FIXEDLIST_SYNCGETRESPONSE_PRODUCT_STATUSES_PRODUCT_PRODUCT_OPTIONS_TYPE
	case "Variable":
		result = VARIABLE_SYNCGETRESPONSE_PRODUCT_STATUSES_PRODUCT_PRODUCT_OPTIONS_TYPE
	default:
		return 0, errors.New("Unknown SyncGetResponse_product_statuses_product_product_options_type value: " + v)
	}
	return &result, nil
}
func SerializeSyncGetResponse_product_statuses_product_product_options_type(values []SyncGetResponse_product_statuses_product_product_options_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i SyncGetResponse_product_statuses_product_product_options_type) isMultiValue() bool {
	return false
}
