package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsItemSequencesItemContactsItemWithContact401Error struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The cause property
	cause CampaignsItemSequencesItemContactsItemWithContact401Error_causeable
	// The localizedMessage property
	localizedMessage *string
	// The message property
	message *string
	// The stackTrace property
	stackTrace []CampaignsItemSequencesItemContactsItemWithContact401Error_stackTraceable
	// The suppressed property
	suppressed []CampaignsItemSequencesItemContactsItemWithContact401Error_suppressedable
}

// NewCampaignsItemSequencesItemContactsItemWithContact401Error instantiates a new CampaignsItemSequencesItemContactsItemWithContact401Error and sets the default values.
func NewCampaignsItemSequencesItemContactsItemWithContact401Error() *CampaignsItemSequencesItemContactsItemWithContact401Error {
	m := &CampaignsItemSequencesItemContactsItemWithContact401Error{
		ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
	}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsItemSequencesItemContactsItemWithContact401ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemSequencesItemContactsItemWithContact401ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemSequencesItemContactsItemWithContact401Error(), nil
}

// Error the primary error message.
// returns a string when successful
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) Error() string {
	return m.ApiError.Error()
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCause gets the cause property value. The cause property
// returns a CampaignsItemSequencesItemContactsItemWithContact401Error_causeable when successful
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) GetCause() CampaignsItemSequencesItemContactsItemWithContact401Error_causeable {
	return m.cause
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["cause"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCampaignsItemSequencesItemContactsItemWithContact401Error_causeFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCause(val.(CampaignsItemSequencesItemContactsItemWithContact401Error_causeable))
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
		val, err := n.GetCollectionOfObjectValues(CreateCampaignsItemSequencesItemContactsItemWithContact401Error_stackTraceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CampaignsItemSequencesItemContactsItemWithContact401Error_stackTraceable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CampaignsItemSequencesItemContactsItemWithContact401Error_stackTraceable)
				}
			}
			m.SetStackTrace(res)
		}
		return nil
	}
	res["suppressed"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCampaignsItemSequencesItemContactsItemWithContact401Error_suppressedFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CampaignsItemSequencesItemContactsItemWithContact401Error_suppressedable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CampaignsItemSequencesItemContactsItemWithContact401Error_suppressedable)
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
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) GetLocalizedMessage() *string {
	return m.localizedMessage
}

// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) GetMessage() *string {
	return m.message
}

// GetStackTrace gets the stackTrace property value. The stackTrace property
// returns a []CampaignsItemSequencesItemContactsItemWithContact401Error_stackTraceable when successful
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) GetStackTrace() []CampaignsItemSequencesItemContactsItemWithContact401Error_stackTraceable {
	return m.stackTrace
}

// GetSuppressed gets the suppressed property value. The suppressed property
// returns a []CampaignsItemSequencesItemContactsItemWithContact401Error_suppressedable when successful
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) GetSuppressed() []CampaignsItemSequencesItemContactsItemWithContact401Error_suppressedable {
	return m.suppressed
}

// Serialize serializes information the current object
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCause sets the cause property value. The cause property
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) SetCause(value CampaignsItemSequencesItemContactsItemWithContact401Error_causeable) {
	m.cause = value
}

// SetLocalizedMessage sets the localizedMessage property value. The localizedMessage property
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) SetLocalizedMessage(value *string) {
	m.localizedMessage = value
}

// SetMessage sets the message property value. The message property
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) SetMessage(value *string) {
	m.message = value
}

// SetStackTrace sets the stackTrace property value. The stackTrace property
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) SetStackTrace(value []CampaignsItemSequencesItemContactsItemWithContact401Error_stackTraceable) {
	m.stackTrace = value
}

// SetSuppressed sets the suppressed property value. The suppressed property
func (m *CampaignsItemSequencesItemContactsItemWithContact401Error) SetSuppressed(value []CampaignsItemSequencesItemContactsItemWithContact401Error_suppressedable) {
	m.suppressed = value
}

type CampaignsItemSequencesItemContactsItemWithContact401Errorable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCause() CampaignsItemSequencesItemContactsItemWithContact401Error_causeable
	GetLocalizedMessage() *string
	GetMessage() *string
	GetStackTrace() []CampaignsItemSequencesItemContactsItemWithContact401Error_stackTraceable
	GetSuppressed() []CampaignsItemSequencesItemContactsItemWithContact401Error_suppressedable
	SetCause(value CampaignsItemSequencesItemContactsItemWithContact401Error_causeable)
	SetLocalizedMessage(value *string)
	SetMessage(value *string)
	SetStackTrace(value []CampaignsItemSequencesItemContactsItemWithContact401Error_stackTraceable)
	SetSuppressed(value []CampaignsItemSequencesItemContactsItemWithContact401Error_suppressedable)
}
