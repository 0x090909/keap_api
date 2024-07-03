package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AutomationCategoryGetResponseable instead.
type AutomationCategoryResponse struct {
	AutomationCategoryGetResponse
}

// NewAutomationCategoryResponse instantiates a new AutomationCategoryResponse and sets the default values.
func NewAutomationCategoryResponse() *AutomationCategoryResponse {
	m := &AutomationCategoryResponse{
		AutomationCategoryGetResponse: *NewAutomationCategoryGetResponse(),
	}
	return m
}

// CreateAutomationCategoryResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAutomationCategoryResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAutomationCategoryResponse(), nil
}

// Deprecated: This class is obsolete. Use AutomationCategoryGetResponseable instead.
type AutomationCategoryResponseable interface {
	AutomationCategoryGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
