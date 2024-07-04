package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersPostResponse_order_items_product_product_options struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The allow_spaces property
	allow_spaces *bool
	// The can_contain_character property
	can_contain_character *bool
	// The can_contain_number property
	can_contain_number *bool
	// The can_end_with_character property
	can_end_with_character *bool
	// The can_end_with_number property
	can_end_with_number *bool
	// The can_start_with_character property
	can_start_with_character *bool
	// The can_start_with_number property
	can_start_with_number *bool
	// The display_index property
	display_index *int32
	// The id property
	id *int64
	// The label property
	label *string
	// The max_chars property
	max_chars *int32
	// The min_chars property
	min_chars *int32
	// The name property
	name *string
	// The required property
	required *bool
	// The text_message property
	text_message *string
	// The values property
	values []OrdersPostResponse_order_items_product_product_options_valuesable
}

// NewOrdersPostResponse_order_items_product_product_options instantiates a new OrdersPostResponse_order_items_product_product_options and sets the default values.
func NewOrdersPostResponse_order_items_product_product_options() *OrdersPostResponse_order_items_product_product_options {
	m := &OrdersPostResponse_order_items_product_product_options{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersPostResponse_order_items_product_product_optionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersPostResponse_order_items_product_product_optionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersPostResponse_order_items_product_product_options(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAllowSpaces gets the allow_spaces property value. The allow_spaces property
// returns a *bool when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetAllowSpaces() *bool {
	return m.allow_spaces
}

// GetCanContainCharacter gets the can_contain_character property value. The can_contain_character property
// returns a *bool when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetCanContainCharacter() *bool {
	return m.can_contain_character
}

// GetCanContainNumber gets the can_contain_number property value. The can_contain_number property
// returns a *bool when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetCanContainNumber() *bool {
	return m.can_contain_number
}

// GetCanEndWithCharacter gets the can_end_with_character property value. The can_end_with_character property
// returns a *bool when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetCanEndWithCharacter() *bool {
	return m.can_end_with_character
}

// GetCanEndWithNumber gets the can_end_with_number property value. The can_end_with_number property
// returns a *bool when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetCanEndWithNumber() *bool {
	return m.can_end_with_number
}

// GetCanStartWithCharacter gets the can_start_with_character property value. The can_start_with_character property
// returns a *bool when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetCanStartWithCharacter() *bool {
	return m.can_start_with_character
}

// GetCanStartWithNumber gets the can_start_with_number property value. The can_start_with_number property
// returns a *bool when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetCanStartWithNumber() *bool {
	return m.can_start_with_number
}

// GetDisplayIndex gets the display_index property value. The display_index property
// returns a *int32 when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetDisplayIndex() *int32 {
	return m.display_index
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["allow_spaces"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAllowSpaces(val)
		}
		return nil
	}
	res["can_contain_character"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCanContainCharacter(val)
		}
		return nil
	}
	res["can_contain_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCanContainNumber(val)
		}
		return nil
	}
	res["can_end_with_character"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCanEndWithCharacter(val)
		}
		return nil
	}
	res["can_end_with_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCanEndWithNumber(val)
		}
		return nil
	}
	res["can_start_with_character"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCanStartWithCharacter(val)
		}
		return nil
	}
	res["can_start_with_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCanStartWithNumber(val)
		}
		return nil
	}
	res["display_index"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDisplayIndex(val)
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["label"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLabel(val)
		}
		return nil
	}
	res["max_chars"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMaxChars(val)
		}
		return nil
	}
	res["min_chars"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMinChars(val)
		}
		return nil
	}
	res["name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetName(val)
		}
		return nil
	}
	res["required"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRequired(val)
		}
		return nil
	}
	res["text_message"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTextMessage(val)
		}
		return nil
	}
	res["values"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateOrdersPostResponse_order_items_product_product_options_valuesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]OrdersPostResponse_order_items_product_product_options_valuesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(OrdersPostResponse_order_items_product_product_options_valuesable)
				}
			}
			m.SetValues(res)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetId() *int64 {
	return m.id
}

// GetLabel gets the label property value. The label property
// returns a *string when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetLabel() *string {
	return m.label
}

// GetMaxChars gets the max_chars property value. The max_chars property
// returns a *int32 when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetMaxChars() *int32 {
	return m.max_chars
}

// GetMinChars gets the min_chars property value. The min_chars property
// returns a *int32 when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetMinChars() *int32 {
	return m.min_chars
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetName() *string {
	return m.name
}

// GetRequired gets the required property value. The required property
// returns a *bool when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetRequired() *bool {
	return m.required
}

// GetTextMessage gets the text_message property value. The text_message property
// returns a *string when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetTextMessage() *string {
	return m.text_message
}

// GetValues gets the values property value. The values property
// returns a []OrdersPostResponse_order_items_product_product_options_valuesable when successful
func (m *OrdersPostResponse_order_items_product_product_options) GetValues() []OrdersPostResponse_order_items_product_product_options_valuesable {
	return m.values
}

// Serialize serializes information the current object
func (m *OrdersPostResponse_order_items_product_product_options) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("allow_spaces", m.GetAllowSpaces())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("can_contain_character", m.GetCanContainCharacter())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("can_contain_number", m.GetCanContainNumber())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("can_end_with_character", m.GetCanEndWithCharacter())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("can_end_with_number", m.GetCanEndWithNumber())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("can_start_with_character", m.GetCanStartWithCharacter())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("can_start_with_number", m.GetCanStartWithNumber())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("display_index", m.GetDisplayIndex())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("label", m.GetLabel())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("max_chars", m.GetMaxChars())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("min_chars", m.GetMinChars())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("name", m.GetName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("required", m.GetRequired())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("text_message", m.GetTextMessage())
		if err != nil {
			return err
		}
	}
	if m.GetValues() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValues()))
		for i, v := range m.GetValues() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("values", cast)
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
func (m *OrdersPostResponse_order_items_product_product_options) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAllowSpaces sets the allow_spaces property value. The allow_spaces property
func (m *OrdersPostResponse_order_items_product_product_options) SetAllowSpaces(value *bool) {
	m.allow_spaces = value
}

// SetCanContainCharacter sets the can_contain_character property value. The can_contain_character property
func (m *OrdersPostResponse_order_items_product_product_options) SetCanContainCharacter(value *bool) {
	m.can_contain_character = value
}

// SetCanContainNumber sets the can_contain_number property value. The can_contain_number property
func (m *OrdersPostResponse_order_items_product_product_options) SetCanContainNumber(value *bool) {
	m.can_contain_number = value
}

// SetCanEndWithCharacter sets the can_end_with_character property value. The can_end_with_character property
func (m *OrdersPostResponse_order_items_product_product_options) SetCanEndWithCharacter(value *bool) {
	m.can_end_with_character = value
}

// SetCanEndWithNumber sets the can_end_with_number property value. The can_end_with_number property
func (m *OrdersPostResponse_order_items_product_product_options) SetCanEndWithNumber(value *bool) {
	m.can_end_with_number = value
}

// SetCanStartWithCharacter sets the can_start_with_character property value. The can_start_with_character property
func (m *OrdersPostResponse_order_items_product_product_options) SetCanStartWithCharacter(value *bool) {
	m.can_start_with_character = value
}

// SetCanStartWithNumber sets the can_start_with_number property value. The can_start_with_number property
func (m *OrdersPostResponse_order_items_product_product_options) SetCanStartWithNumber(value *bool) {
	m.can_start_with_number = value
}

// SetDisplayIndex sets the display_index property value. The display_index property
func (m *OrdersPostResponse_order_items_product_product_options) SetDisplayIndex(value *int32) {
	m.display_index = value
}

// SetId sets the id property value. The id property
func (m *OrdersPostResponse_order_items_product_product_options) SetId(value *int64) {
	m.id = value
}

// SetLabel sets the label property value. The label property
func (m *OrdersPostResponse_order_items_product_product_options) SetLabel(value *string) {
	m.label = value
}

// SetMaxChars sets the max_chars property value. The max_chars property
func (m *OrdersPostResponse_order_items_product_product_options) SetMaxChars(value *int32) {
	m.max_chars = value
}

// SetMinChars sets the min_chars property value. The min_chars property
func (m *OrdersPostResponse_order_items_product_product_options) SetMinChars(value *int32) {
	m.min_chars = value
}

// SetName sets the name property value. The name property
func (m *OrdersPostResponse_order_items_product_product_options) SetName(value *string) {
	m.name = value
}

// SetRequired sets the required property value. The required property
func (m *OrdersPostResponse_order_items_product_product_options) SetRequired(value *bool) {
	m.required = value
}

// SetTextMessage sets the text_message property value. The text_message property
func (m *OrdersPostResponse_order_items_product_product_options) SetTextMessage(value *string) {
	m.text_message = value
}

// SetValues sets the values property value. The values property
func (m *OrdersPostResponse_order_items_product_product_options) SetValues(value []OrdersPostResponse_order_items_product_product_options_valuesable) {
	m.values = value
}

type OrdersPostResponse_order_items_product_product_optionsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAllowSpaces() *bool
	GetCanContainCharacter() *bool
	GetCanContainNumber() *bool
	GetCanEndWithCharacter() *bool
	GetCanEndWithNumber() *bool
	GetCanStartWithCharacter() *bool
	GetCanStartWithNumber() *bool
	GetDisplayIndex() *int32
	GetId() *int64
	GetLabel() *string
	GetMaxChars() *int32
	GetMinChars() *int32
	GetName() *string
	GetRequired() *bool
	GetTextMessage() *string
	GetValues() []OrdersPostResponse_order_items_product_product_options_valuesable
	SetAllowSpaces(value *bool)
	SetCanContainCharacter(value *bool)
	SetCanContainNumber(value *bool)
	SetCanEndWithCharacter(value *bool)
	SetCanEndWithNumber(value *bool)
	SetCanStartWithCharacter(value *bool)
	SetCanStartWithNumber(value *bool)
	SetDisplayIndex(value *int32)
	SetId(value *int64)
	SetLabel(value *string)
	SetMaxChars(value *int32)
	SetMinChars(value *int32)
	SetName(value *string)
	SetRequired(value *bool)
	SetTextMessage(value *string)
	SetValues(value []OrdersPostResponse_order_items_product_product_options_valuesable)
}
