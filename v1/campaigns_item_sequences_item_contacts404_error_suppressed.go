package v1

import (
	ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff "github.com/0x090909/keap_api/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsItemSequencesItemContacts404Error_suppressed struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The cause property
	cause ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable
	// The localizedMessage property
	localizedMessage *string
	// The message property
	message *string
	// The stackTrace property
	stackTrace []CampaignsItemSequencesItemContacts404Error_suppressed_stackTraceable
	// The suppressed property
	suppressed []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable
}

// NewCampaignsItemSequencesItemContacts404Error_suppressed instantiates a new CampaignsItemSequencesItemContacts404Error_suppressed and sets the default values.
func NewCampaignsItemSequencesItemContacts404Error_suppressed() *CampaignsItemSequencesItemContacts404Error_suppressed {
	m := &CampaignsItemSequencesItemContacts404Error_suppressed{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsItemSequencesItemContacts404Error_suppressedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemSequencesItemContacts404Error_suppressedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemSequencesItemContacts404Error_suppressed(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCause gets the cause property value. The cause property
// returns a Throwableable when successful
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) GetCause() ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable {
	return m.cause
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["cause"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.CreateThrowableFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCause(val.(ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable))
		}
		return nil
	}
	res["localizedMessage"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLocalizedMessage(val)
		}
		return nil
	}
	res["message"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMessage(val)
		}
		return nil
	}
	res["stackTrace"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCampaignsItemSequencesItemContacts404Error_suppressed_stackTraceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CampaignsItemSequencesItemContacts404Error_suppressed_stackTraceable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CampaignsItemSequencesItemContacts404Error_suppressed_stackTraceable)
				}
			}
			m.SetStackTrace(res)
		}
		return nil
	}
	res["suppressed"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.CreateThrowableFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable)
				}
			}
			m.SetSuppressed(res)
		}
		return nil
	}
	return res
}

// GetLocalizedMessage gets the localizedMessage property value. The localizedMessage property
// returns a *string when successful
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) GetLocalizedMessage() *string {
	return m.localizedMessage
}

// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) GetMessage() *string {
	return m.message
}

// GetStackTrace gets the stackTrace property value. The stackTrace property
// returns a []CampaignsItemSequencesItemContacts404Error_suppressed_stackTraceable when successful
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) GetStackTrace() []CampaignsItemSequencesItemContacts404Error_suppressed_stackTraceable {
	return m.stackTrace
}

// GetSuppressed gets the suppressed property value. The suppressed property
// returns a []Throwableable when successful
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) GetSuppressed() []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable {
	return m.suppressed
}

// Serialize serializes information the current object
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("cause", m.GetCause())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("localizedMessage", m.GetLocalizedMessage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("message", m.GetMessage())
		if err != nil {
			return err
		}
	}
	if m.GetStackTrace() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStackTrace()))
		for i, v := range m.GetStackTrace() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("stackTrace", cast)
		if err != nil {
			return err
		}
	}
	if m.GetSuppressed() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSuppressed()))
		for i, v := range m.GetSuppressed() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("suppressed", cast)
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
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCause sets the cause property value. The cause property
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) SetCause(value ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable) {
	m.cause = value
}

// SetLocalizedMessage sets the localizedMessage property value. The localizedMessage property
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) SetLocalizedMessage(value *string) {
	m.localizedMessage = value
}

// SetMessage sets the message property value. The message property
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) SetMessage(value *string) {
	m.message = value
}

// SetStackTrace sets the stackTrace property value. The stackTrace property
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) SetStackTrace(value []CampaignsItemSequencesItemContacts404Error_suppressed_stackTraceable) {
	m.stackTrace = value
}

// SetSuppressed sets the suppressed property value. The suppressed property
func (m *CampaignsItemSequencesItemContacts404Error_suppressed) SetSuppressed(value []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable) {
	m.suppressed = value
}

type CampaignsItemSequencesItemContacts404Error_suppressedable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCause() ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable
	GetLocalizedMessage() *string
	GetMessage() *string
	GetStackTrace() []CampaignsItemSequencesItemContacts404Error_suppressed_stackTraceable
	GetSuppressed() []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable
	SetCause(value ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable)
	SetLocalizedMessage(value *string)
	SetMessage(value *string)
	SetStackTrace(value []CampaignsItemSequencesItemContacts404Error_suppressed_stackTraceable)
	SetSuppressed(value []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable)
}
