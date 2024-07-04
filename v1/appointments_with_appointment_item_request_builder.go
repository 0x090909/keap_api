package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AppointmentsWithAppointmentItemRequestBuilder builds and executes requests for operations under \v1\appointments\{appointmentId}
type AppointmentsWithAppointmentItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AppointmentsWithAppointmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppointmentsWithAppointmentItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// AppointmentsWithAppointmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppointmentsWithAppointmentItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// AppointmentsWithAppointmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppointmentsWithAppointmentItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// AppointmentsWithAppointmentItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppointmentsWithAppointmentItemRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewAppointmentsWithAppointmentItemRequestBuilderInternal instantiates a new AppointmentsWithAppointmentItemRequestBuilder and sets the default values.
func NewAppointmentsWithAppointmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AppointmentsWithAppointmentItemRequestBuilder {
	m := &AppointmentsWithAppointmentItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/appointments/{appointmentId}", pathParameters),
	}
	return m
}

// NewAppointmentsWithAppointmentItemRequestBuilder instantiates a new AppointmentsWithAppointmentItemRequestBuilder and sets the default values.
func NewAppointmentsWithAppointmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AppointmentsWithAppointmentItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAppointmentsWithAppointmentItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes the specified appointment
// returns a AppointmentsItemWithAppointment401Error error when the service returns a 401 status code
// returns a AppointmentsItemWithAppointment403Error error when the service returns a 403 status code
// returns a AppointmentsItemWithAppointment404Error error when the service returns a 404 status code
func (m *AppointmentsWithAppointmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppointmentsWithAppointmentItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointmentsItemWithAppointment401ErrorFromDiscriminatorValue,
		"403": CreateAppointmentsItemWithAppointment403ErrorFromDiscriminatorValue,
		"404": CreateAppointmentsItemWithAppointment404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves a specific appointment with respect to user permissions. The authenticated user will need the "can view all records" permission for Task/Appt/Notes
// Deprecated: This method is obsolete. Use GetAsWithAppointmentGetResponse instead.
// returns a AppointmentsItemWithAppointmentResponseable when successful
// returns a AppointmentsItemWithAppointment401Error error when the service returns a 401 status code
// returns a AppointmentsItemWithAppointment403Error error when the service returns a 403 status code
// returns a AppointmentsItemWithAppointment404Error error when the service returns a 404 status code
func (m *AppointmentsWithAppointmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppointmentsWithAppointmentItemRequestBuilderGetRequestConfiguration) (AppointmentsItemWithAppointmentResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointmentsItemWithAppointment401ErrorFromDiscriminatorValue,
		"403": CreateAppointmentsItemWithAppointment403ErrorFromDiscriminatorValue,
		"404": CreateAppointmentsItemWithAppointment404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsItemWithAppointmentResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsItemWithAppointmentResponseable), nil
}

// GetAsWithAppointmentGetResponse retrieves a specific appointment with respect to user permissions. The authenticated user will need the "can view all records" permission for Task/Appt/Notes
// returns a AppointmentsItemWithAppointmentGetResponseable when successful
// returns a AppointmentsItemWithAppointment401Error error when the service returns a 401 status code
// returns a AppointmentsItemWithAppointment403Error error when the service returns a 403 status code
// returns a AppointmentsItemWithAppointment404Error error when the service returns a 404 status code
func (m *AppointmentsWithAppointmentItemRequestBuilder) GetAsWithAppointmentGetResponse(ctx context.Context, requestConfiguration *AppointmentsWithAppointmentItemRequestBuilderGetRequestConfiguration) (AppointmentsItemWithAppointmentGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointmentsItemWithAppointment401ErrorFromDiscriminatorValue,
		"403": CreateAppointmentsItemWithAppointment403ErrorFromDiscriminatorValue,
		"404": CreateAppointmentsItemWithAppointment404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsItemWithAppointmentGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsItemWithAppointmentGetResponseable), nil
}

// Patch updates the provided values of a given appointment
// Deprecated: This method is obsolete. Use PatchAsWithAppointmentPatchResponse instead.
// returns a AppointmentsItemWithAppointmentResponseable when successful
// returns a AppointmentsItemWithAppointment401Error error when the service returns a 401 status code
// returns a AppointmentsItemWithAppointment403Error error when the service returns a 403 status code
// returns a AppointmentsItemWithAppointment404Error error when the service returns a 404 status code
func (m *AppointmentsWithAppointmentItemRequestBuilder) Patch(ctx context.Context, body AppointmentsItemWithAppointmentPatchRequestBodyable, requestConfiguration *AppointmentsWithAppointmentItemRequestBuilderPatchRequestConfiguration) (AppointmentsItemWithAppointmentResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointmentsItemWithAppointment401ErrorFromDiscriminatorValue,
		"403": CreateAppointmentsItemWithAppointment403ErrorFromDiscriminatorValue,
		"404": CreateAppointmentsItemWithAppointment404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsItemWithAppointmentResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsItemWithAppointmentResponseable), nil
}

// PatchAsWithAppointmentPatchResponse updates the provided values of a given appointment
// returns a AppointmentsItemWithAppointmentPatchResponseable when successful
// returns a AppointmentsItemWithAppointment401Error error when the service returns a 401 status code
// returns a AppointmentsItemWithAppointment403Error error when the service returns a 403 status code
// returns a AppointmentsItemWithAppointment404Error error when the service returns a 404 status code
func (m *AppointmentsWithAppointmentItemRequestBuilder) PatchAsWithAppointmentPatchResponse(ctx context.Context, body AppointmentsItemWithAppointmentPatchRequestBodyable, requestConfiguration *AppointmentsWithAppointmentItemRequestBuilderPatchRequestConfiguration) (AppointmentsItemWithAppointmentPatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointmentsItemWithAppointment401ErrorFromDiscriminatorValue,
		"403": CreateAppointmentsItemWithAppointment403ErrorFromDiscriminatorValue,
		"404": CreateAppointmentsItemWithAppointment404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsItemWithAppointmentPatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsItemWithAppointmentPatchResponseable), nil
}

// Put replaces all values of a given appointment
// Deprecated: This method is obsolete. Use PutAsWithAppointmentPutResponse instead.
// returns a AppointmentsItemWithAppointmentResponseable when successful
// returns a AppointmentsItemWithAppointment401Error error when the service returns a 401 status code
// returns a AppointmentsItemWithAppointment403Error error when the service returns a 403 status code
// returns a AppointmentsItemWithAppointment404Error error when the service returns a 404 status code
func (m *AppointmentsWithAppointmentItemRequestBuilder) Put(ctx context.Context, body AppointmentsItemWithAppointmentPutRequestBodyable, requestConfiguration *AppointmentsWithAppointmentItemRequestBuilderPutRequestConfiguration) (AppointmentsItemWithAppointmentResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointmentsItemWithAppointment401ErrorFromDiscriminatorValue,
		"403": CreateAppointmentsItemWithAppointment403ErrorFromDiscriminatorValue,
		"404": CreateAppointmentsItemWithAppointment404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsItemWithAppointmentResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsItemWithAppointmentResponseable), nil
}

// PutAsWithAppointmentPutResponse replaces all values of a given appointment
// returns a AppointmentsItemWithAppointmentPutResponseable when successful
// returns a AppointmentsItemWithAppointment401Error error when the service returns a 401 status code
// returns a AppointmentsItemWithAppointment403Error error when the service returns a 403 status code
// returns a AppointmentsItemWithAppointment404Error error when the service returns a 404 status code
func (m *AppointmentsWithAppointmentItemRequestBuilder) PutAsWithAppointmentPutResponse(ctx context.Context, body AppointmentsItemWithAppointmentPutRequestBodyable, requestConfiguration *AppointmentsWithAppointmentItemRequestBuilderPutRequestConfiguration) (AppointmentsItemWithAppointmentPutResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAppointmentsItemWithAppointment401ErrorFromDiscriminatorValue,
		"403": CreateAppointmentsItemWithAppointment403ErrorFromDiscriminatorValue,
		"404": CreateAppointmentsItemWithAppointment404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAppointmentsItemWithAppointmentPutResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AppointmentsItemWithAppointmentPutResponseable), nil
}

// ToDeleteRequestInformation deletes the specified appointment
// returns a *RequestInformation when successful
func (m *AppointmentsWithAppointmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AppointmentsWithAppointmentItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieves a specific appointment with respect to user permissions. The authenticated user will need the "can view all records" permission for Task/Appt/Notes
// returns a *RequestInformation when successful
func (m *AppointmentsWithAppointmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppointmentsWithAppointmentItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPatchRequestInformation updates the provided values of a given appointment
// returns a *RequestInformation when successful
func (m *AppointmentsWithAppointmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body AppointmentsItemWithAppointmentPatchRequestBodyable, requestConfiguration *AppointmentsWithAppointmentItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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

// ToPutRequestInformation replaces all values of a given appointment
// returns a *RequestInformation when successful
func (m *AppointmentsWithAppointmentItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body AppointmentsItemWithAppointmentPutRequestBodyable, requestConfiguration *AppointmentsWithAppointmentItemRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AppointmentsWithAppointmentItemRequestBuilder when successful
func (m *AppointmentsWithAppointmentItemRequestBuilder) WithUrl(rawUrl string) *AppointmentsWithAppointmentItemRequestBuilder {
	return NewAppointmentsWithAppointmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
