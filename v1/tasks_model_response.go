package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TasksModelGetResponseable instead.
type TasksModelResponse struct {
	TasksModelGetResponse
}

// NewTasksModelResponse instantiates a new TasksModelResponse and sets the default values.
func NewTasksModelResponse() *TasksModelResponse {
	m := &TasksModelResponse{
		TasksModelGetResponse: *NewTasksModelGetResponse(),
	}
	return m
}

// CreateTasksModelResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTasksModelResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTasksModelResponse(), nil
}

// Deprecated: This class is obsolete. Use TasksModelGetResponseable instead.
type TasksModelResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TasksModelGetResponseable
}
