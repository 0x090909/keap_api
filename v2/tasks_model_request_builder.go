package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TasksModelRequestBuilder builds and executes requests for operations under \v2\tasks\model
type TasksModelRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewTasksModelRequestBuilderInternal instantiates a new TasksModelRequestBuilder and sets the default values.
func NewTasksModelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksModelRequestBuilder {
	m := &TasksModelRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/tasks/model", pathParameters),
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
