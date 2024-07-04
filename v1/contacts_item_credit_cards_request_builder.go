package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsItemCreditCardsRequestBuilder builds and executes requests for operations under \v1\contacts\{contactId-id}\creditCards
type ContactsItemCreditCardsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsItemCreditCardsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemCreditCardsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ContactsItemCreditCardsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemCreditCardsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewContactsItemCreditCardsRequestBuilderInternal instantiates a new ContactsItemCreditCardsRequestBuilder and sets the default values.
func NewContactsItemCreditCardsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemCreditCardsRequestBuilder {
	m := &ContactsItemCreditCardsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/contacts/{contactId%2Did}/creditCards", pathParameters),
	}
	return m
}

// NewContactsItemCreditCardsRequestBuilder instantiates a new ContactsItemCreditCardsRequestBuilder and sets the default values.
func NewContactsItemCreditCardsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemCreditCardsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsItemCreditCardsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get list all Credit Cards on a contact
// returns a []ContactsItemCreditCardsable when successful
// returns a ContactsItemCreditCards401Error error when the service returns a 401 status code
// returns a ContactsItemCreditCards403Error error when the service returns a 403 status code
// returns a ContactsItemCreditCards404Error error when the service returns a 404 status code
func (m *ContactsItemCreditCardsRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactsItemCreditCardsRequestBuilderGetRequestConfiguration) ([]ContactsItemCreditCardsable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemCreditCards401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemCreditCards403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemCreditCards404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, CreateContactsItemCreditCardsFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	val := make([]ContactsItemCreditCardsable, len(res))
	for i, v := range res {
		if v != nil {
			val[i] = v.(ContactsItemCreditCardsable)
		}
	}
	return val, nil
}

// Post creates a new credit card associated to a contact
// Deprecated: This method is obsolete. Use PostAsCreditCardsPostResponse instead.
// returns a ContactsItemCreditCardsResponseable when successful
// returns a ContactsItemCreditCards401Error error when the service returns a 401 status code
// returns a ContactsItemCreditCards403Error error when the service returns a 403 status code
func (m *ContactsItemCreditCardsRequestBuilder) Post(ctx context.Context, body ContactsItemCreditCardsPostRequestBodyable, requestConfiguration *ContactsItemCreditCardsRequestBuilderPostRequestConfiguration) (ContactsItemCreditCardsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemCreditCards401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemCreditCards403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemCreditCardsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemCreditCardsResponseable), nil
}

// PostAsCreditCardsPostResponse creates a new credit card associated to a contact
// returns a ContactsItemCreditCardsPostResponseable when successful
// returns a ContactsItemCreditCards401Error error when the service returns a 401 status code
// returns a ContactsItemCreditCards403Error error when the service returns a 403 status code
func (m *ContactsItemCreditCardsRequestBuilder) PostAsCreditCardsPostResponse(ctx context.Context, body ContactsItemCreditCardsPostRequestBodyable, requestConfiguration *ContactsItemCreditCardsRequestBuilderPostRequestConfiguration) (ContactsItemCreditCardsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemCreditCards401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemCreditCards403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemCreditCardsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemCreditCardsPostResponseable), nil
}

// ToGetRequestInformation list all Credit Cards on a contact
// returns a *RequestInformation when successful
func (m *ContactsItemCreditCardsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContactsItemCreditCardsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPostRequestInformation creates a new credit card associated to a contact
// returns a *RequestInformation when successful
func (m *ContactsItemCreditCardsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContactsItemCreditCardsPostRequestBodyable, requestConfiguration *ContactsItemCreditCardsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ContactsItemCreditCardsRequestBuilder when successful
func (m *ContactsItemCreditCardsRequestBuilder) WithUrl(rawUrl string) *ContactsItemCreditCardsRequestBuilder {
	return NewContactsItemCreditCardsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
