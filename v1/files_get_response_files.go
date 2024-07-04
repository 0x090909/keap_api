package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type FilesGetResponse_files struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact_id property
	contact_id *int64
	// The created_by property
	created_by *int64
	// The date_created property
	date_created *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The download_url property
	download_url *string
	// The file_name property
	file_name *string
	// The file_size property
	file_size *int64
	// The id property
	id *int64
	// The last_updated property
	last_updated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The public property
	public *bool
	// The remote_file_key property
	remote_file_key *string
}

// NewFilesGetResponse_files instantiates a new FilesGetResponse_files and sets the default values.
func NewFilesGetResponse_files() *FilesGetResponse_files {
	m := &FilesGetResponse_files{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFilesGetResponse_filesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilesGetResponse_filesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFilesGetResponse_files(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FilesGetResponse_files) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *FilesGetResponse_files) GetContactId() *int64 {
	return m.contact_id
}

// GetCreatedBy gets the created_by property value. The created_by property
// returns a *int64 when successful
func (m *FilesGetResponse_files) GetCreatedBy() *int64 {
	return m.created_by
}

// GetDateCreated gets the date_created property value. The date_created property
// returns a *Time when successful
func (m *FilesGetResponse_files) GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.date_created
}

// GetDownloadUrl gets the download_url property value. The download_url property
// returns a *string when successful
func (m *FilesGetResponse_files) GetDownloadUrl() *string {
	return m.download_url
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FilesGetResponse_files) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["contact_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContactId(val)
		}
		return nil
	}
	res["created_by"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreatedBy(val)
		}
		return nil
	}
	res["date_created"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDateCreated(val)
		}
		return nil
	}
	res["download_url"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDownloadUrl(val)
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
	res["file_size"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFileSize(val)
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
	res["last_updated"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastUpdated(val)
		}
		return nil
	}
	res["public"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPublic(val)
		}
		return nil
	}
	res["remote_file_key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRemoteFileKey(val)
		}
		return nil
	}
	return res
}

// GetFileName gets the file_name property value. The file_name property
// returns a *string when successful
func (m *FilesGetResponse_files) GetFileName() *string {
	return m.file_name
}

// GetFileSize gets the file_size property value. The file_size property
// returns a *int64 when successful
func (m *FilesGetResponse_files) GetFileSize() *int64 {
	return m.file_size
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *FilesGetResponse_files) GetId() *int64 {
	return m.id
}

// GetLastUpdated gets the last_updated property value. The last_updated property
// returns a *Time when successful
func (m *FilesGetResponse_files) GetLastUpdated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.last_updated
}

// GetPublic gets the public property value. The public property
// returns a *bool when successful
func (m *FilesGetResponse_files) GetPublic() *bool {
	return m.public
}

// GetRemoteFileKey gets the remote_file_key property value. The remote_file_key property
// returns a *string when successful
func (m *FilesGetResponse_files) GetRemoteFileKey() *string {
	return m.remote_file_key
}

// Serialize serializes information the current object
func (m *FilesGetResponse_files) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("contact_id", m.GetContactId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("created_by", m.GetCreatedBy())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("date_created", m.GetDateCreated())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("download_url", m.GetDownloadUrl())
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
		err := writer.WriteInt64Value("file_size", m.GetFileSize())
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
		err := writer.WriteTimeValue("last_updated", m.GetLastUpdated())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("public", m.GetPublic())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("remote_file_key", m.GetRemoteFileKey())
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
func (m *FilesGetResponse_files) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *FilesGetResponse_files) SetContactId(value *int64) {
	m.contact_id = value
}

// SetCreatedBy sets the created_by property value. The created_by property
func (m *FilesGetResponse_files) SetCreatedBy(value *int64) {
	m.created_by = value
}

// SetDateCreated sets the date_created property value. The date_created property
func (m *FilesGetResponse_files) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.date_created = value
}

// SetDownloadUrl sets the download_url property value. The download_url property
func (m *FilesGetResponse_files) SetDownloadUrl(value *string) {
	m.download_url = value
}

// SetFileName sets the file_name property value. The file_name property
func (m *FilesGetResponse_files) SetFileName(value *string) {
	m.file_name = value
}

// SetFileSize sets the file_size property value. The file_size property
func (m *FilesGetResponse_files) SetFileSize(value *int64) {
	m.file_size = value
}

// SetId sets the id property value. The id property
func (m *FilesGetResponse_files) SetId(value *int64) {
	m.id = value
}

// SetLastUpdated sets the last_updated property value. The last_updated property
func (m *FilesGetResponse_files) SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.last_updated = value
}

// SetPublic sets the public property value. The public property
func (m *FilesGetResponse_files) SetPublic(value *bool) {
	m.public = value
}

// SetRemoteFileKey sets the remote_file_key property value. The remote_file_key property
func (m *FilesGetResponse_files) SetRemoteFileKey(value *string) {
	m.remote_file_key = value
}

type FilesGetResponse_filesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContactId() *int64
	GetCreatedBy() *int64
	GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetDownloadUrl() *string
	GetFileName() *string
	GetFileSize() *int64
	GetId() *int64
	GetLastUpdated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetPublic() *bool
	GetRemoteFileKey() *string
	SetContactId(value *int64)
	SetCreatedBy(value *int64)
	SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetDownloadUrl(value *string)
	SetFileName(value *string)
	SetFileSize(value *int64)
	SetId(value *int64)
	SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetPublic(value *bool)
	SetRemoteFileKey(value *string)
}
