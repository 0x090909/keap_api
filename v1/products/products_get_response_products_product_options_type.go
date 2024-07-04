package products

import (
	"errors"
)

type ProductsGetResponse_products_product_options_type int

const (
	FIXEDLIST_PRODUCTSGETRESPONSE_PRODUCTS_PRODUCT_OPTIONS_TYPE ProductsGetResponse_products_product_options_type = iota
	VARIABLE_PRODUCTSGETRESPONSE_PRODUCTS_PRODUCT_OPTIONS_TYPE
)

func (i ProductsGetResponse_products_product_options_type) String() string {
	return []string{"FixedList", "Variable"}[i]
}
func ParseProductsGetResponse_products_product_options_type(v string) (any, error) {
	result := FIXEDLIST_PRODUCTSGETRESPONSE_PRODUCTS_PRODUCT_OPTIONS_TYPE
	switch v {
	case "FixedList":
		result = FIXEDLIST_PRODUCTSGETRESPONSE_PRODUCTS_PRODUCT_OPTIONS_TYPE
	case "Variable":
		result = VARIABLE_PRODUCTSGETRESPONSE_PRODUCTS_PRODUCT_OPTIONS_TYPE
	default:
		return 0, errors.New("Unknown ProductsGetResponse_products_product_options_type value: " + v)
	}
	return &result, nil
}
func SerializeProductsGetResponse_products_product_options_type(values []ProductsGetResponse_products_product_options_type) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ProductsGetResponse_products_product_options_type) isMultiValue() bool {
	return false
}
