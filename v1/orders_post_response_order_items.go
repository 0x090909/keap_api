package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersPostResponse_order_items struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The billingCycle property
	billingCycle *int32
	// The cost property
	cost *float64
	// The description property
	description *string
	// The discount property
	discount *float64
	// The frequency property
	frequency *int32
	// The id property
	id *int64
	// The jobRecurringId property
	jobRecurringId *int64
	// The name property
	name *string
	// The notes property
	notes *string
	// The numberOfPayments property
	numberOfPayments *int32
	// The orderItemTaxes property
	orderItemTaxes []OrdersPostResponse_order_items_orderItemTaxesable
	// The price property
	price *float64
	// The product property
	product OrdersPostResponse_order_items_productable
	// The quantity property
	quantity *int32
	// The recurringBilling property
	recurringBilling *bool
	// The recurringCyclesCompleted property
	recurringCyclesCompleted *int32
	// The recurringEndDate property
	recurringEndDate *string
	// The recurringInactive property
	recurringInactive *bool
	// The recurringNextBillDate property
	recurringNextBillDate *string
	// The recurringStartDate property
	recurringStartDate *string
	// The specialAmount property
	specialAmount *float64
	// The specialId property
	specialId *int64
	// The specialPctOrAmt property
	specialPctOrAmt *int32
	// The type property
	typeEscaped *string
}

// NewOrdersPostResponse_order_items instantiates a new OrdersPostResponse_order_items and sets the default values.
func NewOrdersPostResponse_order_items() *OrdersPostResponse_order_items {
	m := &OrdersPostResponse_order_items{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersPostResponse_order_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersPostResponse_order_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersPostResponse_order_items(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersPostResponse_order_items) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBillingCycle gets the billingCycle property value. The billingCycle property
// returns a *int32 when successful
func (m *OrdersPostResponse_order_items) GetBillingCycle() *int32 {
	return m.billingCycle
}

// GetCost gets the cost property value. The cost property
// returns a *float64 when successful
func (m *OrdersPostResponse_order_items) GetCost() *float64 {
	return m.cost
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *OrdersPostResponse_order_items) GetDescription() *string {
	return m.description
}

// GetDiscount gets the discount property value. The discount property
// returns a *float64 when successful
func (m *OrdersPostResponse_order_items) GetDiscount() *float64 {
	return m.discount
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersPostResponse_order_items) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["billingCycle"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBillingCycle(val)
		}
		return nil
	}
	res["cost"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCost(val)
		}
		return nil
	}
	res["description"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDescription(val)
		}
		return nil
	}
	res["discount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDiscount(val)
		}
		return nil
	}
	res["frequency"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFrequency(val)
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
	res["jobRecurringId"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetJobRecurringId(val)
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
	res["notes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNotes(val)
		}
		return nil
	}
	res["numberOfPayments"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNumberOfPayments(val)
		}
		return nil
	}
	res["orderItemTaxes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateOrdersPostResponse_order_items_orderItemTaxesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]OrdersPostResponse_order_items_orderItemTaxesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(OrdersPostResponse_order_items_orderItemTaxesable)
				}
			}
			m.SetOrderItemTaxes(res)
		}
		return nil
	}
	res["price"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPrice(val)
		}
		return nil
	}
	res["product"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOrdersPostResponse_order_items_productFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProduct(val.(OrdersPostResponse_order_items_productable))
		}
		return nil
	}
	res["quantity"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetQuantity(val)
		}
		return nil
	}
	res["recurringBilling"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRecurringBilling(val)
		}
		return nil
	}
	res["recurringCyclesCompleted"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRecurringCyclesCompleted(val)
		}
		return nil
	}
	res["recurringEndDate"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRecurringEndDate(val)
		}
		return nil
	}
	res["recurringInactive"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRecurringInactive(val)
		}
		return nil
	}
	res["recurringNextBillDate"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRecurringNextBillDate(val)
		}
		return nil
	}
	res["recurringStartDate"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRecurringStartDate(val)
		}
		return nil
	}
	res["specialAmount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSpecialAmount(val)
		}
		return nil
	}
	res["specialId"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSpecialId(val)
		}
		return nil
	}
	res["specialPctOrAmt"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSpecialPctOrAmt(val)
		}
		return nil
	}
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val)
		}
		return nil
	}
	return res
}

// GetFrequency gets the frequency property value. The frequency property
// returns a *int32 when successful
func (m *OrdersPostResponse_order_items) GetFrequency() *int32 {
	return m.frequency
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OrdersPostResponse_order_items) GetId() *int64 {
	return m.id
}

// GetJobRecurringId gets the jobRecurringId property value. The jobRecurringId property
// returns a *int64 when successful
func (m *OrdersPostResponse_order_items) GetJobRecurringId() *int64 {
	return m.jobRecurringId
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *OrdersPostResponse_order_items) GetName() *string {
	return m.name
}

// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *OrdersPostResponse_order_items) GetNotes() *string {
	return m.notes
}

// GetNumberOfPayments gets the numberOfPayments property value. The numberOfPayments property
// returns a *int32 when successful
func (m *OrdersPostResponse_order_items) GetNumberOfPayments() *int32 {
	return m.numberOfPayments
}

// GetOrderItemTaxes gets the orderItemTaxes property value. The orderItemTaxes property
// returns a []OrdersPostResponse_order_items_orderItemTaxesable when successful
func (m *OrdersPostResponse_order_items) GetOrderItemTaxes() []OrdersPostResponse_order_items_orderItemTaxesable {
	return m.orderItemTaxes
}

// GetPrice gets the price property value. The price property
// returns a *float64 when successful
func (m *OrdersPostResponse_order_items) GetPrice() *float64 {
	return m.price
}

// GetProduct gets the product property value. The product property
// returns a OrdersPostResponse_order_items_productable when successful
func (m *OrdersPostResponse_order_items) GetProduct() OrdersPostResponse_order_items_productable {
	return m.product
}

// GetQuantity gets the quantity property value. The quantity property
// returns a *int32 when successful
func (m *OrdersPostResponse_order_items) GetQuantity() *int32 {
	return m.quantity
}

// GetRecurringBilling gets the recurringBilling property value. The recurringBilling property
// returns a *bool when successful
func (m *OrdersPostResponse_order_items) GetRecurringBilling() *bool {
	return m.recurringBilling
}

// GetRecurringCyclesCompleted gets the recurringCyclesCompleted property value. The recurringCyclesCompleted property
// returns a *int32 when successful
func (m *OrdersPostResponse_order_items) GetRecurringCyclesCompleted() *int32 {
	return m.recurringCyclesCompleted
}

// GetRecurringEndDate gets the recurringEndDate property value. The recurringEndDate property
// returns a *string when successful
func (m *OrdersPostResponse_order_items) GetRecurringEndDate() *string {
	return m.recurringEndDate
}

// GetRecurringInactive gets the recurringInactive property value. The recurringInactive property
// returns a *bool when successful
func (m *OrdersPostResponse_order_items) GetRecurringInactive() *bool {
	return m.recurringInactive
}

// GetRecurringNextBillDate gets the recurringNextBillDate property value. The recurringNextBillDate property
// returns a *string when successful
func (m *OrdersPostResponse_order_items) GetRecurringNextBillDate() *string {
	return m.recurringNextBillDate
}

// GetRecurringStartDate gets the recurringStartDate property value. The recurringStartDate property
// returns a *string when successful
func (m *OrdersPostResponse_order_items) GetRecurringStartDate() *string {
	return m.recurringStartDate
}

// GetSpecialAmount gets the specialAmount property value. The specialAmount property
// returns a *float64 when successful
func (m *OrdersPostResponse_order_items) GetSpecialAmount() *float64 {
	return m.specialAmount
}

// GetSpecialId gets the specialId property value. The specialId property
// returns a *int64 when successful
func (m *OrdersPostResponse_order_items) GetSpecialId() *int64 {
	return m.specialId
}

// GetSpecialPctOrAmt gets the specialPctOrAmt property value. The specialPctOrAmt property
// returns a *int32 when successful
func (m *OrdersPostResponse_order_items) GetSpecialPctOrAmt() *int32 {
	return m.specialPctOrAmt
}

// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *OrdersPostResponse_order_items) GetTypeEscaped() *string {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *OrdersPostResponse_order_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("billingCycle", m.GetBillingCycle())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("cost", m.GetCost())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("description", m.GetDescription())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("discount", m.GetDiscount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("frequency", m.GetFrequency())
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
		err := writer.WriteInt64Value("jobRecurringId", m.GetJobRecurringId())
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
		err := writer.WriteStringValue("notes", m.GetNotes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("numberOfPayments", m.GetNumberOfPayments())
		if err != nil {
			return err
		}
	}
	if m.GetOrderItemTaxes() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOrderItemTaxes()))
		for i, v := range m.GetOrderItemTaxes() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("orderItemTaxes", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("price", m.GetPrice())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("product", m.GetProduct())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("quantity", m.GetQuantity())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("recurringBilling", m.GetRecurringBilling())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("recurringCyclesCompleted", m.GetRecurringCyclesCompleted())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("recurringEndDate", m.GetRecurringEndDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("recurringInactive", m.GetRecurringInactive())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("recurringNextBillDate", m.GetRecurringNextBillDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("recurringStartDate", m.GetRecurringStartDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("specialAmount", m.GetSpecialAmount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("specialId", m.GetSpecialId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("specialPctOrAmt", m.GetSpecialPctOrAmt())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *OrdersPostResponse_order_items) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBillingCycle sets the billingCycle property value. The billingCycle property
func (m *OrdersPostResponse_order_items) SetBillingCycle(value *int32) {
	m.billingCycle = value
}

// SetCost sets the cost property value. The cost property
func (m *OrdersPostResponse_order_items) SetCost(value *float64) {
	m.cost = value
}

// SetDescription sets the description property value. The description property
func (m *OrdersPostResponse_order_items) SetDescription(value *string) {
	m.description = value
}

// SetDiscount sets the discount property value. The discount property
func (m *OrdersPostResponse_order_items) SetDiscount(value *float64) {
	m.discount = value
}

// SetFrequency sets the frequency property value. The frequency property
func (m *OrdersPostResponse_order_items) SetFrequency(value *int32) {
	m.frequency = value
}

// SetId sets the id property value. The id property
func (m *OrdersPostResponse_order_items) SetId(value *int64) {
	m.id = value
}

// SetJobRecurringId sets the jobRecurringId property value. The jobRecurringId property
func (m *OrdersPostResponse_order_items) SetJobRecurringId(value *int64) {
	m.jobRecurringId = value
}

// SetName sets the name property value. The name property
func (m *OrdersPostResponse_order_items) SetName(value *string) {
	m.name = value
}

// SetNotes sets the notes property value. The notes property
func (m *OrdersPostResponse_order_items) SetNotes(value *string) {
	m.notes = value
}

// SetNumberOfPayments sets the numberOfPayments property value. The numberOfPayments property
func (m *OrdersPostResponse_order_items) SetNumberOfPayments(value *int32) {
	m.numberOfPayments = value
}

// SetOrderItemTaxes sets the orderItemTaxes property value. The orderItemTaxes property
func (m *OrdersPostResponse_order_items) SetOrderItemTaxes(value []OrdersPostResponse_order_items_orderItemTaxesable) {
	m.orderItemTaxes = value
}

// SetPrice sets the price property value. The price property
func (m *OrdersPostResponse_order_items) SetPrice(value *float64) {
	m.price = value
}

// SetProduct sets the product property value. The product property
func (m *OrdersPostResponse_order_items) SetProduct(value OrdersPostResponse_order_items_productable) {
	m.product = value
}

// SetQuantity sets the quantity property value. The quantity property
func (m *OrdersPostResponse_order_items) SetQuantity(value *int32) {
	m.quantity = value
}

// SetRecurringBilling sets the recurringBilling property value. The recurringBilling property
func (m *OrdersPostResponse_order_items) SetRecurringBilling(value *bool) {
	m.recurringBilling = value
}

// SetRecurringCyclesCompleted sets the recurringCyclesCompleted property value. The recurringCyclesCompleted property
func (m *OrdersPostResponse_order_items) SetRecurringCyclesCompleted(value *int32) {
	m.recurringCyclesCompleted = value
}

// SetRecurringEndDate sets the recurringEndDate property value. The recurringEndDate property
func (m *OrdersPostResponse_order_items) SetRecurringEndDate(value *string) {
	m.recurringEndDate = value
}

// SetRecurringInactive sets the recurringInactive property value. The recurringInactive property
func (m *OrdersPostResponse_order_items) SetRecurringInactive(value *bool) {
	m.recurringInactive = value
}

// SetRecurringNextBillDate sets the recurringNextBillDate property value. The recurringNextBillDate property
func (m *OrdersPostResponse_order_items) SetRecurringNextBillDate(value *string) {
	m.recurringNextBillDate = value
}

// SetRecurringStartDate sets the recurringStartDate property value. The recurringStartDate property
func (m *OrdersPostResponse_order_items) SetRecurringStartDate(value *string) {
	m.recurringStartDate = value
}

// SetSpecialAmount sets the specialAmount property value. The specialAmount property
func (m *OrdersPostResponse_order_items) SetSpecialAmount(value *float64) {
	m.specialAmount = value
}

// SetSpecialId sets the specialId property value. The specialId property
func (m *OrdersPostResponse_order_items) SetSpecialId(value *int64) {
	m.specialId = value
}

// SetSpecialPctOrAmt sets the specialPctOrAmt property value. The specialPctOrAmt property
func (m *OrdersPostResponse_order_items) SetSpecialPctOrAmt(value *int32) {
	m.specialPctOrAmt = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *OrdersPostResponse_order_items) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

type OrdersPostResponse_order_itemsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBillingCycle() *int32
	GetCost() *float64
	GetDescription() *string
	GetDiscount() *float64
	GetFrequency() *int32
	GetId() *int64
	GetJobRecurringId() *int64
	GetName() *string
	GetNotes() *string
	GetNumberOfPayments() *int32
	GetOrderItemTaxes() []OrdersPostResponse_order_items_orderItemTaxesable
	GetPrice() *float64
	GetProduct() OrdersPostResponse_order_items_productable
	GetQuantity() *int32
	GetRecurringBilling() *bool
	GetRecurringCyclesCompleted() *int32
	GetRecurringEndDate() *string
	GetRecurringInactive() *bool
	GetRecurringNextBillDate() *string
	GetRecurringStartDate() *string
	GetSpecialAmount() *float64
	GetSpecialId() *int64
	GetSpecialPctOrAmt() *int32
	GetTypeEscaped() *string
	SetBillingCycle(value *int32)
	SetCost(value *float64)
	SetDescription(value *string)
	SetDiscount(value *float64)
	SetFrequency(value *int32)
	SetId(value *int64)
	SetJobRecurringId(value *int64)
	SetName(value *string)
	SetNotes(value *string)
	SetNumberOfPayments(value *int32)
	SetOrderItemTaxes(value []OrdersPostResponse_order_items_orderItemTaxesable)
	SetPrice(value *float64)
	SetProduct(value OrdersPostResponse_order_items_productable)
	SetQuantity(value *int32)
	SetRecurringBilling(value *bool)
	SetRecurringCyclesCompleted(value *int32)
	SetRecurringEndDate(value *string)
	SetRecurringInactive(value *bool)
	SetRecurringNextBillDate(value *string)
	SetRecurringStartDate(value *string)
	SetSpecialAmount(value *float64)
	SetSpecialId(value *int64)
	SetSpecialPctOrAmt(value *int32)
	SetTypeEscaped(value *string)
}
