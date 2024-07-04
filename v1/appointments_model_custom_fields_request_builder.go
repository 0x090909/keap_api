package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AppointmentsModelCustomFieldsRequestBuilder builds and executes requests for operations under \v1\appointments\model\customFields
type AppointmentsModelCustomFieldsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AppointmentsModelCustomFieldsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppointmentsModelCustomFieldsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewAppointmentsModelCustomFieldsRequestBuilderInternal instantiates a new AppointmentsModelCustomFieldsRequestBuilder and sets the default values.
func NewAppointmentsModelCustomFieldsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AppointmentsModelCustomFieldsRequestBuilder {
	m := &AppointmentsModelCustomFieldsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/appointments/model/customFields", pathParameters),
	}
	return m
}

// NewAppointmentsModelCustomFieldsRequestBuilder instantiates a new AppointmentsModelCustomFieldsRequestBuilder and sets the default values.
func NewAppointmentsModelCustomFieldsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AppointmentsModelCustomFieldsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAppointmentsModelCustomFieldsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post adds a custom field of the specified type and options to the Appointment object.
// Deprecated: This method is obsolete. Use PostAsCustomFieldsPostResponse instead.
// returns a AppointmentsModelCustomFieldsResponseable when successful
// returns a AppointmentsModelCustomFields401Error error when the service returns a 401 status code
// returns a AppointmentsModelCustomFields403Error error when the service returns a 403 status code
func (m *AppointmentsModelCustomFieldsRequestBuilder) Post(ctx context.Context, body AppointmentsModelCustomFieldsPostRequestBodyable, requestConfiguration *AppointmentsModelCustomFieldsRequestBuilderPostRequestConfiguration) (AppointmentsModelCustomFieldsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointmentsModelCustomFields401ErrorFromDiscriminatorValue,
		"403": CreateAppointmentsModelCustomFields403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsModelCustomFieldsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsModelCustomFieldsResponseable), nil
}

// PostAsCustomFieldsPostResponse adds a custom field of the specified type and options to the Appointment object.
// returns a AppointmentsModelCustomFieldsPostResponseable when successful
// returns a AppointmentsModelCustomFields401Error error when the service returns a 401 status code
// returns a AppointmentsModelCustomFields403Error error when the service returns a 403 status code
func (m *AppointmentsModelCustomFieldsRequestBuilder) PostAsCustomFieldsPostResponse(ctx context.Context, body AppointmentsModelCustomFieldsPostRequestBodyable, requestConfiguration *AppointmentsModelCustomFieldsRequestBuilderPostRequestConfiguration) (AppointmentsModelCustomFieldsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointmentsModelCustomFields401ErrorFromDiscriminatorValue,
		"403": CreateAppointmentsModelCustomFields403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsModelCustomFieldsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsModelCustomFieldsPostResponseable), nil
}

// ToPostRequestInformation adds a custom field of the specified type and options to the Appointment object.
// returns a *RequestInformation when successful
func (m *AppointmentsModelCustomFieldsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AppointmentsModelCustomFieldsPostRequestBodyable, requestConfiguration *AppointmentsModelCustomFieldsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AppointmentsModelCustomFieldsRequestBuilder when successful
func (m *AppointmentsModelCustomFieldsRequestBuilder) WithUrl(rawUrl string) *AppointmentsModelCustomFieldsRequestBuilder {
	return NewAppointmentsModelCustomFieldsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
