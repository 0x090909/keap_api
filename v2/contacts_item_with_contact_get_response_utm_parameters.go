package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type ContactsItemWithContact_GetResponse_utm_parameters struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The dateCreated property
	dateCreated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The firstTouch property
	firstTouch *bool
	// The id property
	id *int64
	// The keapSourceId property
	keapSourceId *string
	// The lastTouch property
	lastTouch *bool
	// The utmCampaign property
	utmCampaign *string
	// The utmContent property
	utmContent *string
	// The utmMedium property
	utmMedium *string
	// The utmSource property
	utmSource *string
	// The utmTerm property
	utmTerm *string
}

// NewContactsItemWithContact_GetResponse_utm_parameters instantiates a new ContactsItemWithContact_GetResponse_utm_parameters and sets the default values.
func NewContactsItemWithContact_GetResponse_utm_parameters() *ContactsItemWithContact_GetResponse_utm_parameters {
	m := &ContactsItemWithContact_GetResponse_utm_parameters{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemWithContact_GetResponse_utm_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemWithContact_GetResponse_utm_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemWithContact_GetResponse_utm_parameters(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemWithContact_GetResponse_utm_parameters) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDateCreated gets the dateCreated property value. The dateCreated property
// returns a *Time when successful
func (m *ContactsItemWithContact_GetResponse_utm_parameters) GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.dateCreated
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemWithContact_GetResponse_utm_parameters) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["dateCreated"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDateCreated(val)
		}
		return nil
	}
	res["firstTouch"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFirstTouch(val)
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
	res["keapSourceId"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKeapSourceId(val)
		}
		return nil
	}
	res["lastTouch"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastTouch(val)
		}
		return nil
	}
	res["utmCampaign"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUtmCampaign(val)
		}
		return nil
	}
	res["utmContent"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUtmContent(val)
		}
		return nil
	}
	res["utmMedium"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUtmMedium(val)
		}
		return nil
	}
	res["utmSource"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUtmSource(val)
		}
		return nil
	}
	res["utmTerm"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUtmTerm(val)
		}
		return nil
	}
	return res
}

// GetFirstTouch gets the firstTouch property value. The firstTouch property
// returns a *bool when successful
func (m *ContactsItemWithContact_GetResponse_utm_parameters) GetFirstTouch() *bool {
	return m.firstTouch
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *ContactsItemWithContact_GetResponse_utm_parameters) GetId() *int64 {
	return m.id
}

// GetKeapSourceId gets the keapSourceId property value. The keapSourceId property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse_utm_parameters) GetKeapSourceId() *string {
	return m.keapSourceId
}

// GetLastTouch gets the lastTouch property value. The lastTouch property
// returns a *bool when successful
func (m *ContactsItemWithContact_GetResponse_utm_parameters) GetLastTouch() *bool {
	return m.lastTouch
}

// GetUtmCampaign gets the utmCampaign property value. The utmCampaign property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse_utm_parameters) GetUtmCampaign() *string {
	return m.utmCampaign
}

// GetUtmContent gets the utmContent property value. The utmContent property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse_utm_parameters) GetUtmContent() *string {
	return m.utmContent
}

// GetUtmMedium gets the utmMedium property value. The utmMedium property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse_utm_parameters) GetUtmMedium() *string {
	return m.utmMedium
}

// GetUtmSource gets the utmSource property value. The utmSource property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse_utm_parameters) GetUtmSource() *string {
	return m.utmSource
}

// GetUtmTerm gets the utmTerm property value. The utmTerm property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse_utm_parameters) GetUtmTerm() *string {
	return m.utmTerm
}

// Serialize serializes information the current object
func (m *ContactsItemWithContact_GetResponse_utm_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteTimeValue("dateCreated", m.GetDateCreated())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("firstTouch", m.GetFirstTouch())
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
		err := writer.WriteStringValue("keapSourceId", m.GetKeapSourceId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("lastTouch", m.GetLastTouch())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("utmCampaign", m.GetUtmCampaign())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("utmContent", m.GetUtmContent())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("utmMedium", m.GetUtmMedium())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("utmSource", m.GetUtmSource())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("utmTerm", m.GetUtmTerm())
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
func (m *ContactsItemWithContact_GetResponse_utm_parameters) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDateCreated sets the dateCreated property value. The dateCreated property
func (m *ContactsItemWithContact_GetResponse_utm_parameters) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.dateCreated = value
}

// SetFirstTouch sets the firstTouch property value. The firstTouch property
func (m *ContactsItemWithContact_GetResponse_utm_parameters) SetFirstTouch(value *bool) {
	m.firstTouch = value
}

// SetId sets the id property value. The id property
func (m *ContactsItemWithContact_GetResponse_utm_parameters) SetId(value *int64) {
	m.id = value
}

// SetKeapSourceId sets the keapSourceId property value. The keapSourceId property
func (m *ContactsItemWithContact_GetResponse_utm_parameters) SetKeapSourceId(value *string) {
	m.keapSourceId = value
}

// SetLastTouch sets the lastTouch property value. The lastTouch property
func (m *ContactsItemWithContact_GetResponse_utm_parameters) SetLastTouch(value *bool) {
	m.lastTouch = value
}

// SetUtmCampaign sets the utmCampaign property value. The utmCampaign property
func (m *ContactsItemWithContact_GetResponse_utm_parameters) SetUtmCampaign(value *string) {
	m.utmCampaign = value
}

// SetUtmContent sets the utmContent property value. The utmContent property
func (m *ContactsItemWithContact_GetResponse_utm_parameters) SetUtmContent(value *string) {
	m.utmContent = value
}

// SetUtmMedium sets the utmMedium property value. The utmMedium property
func (m *ContactsItemWithContact_GetResponse_utm_parameters) SetUtmMedium(value *string) {
	m.utmMedium = value
}

// SetUtmSource sets the utmSource property value. The utmSource property
func (m *ContactsItemWithContact_GetResponse_utm_parameters) SetUtmSource(value *string) {
	m.utmSource = value
}

// SetUtmTerm sets the utmTerm property value. The utmTerm property
func (m *ContactsItemWithContact_GetResponse_utm_parameters) SetUtmTerm(value *string) {
	m.utmTerm = value
}

type ContactsItemWithContact_GetResponse_utm_parametersable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetFirstTouch() *bool
	GetId() *int64
	GetKeapSourceId() *string
	GetLastTouch() *bool
	GetUtmCampaign() *string
	GetUtmContent() *string
	GetUtmMedium() *string
	GetUtmSource() *string
	GetUtmTerm() *string
	SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetFirstTouch(value *bool)
	SetId(value *int64)
	SetKeapSourceId(value *string)
	SetLastTouch(value *bool)
	SetUtmCampaign(value *string)
	SetUtmContent(value *string)
	SetUtmMedium(value *string)
	SetUtmSource(value *string)
	SetUtmTerm(value *string)
}
