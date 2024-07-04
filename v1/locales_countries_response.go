package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use LocalesCountriesGetResponseable instead.
type LocalesCountriesResponse struct {
	LocalesCountriesGetResponse
}

// NewLocalesCountriesResponse instantiates a new LocalesCountriesResponse and sets the default values.
func NewLocalesCountriesResponse() *LocalesCountriesResponse {
	m := &LocalesCountriesResponse{
		LocalesCountriesGetResponse: *NewLocalesCountriesGetResponse(),
	}
	return m
}

// CreateLocalesCountriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLocalesCountriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLocalesCountriesResponse(), nil
}

// Deprecated: This class is obsolete. Use LocalesCountriesGetResponseable instead.
type LocalesCountriesResponseable interface {
	LocalesCountriesGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
