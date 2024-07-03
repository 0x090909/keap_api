package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TagsItemCompaniesRequestBuilder builds and executes requests for operations under \v2\tags\{tag_id}\companies
type TagsItemCompaniesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TagsItemCompaniesRequestBuilderGetQueryParameters retrieves a list of Companies that have the given Tag applied.To search for `null` or empty fields use `filter==NONE`
type TagsItemCompaniesRequestBuilderGetQueryParameters struct {
	// Search filter to apply to results
	Filter *string `uriparametername:"filter"`
	// Attribute and direction to order items by. E.g. `given_name desc`
	Order_by *string `uriparametername:"order_by"`
	// Total number of items to return per page
	Page_size *int32 `uriparametername:"page_size"`
	// Page token
	Page_token *string `uriparametername:"page_token"`
}

// TagsItemCompaniesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsItemCompaniesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *TagsItemCompaniesRequestBuilderGetQueryParameters
}

// NewTagsItemCompaniesRequestBuilderInternal instantiates a new TagsItemCompaniesRequestBuilder and sets the default values.
func NewTagsItemCompaniesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsItemCompaniesRequestBuilder {
	m := &TagsItemCompaniesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/tags/{tag_id}/companies{?filter*,order_by*,page_size*,page_token*}", pathParameters),
	}
	return m
}

// NewTagsItemCompaniesRequestBuilder instantiates a new TagsItemCompaniesRequestBuilder and sets the default values.
func NewTagsItemCompaniesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsItemCompaniesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTagsItemCompaniesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of Companies that have the given Tag applied.To search for `null` or empty fields use `filter==NONE`
// Deprecated: This method is obsolete. Use GetAsCompaniesGetResponse instead.
// returns a TagsItemCompaniesResponseable when successful
// returns a TagsItemCompanies401Error error when the service returns a 401 status code
// returns a TagsItemCompanies403Error error when the service returns a 403 status code
// returns a TagsItemCompanies404Error error when the service returns a 404 status code
func (m *TagsItemCompaniesRequestBuilder) Get(ctx context.Context, requestConfiguration *TagsItemCompaniesRequestBuilderGetRequestConfiguration) (TagsItemCompaniesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemCompanies401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemCompanies403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemCompanies404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsItemCompaniesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsItemCompaniesResponseable), nil
}

// GetAsCompaniesGetResponse retrieves a list of Companies that have the given Tag applied.To search for `null` or empty fields use `filter==NONE`
// returns a TagsItemCompaniesGetResponseable when successful
// returns a TagsItemCompanies401Error error when the service returns a 401 status code
// returns a TagsItemCompanies403Error error when the service returns a 403 status code
// returns a TagsItemCompanies404Error error when the service returns a 404 status code
func (m *TagsItemCompaniesRequestBuilder) GetAsCompaniesGetResponse(ctx context.Context, requestConfiguration *TagsItemCompaniesRequestBuilderGetRequestConfiguration) (TagsItemCompaniesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemCompanies401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemCompanies403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemCompanies404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsItemCompaniesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsItemCompaniesGetResponseable), nil
}

// ToGetRequestInformation retrieves a list of Companies that have the given Tag applied.To search for `null` or empty fields use `filter==NONE`
// returns a *RequestInformation when successful
func (m *TagsItemCompaniesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TagsItemCompaniesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TagsItemCompaniesRequestBuilder when successful
func (m *TagsItemCompaniesRequestBuilder) WithUrl(rawUrl string) *TagsItemCompaniesRequestBuilder {
	return NewTagsItemCompaniesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
