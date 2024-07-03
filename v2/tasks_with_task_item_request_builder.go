package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TasksWithTask_ItemRequestBuilder builds and executes requests for operations under \v2\tasks\{task_id}
type TasksWithTask_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewTasksWithTask_ItemRequestBuilderInternal instantiates a new TasksWithTask_ItemRequestBuilder and sets the default values.
func NewTasksWithTask_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksWithTask_ItemRequestBuilder {
	m := &TasksWithTask_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/tasks/{task_id}", pathParameters),
	}
	return m
}

// NewTasksWithTask_ItemRequestBuilder instantiates a new TasksWithTask_ItemRequestBuilder and sets the default values.
func NewTasksWithTask_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksWithTask_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTasksWithTask_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
