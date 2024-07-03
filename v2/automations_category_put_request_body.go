package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AutomationsCategoryPutRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The apply_category property
	apply_category *bool
	// The automation_ids property
	automation_ids []int64
	// The category_ids property
	category_ids []int64
}

// NewAutomationsCategoryPutRequestBody instantiates a new AutomationsCategoryPutRequestBody and sets the default values.
func NewAutomationsCategoryPutRequestBody() *AutomationsCategoryPutRequestBody {
	m := &AutomationsCategoryPutRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAutomationsCategoryPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAutomationsCategoryPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAutomationsCategoryPutRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AutomationsCategoryPutRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetApplyCategory gets the apply_category property value. The apply_category property
// returns a *bool when successful
func (m *AutomationsCategoryPutRequestBody) GetApplyCategory() *bool {
	return m.apply_category
}

// GetAutomationIds gets the automation_ids property value. The automation_ids property
// returns a []int64 when successful
func (m *AutomationsCategoryPutRequestBody) GetAutomationIds() []int64 {
	return m.automation_ids
}

// GetCategoryIds gets the category_ids property value. The category_ids property
// returns a []int64 when successful
func (m *AutomationsCategoryPutRequestBody) GetCategoryIds() []int64 {
	return m.category_ids
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AutomationsCategoryPutRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["apply_category"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetApplyCategory(val)
		}
		return nil
	}
	res["automation_ids"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("int64")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]int64, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*int64))
				}
			}
			m.SetAutomationIds(res)
		}
		return nil
	}
	res["category_ids"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("int64")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]int64, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*int64))
				}
			}
			m.SetCategoryIds(res)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *AutomationsCategoryPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("apply_category", m.GetApplyCategory())
		if err != nil {
			return err
		}
	}
	if m.GetAutomationIds() != nil {
		err := writer.WriteCollectionOfInt64Values("automation_ids", m.GetAutomationIds())
		if err != nil {
			return err
		}
	}
	if m.GetCategoryIds() != nil {
		err := writer.WriteCollectionOfInt64Values("category_ids", m.GetCategoryIds())
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
func (m *AutomationsCategoryPutRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetApplyCategory sets the apply_category property value. The apply_category property
func (m *AutomationsCategoryPutRequestBody) SetApplyCategory(value *bool) {
	m.apply_category = value
}

// SetAutomationIds sets the automation_ids property value. The automation_ids property
func (m *AutomationsCategoryPutRequestBody) SetAutomationIds(value []int64) {
	m.automation_ids = value
}

// SetCategoryIds sets the category_ids property value. The category_ids property
func (m *AutomationsCategoryPutRequestBody) SetCategoryIds(value []int64) {
	m.category_ids = value
}

type AutomationsCategoryPutRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetApplyCategory() *bool
	GetAutomationIds() []int64
	GetCategoryIds() []int64
	SetApplyCategory(value *bool)
	SetAutomationIds(value []int64)
	SetCategoryIds(value []int64)
}
