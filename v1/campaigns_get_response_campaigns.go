package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type CampaignsGetResponse_campaigns struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The created_by_global_id property
	created_by_global_id *int64
	// The date_created property
	date_created *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The error_message property
	error_message *string
	// The goals property
	goals []CampaignsGetResponse_campaigns_goalsable
	// The id property
	id *int64
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
	sequences []CampaignsGetResponse_campaigns_sequencesable
	// The time_zone property
	time_zone *string
}

// NewCampaignsGetResponse_campaigns instantiates a new CampaignsGetResponse_campaigns and sets the default values.
func NewCampaignsGetResponse_campaigns() *CampaignsGetResponse_campaigns {
	m := &CampaignsGetResponse_campaigns{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsGetResponse_campaignsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsGetResponse_campaignsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsGetResponse_campaigns(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsGetResponse_campaigns) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCreatedByGlobalId gets the created_by_global_id property value. The created_by_global_id property
// returns a *int64 when successful
func (m *CampaignsGetResponse_campaigns) GetCreatedByGlobalId() *int64 {
	return m.created_by_global_id
}

// GetDateCreated gets the date_created property value. The date_created property
// returns a *Time when successful
func (m *CampaignsGetResponse_campaigns) GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.date_created
}

// GetErrorMessage gets the error_message property value. The error_message property
// returns a *string when successful
func (m *CampaignsGetResponse_campaigns) GetErrorMessage() *string {
	return m.error_message
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsGetResponse_campaigns) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["created_by_global_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreatedByGlobalId(val)
		}
		return nil
	}
	res["date_created"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
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
		val, err := n.GetCollectionOfObjectValues(CreateCampaignsGetResponse_campaigns_goalsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CampaignsGetResponse_campaigns_goalsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CampaignsGetResponse_campaigns_goalsable)
				}
			}
			m.SetGoals(res)
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
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
		val, err := n.GetCollectionOfObjectValues(CreateCampaignsGetResponse_campaigns_sequencesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CampaignsGetResponse_campaigns_sequencesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CampaignsGetResponse_campaigns_sequencesable)
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
// returns a []CampaignsGetResponse_campaigns_goalsable when successful
func (m *CampaignsGetResponse_campaigns) GetGoals() []CampaignsGetResponse_campaigns_goalsable {
	return m.goals
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *CampaignsGetResponse_campaigns) GetId() *int64 {
	return m.id
}

// GetLocked gets the locked property value. The locked property
// returns a *bool when successful
func (m *CampaignsGetResponse_campaigns) GetLocked() *bool {
	return m.locked
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *CampaignsGetResponse_campaigns) GetName() *string {
	return m.name
}

// GetPublishedDate gets the published_date property value. The published_date property
// returns a *Time when successful
func (m *CampaignsGetResponse_campaigns) GetPublishedDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.published_date
}

// GetPublishedStatus gets the published_status property value. The published_status property
// returns a *bool when successful
func (m *CampaignsGetResponse_campaigns) GetPublishedStatus() *bool {
	return m.published_status
}

// GetPublishedTimeZone gets the published_time_zone property value. The published_time_zone property
// returns a *string when successful
func (m *CampaignsGetResponse_campaigns) GetPublishedTimeZone() *string {
	return m.published_time_zone
}

// GetSequences gets the sequences property value. The sequences property
// returns a []CampaignsGetResponse_campaigns_sequencesable when successful
func (m *CampaignsGetResponse_campaigns) GetSequences() []CampaignsGetResponse_campaigns_sequencesable {
	return m.sequences
}

// GetTimeZone gets the time_zone property value. The time_zone property
// returns a *string when successful
func (m *CampaignsGetResponse_campaigns) GetTimeZone() *string {
	return m.time_zone
}

// Serialize serializes information the current object
func (m *CampaignsGetResponse_campaigns) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("created_by_global_id", m.GetCreatedByGlobalId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("date_created", m.GetDateCreated())
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
		err := writer.WriteInt64Value("id", m.GetId())
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

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CampaignsGetResponse_campaigns) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCreatedByGlobalId sets the created_by_global_id property value. The created_by_global_id property
func (m *CampaignsGetResponse_campaigns) SetCreatedByGlobalId(value *int64) {
	m.created_by_global_id = value
}

// SetDateCreated sets the date_created property value. The date_created property
func (m *CampaignsGetResponse_campaigns) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.date_created = value
}

// SetErrorMessage sets the error_message property value. The error_message property
func (m *CampaignsGetResponse_campaigns) SetErrorMessage(value *string) {
	m.error_message = value
}

// SetGoals sets the goals property value. The goals property
func (m *CampaignsGetResponse_campaigns) SetGoals(value []CampaignsGetResponse_campaigns_goalsable) {
	m.goals = value
}

// SetId sets the id property value. The id property
func (m *CampaignsGetResponse_campaigns) SetId(value *int64) {
	m.id = value
}

// SetLocked sets the locked property value. The locked property
func (m *CampaignsGetResponse_campaigns) SetLocked(value *bool) {
	m.locked = value
}

// SetName sets the name property value. The name property
func (m *CampaignsGetResponse_campaigns) SetName(value *string) {
	m.name = value
}

// SetPublishedDate sets the published_date property value. The published_date property
func (m *CampaignsGetResponse_campaigns) SetPublishedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.published_date = value
}

// SetPublishedStatus sets the published_status property value. The published_status property
func (m *CampaignsGetResponse_campaigns) SetPublishedStatus(value *bool) {
	m.published_status = value
}

// SetPublishedTimeZone sets the published_time_zone property value. The published_time_zone property
func (m *CampaignsGetResponse_campaigns) SetPublishedTimeZone(value *string) {
	m.published_time_zone = value
}

// SetSequences sets the sequences property value. The sequences property
func (m *CampaignsGetResponse_campaigns) SetSequences(value []CampaignsGetResponse_campaigns_sequencesable) {
	m.sequences = value
}

// SetTimeZone sets the time_zone property value. The time_zone property
func (m *CampaignsGetResponse_campaigns) SetTimeZone(value *string) {
	m.time_zone = value
}

type CampaignsGetResponse_campaignsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCreatedByGlobalId() *int64
	GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetErrorMessage() *string
	GetGoals() []CampaignsGetResponse_campaigns_goalsable
	GetId() *int64
	GetLocked() *bool
	GetName() *string
	GetPublishedDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetPublishedStatus() *bool
	GetPublishedTimeZone() *string
	GetSequences() []CampaignsGetResponse_campaigns_sequencesable
	GetTimeZone() *string
	SetCreatedByGlobalId(value *int64)
	SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetErrorMessage(value *string)
	SetGoals(value []CampaignsGetResponse_campaigns_goalsable)
	SetId(value *int64)
	SetLocked(value *bool)
	SetName(value *string)
	SetPublishedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetPublishedStatus(value *bool)
	SetPublishedTimeZone(value *string)
	SetSequences(value []CampaignsGetResponse_campaigns_sequencesable)
	SetTimeZone(value *string)
}
