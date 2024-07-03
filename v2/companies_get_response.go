package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CompaniesGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The companies property
	companies []CompaniesGetResponse_companiesable
	// The next_page_token property
	next_page_token *string
}

// NewCompaniesGetResponse instantiates a new CompaniesGetResponse and sets the default values.
func NewCompaniesGetResponse() *CompaniesGetResponse {
	m := &CompaniesGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCompaniesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCompaniesGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CompaniesGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCompanies gets the companies property value. The companies property
// returns a []CompaniesGetResponse_companiesable when successful
func (m *CompaniesGetResponse) GetCompanies() []CompaniesGetResponse_companiesable {
	return m.companies
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompaniesGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["companies"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCompaniesGetResponse_companiesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CompaniesGetResponse_companiesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CompaniesGetResponse_companiesable)
				}
			}
			m.SetCompanies(res)
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
func (m *CompaniesGetResponse) GetNextPageToken() *string {
	return m.next_page_token
}

// Serialize serializes information the current object
func (m *CompaniesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetCompanies() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompanies()))
		for i, v := range m.GetCompanies() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("companies", cast)
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
func (m *CompaniesGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCompanies sets the companies property value. The companies property
func (m *CompaniesGetResponse) SetCompanies(value []CompaniesGetResponse_companiesable) {
	m.companies = value
}

// SetNextPageToken sets the next_page_token property value. The next_page_token property
func (m *CompaniesGetResponse) SetNextPageToken(value *string) {
	m.next_page_token = value
}

type CompaniesGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCompanies() []CompaniesGetResponse_companiesable
	GetNextPageToken() *string
	SetCompanies(value []CompaniesGetResponse_companiesable)
	SetNextPageToken(value *string)
}
