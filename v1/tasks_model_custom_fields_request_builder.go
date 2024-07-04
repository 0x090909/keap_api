package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TasksModelCustomFieldsRequestBuilder builds and executes requests for operations under \v1\tasks\model\customFields
type TasksModelCustomFieldsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TasksModelCustomFieldsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TasksModelCustomFieldsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewTasksModelCustomFieldsRequestBuilderInternal instantiates a new TasksModelCustomFieldsRequestBuilder and sets the default values.
func NewTasksModelCustomFieldsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksModelCustomFieldsRequestBuilder {
	m := &TasksModelCustomFieldsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/tasks/model/customFields", pathParameters),
	}
	return m
}

// NewTasksModelCustomFieldsRequestBuilder instantiates a new TasksModelCustomFieldsRequestBuilder and sets the default values.
func NewTasksModelCustomFieldsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksModelCustomFieldsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTasksModelCustomFieldsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post adds a custom field of the specified type and options to the Task object.
// Deprecated: This method is obsolete. Use PostAsCustomFieldsPostResponse instead.
// returns a TasksModelCustomFieldsResponseable when successful
// returns a TasksModelCustomFields401Error error when the service returns a 401 status code
// returns a TasksModelCustomFields403Error error when the service returns a 403 status code
func (m *TasksModelCustomFieldsRequestBuilder) Post(ctx context.Context, body TasksModelCustomFieldsPostRequestBodyable, requestConfiguration *TasksModelCustomFieldsRequestBuilderPostRequestConfiguration) (TasksModelCustomFieldsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksModelCustomFields401ErrorFromDiscriminatorValue,
		"403": CreateTasksModelCustomFields403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksModelCustomFieldsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksModelCustomFieldsResponseable), nil
}

// PostAsCustomFieldsPostResponse adds a custom field of the specified type and options to the Task object.
// returns a TasksModelCustomFieldsPostResponseable when successful
// returns a TasksModelCustomFields401Error error when the service returns a 401 status code
// returns a TasksModelCustomFields403Error error when the service returns a 403 status code
func (m *TasksModelCustomFieldsRequestBuilder) PostAsCustomFieldsPostResponse(ctx context.Context, body TasksModelCustomFieldsPostRequestBodyable, requestConfiguration *TasksModelCustomFieldsRequestBuilderPostRequestConfiguration) (TasksModelCustomFieldsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksModelCustomFields401ErrorFromDiscriminatorValue,
		"403": CreateTasksModelCustomFields403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksModelCustomFieldsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksModelCustomFieldsPostResponseable), nil
}

// ToPostRequestInformation adds a custom field of the specified type and options to the Task object.
// returns a *RequestInformation when successful
func (m *TasksModelCustomFieldsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TasksModelCustomFieldsPostRequestBodyable, requestConfiguration *TasksModelCustomFieldsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TasksModelCustomFieldsRequestBuilder when successful
func (m *TasksModelCustomFieldsRequestBuilder) WithUrl(rawUrl string) *TasksModelCustomFieldsRequestBuilder {
	return NewTasksModelCustomFieldsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
