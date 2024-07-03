package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailsBatchRemovePostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The results property
	results EmailsBatchRemovePostResponse_resultsable
}

// NewEmailsBatchRemovePostResponse instantiates a new EmailsBatchRemovePostResponse and sets the default values.
func NewEmailsBatchRemovePostResponse() *EmailsBatchRemovePostResponse {
	m := &EmailsBatchRemovePostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailsBatchRemovePostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsBatchRemovePostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsBatchRemovePostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailsBatchRemovePostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsBatchRemovePostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["results"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEmailsBatchRemovePostResponse_resultsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetResults(val.(EmailsBatchRemovePostResponse_resultsable))
		}
		return nil
	}
	return res
}

// GetResults gets the results property value. The results property
// returns a EmailsBatchRemovePostResponse_resultsable when successful
func (m *EmailsBatchRemovePostResponse) GetResults() EmailsBatchRemovePostResponse_resultsable {
	return m.results
}

// Serialize serializes information the current object
func (m *EmailsBatchRemovePostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("results", m.GetResults())
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
func (m *EmailsBatchRemovePostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetResults sets the results property value. The results property
func (m *EmailsBatchRemovePostResponse) SetResults(value EmailsBatchRemovePostResponse_resultsable) {
	m.results = value
}

type EmailsBatchRemovePostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetResults() EmailsBatchRemovePostResponse_resultsable
	SetResults(value EmailsBatchRemovePostResponse_resultsable)
}
