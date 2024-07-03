package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type AutomationsItemWithAutomation_GetResponse struct {
	// The active_contacts property
	active_contacts *int32
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The categories property
	categories []string
	// The error_message property
	error_message *string
	// The id property
	id *string
	// The published_by property
	published_by *string
	// The published_date property
	published_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The published_timezone property
	published_timezone *string
	// The status property
	status *string
	// The title property
	title *string
}

// NewAutomationsItemWithAutomation_GetResponse instantiates a new AutomationsItemWithAutomation_GetResponse and sets the default values.
func NewAutomationsItemWithAutomation_GetResponse() *AutomationsItemWithAutomation_GetResponse {
	m := &AutomationsItemWithAutomation_GetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAutomationsItemWithAutomation_GetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAutomationsItemWithAutomation_GetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAutomationsItemWithAutomation_GetResponse(), nil
}

// GetActiveContacts gets the active_contacts property value. The active_contacts property
// returns a *int32 when successful
func (m *AutomationsItemWithAutomation_GetResponse) GetActiveContacts() *int32 {
	return m.active_contacts
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AutomationsItemWithAutomation_GetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCategories gets the categories property value. The categories property
// returns a []string when successful
func (m *AutomationsItemWithAutomation_GetResponse) GetCategories() []string {
	return m.categories
}

// GetErrorMessage gets the error_message property value. The error_message property
// returns a *string when successful
func (m *AutomationsItemWithAutomation_GetResponse) GetErrorMessage() *string {
	return m.error_message
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AutomationsItemWithAutomation_GetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["active_contacts"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetActiveContacts(val)
		}
		return nil
	}
	res["categories"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetCategories(res)
		}
		return nil
	}
	res["error_message"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetErrorMessage(val)
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["published_by"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPublishedBy(val)
		}
		return nil
	}
	res["published_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPublishedDate(val)
		}
		return nil
	}
	res["published_timezone"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPublishedTimezone(val)
		}
		return nil
	}
	res["status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStatus(val)
		}
		return nil
	}
	res["title"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTitle(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *AutomationsItemWithAutomation_GetResponse) GetId() *string {
	return m.id
}

// GetPublishedBy gets the published_by property value. The published_by property
// returns a *string when successful
func (m *AutomationsItemWithAutomation_GetResponse) GetPublishedBy() *string {
	return m.published_by
}

// GetPublishedDate gets the published_date property value. The published_date property
// returns a *Time when successful
func (m *AutomationsItemWithAutomation_GetResponse) GetPublishedDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.published_date
}

// GetPublishedTimezone gets the published_timezone property value. The published_timezone property
// returns a *string when successful
func (m *AutomationsItemWithAutomation_GetResponse) GetPublishedTimezone() *string {
	return m.published_timezone
}

// GetStatus gets the status property value. The status property
// returns a *string when successful
func (m *AutomationsItemWithAutomation_GetResponse) GetStatus() *string {
	return m.status
}

// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *AutomationsItemWithAutomation_GetResponse) GetTitle() *string {
	return m.title
}

// Serialize serializes information the current object
func (m *AutomationsItemWithAutomation_GetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("active_contacts", m.GetActiveContacts())
		if err != nil {
			return err
		}
	}
	if m.GetCategories() != nil {
		err := writer.WriteCollectionOfStringValues("categories", m.GetCategories())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("error_message", m.GetErrorMessage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("published_by", m.GetPublishedBy())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("published_date", m.GetPublishedDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("published_timezone", m.GetPublishedTimezone())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("status", m.GetStatus())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("title", m.GetTitle())
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

// SetActiveContacts sets the active_contacts property value. The active_contacts property
func (m *AutomationsItemWithAutomation_GetResponse) SetActiveContacts(value *int32) {
	m.active_contacts = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AutomationsItemWithAutomation_GetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCategories sets the categories property value. The categories property
func (m *AutomationsItemWithAutomation_GetResponse) SetCategories(value []string) {
	m.categories = value
}

// SetErrorMessage sets the error_message property value. The error_message property
func (m *AutomationsItemWithAutomation_GetResponse) SetErrorMessage(value *string) {
	m.error_message = value
}

// SetId sets the id property value. The id property
func (m *AutomationsItemWithAutomation_GetResponse) SetId(value *string) {
	m.id = value
}

// SetPublishedBy sets the published_by property value. The published_by property
func (m *AutomationsItemWithAutomation_GetResponse) SetPublishedBy(value *string) {
	m.published_by = value
}

// SetPublishedDate sets the published_date property value. The published_date property
func (m *AutomationsItemWithAutomation_GetResponse) SetPublishedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.published_date = value
}

// SetPublishedTimezone sets the published_timezone property value. The published_timezone property
func (m *AutomationsItemWithAutomation_GetResponse) SetPublishedTimezone(value *string) {
	m.published_timezone = value
}

// SetStatus sets the status property value. The status property
func (m *AutomationsItemWithAutomation_GetResponse) SetStatus(value *string) {
	m.status = value
}

// SetTitle sets the title property value. The title property
func (m *AutomationsItemWithAutomation_GetResponse) SetTitle(value *string) {
	m.title = value
}

type AutomationsItemWithAutomation_GetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetActiveContacts() *int32
	GetCategories() []string
	GetErrorMessage() *string
	GetId() *string
	GetPublishedBy() *string
	GetPublishedDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetPublishedTimezone() *string
	GetStatus() *string
	GetTitle() *string
	SetActiveContacts(value *int32)
	SetCategories(value []string)
	SetErrorMessage(value *string)
	SetId(value *string)
	SetPublishedBy(value *string)
	SetPublishedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetPublishedTimezone(value *string)
	SetStatus(value *string)
	SetTitle(value *string)
}
