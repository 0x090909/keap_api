package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AppointmentsModelRequestBuilder builds and executes requests for operations under \v1\appointments\model
type AppointmentsModelRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AppointmentsModelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppointmentsModelRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewAppointmentsModelRequestBuilderInternal instantiates a new AppointmentsModelRequestBuilder and sets the default values.
func NewAppointmentsModelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AppointmentsModelRequestBuilder {
	m := &AppointmentsModelRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/appointments/model", pathParameters),
	}
	return m
}

// NewAppointmentsModelRequestBuilder instantiates a new AppointmentsModelRequestBuilder and sets the default values.
func NewAppointmentsModelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AppointmentsModelRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAppointmentsModelRequestBuilderInternal(urlParams, requestAdapter)
}

// CustomFields the customFields property
// returns a *AppointmentsModelCustomFieldsRequestBuilder when successful
func (m *AppointmentsModelRequestBuilder) CustomFields() *AppointmentsModelCustomFieldsRequestBuilder {
	return NewAppointmentsModelCustomFieldsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get get the custom fields for the Appointment object
// Deprecated: This method is obsolete. Use GetAsModelGetResponse instead.
// returns a AppointmentsModelResponseable when successful
// returns a AppointmentsModel401Error error when the service returns a 401 status code
// returns a AppointmentsModel403Error error when the service returns a 403 status code
// returns a AppointmentsModel404Error error when the service returns a 404 status code
func (m *AppointmentsModelRequestBuilder) Get(ctx context.Context, requestConfiguration *AppointmentsModelRequestBuilderGetRequestConfiguration) (AppointmentsModelResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointmentsModel401ErrorFromDiscriminatorValue,
		"403": CreateAppointmentsModel403ErrorFromDiscriminatorValue,
		"404": CreateAppointmentsModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsModelResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsModelResponseable), nil
}

// GetAsModelGetResponse get the custom fields for the Appointment object
// returns a AppointmentsModelGetResponseable when successful
// returns a AppointmentsModel401Error error when the service returns a 401 status code
// returns a AppointmentsModel403Error error when the service returns a 403 status code
// returns a AppointmentsModel404Error error when the service returns a 404 status code
func (m *AppointmentsModelRequestBuilder) GetAsModelGetResponse(ctx context.Context, requestConfiguration *AppointmentsModelRequestBuilderGetRequestConfiguration) (AppointmentsModelGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointmentsModel401ErrorFromDiscriminatorValue,
		"403": CreateAppointmentsModel403ErrorFromDiscriminatorValue,
		"404": CreateAppointmentsModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsModelGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsModelGetResponseable), nil
}

// ToGetRequestInformation get the custom fields for the Appointment object
// returns a *RequestInformation when successful
func (m *AppointmentsModelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppointmentsModelRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AppointmentsModelRequestBuilder when successful
func (m *AppointmentsModelRequestBuilder) WithUrl(rawUrl string) *AppointmentsModelRequestBuilder {
	return NewAppointmentsModelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
