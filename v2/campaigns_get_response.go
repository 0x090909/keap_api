package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The campaigns property
	campaigns []CampaignsGetResponse_campaignsable
	// The next_page_token property
	next_page_token *string
}

// NewCampaignsGetResponse instantiates a new CampaignsGetResponse and sets the default values.
func NewCampaignsGetResponse() *CampaignsGetResponse {
	m := &CampaignsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCampaigns gets the campaigns property value. The campaigns property
// returns a []CampaignsGetResponse_campaignsable when successful
func (m *CampaignsGetResponse) GetCampaigns() []CampaignsGetResponse_campaignsable {
	return m.campaigns
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["campaigns"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCampaignsGetResponse_campaignsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CampaignsGetResponse_campaignsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CampaignsGetResponse_campaignsable)
				}
			}
			m.SetCampaigns(res)
		}
		return nil
	}
	res["next_page_token"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNextPageToken(val)
		}
		return nil
	}
	return res
}

// GetNextPageToken gets the next_page_token property value. The next_page_token property
// returns a *string when successful
func (m *CampaignsGetResponse) GetNextPageToken() *string {
	return m.next_page_token
}

// Serialize serializes information the current object
func (m *CampaignsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetCampaigns() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCampaigns()))
		for i, v := range m.GetCampaigns() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("campaigns", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("next_page_token", m.GetNextPageToken())
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
func (m *CampaignsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCampaigns sets the campaigns property value. The campaigns property
func (m *CampaignsGetResponse) SetCampaigns(value []CampaignsGetResponse_campaignsable) {
	m.campaigns = value
}

// SetNextPageToken sets the next_page_token property value. The next_page_token property
func (m *CampaignsGetResponse) SetNextPageToken(value *string) {
	m.next_page_token = value
}

type CampaignsGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCampaigns() []CampaignsGetResponse_campaignsable
	GetNextPageToken() *string
	SetCampaigns(value []CampaignsGetResponse_campaignsable)
	SetNextPageToken(value *string)
}
