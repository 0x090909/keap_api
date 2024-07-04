package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// EmailsRequestBuilder builds and executes requests for operations under \v1\emails
type EmailsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// EmailsRequestBuilderGetQueryParameters retrieve a list of emails that have been sentKeap is currently investigating an issue with degraded performanceof this endpoint with very large (millions of records) record sets
type EmailsRequestBuilderGetQueryParameters struct {
	// Optional Contact Id to find Emails for
	Contact_id *int64 `uriparametername:"contact_id"`
	// Optional email address to query on
	Email *string `uriparametername:"email"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
	// Optional boolean to turn off ORDER BY in SQL query
	Ordered *bool `uriparametername:"ordered"`
	// Optional date to query on, emails sent since the provided date, must be present if untilDate is provided
	Since_sent_date *string `uriparametername:"since_sent_date"`
	// Optional date to query on, email sent until the provided date
	Until_sent_date *string `uriparametername:"until_sent_date"`
}

// EmailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *EmailsRequestBuilderGetQueryParameters
}

// EmailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ById gets an item from the github.com/0x090909/keap_api.v1.emails.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *EmailsEmailsItemRequestBuilder when successful
func (m *EmailsRequestBuilder) ById(id string) *EmailsEmailsItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["id"] = id
	}
	return NewEmailsEmailsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByIdInt64 gets an item from the github.com/0x090909/keap_api.v1.emails.item collection
// returns a *EmailsEmailsItemRequestBuilder when successful
func (m *EmailsRequestBuilder) ByIdInt64(id int64) *EmailsEmailsItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(id, 10)
	return NewEmailsEmailsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewEmailsRequestBuilderInternal instantiates a new EmailsRequestBuilder and sets the default values.
func NewEmailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsRequestBuilder {
	m := &EmailsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/emails{?contact_id*,email*,limit*,offset*,ordered*,since_sent_date*,until_sent_date*}", pathParameters),
	}
	return m
}

// NewEmailsRequestBuilder instantiates a new EmailsRequestBuilder and sets the default values.
func NewEmailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEmailsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieve a list of emails that have been sentKeap is currently investigating an issue with degraded performanceof this endpoint with very large (millions of records) record sets
// Deprecated: This method is obsolete. Use GetAsEmailsGetResponse instead.
// returns a EmailsResponseable when successful
// returns a Emails401Error error when the service returns a 401 status code
// returns a Emails403Error error when the service returns a 403 status code
// returns a Emails404Error error when the service returns a 404 status code
func (m *EmailsRequestBuilder) Get(ctx context.Context, requestConfiguration *EmailsRequestBuilderGetRequestConfiguration) (EmailsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmails401ErrorFromDiscriminatorValue,
		"403": CreateEmails403ErrorFromDiscriminatorValue,
		"404": CreateEmails404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsResponseable), nil
}

// GetAsEmailsGetResponse retrieve a list of emails that have been sentKeap is currently investigating an issue with degraded performanceof this endpoint with very large (millions of records) record sets
// returns a EmailsGetResponseable when successful
// returns a Emails401Error error when the service returns a 401 status code
// returns a Emails403Error error when the service returns a 403 status code
// returns a Emails404Error error when the service returns a 404 status code
func (m *EmailsRequestBuilder) GetAsEmailsGetResponse(ctx context.Context, requestConfiguration *EmailsRequestBuilderGetRequestConfiguration) (EmailsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmails401ErrorFromDiscriminatorValue,
		"403": CreateEmails403ErrorFromDiscriminatorValue,
		"404": CreateEmails404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsGetResponseable), nil
}

// Post create a record of an email sent to a contact
// Deprecated: This method is obsolete. Use PostAsEmailsPostResponse instead.
// returns a EmailsResponseable when successful
// returns a Emails401Error error when the service returns a 401 status code
// returns a Emails403Error error when the service returns a 403 status code
func (m *EmailsRequestBuilder) Post(ctx context.Context, body EmailsPostRequestBodyable, requestConfiguration *EmailsRequestBuilderPostRequestConfiguration) (EmailsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmails401ErrorFromDiscriminatorValue,
		"403": CreateEmails403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsResponseable), nil
}

// PostAsEmailsPostResponse create a record of an email sent to a contact
// returns a EmailsPostResponseable when successful
// returns a Emails401Error error when the service returns a 401 status code
// returns a Emails403Error error when the service returns a 403 status code
func (m *EmailsRequestBuilder) PostAsEmailsPostResponse(ctx context.Context, body EmailsPostRequestBodyable, requestConfiguration *EmailsRequestBuilderPostRequestConfiguration) (EmailsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmails401ErrorFromDiscriminatorValue,
		"403": CreateEmails403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailsPostResponseable), nil
}

// Queue the queue property
// returns a *EmailsQueueRequestBuilder when successful
func (m *EmailsRequestBuilder) Queue() *EmailsQueueRequestBuilder {
	return NewEmailsQueueRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Sync the sync property
// returns a *EmailsSyncRequestBuilder when successful
func (m *EmailsRequestBuilder) Sync() *EmailsSyncRequestBuilder {
	return NewEmailsSyncRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToGetRequestInformation retrieve a list of emails that have been sentKeap is currently investigating an issue with degraded performanceof this endpoint with very large (millions of records) record sets
// returns a *RequestInformation when successful
func (m *EmailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmailsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation create a record of an email sent to a contact
// returns a *RequestInformation when successful
func (m *EmailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body EmailsPostRequestBodyable, requestConfiguration *EmailsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/emails", m.BaseRequestBuilder.PathParameters)
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

// Unsync the unsync property
// returns a *EmailsUnsyncRequestBuilder when successful
func (m *EmailsRequestBuilder) Unsync() *EmailsUnsyncRequestBuilder {
	return NewEmailsUnsyncRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EmailsRequestBuilder when successful
func (m *EmailsRequestBuilder) WithUrl(rawUrl string) *EmailsRequestBuilder {
	return NewEmailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
