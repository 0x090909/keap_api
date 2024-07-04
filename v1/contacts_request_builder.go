package v1

import (
	"context"
	i026efb98b2c0dfcf5505f1b9a78191c252664bac134c05e4afafac4dc8880c6c "github.com/0x090909/keap_api/v1/contacts"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// ContactsRequestBuilder builds and executes requests for operations under \v1\contacts
type ContactsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsRequestBuilderGetQueryParameters retrieves a list of all contacts
type ContactsRequestBuilderGetQueryParameters struct {
	// Optional email address to query on
	Email *string `uriparametername:"email"`
	// Optional last name or surname to query on
	Family_name *string `uriparametername:"family_name"`
	// Optional first name or forename to query on
	Given_name *string `uriparametername:"given_name"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
	// Comma-delimited list of Contact properties to include in the response. (Some fields such as `lead_source_id`, `custom_fields`, and `job_title` aren't included, by default.)
	Optional_properties []string `uriparametername:"optional_properties"`
	// Attribute to order items by
	// Deprecated: This property is deprecated, use OrderAsGetOrderQueryParameterType instead
	Order *string `uriparametername:"order"`
	// How to order the data i.e. ascending (A-Z) or descending (Z-A)
	// Deprecated: This property is deprecated, use Order_directionAsGetOrder_directionQueryParameterType instead
	Order_direction *string `uriparametername:"order_direction"`
	// How to order the data i.e. ascending (A-Z) or descending (Z-A)
	Order_directionAsGetOrder_directionQueryParameterType *i026efb98b2c0dfcf5505f1b9a78191c252664bac134c05e4afafac4dc8880c6c.GetOrder_directionQueryParameterType `uriparametername:"order_direction"`
	// Attribute to order items by
	OrderAsGetOrderQueryParameterType *i026efb98b2c0dfcf5505f1b9a78191c252664bac134c05e4afafac4dc8880c6c.GetOrderQueryParameterType `uriparametername:"order"`
	// Date to start searching from on LastUpdated ex. `2017-01-01T22:17:59.039Z`
	Since *string `uriparametername:"since"`
	// Date to search to on LastUpdated ex. `2017-01-01T22:17:59.039Z`
	Until *string `uriparametername:"until"`
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

// ContactsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByContactIdId gets an item from the github.com/0x090909/keap_api.v1.contacts.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ContactsContactIdItemRequestBuilder when successful
func (m *ContactsRequestBuilder) ByContactIdId(contactIdId string) *ContactsContactIdItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if contactIdId != "" {
		urlTplParams["contactId%2Did"] = contactIdId
	}
	return NewContactsContactIdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByContactIdIdInt64 gets an item from the github.com/0x090909/keap_api.v1.contacts.item collection
// returns a *ContactsContactIdItemRequestBuilder when successful
func (m *ContactsRequestBuilder) ByContactIdIdInt64(contactIdId int64) *ContactsContactIdItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["contactId%2Did"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(contactIdId, 10)
	return NewContactsContactIdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewContactsRequestBuilderInternal instantiates a new ContactsRequestBuilder and sets the default values.
func NewContactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsRequestBuilder {
	m := &ContactsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/contacts{?email*,family_name*,given_name*,limit*,offset*,optional_properties*,order*,order_direction*,since*,until*}", pathParameters),
	}
	return m
}

// NewContactsRequestBuilder instantiates a new ContactsRequestBuilder and sets the default values.
func NewContactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all contacts
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

// GetAsContactsGetResponse retrieves a list of all contacts
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

// Model the model property
// returns a *ContactsModelRequestBuilder when successful
func (m *ContactsRequestBuilder) Model() *ContactsModelRequestBuilder {
	return NewContactsModelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Post creates a new contact as the authenticated user. NB: Contact must contain at least one item in `email_addresses` or `phone_numbers` and `country_code` is required if `region` is specified.Please see the body schema for updates to the postal code field.You may opt-in or mark a Contact as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Contact gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.This API only supports a subset of valid timezones. A list of the supported timezones can be found [here](https://developer.infusionsoft.com/faqs/what-timezones-do-contact-calls-accept/).
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

// PostAsContactsPostResponse creates a new contact as the authenticated user. NB: Contact must contain at least one item in `email_addresses` or `phone_numbers` and `country_code` is required if `region` is specified.Please see the body schema for updates to the postal code field.You may opt-in or mark a Contact as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Contact gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.This API only supports a subset of valid timezones. A list of the supported timezones can be found [here](https://developer.infusionsoft.com/faqs/what-timezones-do-contact-calls-accept/).
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

// Put creates a new contact or updates a contact as the authenticated user. NB: New Contacts must contain at least one item in `email_addresses` or `phone_numbers` and `country_code` is required if `region` is specified. Existing Contacts are updated with only the values provided in the request. Accepts a `duplicate_option` which performs duplicate checking by one of the following options: `Email`, `EmailAndName`, if a match is found using the option provided, the existing contact will be updated. If an existing contact was not found using the `duplicate_option` provided, a new contact record will be created.You may opt-in or mark a Contact as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Contact gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.This API only supports a subset of valid timezones. A list of the supported timezones can be found [here](https://developer.infusionsoft.com/faqs/what-timezones-do-contact-calls-accept/).
// Deprecated: This method is obsolete. Use PutAsContactsPutResponse instead.
// returns a ContactsResponseable when successful
// returns a Contacts401Error error when the service returns a 401 status code
// returns a Contacts403Error error when the service returns a 403 status code
// returns a Contacts404Error error when the service returns a 404 status code
func (m *ContactsRequestBuilder) Put(ctx context.Context, body ContactsPutRequestBodyable, requestConfiguration *ContactsRequestBuilderPutRequestConfiguration) (ContactsResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
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

// PutAsContactsPutResponse creates a new contact or updates a contact as the authenticated user. NB: New Contacts must contain at least one item in `email_addresses` or `phone_numbers` and `country_code` is required if `region` is specified. Existing Contacts are updated with only the values provided in the request. Accepts a `duplicate_option` which performs duplicate checking by one of the following options: `Email`, `EmailAndName`, if a match is found using the option provided, the existing contact will be updated. If an existing contact was not found using the `duplicate_option` provided, a new contact record will be created.You may opt-in or mark a Contact as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Contact gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.This API only supports a subset of valid timezones. A list of the supported timezones can be found [here](https://developer.infusionsoft.com/faqs/what-timezones-do-contact-calls-accept/).
// returns a ContactsPutResponseable when successful
// returns a Contacts401Error error when the service returns a 401 status code
// returns a Contacts403Error error when the service returns a 403 status code
// returns a Contacts404Error error when the service returns a 404 status code
func (m *ContactsRequestBuilder) PutAsContactsPutResponse(ctx context.Context, body ContactsPutRequestBodyable, requestConfiguration *ContactsRequestBuilderPutRequestConfiguration) (ContactsPutResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContacts401ErrorFromDiscriminatorValue,
		"403": CreateContacts403ErrorFromDiscriminatorValue,
		"404": CreateContacts404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsPutResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsPutResponseable), nil
}

// ToGetRequestInformation retrieves a list of all contacts
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

// ToPostRequestInformation creates a new contact as the authenticated user. NB: Contact must contain at least one item in `email_addresses` or `phone_numbers` and `country_code` is required if `region` is specified.Please see the body schema for updates to the postal code field.You may opt-in or mark a Contact as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Contact gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.This API only supports a subset of valid timezones. A list of the supported timezones can be found [here](https://developer.infusionsoft.com/faqs/what-timezones-do-contact-calls-accept/).
// returns a *RequestInformation when successful
func (m *ContactsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContactsPostRequestBodyable, requestConfiguration *ContactsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/contacts", m.BaseRequestBuilder.PathParameters)
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

// ToPutRequestInformation creates a new contact or updates a contact as the authenticated user. NB: New Contacts must contain at least one item in `email_addresses` or `phone_numbers` and `country_code` is required if `region` is specified. Existing Contacts are updated with only the values provided in the request. Accepts a `duplicate_option` which performs duplicate checking by one of the following options: `Email`, `EmailAndName`, if a match is found using the option provided, the existing contact will be updated. If an existing contact was not found using the `duplicate_option` provided, a new contact record will be created.You may opt-in or mark a Contact as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Contact gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.This API only supports a subset of valid timezones. A list of the supported timezones can be found [here](https://developer.infusionsoft.com/faqs/what-timezones-do-contact-calls-accept/).
// returns a *RequestInformation when successful
func (m *ContactsRequestBuilder) ToPutRequestInformation(ctx context.Context, body ContactsPutRequestBodyable, requestConfiguration *ContactsRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, "{+baseurl}/v1/contacts", m.BaseRequestBuilder.PathParameters)
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
