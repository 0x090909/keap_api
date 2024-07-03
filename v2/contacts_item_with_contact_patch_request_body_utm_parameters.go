package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemWithContact_PatchRequestBody_utm_parameters struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The formId
	keap_source_id *string
	// UTM campaign information
	utm_campaign *string
	// UTM content information
	utm_content *string
	// UTM medium information
	utm_medium *string
	// UTM source information
	utm_source *string
	// UTM term information
	utm_term *string
}

// NewContactsItemWithContact_PatchRequestBody_utm_parameters instantiates a new ContactsItemWithContact_PatchRequestBody_utm_parameters and sets the default values.
func NewContactsItemWithContact_PatchRequestBody_utm_parameters() *ContactsItemWithContact_PatchRequestBody_utm_parameters {
	m := &ContactsItemWithContact_PatchRequestBody_utm_parameters{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemWithContact_PatchRequestBody_utm_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemWithContact_PatchRequestBody_utm_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemWithContact_PatchRequestBody_utm_parameters(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["keap_source_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKeapSourceId(val)
		}
		return nil
	}
	res["utm_campaign"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUtmCampaign(val)
		}
		return nil
	}
	res["utm_content"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUtmContent(val)
		}
		return nil
	}
	res["utm_medium"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUtmMedium(val)
		}
		return nil
	}
	res["utm_source"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUtmSource(val)
		}
		return nil
	}
	res["utm_term"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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

// GetKeapSourceId gets the keap_source_id property value. The formId
// returns a *string when successful
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) GetKeapSourceId() *string {
	return m.keap_source_id
}

// GetUtmCampaign gets the utm_campaign property value. UTM campaign information
// returns a *string when successful
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) GetUtmCampaign() *string {
	return m.utm_campaign
}

// GetUtmContent gets the utm_content property value. UTM content information
// returns a *string when successful
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) GetUtmContent() *string {
	return m.utm_content
}

// GetUtmMedium gets the utm_medium property value. UTM medium information
// returns a *string when successful
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) GetUtmMedium() *string {
	return m.utm_medium
}

// GetUtmSource gets the utm_source property value. UTM source information
// returns a *string when successful
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) GetUtmSource() *string {
	return m.utm_source
}

// GetUtmTerm gets the utm_term property value. UTM term information
// returns a *string when successful
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) GetUtmTerm() *string {
	return m.utm_term
}

// Serialize serializes information the current object
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("keap_source_id", m.GetKeapSourceId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("utm_campaign", m.GetUtmCampaign())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("utm_content", m.GetUtmContent())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("utm_medium", m.GetUtmMedium())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("utm_source", m.GetUtmSource())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("utm_term", m.GetUtmTerm())
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
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetKeapSourceId sets the keap_source_id property value. The formId
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) SetKeapSourceId(value *string) {
	m.keap_source_id = value
}

// SetUtmCampaign sets the utm_campaign property value. UTM campaign information
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) SetUtmCampaign(value *string) {
	m.utm_campaign = value
}

// SetUtmContent sets the utm_content property value. UTM content information
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) SetUtmContent(value *string) {
	m.utm_content = value
}

// SetUtmMedium sets the utm_medium property value. UTM medium information
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) SetUtmMedium(value *string) {
	m.utm_medium = value
}

// SetUtmSource sets the utm_source property value. UTM source information
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) SetUtmSource(value *string) {
	m.utm_source = value
}

// SetUtmTerm sets the utm_term property value. UTM term information
func (m *ContactsItemWithContact_PatchRequestBody_utm_parameters) SetUtmTerm(value *string) {
	m.utm_term = value
}

type ContactsItemWithContact_PatchRequestBody_utm_parametersable interface {
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
