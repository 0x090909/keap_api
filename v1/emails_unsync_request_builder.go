package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EmailsUnsyncRequestBuilder builds and executes requests for operations under \v1\emails\unsync
type EmailsUnsyncRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// EmailsUnsyncRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsUnsyncRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewEmailsUnsyncRequestBuilderInternal instantiates a new EmailsUnsyncRequestBuilder and sets the default values.
func NewEmailsUnsyncRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsUnsyncRequestBuilder {
	m := &EmailsUnsyncRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/emails/unsync", pathParameters),
	}
	return m
}

// NewEmailsUnsyncRequestBuilder instantiates a new EmailsUnsyncRequestBuilder and sets the default values.
func NewEmailsUnsyncRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsUnsyncRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEmailsUnsyncRequestBuilderInternal(urlParams, requestAdapter)
}

// Post un-syncs a batch of email records
// Deprecated: This method is obsolete. Use PostAsUnsyncPostResponse instead.
// returns a EmailsUnsyncResponseable when successful
// returns a EmailsUnsync401Error error when the service returns a 401 status code
// returns a EmailsUnsync403Error error when the service returns a 403 status code
func (m *EmailsUnsyncRequestBuilder) Post(ctx context.Context, body EmailsUnsyncPostRequestBodyable, requestConfiguration *EmailsUnsyncRequestBuilderPostRequestConfiguration) (EmailsUnsyncResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsUnsync401ErrorFromDiscriminatorValue,
		"403": CreateEmailsUnsync403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsUnsyncResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsUnsyncResponseable), nil
}

// PostAsUnsyncPostResponse un-syncs a batch of email records
// returns a EmailsUnsyncPostResponseable when successful
// returns a EmailsUnsync401Error error when the service returns a 401 status code
// returns a EmailsUnsync403Error error when the service returns a 403 status code
func (m *EmailsUnsyncRequestBuilder) PostAsUnsyncPostResponse(ctx context.Context, body EmailsUnsyncPostRequestBodyable, requestConfiguration *EmailsUnsyncRequestBuilderPostRequestConfiguration) (EmailsUnsyncPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsUnsync401ErrorFromDiscriminatorValue,
		"403": CreateEmailsUnsync403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsUnsyncPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsUnsyncPostResponseable), nil
}

// ToPostRequestInformation un-syncs a batch of email records
// returns a *RequestInformation when successful
func (m *EmailsUnsyncRequestBuilder) ToPostRequestInformation(ctx context.Context, body EmailsUnsyncPostRequestBodyable, requestConfiguration *EmailsUnsyncRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmailsUnsyncRequestBuilder when successful
func (m *EmailsUnsyncRequestBuilder) WithUrl(rawUrl string) *EmailsUnsyncRequestBuilder {
	return NewEmailsUnsyncRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
