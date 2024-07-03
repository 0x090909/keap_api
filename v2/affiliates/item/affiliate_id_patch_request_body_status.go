package item

import (
	"errors"
)

// The Affiliate Status
type Affiliate_idPatchRequestBody_status int

const (
	ACTIVE_AFFILIATE_IDPATCHREQUESTBODY_STATUS Affiliate_idPatchRequestBody_status = iota
	INACTIVE_AFFILIATE_IDPATCHREQUESTBODY_STATUS
)

func (i Affiliate_idPatchRequestBody_status) String() string {
	return []string{"active", "inactive"}[i]
}
func ParseAffiliate_idPatchRequestBody_status(v string) (any, error) {
	result := ACTIVE_AFFILIATE_IDPATCHREQUESTBODY_STATUS
	switch v {
	case "active":
		result = ACTIVE_AFFILIATE_IDPATCHREQUESTBODY_STATUS
	case "inactive":
		result = INACTIVE_AFFILIATE_IDPATCHREQUESTBODY_STATUS
	default:
		return 0, errors.New("Unknown Affiliate_idPatchRequestBody_status value: " + v)
	}
	return &result, nil
}
func SerializeAffiliate_idPatchRequestBody_status(values []Affiliate_idPatchRequestBody_status) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i Affiliate_idPatchRequestBody_status) isMultiValue() bool {
	return false
}
