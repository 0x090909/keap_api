package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsItemLinksRequestBuilder builds and executes requests for operations under \v2\contacts\{contact_id}\links
type ContactsItemLinksRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsItemLinksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemLinksRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewContactsItemLinksRequestBuilderInternal instantiates a new ContactsItemLinksRequestBuilder and sets the default values.
func NewContactsItemLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemLinksRequestBuilder {
	m := &ContactsItemLinksRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/contacts/{contact_id}/links", pathParameters),
	}
	return m
}

// NewContactsItemLinksRequestBuilder instantiates a new ContactsItemLinksRequestBuilder and sets the default values.
func NewContactsItemLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemLinksRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsItemLinksRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of Linked Contacts for a given Contact
// Deprecated: This method is obsolete. Use GetAsLinksGetResponse instead.
// returns a ContactsItemLinksResponseable when successful
// returns a ContactsItemLinks401Error error when the service returns a 401 status code
// returns a ContactsItemLinks403Error error when the service returns a 403 status code
// returns a ContactsItemLinks404Error error when the service returns a 404 status code
func (m *ContactsItemLinksRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactsItemLinksRequestBuilderGetRequestConfiguration) (ContactsItemLinksResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemLinks401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemLinks403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemLinks404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemLinksResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemLinksResponseable), nil
}

// GetAsLinksGetResponse retrieves a list of Linked Contacts for a given Contact
// returns a ContactsItemLinksGetResponseable when successful
// returns a ContactsItemLinks401Error error when the service returns a 401 status code
// returns a ContactsItemLinks403Error error when the service returns a 403 status code
// returns a ContactsItemLinks404Error error when the service returns a 404 status code
func (m *ContactsItemLinksRequestBuilder) GetAsLinksGetResponse(ctx context.Context, requestConfiguration *ContactsItemLinksRequestBuilderGetRequestConfiguration) (ContactsItemLinksGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemLinks401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemLinks403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemLinks404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemLinksGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemLinksGetResponseable), nil
}

// ToGetRequestInformation retrieves a list of Linked Contacts for a given Contact
// returns a *RequestInformation when successful
func (m *ContactsItemLinksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContactsItemLinksRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ContactsItemLinksRequestBuilder when successful
func (m *ContactsItemLinksRequestBuilder) WithUrl(rawUrl string) *ContactsItemLinksRequestBuilder {
	return NewContactsItemLinksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
