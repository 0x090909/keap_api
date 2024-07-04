package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1RequestBuilder builds and executes requests for operations under \v1
type V1RequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// Account the account property
// returns a *AccountRequestBuilder when successful
func (m *V1RequestBuilder) Account() *AccountRequestBuilder {
	return NewAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Affiliates the affiliates property
// returns a *AffiliatesRequestBuilder when successful
func (m *V1RequestBuilder) Affiliates() *AffiliatesRequestBuilder {
	return NewAffiliatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Appointments the appointments property
// returns a *AppointmentsRequestBuilder when successful
func (m *V1RequestBuilder) Appointments() *AppointmentsRequestBuilder {
	return NewAppointmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Campaigns the campaigns property
// returns a *CampaignsRequestBuilder when successful
func (m *V1RequestBuilder) Campaigns() *CampaignsRequestBuilder {
	return NewCampaignsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Companies the companies property
// returns a *CompaniesRequestBuilder when successful
func (m *V1RequestBuilder) Companies() *CompaniesRequestBuilder {
	return NewCompaniesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewV1RequestBuilderInternal instantiates a new V1RequestBuilder and sets the default values.
func NewV1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *V1RequestBuilder {
	m := &V1RequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1", pathParameters),
	}
	return m
}

// NewV1RequestBuilder instantiates a new V1RequestBuilder and sets the default values.
func NewV1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *V1RequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewV1RequestBuilderInternal(urlParams, requestAdapter)
}

// Contacts the contacts property
// returns a *ContactsRequestBuilder when successful
func (m *V1RequestBuilder) Contacts() *ContactsRequestBuilder {
	return NewContactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// EmailAddresses the emailAddresses property
// returns a *EmailAddressesRequestBuilder when successful
func (m *V1RequestBuilder) EmailAddresses() *EmailAddressesRequestBuilder {
	return NewEmailAddressesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Emails the emails property
// returns a *EmailsRequestBuilder when successful
func (m *V1RequestBuilder) Emails() *EmailsRequestBuilder {
	return NewEmailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Files the files property
// returns a *FilesRequestBuilder when successful
func (m *V1RequestBuilder) Files() *FilesRequestBuilder {
	return NewFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Hooks the hooks property
// returns a *HooksRequestBuilder when successful
func (m *V1RequestBuilder) Hooks() *HooksRequestBuilder {
	return NewHooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Locales the locales property
// returns a *LocalesRequestBuilder when successful
func (m *V1RequestBuilder) Locales() *LocalesRequestBuilder {
	return NewLocalesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Merchants the merchants property
// returns a *MerchantsRequestBuilder when successful
func (m *V1RequestBuilder) Merchants() *MerchantsRequestBuilder {
	return NewMerchantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Notes the notes property
// returns a *NotesRequestBuilder when successful
func (m *V1RequestBuilder) Notes() *NotesRequestBuilder {
	return NewNotesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Oauth the oauth property
// returns a *OauthRequestBuilder when successful
func (m *V1RequestBuilder) Oauth() *OauthRequestBuilder {
	return NewOauthRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Opportunities the opportunities property
// returns a *OpportunitiesRequestBuilder when successful
func (m *V1RequestBuilder) Opportunities() *OpportunitiesRequestBuilder {
	return NewOpportunitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Opportunity the opportunity property
// returns a *OpportunityRequestBuilder when successful
func (m *V1RequestBuilder) Opportunity() *OpportunityRequestBuilder {
	return NewOpportunityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Orders the orders property
// returns a *OrdersRequestBuilder when successful
func (m *V1RequestBuilder) Orders() *OrdersRequestBuilder {
	return NewOrdersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Products the products property
// returns a *ProductsRequestBuilder when successful
func (m *V1RequestBuilder) Products() *ProductsRequestBuilder {
	return NewProductsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Setting the setting property
// returns a *SettingRequestBuilder when successful
func (m *V1RequestBuilder) Setting() *SettingRequestBuilder {
	return NewSettingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Subscriptions the subscriptions property
// returns a *SubscriptionsRequestBuilder when successful
func (m *V1RequestBuilder) Subscriptions() *SubscriptionsRequestBuilder {
	return NewSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Tags the tags property
// returns a *TagsRequestBuilder when successful
func (m *V1RequestBuilder) Tags() *TagsRequestBuilder {
	return NewTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Tasks the tasks property
// returns a *TasksRequestBuilder when successful
func (m *V1RequestBuilder) Tasks() *TasksRequestBuilder {
	return NewTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Transactions the transactions property
// returns a *TransactionsRequestBuilder when successful
func (m *V1RequestBuilder) Transactions() *TransactionsRequestBuilder {
	return NewTransactionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Users the users property
// returns a *UsersRequestBuilder when successful
func (m *V1RequestBuilder) Users() *UsersRequestBuilder {
	return NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
