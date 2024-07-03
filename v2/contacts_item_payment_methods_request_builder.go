package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsItemPaymentMethodsRequestBuilder builds and executes requests for operations under \v2\contacts\{contact_id}\paymentMethods
type ContactsItemPaymentMethodsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsItemPaymentMethodsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemPaymentMethodsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewContactsItemPaymentMethodsRequestBuilderInternal instantiates a new ContactsItemPaymentMethodsRequestBuilder and sets the default values.
func NewContactsItemPaymentMethodsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemPaymentMethodsRequestBuilder {
	m := &ContactsItemPaymentMethodsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/contacts/{contact_id}/paymentMethods", pathParameters),
	}
	return m
}

// NewContactsItemPaymentMethodsRequestBuilder instantiates a new ContactsItemPaymentMethodsRequestBuilder and sets the default values.
func NewContactsItemPaymentMethodsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemPaymentMethodsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsItemPaymentMethodsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get list all Payment Methods for a Contact.
// Deprecated: This method is obsolete. Use GetAsPaymentMethodsGetResponse instead.
// returns a ContactsItemPaymentMethodsResponseable when successful
// returns a ContactsItemPaymentMethods401Error error when the service returns a 401 status code
// returns a ContactsItemPaymentMethods403Error error when the service returns a 403 status code
// returns a ContactsItemPaymentMethods404Error error when the service returns a 404 status code
func (m *ContactsItemPaymentMethodsRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactsItemPaymentMethodsRequestBuilderGetRequestConfiguration) (ContactsItemPaymentMethodsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemPaymentMethods401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemPaymentMethods403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemPaymentMethods404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemPaymentMethodsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemPaymentMethodsResponseable), nil
}

// GetAsPaymentMethodsGetResponse list all Payment Methods for a Contact.
// returns a ContactsItemPaymentMethodsGetResponseable when successful
// returns a ContactsItemPaymentMethods401Error error when the service returns a 401 status code
// returns a ContactsItemPaymentMethods403Error error when the service returns a 403 status code
// returns a ContactsItemPaymentMethods404Error error when the service returns a 404 status code
func (m *ContactsItemPaymentMethodsRequestBuilder) GetAsPaymentMethodsGetResponse(ctx context.Context, requestConfiguration *ContactsItemPaymentMethodsRequestBuilderGetRequestConfiguration) (ContactsItemPaymentMethodsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemPaymentMethods401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemPaymentMethods403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemPaymentMethods404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemPaymentMethodsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemPaymentMethodsGetResponseable), nil
}

// ToGetRequestInformation list all Payment Methods for a Contact.
// returns a *RequestInformation when successful
func (m *ContactsItemPaymentMethodsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContactsItemPaymentMethodsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ContactsItemPaymentMethodsRequestBuilder when successful
func (m *ContactsItemPaymentMethodsRequestBuilder) WithUrl(rawUrl string) *ContactsItemPaymentMethodsRequestBuilder {
	return NewContactsItemPaymentMethodsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
