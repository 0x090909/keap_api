package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsWithContact_ItemRequestBuilder builds and executes requests for operations under \v2\contacts\{contact_id}
type ContactsWithContact_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsWithContact_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsWithContact_ItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ContactsWithContact_ItemRequestBuilderGetQueryParameters retrieves a single Contact
type ContactsWithContact_ItemRequestBuilderGetQueryParameters struct {
	// Comma-delimited list of Contact properties to include in the response. (Available fields are: `score_value`, `addresses`, `anniversary`, `birthday`, `company`, `contact_type`, `custom_fields`, `create_time`, `email_addresses`, `fax_numbers`, `job_title`, `update_time`, `create_time`, `leadsource_id`,`middle_name`, `origin`, `owner_id`, `phone_numbers`, `preferred_locale`, `preferred_name`,`prefix`, `relationships`, `social_accounts`, `source_type`, `spouse_name`, `suffix`, `time_zone`,`website`, `tag_ids`, `utm_parameters`)
	Fields []string `uriparametername:"fields"`
}

// ContactsWithContact_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsWithContact_ItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ContactsWithContact_ItemRequestBuilderGetQueryParameters
}

// ContactsWithContact_ItemRequestBuilderPatchQueryParameters updates a Contact
type ContactsWithContact_ItemRequestBuilderPatchQueryParameters struct {
	// An optional list of properties to be updated. If set, only the provided properties will be updated and others will be skipped.
	Update_mask []string `uriparametername:"update_mask"`
}

// ContactsWithContact_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsWithContact_ItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ContactsWithContact_ItemRequestBuilderPatchQueryParameters
}

// NewContactsWithContact_ItemRequestBuilderInternal instantiates a new ContactsWithContact_ItemRequestBuilder and sets the default values.
func NewContactsWithContact_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsWithContact_ItemRequestBuilder {
	m := &ContactsWithContact_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/contacts/{contact_id}{?fields*}", pathParameters),
	}
	return m
}

// NewContactsWithContact_ItemRequestBuilder instantiates a new ContactsWithContact_ItemRequestBuilder and sets the default values.
func NewContactsWithContact_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsWithContact_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsWithContact_ItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes the specified Contact.
// returns a ContactsItemWithContact_401Error error when the service returns a 401 status code
// returns a ContactsItemWithContact_403Error error when the service returns a 403 status code
// returns a ContactsItemWithContact_404Error error when the service returns a 404 status code
func (m *ContactsWithContact_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ContactsWithContact_ItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemWithContact_401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemWithContact_403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemWithContact_404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves a single Contact
// Deprecated: This method is obsolete. Use GetAsWithContact_GetResponse instead.
// returns a ContactsItemWithContact_Responseable when successful
// returns a ContactsItemWithContact_401Error error when the service returns a 401 status code
// returns a ContactsItemWithContact_403Error error when the service returns a 403 status code
// returns a ContactsItemWithContact_404Error error when the service returns a 404 status code
func (m *ContactsWithContact_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactsWithContact_ItemRequestBuilderGetRequestConfiguration) (ContactsItemWithContact_Responseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemWithContact_401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemWithContact_403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemWithContact_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemWithContact_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemWithContact_Responseable), nil
}

// GetAsWithContact_GetResponse retrieves a single Contact
// returns a ContactsItemWithContact_GetResponseable when successful
// returns a ContactsItemWithContact_401Error error when the service returns a 401 status code
// returns a ContactsItemWithContact_403Error error when the service returns a 403 status code
// returns a ContactsItemWithContact_404Error error when the service returns a 404 status code
func (m *ContactsWithContact_ItemRequestBuilder) GetAsWithContact_GetResponse(ctx context.Context, requestConfiguration *ContactsWithContact_ItemRequestBuilderGetRequestConfiguration) (ContactsItemWithContact_GetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemWithContact_401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemWithContact_403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemWithContact_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemWithContact_GetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemWithContact_GetResponseable), nil
}

// Links the links property
// returns a *ContactsItemLinksRequestBuilder when successful
func (m *ContactsWithContact_ItemRequestBuilder) Links() *ContactsItemLinksRequestBuilder {
	return NewContactsItemLinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Notes the notes property
// returns a *ContactsItemNotesRequestBuilder when successful
func (m *ContactsWithContact_ItemRequestBuilder) Notes() *ContactsItemNotesRequestBuilder {
	return NewContactsItemNotesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Patch updates a Contact
// Deprecated: This method is obsolete. Use PatchAsWithContact_PatchResponse instead.
// returns a ContactsItemWithContact_Responseable when successful
// returns a ContactsItemWithContact_401Error error when the service returns a 401 status code
// returns a ContactsItemWithContact_403Error error when the service returns a 403 status code
// returns a ContactsItemWithContact_404Error error when the service returns a 404 status code
func (m *ContactsWithContact_ItemRequestBuilder) Patch(ctx context.Context, body ContactsItemWithContact_PatchRequestBodyable, requestConfiguration *ContactsWithContact_ItemRequestBuilderPatchRequestConfiguration) (ContactsItemWithContact_Responseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemWithContact_401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemWithContact_403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemWithContact_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemWithContact_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemWithContact_Responseable), nil
}

// PatchAsWithContact_PatchResponse updates a Contact
// returns a ContactsItemWithContact_PatchResponseable when successful
// returns a ContactsItemWithContact_401Error error when the service returns a 401 status code
// returns a ContactsItemWithContact_403Error error when the service returns a 403 status code
// returns a ContactsItemWithContact_404Error error when the service returns a 404 status code
func (m *ContactsWithContact_ItemRequestBuilder) PatchAsWithContact_PatchResponse(ctx context.Context, body ContactsItemWithContact_PatchRequestBodyable, requestConfiguration *ContactsWithContact_ItemRequestBuilderPatchRequestConfiguration) (ContactsItemWithContact_PatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemWithContact_401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemWithContact_403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemWithContact_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemWithContact_PatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemWithContact_PatchResponseable), nil
}

// PaymentMethods the paymentMethods property
// returns a *ContactsItemPaymentMethodsRequestBuilder when successful
func (m *ContactsWithContact_ItemRequestBuilder) PaymentMethods() *ContactsItemPaymentMethodsRequestBuilder {
	return NewContactsItemPaymentMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToDeleteRequestInformation deletes the specified Contact.
// returns a *RequestInformation when successful
func (m *ContactsWithContact_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ContactsWithContact_ItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/v2/contacts/{contact_id}", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieves a single Contact
// returns a *RequestInformation when successful
func (m *ContactsWithContact_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContactsWithContact_ItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation updates a Contact
// returns a *RequestInformation when successful
func (m *ContactsWithContact_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ContactsItemWithContact_PatchRequestBodyable, requestConfiguration *ContactsWithContact_ItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/v2/contacts/{contact_id}{?update_mask*}", m.BaseRequestBuilder.PathParameters)
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

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ContactsWithContact_ItemRequestBuilder when successful
func (m *ContactsWithContact_ItemRequestBuilder) WithUrl(rawUrl string) *ContactsWithContact_ItemRequestBuilder {
	return NewContactsWithContact_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
