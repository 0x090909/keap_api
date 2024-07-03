package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type CampaignsItemWithCampaign_GetResponse struct {
	// The active_contact_count property
	active_contact_count *int32
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The completed_contact_count property
	completed_contact_count *int32
	// The created_by_global_id property
	created_by_global_id *string
	// The date_created property
	date_created *string
	// The error_message property
	error_message *string
	// The goals property
	goals []CampaignsItemWithCampaign_GetResponse_goalsable
	// The id property
	id *string
	// The locked property
	locked *bool
	// The name property
	name *string
	// The published_date property
	published_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The published_status property
	published_status *bool
	// The published_time_zone property
	published_time_zone *string
	// The sequences property
	sequences []CampaignsItemWithCampaign_GetResponse_sequencesable
	// The time_zone property
	time_zone *string
}

// NewCampaignsItemWithCampaign_GetResponse instantiates a new CampaignsItemWithCampaign_GetResponse and sets the default values.
func NewCampaignsItemWithCampaign_GetResponse() *CampaignsItemWithCampaign_GetResponse {
	m := &CampaignsItemWithCampaign_GetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsItemWithCampaign_GetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemWithCampaign_GetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemWithCampaign_GetResponse(), nil
}

// GetActiveContactCount gets the active_contact_count property value. The active_contact_count property
// returns a *int32 when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetActiveContactCount() *int32 {
	return m.active_contact_count
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCompletedContactCount gets the completed_contact_count property value. The completed_contact_count property
// returns a *int32 when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetCompletedContactCount() *int32 {
	return m.completed_contact_count
}

// GetCreatedByGlobalId gets the created_by_global_id property value. The created_by_global_id property
// returns a *string when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetCreatedByGlobalId() *string {
	return m.created_by_global_id
}

// GetDateCreated gets the date_created property value. The date_created property
// returns a *string when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetDateCreated() *string {
	return m.date_created
}

// GetErrorMessage gets the error_message property value. The error_message property
// returns a *string when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetErrorMessage() *string {
	return m.error_message
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["active_contact_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetActiveContactCount(val)
		}
		return nil
	}
	res["completed_contact_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompletedContactCount(val)
		}
		return nil
	}
	res["created_by_global_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreatedByGlobalId(val)
		}
		return nil
	}
	res["date_created"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDateCreated(val)
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
	res["goals"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCampaignsItemWithCampaign_GetResponse_goalsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CampaignsItemWithCampaign_GetResponse_goalsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CampaignsItemWithCampaign_GetResponse_goalsable)
				}
			}
			m.SetGoals(res)
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
	res["locked"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLocked(val)
		}
		return nil
	}
	res["name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetName(val)
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
	res["published_status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPublishedStatus(val)
		}
		return nil
	}
	res["published_time_zone"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPublishedTimeZone(val)
		}
		return nil
	}
	res["sequences"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCampaignsItemWithCampaign_GetResponse_sequencesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CampaignsItemWithCampaign_GetResponse_sequencesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CampaignsItemWithCampaign_GetResponse_sequencesable)
				}
			}
			m.SetSequences(res)
		}
		return nil
	}
	res["time_zone"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTimeZone(val)
		}
		return nil
	}
	return res
}

// GetGoals gets the goals property value. The goals property
// returns a []CampaignsItemWithCampaign_GetResponse_goalsable when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetGoals() []CampaignsItemWithCampaign_GetResponse_goalsable {
	return m.goals
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetId() *string {
	return m.id
}

// GetLocked gets the locked property value. The locked property
// returns a *bool when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetLocked() *bool {
	return m.locked
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetName() *string {
	return m.name
}

// GetPublishedDate gets the published_date property value. The published_date property
// returns a *Time when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetPublishedDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.published_date
}

// GetPublishedStatus gets the published_status property value. The published_status property
// returns a *bool when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetPublishedStatus() *bool {
	return m.published_status
}

// GetPublishedTimeZone gets the published_time_zone property value. The published_time_zone property
// returns a *string when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetPublishedTimeZone() *string {
	return m.published_time_zone
}

// GetSequences gets the sequences property value. The sequences property
// returns a []CampaignsItemWithCampaign_GetResponse_sequencesable when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetSequences() []CampaignsItemWithCampaign_GetResponse_sequencesable {
	return m.sequences
}

// GetTimeZone gets the time_zone property value. The time_zone property
// returns a *string when successful
func (m *CampaignsItemWithCampaign_GetResponse) GetTimeZone() *string {
	return m.time_zone
}

// Serialize serializes information the current object
func (m *CampaignsItemWithCampaign_GetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("active_contact_count", m.GetActiveContactCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("completed_contact_count", m.GetCompletedContactCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("created_by_global_id", m.GetCreatedByGlobalId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("date_created", m.GetDateCreated())
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
	if m.GetGoals() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGoals()))
		for i, v := range m.GetGoals() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("goals", cast)
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
		err := writer.WriteBoolValue("locked", m.GetLocked())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("name", m.GetName())
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
		err := writer.WriteBoolValue("published_status", m.GetPublishedStatus())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("published_time_zone", m.GetPublishedTimeZone())
		if err != nil {
			return err
		}
	}
	if m.GetSequences() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSequences()))
		for i, v := range m.GetSequences() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("sequences", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("time_zone", m.GetTimeZone())
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

// SetActiveContactCount sets the active_contact_count property value. The active_contact_count property
func (m *CampaignsItemWithCampaign_GetResponse) SetActiveContactCount(value *int32) {
	m.active_contact_count = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CampaignsItemWithCampaign_GetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCompletedContactCount sets the completed_contact_count property value. The completed_contact_count property
func (m *CampaignsItemWithCampaign_GetResponse) SetCompletedContactCount(value *int32) {
	m.completed_contact_count = value
}

// SetCreatedByGlobalId sets the created_by_global_id property value. The created_by_global_id property
func (m *CampaignsItemWithCampaign_GetResponse) SetCreatedByGlobalId(value *string) {
	m.created_by_global_id = value
}

// SetDateCreated sets the date_created property value. The date_created property
func (m *CampaignsItemWithCampaign_GetResponse) SetDateCreated(value *string) {
	m.date_created = value
}

// SetErrorMessage sets the error_message property value. The error_message property
func (m *CampaignsItemWithCampaign_GetResponse) SetErrorMessage(value *string) {
	m.error_message = value
}

// SetGoals sets the goals property value. The goals property
func (m *CampaignsItemWithCampaign_GetResponse) SetGoals(value []CampaignsItemWithCampaign_GetResponse_goalsable) {
	m.goals = value
}

// SetId sets the id property value. The id property
func (m *CampaignsItemWithCampaign_GetResponse) SetId(value *string) {
	m.id = value
}

// SetLocked sets the locked property value. The locked property
func (m *CampaignsItemWithCampaign_GetResponse) SetLocked(value *bool) {
	m.locked = value
}

// SetName sets the name property value. The name property
func (m *CampaignsItemWithCampaign_GetResponse) SetName(value *string) {
	m.name = value
}

// SetPublishedDate sets the published_date property value. The published_date property
func (m *CampaignsItemWithCampaign_GetResponse) SetPublishedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.published_date = value
}

// SetPublishedStatus sets the published_status property value. The published_status property
func (m *CampaignsItemWithCampaign_GetResponse) SetPublishedStatus(value *bool) {
	m.published_status = value
}

// SetPublishedTimeZone sets the published_time_zone property value. The published_time_zone property
func (m *CampaignsItemWithCampaign_GetResponse) SetPublishedTimeZone(value *string) {
	m.published_time_zone = value
}

// SetSequences sets the sequences property value. The sequences property
func (m *CampaignsItemWithCampaign_GetResponse) SetSequences(value []CampaignsItemWithCampaign_GetResponse_sequencesable) {
	m.sequences = value
}

// SetTimeZone sets the time_zone property value. The time_zone property
func (m *CampaignsItemWithCampaign_GetResponse) SetTimeZone(value *string) {
	m.time_zone = value
}

type CampaignsItemWithCampaign_GetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetActiveContactCount() *int32
	GetCompletedContactCount() *int32
	GetCreatedByGlobalId() *string
	GetDateCreated() *string
	GetErrorMessage() *string
	GetGoals() []CampaignsItemWithCampaign_GetResponse_goalsable
	GetId() *string
	GetLocked() *bool
	GetName() *string
	GetPublishedDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetPublishedStatus() *bool
	GetPublishedTimeZone() *string
	GetSequences() []CampaignsItemWithCampaign_GetResponse_sequencesable
	GetTimeZone() *string
	SetActiveContactCount(value *int32)
	SetCompletedContactCount(value *int32)
	SetCreatedByGlobalId(value *string)
	SetDateCreated(value *string)
	SetErrorMessage(value *string)
	SetGoals(value []CampaignsItemWithCampaign_GetResponse_goalsable)
	SetId(value *string)
	SetLocked(value *bool)
	SetName(value *string)
	SetPublishedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetPublishedStatus(value *bool)
	SetPublishedTimeZone(value *string)
	SetSequences(value []CampaignsItemWithCampaign_GetResponse_sequencesable)
	SetTimeZone(value *string)
}
