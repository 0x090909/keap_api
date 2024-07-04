package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsGetResponse_campaigns_sequences_paths struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *int64
	// The items property
	items []CampaignsGetResponse_campaigns_sequences_paths_itemsable
}

// NewCampaignsGetResponse_campaigns_sequences_paths instantiates a new CampaignsGetResponse_campaigns_sequences_paths and sets the default values.
func NewCampaignsGetResponse_campaigns_sequences_paths() *CampaignsGetResponse_campaigns_sequences_paths {
	m := &CampaignsGetResponse_campaigns_sequences_paths{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsGetResponse_campaigns_sequences_pathsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsGetResponse_campaigns_sequences_pathsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsGetResponse_campaigns_sequences_paths(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsGetResponse_campaigns_sequences_paths) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsGetResponse_campaigns_sequences_paths) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["items"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCampaignsGetResponse_campaigns_sequences_paths_itemsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CampaignsGetResponse_campaigns_sequences_paths_itemsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CampaignsGetResponse_campaigns_sequences_paths_itemsable)
				}
			}
			m.SetItems(res)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *CampaignsGetResponse_campaigns_sequences_paths) GetId() *int64 {
	return m.id
}

// GetItems gets the items property value. The items property
// returns a []CampaignsGetResponse_campaigns_sequences_paths_itemsable when successful
func (m *CampaignsGetResponse_campaigns_sequences_paths) GetItems() []CampaignsGetResponse_campaigns_sequences_paths_itemsable {
	return m.items
}

// Serialize serializes information the current object
func (m *CampaignsGetResponse_campaigns_sequences_paths) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	if m.GetItems() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
		for i, v := range m.GetItems() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("items", cast)
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
func (m *CampaignsGetResponse_campaigns_sequences_paths) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *CampaignsGetResponse_campaigns_sequences_paths) SetId(value *int64) {
	m.id = value
}

// SetItems sets the items property value. The items property
func (m *CampaignsGetResponse_campaigns_sequences_paths) SetItems(value []CampaignsGetResponse_campaigns_sequences_paths_itemsable) {
	m.items = value
}

type CampaignsGetResponse_campaigns_sequences_pathsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetId() *int64
	GetItems() []CampaignsGetResponse_campaigns_sequences_paths_itemsable
	SetId(value *int64)
	SetItems(value []CampaignsGetResponse_campaigns_sequences_paths_itemsable)
}
