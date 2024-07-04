package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TransactionsItemWithTransactionGetResponseable instead.
type TransactionsItemWithTransactionResponse struct {
	TransactionsItemWithTransactionGetResponse
}

// NewTransactionsItemWithTransactionResponse instantiates a new TransactionsItemWithTransactionResponse and sets the default values.
func NewTransactionsItemWithTransactionResponse() *TransactionsItemWithTransactionResponse {
	m := &TransactionsItemWithTransactionResponse{
		TransactionsItemWithTransactionGetResponse: *NewTransactionsItemWithTransactionGetResponse(),
	}
	return m
}

// CreateTransactionsItemWithTransactionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionsItemWithTransactionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionsItemWithTransactionResponse(), nil
}

// Deprecated: This class is obsolete. Use TransactionsItemWithTransactionGetResponseable instead.
type TransactionsItemWithTransactionResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TransactionsItemWithTransactionGetResponseable
}
