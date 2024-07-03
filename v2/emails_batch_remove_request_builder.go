package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EmailsBatchRemoveRequestBuilder builds and executes requests for operations under \v2\emails:batchRemove
type EmailsBatchRemoveRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// EmailsBatchRemoveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsBatchRemoveRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewEmailsBatchRemoveRequestBuilderInternal instantiates a new EmailsBatchRemoveRequestBuilder and sets the default values.
func NewEmailsBatchRemoveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsBatchRemoveRequestBuilder {
	m := &EmailsBatchRemoveRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/emails:batchRemove", pathParameters),
	}
	return m
}

// NewEmailsBatchRemoveRequestBuilder instantiates a new EmailsBatchRemoveRequestBuilder and sets the default values.
func NewEmailsBatchRemoveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsBatchRemoveRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEmailsBatchRemoveRequestBuilderInternal(urlParams, requestAdapter)
}

// Post removes a set of Email Records
// Deprecated: This method is obsolete. Use PostAsEmailsBatchRemovePostResponse instead.
// returns a EmailsBatchRemoveResponseable when successful
// returns a EmailsBatchRemove401Error error when the service returns a 401 status code
// returns a EmailsBatchRemove403Error error when the service returns a 403 status code
func (m *EmailsBatchRemoveRequestBuilder) Post(ctx context.Context, body EmailsBatchRemovePostRequestBodyable, requestConfiguration *EmailsBatchRemoveRequestBuilderPostRequestConfiguration) (EmailsBatchRemoveResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsBatchRemove401ErrorFromDiscriminatorValue,
		"403": CreateEmailsBatchRemove403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsBatchRemoveResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsBatchRemoveResponseable), nil
}

// PostAsEmailsBatchRemovePostResponse removes a set of Email Records
// returns a EmailsBatchRemovePostResponseable when successful
// returns a EmailsBatchRemove401Error error when the service returns a 401 status code
// returns a EmailsBatchRemove403Error error when the service returns a 403 status code
func (m *EmailsBatchRemoveRequestBuilder) PostAsEmailsBatchRemovePostResponse(ctx context.Context, body EmailsBatchRemovePostRequestBodyable, requestConfiguration *EmailsBatchRemoveRequestBuilderPostRequestConfiguration) (EmailsBatchRemovePostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsBatchRemove401ErrorFromDiscriminatorValue,
		"403": CreateEmailsBatchRemove403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsBatchRemovePostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsBatchRemovePostResponseable), nil
}

// ToPostRequestInformation removes a set of Email Records
// returns a *RequestInformation when successful
func (m *EmailsBatchRemoveRequestBuilder) ToPostRequestInformation(ctx context.Context, body EmailsBatchRemovePostRequestBodyable, requestConfiguration *EmailsBatchRemoveRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *EmailsBatchRemoveRequestBuilder when successful
func (m *EmailsBatchRemoveRequestBuilder) WithUrl(rawUrl string) *EmailsBatchRemoveRequestBuilder {
	return NewEmailsBatchRemoveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
