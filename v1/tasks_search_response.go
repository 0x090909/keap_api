package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TasksSearchGetResponseable instead.
type TasksSearchResponse struct {
	TasksSearchGetResponse
}

// NewTasksSearchResponse instantiates a new TasksSearchResponse and sets the default values.
func NewTasksSearchResponse() *TasksSearchResponse {
	m := &TasksSearchResponse{
		TasksSearchGetResponse: *NewTasksSearchGetResponse(),
	}
	return m
}

// CreateTasksSearchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTasksSearchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTasksSearchResponse(), nil
}

// Deprecated: This class is obsolete. Use TasksSearchGetResponseable instead.
type TasksSearchResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TasksSearchGetResponseable
}
