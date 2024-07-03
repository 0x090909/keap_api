package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EmailsBatchAddRequestBuilder builds and executes requests for operations under \v2\emails:batchAdd
type EmailsBatchAddRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// EmailsBatchAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsBatchAddRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewEmailsBatchAddRequestBuilderInternal instantiates a new EmailsBatchAddRequestBuilder and sets the default values.
func NewEmailsBatchAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsBatchAddRequestBuilder {
	m := &EmailsBatchAddRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/emails:batchAdd", pathParameters),
	}
	return m
}

// NewEmailsBatchAddRequestBuilder instantiates a new EmailsBatchAddRequestBuilder and sets the default values.
func NewEmailsBatchAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsBatchAddRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEmailsBatchAddRequestBuilderInternal(urlParams, requestAdapter)
}

// Post creates a set of Records of Emails sent to Contacts, maximum 1000 per transaction.
// Deprecated: This method is obsolete. Use PostAsEmailsBatchAddPostResponse instead.
// returns a EmailsBatchAddResponseable when successful
// returns a EmailsBatchAdd401Error error when the service returns a 401 status code
// returns a EmailsBatchAdd403Error error when the service returns a 403 status code
func (m *EmailsBatchAddRequestBuilder) Post(ctx context.Context, body EmailsBatchAddPostRequestBodyable, requestConfiguration *EmailsBatchAddRequestBuilderPostRequestConfiguration) (EmailsBatchAddResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsBatchAdd401ErrorFromDiscriminatorValue,
		"403": CreateEmailsBatchAdd403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsBatchAddResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsBatchAddResponseable), nil
}

// PostAsEmailsBatchAddPostResponse creates a set of Records of Emails sent to Contacts, maximum 1000 per transaction.
// returns a EmailsBatchAddPostResponseable when successful
// returns a EmailsBatchAdd401Error error when the service returns a 401 status code
// returns a EmailsBatchAdd403Error error when the service returns a 403 status code
func (m *EmailsBatchAddRequestBuilder) PostAsEmailsBatchAddPostResponse(ctx context.Context, body EmailsBatchAddPostRequestBodyable, requestConfiguration *EmailsBatchAddRequestBuilderPostRequestConfiguration) (EmailsBatchAddPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsBatchAdd401ErrorFromDiscriminatorValue,
		"403": CreateEmailsBatchAdd403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsBatchAddPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsBatchAddPostResponseable), nil
}

// ToPostRequestInformation creates a set of Records of Emails sent to Contacts, maximum 1000 per transaction.
// returns a *RequestInformation when successful
func (m *EmailsBatchAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body EmailsBatchAddPostRequestBodyable, requestConfiguration *EmailsBatchAddRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmailsBatchAddRequestBuilder when successful
func (m *EmailsBatchAddRequestBuilder) WithUrl(rawUrl string) *EmailsBatchAddRequestBuilder {
	return NewEmailsBatchAddRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
