package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EmailsSyncRequestBuilder builds and executes requests for operations under \v1\emails\sync
type EmailsSyncRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// EmailsSyncRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsSyncRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewEmailsSyncRequestBuilderInternal instantiates a new EmailsSyncRequestBuilder and sets the default values.
func NewEmailsSyncRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsSyncRequestBuilder {
	m := &EmailsSyncRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/emails/sync", pathParameters),
	}
	return m
}

// NewEmailsSyncRequestBuilder instantiates a new EmailsSyncRequestBuilder and sets the default values.
func NewEmailsSyncRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsSyncRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEmailsSyncRequestBuilderInternal(urlParams, requestAdapter)
}

// Post create a set of records of emails sent to contacts, maximum 1000 per transaction.
// Deprecated: This method is obsolete. Use PostAsSyncPostResponse instead.
// returns a EmailsSyncResponseable when successful
// returns a EmailsSync401Error error when the service returns a 401 status code
// returns a EmailsSync403Error error when the service returns a 403 status code
func (m *EmailsSyncRequestBuilder) Post(ctx context.Context, body EmailsSyncPostRequestBodyable, requestConfiguration *EmailsSyncRequestBuilderPostRequestConfiguration) (EmailsSyncResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsSync401ErrorFromDiscriminatorValue,
		"403": CreateEmailsSync403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsSyncResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsSyncResponseable), nil
}

// PostAsSyncPostResponse create a set of records of emails sent to contacts, maximum 1000 per transaction.
// returns a EmailsSyncPostResponseable when successful
// returns a EmailsSync401Error error when the service returns a 401 status code
// returns a EmailsSync403Error error when the service returns a 403 status code
func (m *EmailsSyncRequestBuilder) PostAsSyncPostResponse(ctx context.Context, body EmailsSyncPostRequestBodyable, requestConfiguration *EmailsSyncRequestBuilderPostRequestConfiguration) (EmailsSyncPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsSync401ErrorFromDiscriminatorValue,
		"403": CreateEmailsSync403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsSyncPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsSyncPostResponseable), nil
}

// ToPostRequestInformation create a set of records of emails sent to contacts, maximum 1000 per transaction.
// returns a *RequestInformation when successful
func (m *EmailsSyncRequestBuilder) ToPostRequestInformation(ctx context.Context, body EmailsSyncPostRequestBodyable, requestConfiguration *EmailsSyncRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmailsSyncRequestBuilder when successful
func (m *EmailsSyncRequestBuilder) WithUrl(rawUrl string) *EmailsSyncRequestBuilder {
	return NewEmailsSyncRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
