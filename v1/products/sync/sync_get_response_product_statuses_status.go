package sync

import (
	"errors"
)

type SyncGetResponse_product_statuses_status int

const (
	CREATED_SYNCGETRESPONSE_PRODUCT_STATUSES_STATUS SyncGetResponse_product_statuses_status = iota
	UPDATED_SYNCGETRESPONSE_PRODUCT_STATUSES_STATUS
	DELETED_SYNCGETRESPONSE_PRODUCT_STATUSES_STATUS
)

func (i SyncGetResponse_product_statuses_status) String() string {
	return []string{"CREATED", "UPDATED", "DELETED"}[i]
}
func ParseSyncGetResponse_product_statuses_status(v string) (any, error) {
	result := CREATED_SYNCGETRESPONSE_PRODUCT_STATUSES_STATUS
	switch v {
	case "CREATED":
		result = CREATED_SYNCGETRESPONSE_PRODUCT_STATUSES_STATUS
	case "UPDATED":
		result = UPDATED_SYNCGETRESPONSE_PRODUCT_STATUSES_STATUS
	case "DELETED":
		result = DELETED_SYNCGETRESPONSE_PRODUCT_STATUSES_STATUS
	default:
		return 0, errors.New("Unknown SyncGetResponse_product_statuses_status value: " + v)
	}
	return &result, nil
}
func SerializeSyncGetResponse_product_statuses_status(values []SyncGetResponse_product_statuses_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i SyncGetResponse_product_statuses_status) isMultiValue() bool {
	return false
}
