package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AutomationCategoryGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The automation_categories property
	automation_categories []AutomationCategoryGetResponse_automation_categoriesable
	// The next_page_token property
	next_page_token *string
}

// NewAutomationCategoryGetResponse instantiates a new AutomationCategoryGetResponse and sets the default values.
func NewAutomationCategoryGetResponse() *AutomationCategoryGetResponse {
	m := &AutomationCategoryGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAutomationCategoryGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAutomationCategoryGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAutomationCategoryGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AutomationCategoryGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAutomationCategories gets the automation_categories property value. The automation_categories property
// returns a []AutomationCategoryGetResponse_automation_categoriesable when successful
func (m *AutomationCategoryGetResponse) GetAutomationCategories() []AutomationCategoryGetResponse_automation_categoriesable {
	return m.automation_categories
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AutomationCategoryGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["automation_categories"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateAutomationCategoryGetResponse_automation_categoriesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]AutomationCategoryGetResponse_automation_categoriesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(AutomationCategoryGetResponse_automation_categoriesable)
				}
			}
			m.SetAutomationCategories(res)
		}
		return nil
	}
	res["next_page_token"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNextPageToken(val)
		}
		return nil
	}
	return res
}

// GetNextPageToken gets the next_page_token property value. The next_page_token property
// returns a *string when successful
func (m *AutomationCategoryGetResponse) GetNextPageToken() *string {
	return m.next_page_token
}

// Serialize serializes information the current object
func (m *AutomationCategoryGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetAutomationCategories() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAutomationCategories()))
		for i, v := range m.GetAutomationCategories() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("automation_categories", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("next_page_token", m.GetNextPageToken())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AutomationCategoryGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAutomationCategories sets the automation_categories property value. The automation_categories property
func (m *AutomationCategoryGetResponse) SetAutomationCategories(value []AutomationCategoryGetResponse_automation_categoriesable) {
	m.automation_categories = value
}

// SetNextPageToken sets the next_page_token property value. The next_page_token property
func (m *AutomationCategoryGetResponse) SetNextPageToken(value *string) {
	m.next_page_token = value
}

type AutomationCategoryGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAutomationCategories() []AutomationCategoryGetResponse_automation_categoriesable
	GetNextPageToken() *string
	SetAutomationCategories(value []AutomationCategoryGetResponse_automation_categoriesable)
	SetNextPageToken(value *string)
}
