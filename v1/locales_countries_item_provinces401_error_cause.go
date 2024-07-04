package v1

import (
	ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff "github.com/0x090909/keap_api/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LocalesCountriesItemProvinces401Error_cause struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The cause property
	cause ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable
	// The localizedMessage property
	localizedMessage *string
	// The message property
	message *string
	// The stackTrace property
	stackTrace []LocalesCountriesItemProvinces401Error_cause_stackTraceable
	// The suppressed property
	suppressed []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable
}

// NewLocalesCountriesItemProvinces401Error_cause instantiates a new LocalesCountriesItemProvinces401Error_cause and sets the default values.
func NewLocalesCountriesItemProvinces401Error_cause() *LocalesCountriesItemProvinces401Error_cause {
	m := &LocalesCountriesItemProvinces401Error_cause{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLocalesCountriesItemProvinces401Error_causeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLocalesCountriesItemProvinces401Error_causeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLocalesCountriesItemProvinces401Error_cause(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LocalesCountriesItemProvinces401Error_cause) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCause gets the cause property value. The cause property
// returns a Throwableable when successful
func (m *LocalesCountriesItemProvinces401Error_cause) GetCause() ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable {
	return m.cause
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LocalesCountriesItemProvinces401Error_cause) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
		val, err := n.GetCollectionOfObjectValues(CreateLocalesCountriesItemProvinces401Error_cause_stackTraceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]LocalesCountriesItemProvinces401Error_cause_stackTraceable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(LocalesCountriesItemProvinces401Error_cause_stackTraceable)
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
func (m *LocalesCountriesItemProvinces401Error_cause) GetLocalizedMessage() *string {
	return m.localizedMessage
}

// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *LocalesCountriesItemProvinces401Error_cause) GetMessage() *string {
	return m.message
}

// GetStackTrace gets the stackTrace property value. The stackTrace property
// returns a []LocalesCountriesItemProvinces401Error_cause_stackTraceable when successful
func (m *LocalesCountriesItemProvinces401Error_cause) GetStackTrace() []LocalesCountriesItemProvinces401Error_cause_stackTraceable {
	return m.stackTrace
}

// GetSuppressed gets the suppressed property value. The suppressed property
// returns a []Throwableable when successful
func (m *LocalesCountriesItemProvinces401Error_cause) GetSuppressed() []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable {
	return m.suppressed
}

// Serialize serializes information the current object
func (m *LocalesCountriesItemProvinces401Error_cause) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *LocalesCountriesItemProvinces401Error_cause) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCause sets the cause property value. The cause property
func (m *LocalesCountriesItemProvinces401Error_cause) SetCause(value ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable) {
	m.cause = value
}

// SetLocalizedMessage sets the localizedMessage property value. The localizedMessage property
func (m *LocalesCountriesItemProvinces401Error_cause) SetLocalizedMessage(value *string) {
	m.localizedMessage = value
}

// SetMessage sets the message property value. The message property
func (m *LocalesCountriesItemProvinces401Error_cause) SetMessage(value *string) {
	m.message = value
}

// SetStackTrace sets the stackTrace property value. The stackTrace property
func (m *LocalesCountriesItemProvinces401Error_cause) SetStackTrace(value []LocalesCountriesItemProvinces401Error_cause_stackTraceable) {
	m.stackTrace = value
}

// SetSuppressed sets the suppressed property value. The suppressed property
func (m *LocalesCountriesItemProvinces401Error_cause) SetSuppressed(value []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable) {
	m.suppressed = value
}

type LocalesCountriesItemProvinces401Error_causeable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCause() ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable
	GetLocalizedMessage() *string
	GetMessage() *string
	GetStackTrace() []LocalesCountriesItemProvinces401Error_cause_stackTraceable
	GetSuppressed() []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable
	SetCause(value ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable)
	SetLocalizedMessage(value *string)
	SetMessage(value *string)
	SetStackTrace(value []LocalesCountriesItemProvinces401Error_cause_stackTraceable)
	SetSuppressed(value []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.Throwableable)
}
