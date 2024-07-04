package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EmailsQueueRequestBuilder builds and executes requests for operations under \v1\emails\queue
type EmailsQueueRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// EmailsQueueRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsQueueRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewEmailsQueueRequestBuilderInternal instantiates a new EmailsQueueRequestBuilder and sets the default values.
func NewEmailsQueueRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsQueueRequestBuilder {
	m := &EmailsQueueRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/emails/queue", pathParameters),
	}
	return m
}

// NewEmailsQueueRequestBuilder instantiates a new EmailsQueueRequestBuilder and sets the default values.
func NewEmailsQueueRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsQueueRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEmailsQueueRequestBuilderInternal(urlParams, requestAdapter)
}

// Post send an Email to a list of Contacts
// returns a EmailsQueue401Error error when the service returns a 401 status code
// returns a EmailsQueue403Error error when the service returns a 403 status code
func (m *EmailsQueueRequestBuilder) Post(ctx context.Context, body EmailsQueuePostRequestBodyable, requestConfiguration *EmailsQueueRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsQueue401ErrorFromDiscriminatorValue,
		"403": CreateEmailsQueue403ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// ToPostRequestInformation send an Email to a list of Contacts
// returns a *RequestInformation when successful
func (m *EmailsQueueRequestBuilder) ToPostRequestInformation(ctx context.Context, body EmailsQueuePostRequestBodyable, requestConfiguration *EmailsQueueRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmailsQueueRequestBuilder when successful
func (m *EmailsQueueRequestBuilder) WithUrl(rawUrl string) *EmailsQueueRequestBuilder {
	return NewEmailsQueueRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
