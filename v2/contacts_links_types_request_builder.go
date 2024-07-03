package v2

import (
	"context"
	i6c05cc2e03524da0fd5934fe2ac36bba238943fa63d1be33d8769580d294da4a "github.com/0x090909/keap_api/v2/contacts/links/types"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsLinksTypesRequestBuilder builds and executes requests for operations under \v2\contacts\links\types
type ContactsLinksTypesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsLinksTypesRequestBuilderGetQueryParameters retrieves a list of Contact Link types.
type ContactsLinksTypesRequestBuilderGetQueryParameters struct {
	// Search filter to apply to results. Formatted as (unencoded) ?filter=name==expectedValue
	// Deprecated: This property is deprecated, use FilterAsGetFilterQueryParameterType instead
	Filter *string `uriparametername:"filter"`
	// Search filter to apply to results. Formatted as (unencoded) ?filter=name==expectedValue
	FilterAsGetFilterQueryParameterType *i6c05cc2e03524da0fd5934fe2ac36bba238943fa63d1be33d8769580d294da4a.GetFilterQueryParameterType `uriparametername:"filter"`
	OrderBy                             *string                                                                                        `uriparametername:"orderBy"`
	PageSize                            *int32                                                                                         `uriparametername:"pageSize"`
	PageToken                           *string                                                                                        `uriparametername:"pageToken"`
}

// ContactsLinksTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsLinksTypesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ContactsLinksTypesRequestBuilderGetQueryParameters
}

// ContactsLinksTypesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsLinksTypesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewContactsLinksTypesRequestBuilderInternal instantiates a new ContactsLinksTypesRequestBuilder and sets the default values.
func NewContactsLinksTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsLinksTypesRequestBuilder {
	m := &ContactsLinksTypesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/contacts/links/types{?filter*,orderBy*,pageSize*,pageToken*}", pathParameters),
	}
	return m
}

// NewContactsLinksTypesRequestBuilder instantiates a new ContactsLinksTypesRequestBuilder and sets the default values.
func NewContactsLinksTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsLinksTypesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsLinksTypesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of Contact Link types.
// Deprecated: This method is obsolete. Use GetAsTypesGetResponse instead.
// returns a ContactsLinksTypesResponseable when successful
// returns a ContactsLinksTypes401Error error when the service returns a 401 status code
// returns a ContactsLinksTypes403Error error when the service returns a 403 status code
// returns a ContactsLinksTypes404Error error when the service returns a 404 status code
func (m *ContactsLinksTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactsLinksTypesRequestBuilderGetRequestConfiguration) (ContactsLinksTypesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsLinksTypes401ErrorFromDiscriminatorValue,
		"403": CreateContactsLinksTypes403ErrorFromDiscriminatorValue,
		"404": CreateContactsLinksTypes404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsLinksTypesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsLinksTypesResponseable), nil
}

// GetAsTypesGetResponse retrieves a list of Contact Link types.
// returns a ContactsLinksTypesGetResponseable when successful
// returns a ContactsLinksTypes401Error error when the service returns a 401 status code
// returns a ContactsLinksTypes403Error error when the service returns a 403 status code
// returns a ContactsLinksTypes404Error error when the service returns a 404 status code
func (m *ContactsLinksTypesRequestBuilder) GetAsTypesGetResponse(ctx context.Context, requestConfiguration *ContactsLinksTypesRequestBuilderGetRequestConfiguration) (ContactsLinksTypesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsLinksTypes401ErrorFromDiscriminatorValue,
		"403": CreateContactsLinksTypes403ErrorFromDiscriminatorValue,
		"404": CreateContactsLinksTypes404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsLinksTypesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsLinksTypesGetResponseable), nil
}

// Post creates a new type of Contact Link.
// Deprecated: This method is obsolete. Use PostAsTypesPostResponse instead.
// returns a ContactsLinksTypesResponseable when successful
// returns a ContactsLinksTypes401Error error when the service returns a 401 status code
// returns a ContactsLinksTypes403Error error when the service returns a 403 status code
func (m *ContactsLinksTypesRequestBuilder) Post(ctx context.Context, body ContactsLinksTypesPostRequestBodyable, requestConfiguration *ContactsLinksTypesRequestBuilderPostRequestConfiguration) (ContactsLinksTypesResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsLinksTypes401ErrorFromDiscriminatorValue,
		"403": CreateContactsLinksTypes403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsLinksTypesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsLinksTypesResponseable), nil
}

// PostAsTypesPostResponse creates a new type of Contact Link.
// returns a ContactsLinksTypesPostResponseable when successful
// returns a ContactsLinksTypes401Error error when the service returns a 401 status code
// returns a ContactsLinksTypes403Error error when the service returns a 403 status code
func (m *ContactsLinksTypesRequestBuilder) PostAsTypesPostResponse(ctx context.Context, body ContactsLinksTypesPostRequestBodyable, requestConfiguration *ContactsLinksTypesRequestBuilderPostRequestConfiguration) (ContactsLinksTypesPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsLinksTypes401ErrorFromDiscriminatorValue,
		"403": CreateContactsLinksTypes403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsLinksTypesPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsLinksTypesPostResponseable), nil
}

// ToGetRequestInformation retrieves a list of Contact Link types.
// returns a *RequestInformation when successful
func (m *ContactsLinksTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContactsLinksTypesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPostRequestInformation creates a new type of Contact Link.
// returns a *RequestInformation when successful
func (m *ContactsLinksTypesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContactsLinksTypesPostRequestBodyable, requestConfiguration *ContactsLinksTypesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v2/contacts/links/types", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ContactsLinksTypesRequestBuilder when successful
func (m *ContactsLinksTypesRequestBuilder) WithUrl(rawUrl string) *ContactsLinksTypesRequestBuilder {
	return NewContactsLinksTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
