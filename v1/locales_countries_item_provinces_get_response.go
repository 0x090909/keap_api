package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LocalesCountriesItemProvincesGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// A key-value pair of province code and province name.
	provinces LocalesCountriesItemProvincesGetResponse_provincesable
}

// NewLocalesCountriesItemProvincesGetResponse instantiates a new LocalesCountriesItemProvincesGetResponse and sets the default values.
func NewLocalesCountriesItemProvincesGetResponse() *LocalesCountriesItemProvincesGetResponse {
	m := &LocalesCountriesItemProvincesGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLocalesCountriesItemProvincesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLocalesCountriesItemProvincesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLocalesCountriesItemProvincesGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LocalesCountriesItemProvincesGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LocalesCountriesItemProvincesGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["provinces"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLocalesCountriesItemProvincesGetResponse_provincesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProvinces(val.(LocalesCountriesItemProvincesGetResponse_provincesable))
		}
		return nil
	}
	return res
}

// GetProvinces gets the provinces property value. A key-value pair of province code and province name.
// returns a LocalesCountriesItemProvincesGetResponse_provincesable when successful
func (m *LocalesCountriesItemProvincesGetResponse) GetProvinces() LocalesCountriesItemProvincesGetResponse_provincesable {
	return m.provinces
}

// Serialize serializes information the current object
func (m *LocalesCountriesItemProvincesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("provinces", m.GetProvinces())
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
func (m *LocalesCountriesItemProvincesGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetProvinces sets the provinces property value. A key-value pair of province code and province name.
func (m *LocalesCountriesItemProvincesGetResponse) SetProvinces(value LocalesCountriesItemProvincesGetResponse_provincesable) {
	m.provinces = value
}

type LocalesCountriesItemProvincesGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetProvinces() LocalesCountriesItemProvincesGetResponse_provincesable
	SetProvinces(value LocalesCountriesItemProvincesGetResponse_provincesable)
}
