package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemUtmPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The formId
	keapSourceId *string
	// UTM campaign information
	utmCampaign *string
	// UTM content information
	utmContent *string
	// UTM medium information
	utmMedium *string
	// UTM source information
	utmSource *string
	// UTM term information
	utmTerm *string
}

// NewContactsItemUtmPostRequestBody instantiates a new ContactsItemUtmPostRequestBody and sets the default values.
func NewContactsItemUtmPostRequestBody() *ContactsItemUtmPostRequestBody {
	m := &ContactsItemUtmPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemUtmPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemUtmPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemUtmPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemUtmPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemUtmPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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

// GetKeapSourceId gets the keapSourceId property value. The formId
// returns a *string when successful
func (m *ContactsItemUtmPostRequestBody) GetKeapSourceId() *string {
	return m.keapSourceId
}

// GetUtmCampaign gets the utmCampaign property value. UTM campaign information
// returns a *string when successful
func (m *ContactsItemUtmPostRequestBody) GetUtmCampaign() *string {
	return m.utmCampaign
}

// GetUtmContent gets the utmContent property value. UTM content information
// returns a *string when successful
func (m *ContactsItemUtmPostRequestBody) GetUtmContent() *string {
	return m.utmContent
}

// GetUtmMedium gets the utmMedium property value. UTM medium information
// returns a *string when successful
func (m *ContactsItemUtmPostRequestBody) GetUtmMedium() *string {
	return m.utmMedium
}

// GetUtmSource gets the utmSource property value. UTM source information
// returns a *string when successful
func (m *ContactsItemUtmPostRequestBody) GetUtmSource() *string {
	return m.utmSource
}

// GetUtmTerm gets the utmTerm property value. UTM term information
// returns a *string when successful
func (m *ContactsItemUtmPostRequestBody) GetUtmTerm() *string {
	return m.utmTerm
}

// Serialize serializes information the current object
func (m *ContactsItemUtmPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("keapSourceId", m.GetKeapSourceId())
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
func (m *ContactsItemUtmPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetKeapSourceId sets the keapSourceId property value. The formId
func (m *ContactsItemUtmPostRequestBody) SetKeapSourceId(value *string) {
	m.keapSourceId = value
}

// SetUtmCampaign sets the utmCampaign property value. UTM campaign information
func (m *ContactsItemUtmPostRequestBody) SetUtmCampaign(value *string) {
	m.utmCampaign = value
}

// SetUtmContent sets the utmContent property value. UTM content information
func (m *ContactsItemUtmPostRequestBody) SetUtmContent(value *string) {
	m.utmContent = value
}

// SetUtmMedium sets the utmMedium property value. UTM medium information
func (m *ContactsItemUtmPostRequestBody) SetUtmMedium(value *string) {
	m.utmMedium = value
}

// SetUtmSource sets the utmSource property value. UTM source information
func (m *ContactsItemUtmPostRequestBody) SetUtmSource(value *string) {
	m.utmSource = value
}

// SetUtmTerm sets the utmTerm property value. UTM term information
func (m *ContactsItemUtmPostRequestBody) SetUtmTerm(value *string) {
	m.utmTerm = value
}

type ContactsItemUtmPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetKeapSourceId() *string
	GetUtmCampaign() *string
	GetUtmContent() *string
	GetUtmMedium() *string
	GetUtmSource() *string
	GetUtmTerm() *string
	SetKeapSourceId(value *string)
	SetUtmCampaign(value *string)
	SetUtmContent(value *string)
	SetUtmMedium(value *string)
	SetUtmSource(value *string)
	SetUtmTerm(value *string)
}
