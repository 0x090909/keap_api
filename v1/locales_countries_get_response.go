package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LocalesCountriesGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// A key-value pair of country code and country name.
	countries LocalesCountriesGetResponse_countriesable
}

// NewLocalesCountriesGetResponse instantiates a new LocalesCountriesGetResponse and sets the default values.
func NewLocalesCountriesGetResponse() *LocalesCountriesGetResponse {
	m := &LocalesCountriesGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLocalesCountriesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLocalesCountriesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLocalesCountriesGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LocalesCountriesGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCountries gets the countries property value. A key-value pair of country code and country name.
// returns a LocalesCountriesGetResponse_countriesable when successful
func (m *LocalesCountriesGetResponse) GetCountries() LocalesCountriesGetResponse_countriesable {
	return m.countries
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LocalesCountriesGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["countries"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateLocalesCountriesGetResponse_countriesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCountries(val.(LocalesCountriesGetResponse_countriesable))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *LocalesCountriesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("countries", m.GetCountries())
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
func (m *LocalesCountriesGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCountries sets the countries property value. A key-value pair of country code and country name.
func (m *LocalesCountriesGetResponse) SetCountries(value LocalesCountriesGetResponse_countriesable) {
	m.countries = value
}

type LocalesCountriesGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCountries() LocalesCountriesGetResponse_countriesable
	SetCountries(value LocalesCountriesGetResponse_countriesable)
}
