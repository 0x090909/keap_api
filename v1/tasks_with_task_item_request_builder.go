package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TasksWithTaskItemRequestBuilder builds and executes requests for operations under \v1\tasks\{taskId}
type TasksWithTaskItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TasksWithTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TasksWithTaskItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// TasksWithTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TasksWithTaskItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// TasksWithTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TasksWithTaskItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// TasksWithTaskItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TasksWithTaskItemRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewTasksWithTaskItemRequestBuilderInternal instantiates a new TasksWithTaskItemRequestBuilder and sets the default values.
func NewTasksWithTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksWithTaskItemRequestBuilder {
	m := &TasksWithTaskItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/tasks/{taskId}", pathParameters),
	}
	return m
}

// NewTasksWithTaskItemRequestBuilder instantiates a new TasksWithTaskItemRequestBuilder and sets the default values.
func NewTasksWithTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksWithTaskItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTasksWithTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete permanently deletes a task
// returns a TasksItemWithTask401Error error when the service returns a 401 status code
// returns a TasksItemWithTask403Error error when the service returns a 403 status code
// returns a TasksItemWithTask404Error error when the service returns a 404 status code
func (m *TasksWithTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TasksWithTaskItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksItemWithTask401ErrorFromDiscriminatorValue,
		"403": CreateTasksItemWithTask403ErrorFromDiscriminatorValue,
		"404": CreateTasksItemWithTask404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves a single task
// Deprecated: This method is obsolete. Use GetAsWithTaskGetResponse instead.
// returns a TasksItemWithTaskResponseable when successful
// returns a TasksItemWithTask401Error error when the service returns a 401 status code
// returns a TasksItemWithTask403Error error when the service returns a 403 status code
// returns a TasksItemWithTask404Error error when the service returns a 404 status code
func (m *TasksWithTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TasksWithTaskItemRequestBuilderGetRequestConfiguration) (TasksItemWithTaskResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksItemWithTask401ErrorFromDiscriminatorValue,
		"403": CreateTasksItemWithTask403ErrorFromDiscriminatorValue,
		"404": CreateTasksItemWithTask404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksItemWithTaskResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksItemWithTaskResponseable), nil
}

// GetAsWithTaskGetResponse retrieves a single task
// returns a TasksItemWithTaskGetResponseable when successful
// returns a TasksItemWithTask401Error error when the service returns a 401 status code
// returns a TasksItemWithTask403Error error when the service returns a 403 status code
// returns a TasksItemWithTask404Error error when the service returns a 404 status code
func (m *TasksWithTaskItemRequestBuilder) GetAsWithTaskGetResponse(ctx context.Context, requestConfiguration *TasksWithTaskItemRequestBuilderGetRequestConfiguration) (TasksItemWithTaskGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksItemWithTask401ErrorFromDiscriminatorValue,
		"403": CreateTasksItemWithTask403ErrorFromDiscriminatorValue,
		"404": CreateTasksItemWithTask404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksItemWithTaskGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksItemWithTaskGetResponseable), nil
}

// Patch updates a task with only the values provided in the request. This endpoint does not currently support updating Custom Field values.
// Deprecated: This method is obsolete. Use PatchAsWithTaskPatchResponse instead.
// returns a TasksItemWithTaskResponseable when successful
// returns a TasksItemWithTask401Error error when the service returns a 401 status code
// returns a TasksItemWithTask403Error error when the service returns a 403 status code
// returns a TasksItemWithTask404Error error when the service returns a 404 status code
func (m *TasksWithTaskItemRequestBuilder) Patch(ctx context.Context, body TasksItemWithTaskPatchRequestBodyable, requestConfiguration *TasksWithTaskItemRequestBuilderPatchRequestConfiguration) (TasksItemWithTaskResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksItemWithTask401ErrorFromDiscriminatorValue,
		"403": CreateTasksItemWithTask403ErrorFromDiscriminatorValue,
		"404": CreateTasksItemWithTask404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksItemWithTaskResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksItemWithTaskResponseable), nil
}

// PatchAsWithTaskPatchResponse updates a task with only the values provided in the request. This endpoint does not currently support updating Custom Field values.
// returns a TasksItemWithTaskPatchResponseable when successful
// returns a TasksItemWithTask401Error error when the service returns a 401 status code
// returns a TasksItemWithTask403Error error when the service returns a 403 status code
// returns a TasksItemWithTask404Error error when the service returns a 404 status code
func (m *TasksWithTaskItemRequestBuilder) PatchAsWithTaskPatchResponse(ctx context.Context, body TasksItemWithTaskPatchRequestBodyable, requestConfiguration *TasksWithTaskItemRequestBuilderPatchRequestConfiguration) (TasksItemWithTaskPatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksItemWithTask401ErrorFromDiscriminatorValue,
		"403": CreateTasksItemWithTask403ErrorFromDiscriminatorValue,
		"404": CreateTasksItemWithTask404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksItemWithTaskPatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksItemWithTaskPatchResponseable), nil
}

// Put replaces a task with the values provided in the request. This endpoint does not currently support updating Custom Field values.
// Deprecated: This method is obsolete. Use PutAsWithTaskPutResponse instead.
// returns a TasksItemWithTaskResponseable when successful
// returns a TasksItemWithTask401Error error when the service returns a 401 status code
// returns a TasksItemWithTask403Error error when the service returns a 403 status code
// returns a TasksItemWithTask404Error error when the service returns a 404 status code
func (m *TasksWithTaskItemRequestBuilder) Put(ctx context.Context, body TasksItemWithTaskPutRequestBodyable, requestConfiguration *TasksWithTaskItemRequestBuilderPutRequestConfiguration) (TasksItemWithTaskResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksItemWithTask401ErrorFromDiscriminatorValue,
		"403": CreateTasksItemWithTask403ErrorFromDiscriminatorValue,
		"404": CreateTasksItemWithTask404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksItemWithTaskResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksItemWithTaskResponseable), nil
}

// PutAsWithTaskPutResponse replaces a task with the values provided in the request. This endpoint does not currently support updating Custom Field values.
// returns a TasksItemWithTaskPutResponseable when successful
// returns a TasksItemWithTask401Error error when the service returns a 401 status code
// returns a TasksItemWithTask403Error error when the service returns a 403 status code
// returns a TasksItemWithTask404Error error when the service returns a 404 status code
func (m *TasksWithTaskItemRequestBuilder) PutAsWithTaskPutResponse(ctx context.Context, body TasksItemWithTaskPutRequestBodyable, requestConfiguration *TasksWithTaskItemRequestBuilderPutRequestConfiguration) (TasksItemWithTaskPutResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksItemWithTask401ErrorFromDiscriminatorValue,
		"403": CreateTasksItemWithTask403ErrorFromDiscriminatorValue,
		"404": CreateTasksItemWithTask404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksItemWithTaskPutResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksItemWithTaskPutResponseable), nil
}

// ToDeleteRequestInformation permanently deletes a task
// returns a *RequestInformation when successful
func (m *TasksWithTaskItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TasksWithTaskItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieves a single task
// returns a *RequestInformation when successful
func (m *TasksWithTaskItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TasksWithTaskItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPatchRequestInformation updates a task with only the values provided in the request. This endpoint does not currently support updating Custom Field values.
// returns a *RequestInformation when successful
func (m *TasksWithTaskItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body TasksItemWithTaskPatchRequestBodyable, requestConfiguration *TasksWithTaskItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPutRequestInformation replaces a task with the values provided in the request. This endpoint does not currently support updating Custom Field values.
// returns a *RequestInformation when successful
func (m *TasksWithTaskItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body TasksItemWithTaskPutRequestBodyable, requestConfiguration *TasksWithTaskItemRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TasksWithTaskItemRequestBuilder when successful
func (m *TasksWithTaskItemRequestBuilder) WithUrl(rawUrl string) *TasksWithTaskItemRequestBuilder {
	return NewTasksWithTaskItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
