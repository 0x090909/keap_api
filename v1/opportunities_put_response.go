package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type OpportunitiesPutResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The affiliate_id property
	affiliate_id *int64
	// The contact property
	contact OpportunitiesPutResponse_contactable
	// The custom_fields property
	custom_fields []OpportunitiesPutResponse_custom_fieldsable
	// The date_created property
	date_created *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The estimated_close_date property
	estimated_close_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The id property
	id *int64
	// The include_in_forecast property
	include_in_forecast *int32
	// The last_updated property
	last_updated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The next_action_date property
	next_action_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The next_action_notes property
	next_action_notes *string
	// The opportunity_notes property
	opportunity_notes *string
	// The opportunity_title property
	opportunity_title *string
	// The projected_revenue_high property
	projected_revenue_high *float64
	// The projected_revenue_low property
	projected_revenue_low *float64
	// The stage property
	stage OpportunitiesPutResponse_stageable
	// The user property
	user OpportunitiesPutResponse_userable
}

// NewOpportunitiesPutResponse instantiates a new OpportunitiesPutResponse and sets the default values.
func NewOpportunitiesPutResponse() *OpportunitiesPutResponse {
	m := &OpportunitiesPutResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOpportunitiesPutResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOpportunitiesPutResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOpportunitiesPutResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OpportunitiesPutResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAffiliateId gets the affiliate_id property value. The affiliate_id property
// returns a *int64 when successful
func (m *OpportunitiesPutResponse) GetAffiliateId() *int64 {
	return m.affiliate_id
}

// GetContact gets the contact property value. The contact property
// returns a OpportunitiesPutResponse_contactable when successful
func (m *OpportunitiesPutResponse) GetContact() OpportunitiesPutResponse_contactable {
	return m.contact
}

// GetCustomFields gets the custom_fields property value. The custom_fields property
// returns a []OpportunitiesPutResponse_custom_fieldsable when successful
func (m *OpportunitiesPutResponse) GetCustomFields() []OpportunitiesPutResponse_custom_fieldsable {
	return m.custom_fields
}

// GetDateCreated gets the date_created property value. The date_created property
// returns a *Time when successful
func (m *OpportunitiesPutResponse) GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.date_created
}

// GetEstimatedCloseDate gets the estimated_close_date property value. The estimated_close_date property
// returns a *Time when successful
func (m *OpportunitiesPutResponse) GetEstimatedCloseDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.estimated_close_date
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OpportunitiesPutResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["affiliate_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAffiliateId(val)
		}
		return nil
	}
	res["contact"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOpportunitiesPutResponse_contactFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContact(val.(OpportunitiesPutResponse_contactable))
		}
		return nil
	}
	res["custom_fields"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateOpportunitiesPutResponse_custom_fieldsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]OpportunitiesPutResponse_custom_fieldsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(OpportunitiesPutResponse_custom_fieldsable)
				}
			}
			m.SetCustomFields(res)
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
	res["estimated_close_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEstimatedCloseDate(val)
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
	res["include_in_forecast"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIncludeInForecast(val)
		}
		return nil
	}
	res["last_updated"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastUpdated(val)
		}
		return nil
	}
	res["next_action_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNextActionDate(val)
		}
		return nil
	}
	res["next_action_notes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNextActionNotes(val)
		}
		return nil
	}
	res["opportunity_notes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOpportunityNotes(val)
		}
		return nil
	}
	res["opportunity_title"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOpportunityTitle(val)
		}
		return nil
	}
	res["projected_revenue_high"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProjectedRevenueHigh(val)
		}
		return nil
	}
	res["projected_revenue_low"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProjectedRevenueLow(val)
		}
		return nil
	}
	res["stage"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOpportunitiesPutResponse_stageFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStage(val.(OpportunitiesPutResponse_stageable))
		}
		return nil
	}
	res["user"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOpportunitiesPutResponse_userFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUser(val.(OpportunitiesPutResponse_userable))
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OpportunitiesPutResponse) GetId() *int64 {
	return m.id
}

// GetIncludeInForecast gets the include_in_forecast property value. The include_in_forecast property
// returns a *int32 when successful
func (m *OpportunitiesPutResponse) GetIncludeInForecast() *int32 {
	return m.include_in_forecast
}

// GetLastUpdated gets the last_updated property value. The last_updated property
// returns a *Time when successful
func (m *OpportunitiesPutResponse) GetLastUpdated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.last_updated
}

// GetNextActionDate gets the next_action_date property value. The next_action_date property
// returns a *Time when successful
func (m *OpportunitiesPutResponse) GetNextActionDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.next_action_date
}

// GetNextActionNotes gets the next_action_notes property value. The next_action_notes property
// returns a *string when successful
func (m *OpportunitiesPutResponse) GetNextActionNotes() *string {
	return m.next_action_notes
}

// GetOpportunityNotes gets the opportunity_notes property value. The opportunity_notes property
// returns a *string when successful
func (m *OpportunitiesPutResponse) GetOpportunityNotes() *string {
	return m.opportunity_notes
}

// GetOpportunityTitle gets the opportunity_title property value. The opportunity_title property
// returns a *string when successful
func (m *OpportunitiesPutResponse) GetOpportunityTitle() *string {
	return m.opportunity_title
}

// GetProjectedRevenueHigh gets the projected_revenue_high property value. The projected_revenue_high property
// returns a *float64 when successful
func (m *OpportunitiesPutResponse) GetProjectedRevenueHigh() *float64 {
	return m.projected_revenue_high
}

// GetProjectedRevenueLow gets the projected_revenue_low property value. The projected_revenue_low property
// returns a *float64 when successful
func (m *OpportunitiesPutResponse) GetProjectedRevenueLow() *float64 {
	return m.projected_revenue_low
}

// GetStage gets the stage property value. The stage property
// returns a OpportunitiesPutResponse_stageable when successful
func (m *OpportunitiesPutResponse) GetStage() OpportunitiesPutResponse_stageable {
	return m.stage
}

// GetUser gets the user property value. The user property
// returns a OpportunitiesPutResponse_userable when successful
func (m *OpportunitiesPutResponse) GetUser() OpportunitiesPutResponse_userable {
	return m.user
}

// Serialize serializes information the current object
func (m *OpportunitiesPutResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("affiliate_id", m.GetAffiliateId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("contact", m.GetContact())
		if err != nil {
			return err
		}
	}
	if m.GetCustomFields() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomFields()))
		for i, v := range m.GetCustomFields() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("custom_fields", cast)
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
		err := writer.WriteTimeValue("estimated_close_date", m.GetEstimatedCloseDate())
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
		err := writer.WriteInt32Value("include_in_forecast", m.GetIncludeInForecast())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("last_updated", m.GetLastUpdated())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("next_action_date", m.GetNextActionDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("next_action_notes", m.GetNextActionNotes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("opportunity_notes", m.GetOpportunityNotes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("opportunity_title", m.GetOpportunityTitle())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("projected_revenue_high", m.GetProjectedRevenueHigh())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("projected_revenue_low", m.GetProjectedRevenueLow())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("stage", m.GetStage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *OpportunitiesPutResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAffiliateId sets the affiliate_id property value. The affiliate_id property
func (m *OpportunitiesPutResponse) SetAffiliateId(value *int64) {
	m.affiliate_id = value
}

// SetContact sets the contact property value. The contact property
func (m *OpportunitiesPutResponse) SetContact(value OpportunitiesPutResponse_contactable) {
	m.contact = value
}

// SetCustomFields sets the custom_fields property value. The custom_fields property
func (m *OpportunitiesPutResponse) SetCustomFields(value []OpportunitiesPutResponse_custom_fieldsable) {
	m.custom_fields = value
}

// SetDateCreated sets the date_created property value. The date_created property
func (m *OpportunitiesPutResponse) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.date_created = value
}

// SetEstimatedCloseDate sets the estimated_close_date property value. The estimated_close_date property
func (m *OpportunitiesPutResponse) SetEstimatedCloseDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.estimated_close_date = value
}

// SetId sets the id property value. The id property
func (m *OpportunitiesPutResponse) SetId(value *int64) {
	m.id = value
}

// SetIncludeInForecast sets the include_in_forecast property value. The include_in_forecast property
func (m *OpportunitiesPutResponse) SetIncludeInForecast(value *int32) {
	m.include_in_forecast = value
}

// SetLastUpdated sets the last_updated property value. The last_updated property
func (m *OpportunitiesPutResponse) SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.last_updated = value
}

// SetNextActionDate sets the next_action_date property value. The next_action_date property
func (m *OpportunitiesPutResponse) SetNextActionDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.next_action_date = value
}

// SetNextActionNotes sets the next_action_notes property value. The next_action_notes property
func (m *OpportunitiesPutResponse) SetNextActionNotes(value *string) {
	m.next_action_notes = value
}

// SetOpportunityNotes sets the opportunity_notes property value. The opportunity_notes property
func (m *OpportunitiesPutResponse) SetOpportunityNotes(value *string) {
	m.opportunity_notes = value
}

// SetOpportunityTitle sets the opportunity_title property value. The opportunity_title property
func (m *OpportunitiesPutResponse) SetOpportunityTitle(value *string) {
	m.opportunity_title = value
}

// SetProjectedRevenueHigh sets the projected_revenue_high property value. The projected_revenue_high property
func (m *OpportunitiesPutResponse) SetProjectedRevenueHigh(value *float64) {
	m.projected_revenue_high = value
}

// SetProjectedRevenueLow sets the projected_revenue_low property value. The projected_revenue_low property
func (m *OpportunitiesPutResponse) SetProjectedRevenueLow(value *float64) {
	m.projected_revenue_low = value
}

// SetStage sets the stage property value. The stage property
func (m *OpportunitiesPutResponse) SetStage(value OpportunitiesPutResponse_stageable) {
	m.stage = value
}

// SetUser sets the user property value. The user property
func (m *OpportunitiesPutResponse) SetUser(value OpportunitiesPutResponse_userable) {
	m.user = value
}

type OpportunitiesPutResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAffiliateId() *int64
	GetContact() OpportunitiesPutResponse_contactable
	GetCustomFields() []OpportunitiesPutResponse_custom_fieldsable
	GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetEstimatedCloseDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetId() *int64
	GetIncludeInForecast() *int32
	GetLastUpdated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetNextActionDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetNextActionNotes() *string
	GetOpportunityNotes() *string
	GetOpportunityTitle() *string
	GetProjectedRevenueHigh() *float64
	GetProjectedRevenueLow() *float64
	GetStage() OpportunitiesPutResponse_stageable
	GetUser() OpportunitiesPutResponse_userable
	SetAffiliateId(value *int64)
	SetContact(value OpportunitiesPutResponse_contactable)
	SetCustomFields(value []OpportunitiesPutResponse_custom_fieldsable)
	SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetEstimatedCloseDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetId(value *int64)
	SetIncludeInForecast(value *int32)
	SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetNextActionDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetNextActionNotes(value *string)
	SetOpportunityNotes(value *string)
	SetOpportunityTitle(value *string)
	SetProjectedRevenueHigh(value *float64)
	SetProjectedRevenueLow(value *float64)
	SetStage(value OpportunitiesPutResponse_stageable)
	SetUser(value OpportunitiesPutResponse_userable)
}
