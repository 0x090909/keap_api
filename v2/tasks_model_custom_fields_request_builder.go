package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TasksModelCustomFieldsRequestBuilder builds and executes requests for operations under \v2\tasks\model\customFields
type TasksModelCustomFieldsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewTasksModelCustomFieldsRequestBuilderInternal instantiates a new TasksModelCustomFieldsRequestBuilder and sets the default values.
func NewTasksModelCustomFieldsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksModelCustomFieldsRequestBuilder {
	m := &TasksModelCustomFieldsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/tasks/model/customFields", pathParameters),
	}
	return m
}

// NewTasksModelCustomFieldsRequestBuilder instantiates a new TasksModelCustomFieldsRequestBuilder and sets the default values.
func NewTasksModelCustomFieldsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksModelCustomFieldsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTasksModelCustomFieldsRequestBuilderInternal(urlParams, requestAdapter)
}
