package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TasksModelRequestBuilder builds and executes requests for operations under \v1\tasks\model
type TasksModelRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TasksModelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TasksModelRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewTasksModelRequestBuilderInternal instantiates a new TasksModelRequestBuilder and sets the default values.
func NewTasksModelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksModelRequestBuilder {
	m := &TasksModelRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/tasks/model", pathParameters),
	}
	return m
}

// NewTasksModelRequestBuilder instantiates a new TasksModelRequestBuilder and sets the default values.
func NewTasksModelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksModelRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTasksModelRequestBuilderInternal(urlParams, requestAdapter)
}

// CustomFields the customFields property
// returns a *TasksModelCustomFieldsRequestBuilder when successful
func (m *TasksModelRequestBuilder) CustomFields() *TasksModelCustomFieldsRequestBuilder {
	return NewTasksModelCustomFieldsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get get the custom fields for the Task object
// Deprecated: This method is obsolete. Use GetAsModelGetResponse instead.
// returns a TasksModelResponseable when successful
// returns a TasksModel401Error error when the service returns a 401 status code
// returns a TasksModel403Error error when the service returns a 403 status code
// returns a TasksModel404Error error when the service returns a 404 status code
func (m *TasksModelRequestBuilder) Get(ctx context.Context, requestConfiguration *TasksModelRequestBuilderGetRequestConfiguration) (TasksModelResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksModel401ErrorFromDiscriminatorValue,
		"403": CreateTasksModel403ErrorFromDiscriminatorValue,
		"404": CreateTasksModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksModelResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksModelResponseable), nil
}

// GetAsModelGetResponse get the custom fields for the Task object
// returns a TasksModelGetResponseable when successful
// returns a TasksModel401Error error when the service returns a 401 status code
// returns a TasksModel403Error error when the service returns a 403 status code
// returns a TasksModel404Error error when the service returns a 404 status code
func (m *TasksModelRequestBuilder) GetAsModelGetResponse(ctx context.Context, requestConfiguration *TasksModelRequestBuilderGetRequestConfiguration) (TasksModelGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksModel401ErrorFromDiscriminatorValue,
		"403": CreateTasksModel403ErrorFromDiscriminatorValue,
		"404": CreateTasksModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksModelGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksModelGetResponseable), nil
}

// ToGetRequestInformation get the custom fields for the Task object
// returns a *RequestInformation when successful
func (m *TasksModelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TasksModelRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TasksModelRequestBuilder when successful
func (m *TasksModelRequestBuilder) WithUrl(rawUrl string) *TasksModelRequestBuilder {
	return NewTasksModelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
