package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// AppointmentsRequestBuilder builds and executes requests for operations under \v1\appointments
type AppointmentsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AppointmentsRequestBuilderGetQueryParameters retrieves all appointments for the authenticated user
type AppointmentsRequestBuilderGetQueryParameters struct {
	// Optionally find appointments with a contact
	Contact_id *int64 `uriparametername:"contact_id"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
	// Date to start searching from ex. `2017-01-01T22:17:59.039Z`
	Since *string `uriparametername:"since"`
	// Date to search to ex. `2017-01-01T22:17:59.039Z`
	Until *string `uriparametername:"until"`
}

// AppointmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppointmentsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AppointmentsRequestBuilderGetQueryParameters
}

// AppointmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppointmentsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByAppointmentId gets an item from the github.com/0x090909/keap_api.v1.appointments.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *AppointmentsWithAppointmentItemRequestBuilder when successful
func (m *AppointmentsRequestBuilder) ByAppointmentId(appointmentId string) *AppointmentsWithAppointmentItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if appointmentId != "" {
		urlTplParams["appointmentId"] = appointmentId
	}
	return NewAppointmentsWithAppointmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByAppointmentIdInt64 gets an item from the github.com/0x090909/keap_api.v1.appointments.item collection
// returns a *AppointmentsWithAppointmentItemRequestBuilder when successful
func (m *AppointmentsRequestBuilder) ByAppointmentIdInt64(appointmentId int64) *AppointmentsWithAppointmentItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["appointmentId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(appointmentId, 10)
	return NewAppointmentsWithAppointmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewAppointmentsRequestBuilderInternal instantiates a new AppointmentsRequestBuilder and sets the default values.
func NewAppointmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AppointmentsRequestBuilder {
	m := &AppointmentsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/appointments{?contact_id*,limit*,offset*,since*,until*}", pathParameters),
	}
	return m
}

// NewAppointmentsRequestBuilder instantiates a new AppointmentsRequestBuilder and sets the default values.
func NewAppointmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AppointmentsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAppointmentsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves all appointments for the authenticated user
// Deprecated: This method is obsolete. Use GetAsAppointmentsGetResponse instead.
// returns a AppointmentsResponseable when successful
// returns a Appointments401Error error when the service returns a 401 status code
// returns a Appointments403Error error when the service returns a 403 status code
// returns a Appointments404Error error when the service returns a 404 status code
func (m *AppointmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *AppointmentsRequestBuilderGetRequestConfiguration) (AppointmentsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointments401ErrorFromDiscriminatorValue,
		"403": CreateAppointments403ErrorFromDiscriminatorValue,
		"404": CreateAppointments404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsResponseable), nil
}

// GetAsAppointmentsGetResponse retrieves all appointments for the authenticated user
// returns a AppointmentsGetResponseable when successful
// returns a Appointments401Error error when the service returns a 401 status code
// returns a Appointments403Error error when the service returns a 403 status code
// returns a Appointments404Error error when the service returns a 404 status code
func (m *AppointmentsRequestBuilder) GetAsAppointmentsGetResponse(ctx context.Context, requestConfiguration *AppointmentsRequestBuilderGetRequestConfiguration) (AppointmentsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointments401ErrorFromDiscriminatorValue,
		"403": CreateAppointments403ErrorFromDiscriminatorValue,
		"404": CreateAppointments404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsGetResponseable), nil
}

// Model the model property
// returns a *AppointmentsModelRequestBuilder when successful
func (m *AppointmentsRequestBuilder) Model() *AppointmentsModelRequestBuilder {
	return NewAppointmentsModelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Post creates a new appointment as the authenticated user
// Deprecated: This method is obsolete. Use PostAsAppointmentsPostResponse instead.
// returns a AppointmentsResponseable when successful
// returns a Appointments401Error error when the service returns a 401 status code
// returns a Appointments403Error error when the service returns a 403 status code
func (m *AppointmentsRequestBuilder) Post(ctx context.Context, body AppointmentsPostRequestBodyable, requestConfiguration *AppointmentsRequestBuilderPostRequestConfiguration) (AppointmentsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointments401ErrorFromDiscriminatorValue,
		"403": CreateAppointments403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsResponseable), nil
}

// PostAsAppointmentsPostResponse creates a new appointment as the authenticated user
// returns a AppointmentsPostResponseable when successful
// returns a Appointments401Error error when the service returns a 401 status code
// returns a Appointments403Error error when the service returns a 403 status code
func (m *AppointmentsRequestBuilder) PostAsAppointmentsPostResponse(ctx context.Context, body AppointmentsPostRequestBodyable, requestConfiguration *AppointmentsRequestBuilderPostRequestConfiguration) (AppointmentsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointments401ErrorFromDiscriminatorValue,
		"403": CreateAppointments403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsPostResponseable), nil
}

// ToGetRequestInformation retrieves all appointments for the authenticated user
// returns a *RequestInformation when successful
func (m *AppointmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppointmentsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation creates a new appointment as the authenticated user
// returns a *RequestInformation when successful
func (m *AppointmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AppointmentsPostRequestBodyable, requestConfiguration *AppointmentsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/appointments", m.BaseRequestBuilder.PathParameters)
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
// returns a *AppointmentsRequestBuilder when successful
func (m *AppointmentsRequestBuilder) WithUrl(rawUrl string) *AppointmentsRequestBuilder {
	return NewAppointmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
