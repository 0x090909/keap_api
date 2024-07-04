package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LocalesCountriesItemProvincesGetResponse_provinces a key-value pair of province code and province name.
type LocalesCountriesItemProvincesGetResponse_provinces struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
}

// NewLocalesCountriesItemProvincesGetResponse_provinces instantiates a new LocalesCountriesItemProvincesGetResponse_provinces and sets the default values.
func NewLocalesCountriesItemProvincesGetResponse_provinces() *LocalesCountriesItemProvincesGetResponse_provinces {
	m := &LocalesCountriesItemProvincesGetResponse_provinces{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLocalesCountriesItemProvincesGetResponse_provincesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLocalesCountriesItemProvincesGetResponse_provincesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLocalesCountriesItemProvincesGetResponse_provinces(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LocalesCountriesItemProvincesGetResponse_provinces) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LocalesCountriesItemProvincesGetResponse_provinces) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	return res
}

// Serialize serializes information the current object
func (m *LocalesCountriesItemProvincesGetResponse_provinces) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LocalesCountriesItemProvincesGetResponse_provinces) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

type LocalesCountriesItemProvincesGetResponse_provincesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
