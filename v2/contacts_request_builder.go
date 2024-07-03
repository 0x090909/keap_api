package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsRequestBuilder builds and executes requests for operations under \v2\contacts
type ContactsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsRequestBuilderGetQueryParameters retrieves a list of Contacts
type ContactsRequestBuilderGetQueryParameters struct {
	// Filter to apply, allowed fields are:  - (String) email  - (String) given_name  - (String) family_name  - (String) company_id  - (Set[String]) contact_ids  - (String) start_update_time  - (String) end_update_time You will need to apply the `==` operator to check the equality of one of the filters with your searched word, in the encoded form `%3D%3D`. For the filters listed above, here are some examples:  - `filter=given_name%3D%3DMary` - `filter=company_id%3D%3D123` - `filter=company_id%3D%3D123;family_name=Smith`
	Filter *string `uriparametername:"filter"`
	// Attribute and direction to order items.  One of the following fields:  - id  - date_created  - email  One of the following directions:  - asc  - desc
	Order_by *string `uriparametername:"order_by"`
	// Total number of items to return per page
	Page_size *int32 `uriparametername:"page_size"`
	// Page token
	Page_token *string `uriparametername:"page_token"`
}

// ContactsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ContactsRequestBuilderGetQueryParameters
}

// ContactsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByContact_id gets an item from the github.com/0x090909/keap_api.v2.contacts.item collection
// returns a *ContactsWithContact_ItemRequestBuilder when successful
func (m *ContactsRequestBuilder) ByContact_id(contact_id string) *ContactsWithContact_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if contact_id != "" {
		urlTplParams["contact_id"] = contact_id
	}
	return NewContactsWithContact_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewContactsRequestBuilderInternal instantiates a new ContactsRequestBuilder and sets the default values.
func NewContactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsRequestBuilder {
	m := &ContactsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/contacts{?filter*,order_by*,page_size*,page_token*}", pathParameters),
	}
	return m
}

// NewContactsRequestBuilder instantiates a new ContactsRequestBuilder and sets the default values.
func NewContactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of Contacts
// Deprecated: This method is obsolete. Use GetAsContactsGetResponse instead.
// returns a ContactsResponseable when successful
// returns a Contacts401Error error when the service returns a 401 status code
// returns a Contacts403Error error when the service returns a 403 status code
// returns a Contacts404Error error when the service returns a 404 status code
func (m *ContactsRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactsRequestBuilderGetRequestConfiguration) (ContactsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContacts401ErrorFromDiscriminatorValue,
		"403": CreateContacts403ErrorFromDiscriminatorValue,
		"404": CreateContacts404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsResponseable), nil
}

// GetAsContactsGetResponse retrieves a list of Contacts
// returns a ContactsGetResponseable when successful
// returns a Contacts401Error error when the service returns a 401 status code
// returns a Contacts403Error error when the service returns a 403 status code
// returns a Contacts404Error error when the service returns a 404 status code
func (m *ContactsRequestBuilder) GetAsContactsGetResponse(ctx context.Context, requestConfiguration *ContactsRequestBuilderGetRequestConfiguration) (ContactsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContacts401ErrorFromDiscriminatorValue,
		"403": CreateContacts403ErrorFromDiscriminatorValue,
		"404": CreateContacts404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsGetResponseable), nil
}

// Links the links property
// returns a *ContactsLinksRequestBuilder when successful
func (m *ContactsRequestBuilder) Links() *ContactsLinksRequestBuilder {
	return NewContactsLinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Model the model property
// returns a *ContactsModelRequestBuilder when successful
func (m *ContactsRequestBuilder) Model() *ContactsModelRequestBuilder {
	return NewContactsModelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Post creates a new Contact.*Note:* Contact must contain at least one item in `email_addresses` or `phone_numbers` and `country_code` is required if `region` is specified.
// Deprecated: This method is obsolete. Use PostAsContactsPostResponse instead.
// returns a ContactsResponseable when successful
// returns a Contacts401Error error when the service returns a 401 status code
// returns a Contacts403Error error when the service returns a 403 status code
func (m *ContactsRequestBuilder) Post(ctx context.Context, body ContactsPostRequestBodyable, requestConfiguration *ContactsRequestBuilderPostRequestConfiguration) (ContactsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContacts401ErrorFromDiscriminatorValue,
		"403": CreateContacts403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsResponseable), nil
}

// PostAsContactsPostResponse creates a new Contact.*Note:* Contact must contain at least one item in `email_addresses` or `phone_numbers` and `country_code` is required if `region` is specified.
// returns a ContactsPostResponseable when successful
// returns a Contacts401Error error when the service returns a 401 status code
// returns a Contacts403Error error when the service returns a 403 status code
func (m *ContactsRequestBuilder) PostAsContactsPostResponse(ctx context.Context, body ContactsPostRequestBodyable, requestConfiguration *ContactsRequestBuilderPostRequestConfiguration) (ContactsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContacts401ErrorFromDiscriminatorValue,
		"403": CreateContacts403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsPostResponseable), nil
}

// ToGetRequestInformation retrieves a list of Contacts
// returns a *RequestInformation when successful
func (m *ContactsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContactsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation creates a new Contact.*Note:* Contact must contain at least one item in `email_addresses` or `phone_numbers` and `country_code` is required if `region` is specified.
// returns a *RequestInformation when successful
func (m *ContactsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContactsPostRequestBodyable, requestConfiguration *ContactsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v2/contacts", m.BaseRequestBuilder.PathParameters)
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
// returns a *ContactsRequestBuilder when successful
func (m *ContactsRequestBuilder) WithUrl(rawUrl string) *ContactsRequestBuilder {
	return NewContactsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
