package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesRedirectlinksGetResponse_redirects struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The affiliate_id property
	affiliate_id *int64
	// The id property
	id *int64
	// The local_url_code property
	local_url_code *string
	// The name property
	name *string
	// The program_ids property
	program_ids []int64
	// The redirect_url property
	redirect_url *string
}

// NewAffiliatesRedirectlinksGetResponse_redirects instantiates a new AffiliatesRedirectlinksGetResponse_redirects and sets the default values.
func NewAffiliatesRedirectlinksGetResponse_redirects() *AffiliatesRedirectlinksGetResponse_redirects {
	m := &AffiliatesRedirectlinksGetResponse_redirects{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesRedirectlinksGetResponse_redirectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesRedirectlinksGetResponse_redirectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesRedirectlinksGetResponse_redirects(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesRedirectlinksGetResponse_redirects) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAffiliateId gets the affiliate_id property value. The affiliate_id property
// returns a *int64 when successful
func (m *AffiliatesRedirectlinksGetResponse_redirects) GetAffiliateId() *int64 {
	return m.affiliate_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesRedirectlinksGetResponse_redirects) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["affiliate_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAffiliateId(val)
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
	res["local_url_code"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLocalUrlCode(val)
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
	res["program_ids"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("int64")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]int64, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*int64))
				}
			}
			m.SetProgramIds(res)
		}
		return nil
	}
	res["redirect_url"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRedirectUrl(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *AffiliatesRedirectlinksGetResponse_redirects) GetId() *int64 {
	return m.id
}

// GetLocalUrlCode gets the local_url_code property value. The local_url_code property
// returns a *string when successful
func (m *AffiliatesRedirectlinksGetResponse_redirects) GetLocalUrlCode() *string {
	return m.local_url_code
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *AffiliatesRedirectlinksGetResponse_redirects) GetName() *string {
	return m.name
}

// GetProgramIds gets the program_ids property value. The program_ids property
// returns a []int64 when successful
func (m *AffiliatesRedirectlinksGetResponse_redirects) GetProgramIds() []int64 {
	return m.program_ids
}

// GetRedirectUrl gets the redirect_url property value. The redirect_url property
// returns a *string when successful
func (m *AffiliatesRedirectlinksGetResponse_redirects) GetRedirectUrl() *string {
	return m.redirect_url
}

// Serialize serializes information the current object
func (m *AffiliatesRedirectlinksGetResponse_redirects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("affiliate_id", m.GetAffiliateId())
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
		err := writer.WriteStringValue("local_url_code", m.GetLocalUrlCode())
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
	if m.GetProgramIds() != nil {
		err := writer.WriteCollectionOfInt64Values("program_ids", m.GetProgramIds())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("redirect_url", m.GetRedirectUrl())
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
func (m *AffiliatesRedirectlinksGetResponse_redirects) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAffiliateId sets the affiliate_id property value. The affiliate_id property
func (m *AffiliatesRedirectlinksGetResponse_redirects) SetAffiliateId(value *int64) {
	m.affiliate_id = value
}

// SetId sets the id property value. The id property
func (m *AffiliatesRedirectlinksGetResponse_redirects) SetId(value *int64) {
	m.id = value
}

// SetLocalUrlCode sets the local_url_code property value. The local_url_code property
func (m *AffiliatesRedirectlinksGetResponse_redirects) SetLocalUrlCode(value *string) {
	m.local_url_code = value
}

// SetName sets the name property value. The name property
func (m *AffiliatesRedirectlinksGetResponse_redirects) SetName(value *string) {
	m.name = value
}

// SetProgramIds sets the program_ids property value. The program_ids property
func (m *AffiliatesRedirectlinksGetResponse_redirects) SetProgramIds(value []int64) {
	m.program_ids = value
}

// SetRedirectUrl sets the redirect_url property value. The redirect_url property
func (m *AffiliatesRedirectlinksGetResponse_redirects) SetRedirectUrl(value *string) {
	m.redirect_url = value
}

type AffiliatesRedirectlinksGetResponse_redirectsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAffiliateId() *int64
	GetId() *int64
	GetLocalUrlCode() *string
	GetName() *string
	GetProgramIds() []int64
	GetRedirectUrl() *string
	SetAffiliateId(value *int64)
	SetId(value *int64)
	SetLocalUrlCode(value *string)
	SetName(value *string)
	SetProgramIds(value []int64)
	SetRedirectUrl(value *string)
}
