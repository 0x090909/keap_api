package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EmailsSendRequestBuilder builds and executes requests for operations under \v2\emails:send
type EmailsSendRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// EmailsSendRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsSendRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewEmailsSendRequestBuilderInternal instantiates a new EmailsSendRequestBuilder and sets the default values.
func NewEmailsSendRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsSendRequestBuilder {
	m := &EmailsSendRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/emails:send", pathParameters),
	}
	return m
}

// NewEmailsSendRequestBuilder instantiates a new EmailsSendRequestBuilder and sets the default values.
func NewEmailsSendRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsSendRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEmailsSendRequestBuilderInternal(urlParams, requestAdapter)
}

// Post sends an Email to a list of Contacts
// returns a EmailsSend401Error error when the service returns a 401 status code
// returns a EmailsSend403Error error when the service returns a 403 status code
func (m *EmailsSendRequestBuilder) Post(ctx context.Context, body EmailsSendPostRequestBodyable, requestConfiguration *EmailsSendRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsSend401ErrorFromDiscriminatorValue,
		"403": CreateEmailsSend403ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// ToPostRequestInformation sends an Email to a list of Contacts
// returns a *RequestInformation when successful
func (m *EmailsSendRequestBuilder) ToPostRequestInformation(ctx context.Context, body EmailsSendPostRequestBodyable, requestConfiguration *EmailsSendRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmailsSendRequestBuilder when successful
func (m *EmailsSendRequestBuilder) WithUrl(rawUrl string) *EmailsSendRequestBuilder {
	return NewEmailsSendRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
