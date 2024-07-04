package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EmailAddressesWithEmailItemRequestBuilder builds and executes requests for operations under \v1\emailAddresses\{email}
type EmailAddressesWithEmailItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// EmailAddressesWithEmailItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailAddressesWithEmailItemRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewEmailAddressesWithEmailItemRequestBuilderInternal instantiates a new EmailAddressesWithEmailItemRequestBuilder and sets the default values.
func NewEmailAddressesWithEmailItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailAddressesWithEmailItemRequestBuilder {
	m := &EmailAddressesWithEmailItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/emailAddresses/{email}", pathParameters),
	}
	return m
}

// NewEmailAddressesWithEmailItemRequestBuilder instantiates a new EmailAddressesWithEmailItemRequestBuilder and sets the default values.
func NewEmailAddressesWithEmailItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailAddressesWithEmailItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEmailAddressesWithEmailItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Put replaces all of the values of a given email addressYou may opt-in or mark an email address as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Company gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
// Deprecated: This method is obsolete. Use PutAsWithEmailPutResponse instead.
// returns a EmailAddressesItemWithEmailResponseable when successful
// returns a EmailAddressesItemWithEmail401Error error when the service returns a 401 status code
// returns a EmailAddressesItemWithEmail403Error error when the service returns a 403 status code
// returns a EmailAddressesItemWithEmail404Error error when the service returns a 404 status code
func (m *EmailAddressesWithEmailItemRequestBuilder) Put(ctx context.Context, body EmailAddressesItemWithEmailPutRequestBodyable, requestConfiguration *EmailAddressesWithEmailItemRequestBuilderPutRequestConfiguration) (EmailAddressesItemWithEmailResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailAddressesItemWithEmail401ErrorFromDiscriminatorValue,
		"403": CreateEmailAddressesItemWithEmail403ErrorFromDiscriminatorValue,
		"404": CreateEmailAddressesItemWithEmail404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailAddressesItemWithEmailResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailAddressesItemWithEmailResponseable), nil
}

// PutAsWithEmailPutResponse replaces all of the values of a given email addressYou may opt-in or mark an email address as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Company gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
// returns a EmailAddressesItemWithEmailPutResponseable when successful
// returns a EmailAddressesItemWithEmail401Error error when the service returns a 401 status code
// returns a EmailAddressesItemWithEmail403Error error when the service returns a 403 status code
// returns a EmailAddressesItemWithEmail404Error error when the service returns a 404 status code
func (m *EmailAddressesWithEmailItemRequestBuilder) PutAsWithEmailPutResponse(ctx context.Context, body EmailAddressesItemWithEmailPutRequestBodyable, requestConfiguration *EmailAddressesWithEmailItemRequestBuilderPutRequestConfiguration) (EmailAddressesItemWithEmailPutResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateEmailAddressesItemWithEmail401ErrorFromDiscriminatorValue,
		"403": CreateEmailAddressesItemWithEmail403ErrorFromDiscriminatorValue,
		"404": CreateEmailAddressesItemWithEmail404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmailAddressesItemWithEmailPutResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EmailAddressesItemWithEmailPutResponseable), nil
}

// ToPutRequestInformation replaces all of the values of a given email addressYou may opt-in or mark an email address as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Company gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
// returns a *RequestInformation when successful
func (m *EmailAddressesWithEmailItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body EmailAddressesItemWithEmailPutRequestBodyable, requestConfiguration *EmailAddressesWithEmailItemRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *EmailAddressesWithEmailItemRequestBuilder when successful
func (m *EmailAddressesWithEmailItemRequestBuilder) WithUrl(rawUrl string) *EmailAddressesWithEmailItemRequestBuilder {
	return NewEmailAddressesWithEmailItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
