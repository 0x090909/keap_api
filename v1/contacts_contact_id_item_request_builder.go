package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsContactIdItemRequestBuilder builds and executes requests for operations under \v1\contacts\{contactId-id}
type ContactsContactIdItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsContactIdItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsContactIdItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ContactsContactIdItemRequestBuilderGetQueryParameters retrieves a single contact
type ContactsContactIdItemRequestBuilderGetQueryParameters struct {
	// Comma-delimited list of Contact properties to include in the response. (Some fields such as `lead_source_id`, `custom_fields`, and `job_title` aren't included, by default.)
	Optional_properties []string `uriparametername:"optional_properties"`
}

// ContactsContactIdItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsContactIdItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ContactsContactIdItemRequestBuilderGetQueryParameters
}

// ContactsContactIdItemRequestBuilderPatchQueryParameters updates a contact with only the values provided in the request.You may opt-in or mark a Contact as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Contact gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
type ContactsContactIdItemRequestBuilderPatchQueryParameters struct {
	// An optional list of properties to be updated. If set, only the provided properties will be updated and others will be skipped.
	Update_mask []string `uriparametername:"update_mask"`
}

// ContactsContactIdItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsContactIdItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ContactsContactIdItemRequestBuilderPatchQueryParameters
}

// NewContactsContactIdItemRequestBuilderInternal instantiates a new ContactsContactIdItemRequestBuilder and sets the default values.
func NewContactsContactIdItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsContactIdItemRequestBuilder {
	m := &ContactsContactIdItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/contacts/{contactId%2Did}{?optional_properties*}", pathParameters),
	}
	return m
}

// NewContactsContactIdItemRequestBuilder instantiates a new ContactsContactIdItemRequestBuilder and sets the default values.
func NewContactsContactIdItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsContactIdItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsContactIdItemRequestBuilderInternal(urlParams, requestAdapter)
}

// CreditCards the creditCards property
// returns a *ContactsItemCreditCardsRequestBuilder when successful
func (m *ContactsContactIdItemRequestBuilder) CreditCards() *ContactsItemCreditCardsRequestBuilder {
	return NewContactsItemCreditCardsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Delete deletes the specified contact.
// returns a ContactsItemContactId401Error error when the service returns a 401 status code
// returns a ContactsItemContactId403Error error when the service returns a 403 status code
// returns a ContactsItemContactId404Error error when the service returns a 404 status code
func (m *ContactsContactIdItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ContactsContactIdItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemContactId401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemContactId403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemContactId404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Emails the emails property
// returns a *ContactsItemEmailsRequestBuilder when successful
func (m *ContactsContactIdItemRequestBuilder) Emails() *ContactsItemEmailsRequestBuilder {
	return NewContactsItemEmailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get retrieves a single contact
// Deprecated: This method is obsolete. Use GetAsContactIdGetResponse instead.
// returns a ContactsItemContactIdResponseable when successful
// returns a ContactsItemContactId401Error error when the service returns a 401 status code
// returns a ContactsItemContactId403Error error when the service returns a 403 status code
// returns a ContactsItemContactId404Error error when the service returns a 404 status code
func (m *ContactsContactIdItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactsContactIdItemRequestBuilderGetRequestConfiguration) (ContactsItemContactIdResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemContactId401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemContactId403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemContactId404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemContactIdResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemContactIdResponseable), nil
}

// GetAsContactIdGetResponse retrieves a single contact
// returns a ContactsItemContactIdGetResponseable when successful
// returns a ContactsItemContactId401Error error when the service returns a 401 status code
// returns a ContactsItemContactId403Error error when the service returns a 403 status code
// returns a ContactsItemContactId404Error error when the service returns a 404 status code
func (m *ContactsContactIdItemRequestBuilder) GetAsContactIdGetResponse(ctx context.Context, requestConfiguration *ContactsContactIdItemRequestBuilderGetRequestConfiguration) (ContactsItemContactIdGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemContactId401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemContactId403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemContactId404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemContactIdGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemContactIdGetResponseable), nil
}

// Patch updates a contact with only the values provided in the request.You may opt-in or mark a Contact as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Contact gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
// Deprecated: This method is obsolete. Use PatchAsContactIdPatchResponse instead.
// returns a ContactsItemContactIdResponseable when successful
// returns a ContactsItemContactId401Error error when the service returns a 401 status code
// returns a ContactsItemContactId403Error error when the service returns a 403 status code
// returns a ContactsItemContactId404Error error when the service returns a 404 status code
func (m *ContactsContactIdItemRequestBuilder) Patch(ctx context.Context, body ContactsItemContactIdPatchRequestBodyable, requestConfiguration *ContactsContactIdItemRequestBuilderPatchRequestConfiguration) (ContactsItemContactIdResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemContactId401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemContactId403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemContactId404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemContactIdResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemContactIdResponseable), nil
}

// PatchAsContactIdPatchResponse updates a contact with only the values provided in the request.You may opt-in or mark a Contact as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Contact gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
// returns a ContactsItemContactIdPatchResponseable when successful
// returns a ContactsItemContactId401Error error when the service returns a 401 status code
// returns a ContactsItemContactId403Error error when the service returns a 403 status code
// returns a ContactsItemContactId404Error error when the service returns a 404 status code
func (m *ContactsContactIdItemRequestBuilder) PatchAsContactIdPatchResponse(ctx context.Context, body ContactsItemContactIdPatchRequestBodyable, requestConfiguration *ContactsContactIdItemRequestBuilderPatchRequestConfiguration) (ContactsItemContactIdPatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemContactId401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemContactId403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemContactId404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemContactIdPatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemContactIdPatchResponseable), nil
}

// Tags the tags property
// returns a *ContactsItemTagsRequestBuilder when successful
func (m *ContactsContactIdItemRequestBuilder) Tags() *ContactsItemTagsRequestBuilder {
	return NewContactsItemTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToDeleteRequestInformation deletes the specified contact.
// returns a *RequestInformation when successful
func (m *ContactsContactIdItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ContactsContactIdItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/v1/contacts/{contactId%2Did}", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieves a single contact
// returns a *RequestInformation when successful
func (m *ContactsContactIdItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContactsContactIdItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation updates a contact with only the values provided in the request.You may opt-in or mark a Contact as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Contact gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
// returns a *RequestInformation when successful
func (m *ContactsContactIdItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ContactsItemContactIdPatchRequestBodyable, requestConfiguration *ContactsContactIdItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/v1/contacts/{contactId%2Did}{?update_mask*}", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
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

// Utm the utm property
// returns a *ContactsItemUtmRequestBuilder when successful
func (m *ContactsContactIdItemRequestBuilder) Utm() *ContactsItemUtmRequestBuilder {
	return NewContactsItemUtmRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ContactsContactIdItemRequestBuilder when successful
func (m *ContactsContactIdItemRequestBuilder) WithUrl(rawUrl string) *ContactsContactIdItemRequestBuilder {
	return NewContactsContactIdItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
