package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProductsSyncGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The count property
	count *int32
	// The next property
	next *string
	// The previous property
	previous *string
	// The product_statuses property
	product_statuses []ProductsSyncGetResponse_product_statusesable
	// The sync_token property
	sync_token *string
}

// NewProductsSyncGetResponse instantiates a new ProductsSyncGetResponse and sets the default values.
func NewProductsSyncGetResponse() *ProductsSyncGetResponse {
	m := &ProductsSyncGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateProductsSyncGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProductsSyncGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewProductsSyncGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProductsSyncGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *ProductsSyncGetResponse) GetCount() *int32 {
	return m.count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProductsSyncGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCount(val)
		}
		return nil
	}
	res["next"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNext(val)
		}
		return nil
	}
	res["previous"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPrevious(val)
		}
		return nil
	}
	res["product_statuses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateProductsSyncGetResponse_product_statusesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ProductsSyncGetResponse_product_statusesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ProductsSyncGetResponse_product_statusesable)
				}
			}
			m.SetProductStatuses(res)
		}
		return nil
	}
	res["sync_token"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSyncToken(val)
		}
		return nil
	}
	return res
}

// GetNext gets the next property value. The next property
// returns a *string when successful
func (m *ProductsSyncGetResponse) GetNext() *string {
	return m.next
}

// GetPrevious gets the previous property value. The previous property
// returns a *string when successful
func (m *ProductsSyncGetResponse) GetPrevious() *string {
	return m.previous
}

// GetProductStatuses gets the product_statuses property value. The product_statuses property
// returns a []ProductsSyncGetResponse_product_statusesable when successful
func (m *ProductsSyncGetResponse) GetProductStatuses() []ProductsSyncGetResponse_product_statusesable {
	return m.product_statuses
}

// GetSyncToken gets the sync_token property value. The sync_token property
// returns a *string when successful
func (m *ProductsSyncGetResponse) GetSyncToken() *string {
	return m.sync_token
}

// Serialize serializes information the current object
func (m *ProductsSyncGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("count", m.GetCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("next", m.GetNext())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("previous", m.GetPrevious())
		if err != nil {
			return err
		}
	}
	if m.GetProductStatuses() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProductStatuses()))
		for i, v := range m.GetProductStatuses() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("product_statuses", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sync_token", m.GetSyncToken())
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
func (m *ProductsSyncGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCount sets the count property value. The count property
func (m *ProductsSyncGetResponse) SetCount(value *int32) {
	m.count = value
}

// SetNext sets the next property value. The next property
func (m *ProductsSyncGetResponse) SetNext(value *string) {
	m.next = value
}

// SetPrevious sets the previous property value. The previous property
func (m *ProductsSyncGetResponse) SetPrevious(value *string) {
	m.previous = value
}

// SetProductStatuses sets the product_statuses property value. The product_statuses property
func (m *ProductsSyncGetResponse) SetProductStatuses(value []ProductsSyncGetResponse_product_statusesable) {
	m.product_statuses = value
}

// SetSyncToken sets the sync_token property value. The sync_token property
func (m *ProductsSyncGetResponse) SetSyncToken(value *string) {
	m.sync_token = value
}

type ProductsSyncGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCount() *int32
	GetNext() *string
	GetPrevious() *string
	GetProductStatuses() []ProductsSyncGetResponse_product_statusesable
	GetSyncToken() *string
	SetCount(value *int32)
	SetNext(value *string)
	SetPrevious(value *string)
	SetProductStatuses(value []ProductsSyncGetResponse_product_statusesable)
	SetSyncToken(value *string)
}
