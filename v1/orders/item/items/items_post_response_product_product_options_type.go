package items

import (
	"errors"
)

type ItemsPostResponse_product_product_options_type int

const (
	FIXEDLIST_ITEMSPOSTRESPONSE_PRODUCT_PRODUCT_OPTIONS_TYPE ItemsPostResponse_product_product_options_type = iota
	VARIABLE_ITEMSPOSTRESPONSE_PRODUCT_PRODUCT_OPTIONS_TYPE
)

func (i ItemsPostResponse_product_product_options_type) String() string {
	return []string{"FixedList", "Variable"}[i]
}
func ParseItemsPostResponse_product_product_options_type(v string) (any, error) {
	result := FIXEDLIST_ITEMSPOSTRESPONSE_PRODUCT_PRODUCT_OPTIONS_TYPE
	switch v {
	case "FixedList":
		result = FIXEDLIST_ITEMSPOSTRESPONSE_PRODUCT_PRODUCT_OPTIONS_TYPE
	case "Variable":
		result = VARIABLE_ITEMSPOSTRESPONSE_PRODUCT_PRODUCT_OPTIONS_TYPE
	default:
		return 0, errors.New("Unknown ItemsPostResponse_product_product_options_type value: " + v)
	}
	return &result, nil
}
func SerializeItemsPostResponse_product_product_options_type(values []ItemsPostResponse_product_product_options_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ItemsPostResponse_product_product_options_type) isMultiValue() bool {
	return false
}
