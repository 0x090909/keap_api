package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsItemWithCampaign_GetResponse_sequences_paths struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *string
	// The items property
	items []CampaignsItemWithCampaign_GetResponse_sequences_paths_itemsable
}

// NewCampaignsItemWithCampaign_GetResponse_sequences_paths instantiates a new CampaignsItemWithCampaign_GetResponse_sequences_paths and sets the default values.
func NewCampaignsItemWithCampaign_GetResponse_sequences_paths() *CampaignsItemWithCampaign_GetResponse_sequences_paths {
	m := &CampaignsItemWithCampaign_GetResponse_sequences_paths{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsItemWithCampaign_GetResponse_sequences_pathsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemWithCampaign_GetResponse_sequences_pathsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemWithCampaign_GetResponse_sequences_paths(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsItemWithCampaign_GetResponse_sequences_paths) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsItemWithCampaign_GetResponse_sequences_paths) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["items"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCampaignsItemWithCampaign_GetResponse_sequences_paths_itemsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CampaignsItemWithCampaign_GetResponse_sequences_paths_itemsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CampaignsItemWithCampaign_GetResponse_sequences_paths_itemsable)
				}
			}
			m.SetItems(res)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *CampaignsItemWithCampaign_GetResponse_sequences_paths) GetId() *string {
	return m.id
}

// GetItems gets the items property value. The items property
// returns a []CampaignsItemWithCampaign_GetResponse_sequences_paths_itemsable when successful
func (m *CampaignsItemWithCampaign_GetResponse_sequences_paths) GetItems() []CampaignsItemWithCampaign_GetResponse_sequences_paths_itemsable {
	return m.items
}

// Serialize serializes information the current object
func (m *CampaignsItemWithCampaign_GetResponse_sequences_paths) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("id", m.GetId())
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
func (m *CampaignsItemWithCampaign_GetResponse_sequences_paths) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *CampaignsItemWithCampaign_GetResponse_sequences_paths) SetId(value *string) {
	m.id = value
}

// SetItems sets the items property value. The items property
func (m *CampaignsItemWithCampaign_GetResponse_sequences_paths) SetItems(value []CampaignsItemWithCampaign_GetResponse_sequences_paths_itemsable) {
	m.items = value
}

type CampaignsItemWithCampaign_GetResponse_sequences_pathsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetId() *string
	GetItems() []CampaignsItemWithCampaign_GetResponse_sequences_paths_itemsable
	SetId(value *string)
	SetItems(value []CampaignsItemWithCampaign_GetResponse_sequences_paths_itemsable)
}
