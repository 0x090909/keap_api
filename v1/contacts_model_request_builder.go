package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsModelRequestBuilder builds and executes requests for operations under \v1\contacts\model
type ContactsModelRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsModelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsModelRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewContactsModelRequestBuilderInternal instantiates a new ContactsModelRequestBuilder and sets the default values.
func NewContactsModelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsModelRequestBuilder {
	m := &ContactsModelRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/contacts/model", pathParameters),
	}
	return m
}

// NewContactsModelRequestBuilder instantiates a new ContactsModelRequestBuilder and sets the default values.
func NewContactsModelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsModelRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsModelRequestBuilderInternal(urlParams, requestAdapter)
}

// CustomFields the customFields property
// returns a *ContactsModelCustomFieldsRequestBuilder when successful
func (m *ContactsModelRequestBuilder) CustomFields() *ContactsModelCustomFieldsRequestBuilder {
	return NewContactsModelCustomFieldsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get get the custom fields and optional properties for the Contact object
// Deprecated: This method is obsolete. Use GetAsModelGetResponse instead.
// returns a ContactsModelResponseable when successful
// returns a ContactsModel401Error error when the service returns a 401 status code
// returns a ContactsModel403Error error when the service returns a 403 status code
// returns a ContactsModel404Error error when the service returns a 404 status code
func (m *ContactsModelRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactsModelRequestBuilderGetRequestConfiguration) (ContactsModelResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsModel401ErrorFromDiscriminatorValue,
		"403": CreateContactsModel403ErrorFromDiscriminatorValue,
		"404": CreateContactsModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsModelResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsModelResponseable), nil
}

// GetAsModelGetResponse get the custom fields and optional properties for the Contact object
// returns a ContactsModelGetResponseable when successful
// returns a ContactsModel401Error error when the service returns a 401 status code
// returns a ContactsModel403Error error when the service returns a 403 status code
// returns a ContactsModel404Error error when the service returns a 404 status code
func (m *ContactsModelRequestBuilder) GetAsModelGetResponse(ctx context.Context, requestConfiguration *ContactsModelRequestBuilderGetRequestConfiguration) (ContactsModelGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsModel401ErrorFromDiscriminatorValue,
		"403": CreateContactsModel403ErrorFromDiscriminatorValue,
		"404": CreateContactsModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsModelGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsModelGetResponseable), nil
}

// ToGetRequestInformation get the custom fields and optional properties for the Contact object
// returns a *RequestInformation when successful
func (m *ContactsModelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContactsModelRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ContactsModelRequestBuilder when successful
func (m *ContactsModelRequestBuilder) WithUrl(rawUrl string) *ContactsModelRequestBuilder {
	return NewContactsModelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
