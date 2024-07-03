package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TasksRequestBuilder builds and executes requests for operations under \v2\tasks
type TasksRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByTask_id gets an item from the keapapi.v2.tasks.item collection
// returns a *TasksWithTask_ItemRequestBuilder when successful
func (m *TasksRequestBuilder) ByTask_id(task_id string) *TasksWithTask_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if task_id != "" {
		urlTplParams["task_id"] = task_id
	}
	return NewTasksWithTask_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewTasksRequestBuilderInternal instantiates a new TasksRequestBuilder and sets the default values.
func NewTasksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksRequestBuilder {
	m := &TasksRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/tasks", pathParameters),
	}
	return m
}

// NewTasksRequestBuilder instantiates a new TasksRequestBuilder and sets the default values.
func NewTasksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTasksRequestBuilderInternal(urlParams, requestAdapter)
}

// Model the model property
// returns a *TasksModelRequestBuilder when successful
func (m *TasksRequestBuilder) Model() *TasksModelRequestBuilder {
	return NewTasksModelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
