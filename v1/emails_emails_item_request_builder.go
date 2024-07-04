package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EmailsEmailsItemRequestBuilder builds and executes requests for operations under \v1\emails\{id}
type EmailsEmailsItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// EmailsEmailsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsEmailsItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// EmailsEmailsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsEmailsItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewEmailsEmailsItemRequestBuilderInternal instantiates a new EmailsEmailsItemRequestBuilder and sets the default values.
func NewEmailsEmailsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsEmailsItemRequestBuilder {
	m := &EmailsEmailsItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/emails/{id}", pathParameters),
	}
	return m
}

// NewEmailsEmailsItemRequestBuilder instantiates a new EmailsEmailsItemRequestBuilder and sets the default values.
func NewEmailsEmailsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsEmailsItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEmailsEmailsItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete a specific email record
// returns a EmailsItemEmails401Error error when the service returns a 401 status code
// returns a EmailsItemEmails403Error error when the service returns a 403 status code
// returns a EmailsItemEmails404Error error when the service returns a 404 status code
func (m *EmailsEmailsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EmailsEmailsItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsItemEmails401ErrorFromDiscriminatorValue,
		"403": CreateEmailsItemEmails403ErrorFromDiscriminatorValue,
		"404": CreateEmailsItemEmails404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves a single email that has been sent
// Deprecated: This method is obsolete. Use GetAsEmailsGetResponse instead.
// returns a EmailsItemEmailsResponseable when successful
// returns a EmailsItemEmails401Error error when the service returns a 401 status code
// returns a EmailsItemEmails403Error error when the service returns a 403 status code
// returns a EmailsItemEmails404Error error when the service returns a 404 status code
func (m *EmailsEmailsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EmailsEmailsItemRequestBuilderGetRequestConfiguration) (EmailsItemEmailsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsItemEmails401ErrorFromDiscriminatorValue,
		"403": CreateEmailsItemEmails403ErrorFromDiscriminatorValue,
		"404": CreateEmailsItemEmails404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsItemEmailsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsItemEmailsResponseable), nil
}

// GetAsEmailsGetResponse retrieves a single email that has been sent
// returns a EmailsItemEmailsGetResponseable when successful
// returns a EmailsItemEmails401Error error when the service returns a 401 status code
// returns a EmailsItemEmails403Error error when the service returns a 403 status code
// returns a EmailsItemEmails404Error error when the service returns a 404 status code
func (m *EmailsEmailsItemRequestBuilder) GetAsEmailsGetResponse(ctx context.Context, requestConfiguration *EmailsEmailsItemRequestBuilderGetRequestConfiguration) (EmailsItemEmailsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailsItemEmails401ErrorFromDiscriminatorValue,
		"403": CreateEmailsItemEmails403ErrorFromDiscriminatorValue,
		"404": CreateEmailsItemEmails404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsItemEmailsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsItemEmailsGetResponseable), nil
}

// ToDeleteRequestInformation delete a specific email record
// returns a *RequestInformation when successful
func (m *EmailsEmailsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EmailsEmailsItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieves a single email that has been sent
// returns a *RequestInformation when successful
func (m *EmailsEmailsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmailsEmailsItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EmailsEmailsItemRequestBuilder when successful
func (m *EmailsEmailsItemRequestBuilder) WithUrl(rawUrl string) *EmailsEmailsItemRequestBuilder {
	return NewEmailsEmailsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
