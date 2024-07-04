package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProductsItemImagePostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// SHA256 checksum of image in Hex.
	checksum *string
	// The image data, base64 encoded.
	file_data *string
	// The name of the file with extension, must match file data.  Acceptable file types:  [.png, .gif, .jpg, .jpeg].
	file_name *string
}

// NewProductsItemImagePostRequestBody instantiates a new ProductsItemImagePostRequestBody and sets the default values.
func NewProductsItemImagePostRequestBody() *ProductsItemImagePostRequestBody {
	m := &ProductsItemImagePostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateProductsItemImagePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProductsItemImagePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewProductsItemImagePostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProductsItemImagePostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetChecksum gets the checksum property value. SHA256 checksum of image in Hex.
// returns a *string when successful
func (m *ProductsItemImagePostRequestBody) GetChecksum() *string {
	return m.checksum
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProductsItemImagePostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["checksum"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetChecksum(val)
		}
		return nil
	}
	res["file_data"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFileData(val)
		}
		return nil
	}
	res["file_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFileName(val)
		}
		return nil
	}
	return res
}

// GetFileData gets the file_data property value. The image data, base64 encoded.
// returns a *string when successful
func (m *ProductsItemImagePostRequestBody) GetFileData() *string {
	return m.file_data
}

// GetFileName gets the file_name property value. The name of the file with extension, must match file data.  Acceptable file types:  [.png, .gif, .jpg, .jpeg].
// returns a *string when successful
func (m *ProductsItemImagePostRequestBody) GetFileName() *string {
	return m.file_name
}

// Serialize serializes information the current object
func (m *ProductsItemImagePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("checksum", m.GetChecksum())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("file_data", m.GetFileData())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("file_name", m.GetFileName())
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
func (m *ProductsItemImagePostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetChecksum sets the checksum property value. SHA256 checksum of image in Hex.
func (m *ProductsItemImagePostRequestBody) SetChecksum(value *string) {
	m.checksum = value
}

// SetFileData sets the file_data property value. The image data, base64 encoded.
func (m *ProductsItemImagePostRequestBody) SetFileData(value *string) {
	m.file_data = value
}

// SetFileName sets the file_name property value. The name of the file with extension, must match file data.  Acceptable file types:  [.png, .gif, .jpg, .jpeg].
func (m *ProductsItemImagePostRequestBody) SetFileName(value *string) {
	m.file_name = value
}

type ProductsItemImagePostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetChecksum() *string
	GetFileData() *string
	GetFileName() *string
	SetChecksum(value *string)
	SetFileData(value *string)
	SetFileName(value *string)
}
