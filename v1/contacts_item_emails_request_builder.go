package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsItemEmailsRequestBuilder builds and executes requests for operations under \v1\contacts\{contactId-id}\emails
type ContactsItemEmailsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsItemEmailsRequestBuilderGetQueryParameters list Emails that have been sent to a Contact
type ContactsItemEmailsRequestBuilderGetQueryParameters struct {
	// Optional email address to query on
	Email *string `uriparametername:"email"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
}

// ContactsItemEmailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemEmailsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ContactsItemEmailsRequestBuilderGetQueryParameters
}

// ContactsItemEmailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemEmailsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewContactsItemEmailsRequestBuilderInternal instantiates a new ContactsItemEmailsRequestBuilder and sets the default values.
func NewContactsItemEmailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemEmailsRequestBuilder {
	m := &ContactsItemEmailsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/contacts/{contactId%2Did}/emails{?email*,limit*,offset*}", pathParameters),
	}
	return m
}

// NewContactsItemEmailsRequestBuilder instantiates a new ContactsItemEmailsRequestBuilder and sets the default values.
func NewContactsItemEmailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemEmailsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsItemEmailsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get list Emails that have been sent to a Contact
// Deprecated: This method is obsolete. Use GetAsEmailsGetResponse instead.
// returns a ContactsItemEmailsResponseable when successful
// returns a ContactsItemEmails401Error error when the service returns a 401 status code
// returns a ContactsItemEmails403Error error when the service returns a 403 status code
// returns a ContactsItemEmails404Error error when the service returns a 404 status code
func (m *ContactsItemEmailsRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactsItemEmailsRequestBuilderGetRequestConfiguration) (ContactsItemEmailsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemEmails401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemEmails403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemEmails404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemEmailsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemEmailsResponseable), nil
}

// GetAsEmailsGetResponse list Emails that have been sent to a Contact
// returns a ContactsItemEmailsGetResponseable when successful
// returns a ContactsItemEmails401Error error when the service returns a 401 status code
// returns a ContactsItemEmails403Error error when the service returns a 403 status code
// returns a ContactsItemEmails404Error error when the service returns a 404 status code
func (m *ContactsItemEmailsRequestBuilder) GetAsEmailsGetResponse(ctx context.Context, requestConfiguration *ContactsItemEmailsRequestBuilderGetRequestConfiguration) (ContactsItemEmailsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemEmails401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemEmails403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemEmails404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemEmailsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemEmailsGetResponseable), nil
}

// Post create a record of an email sent to a contact
// Deprecated: This method is obsolete. Use PostAsEmailsPostResponse instead.
// returns a ContactsItemEmailsResponseable when successful
// returns a ContactsItemEmails401Error error when the service returns a 401 status code
// returns a ContactsItemEmails403Error error when the service returns a 403 status code
func (m *ContactsItemEmailsRequestBuilder) Post(ctx context.Context, body ContactsItemEmailsPostRequestBodyable, requestConfiguration *ContactsItemEmailsRequestBuilderPostRequestConfiguration) (ContactsItemEmailsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemEmails401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemEmails403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemEmailsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemEmailsResponseable), nil
}

// PostAsEmailsPostResponse create a record of an email sent to a contact
// returns a ContactsItemEmailsPostResponseable when successful
// returns a ContactsItemEmails401Error error when the service returns a 401 status code
// returns a ContactsItemEmails403Error error when the service returns a 403 status code
func (m *ContactsItemEmailsRequestBuilder) PostAsEmailsPostResponse(ctx context.Context, body ContactsItemEmailsPostRequestBodyable, requestConfiguration *ContactsItemEmailsRequestBuilderPostRequestConfiguration) (ContactsItemEmailsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemEmails401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemEmails403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemEmailsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemEmailsPostResponseable), nil
}

// ToGetRequestInformation list Emails that have been sent to a Contact
// returns a *RequestInformation when successful
func (m *ContactsItemEmailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContactsItemEmailsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ContactsItemEmailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContactsItemEmailsPostRequestBodyable, requestConfiguration *ContactsItemEmailsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/contacts/{contactId%2Did}/emails", m.BaseRequestBuilder.PathParameters)
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
// returns a *ContactsItemEmailsRequestBuilder when successful
func (m *ContactsItemEmailsRequestBuilder) WithUrl(rawUrl string) *ContactsItemEmailsRequestBuilder {
	return NewContactsItemEmailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
