package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V2RequestBuilder builds and executes requests for operations under \v2
type V2RequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// Affiliates the affiliates property
// returns a *AffiliatesRequestBuilder when successful
func (m *V2RequestBuilder) Affiliates() *AffiliatesRequestBuilder {
	return NewAffiliatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// AutomationCategory the automationCategory property
// returns a *AutomationCategoryRequestBuilder when successful
func (m *V2RequestBuilder) AutomationCategory() *AutomationCategoryRequestBuilder {
	return NewAutomationCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Automations the automations property
// returns a *AutomationsRequestBuilder when successful
func (m *V2RequestBuilder) Automations() *AutomationsRequestBuilder {
	return NewAutomationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// BusinessProfile the businessProfile property
// returns a *BusinessProfileRequestBuilder when successful
func (m *V2RequestBuilder) BusinessProfile() *BusinessProfileRequestBuilder {
	return NewBusinessProfileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Campaigns the campaigns property
// returns a *CampaignsRequestBuilder when successful
func (m *V2RequestBuilder) Campaigns() *CampaignsRequestBuilder {
	return NewCampaignsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Companies the companies property
// returns a *CompaniesRequestBuilder when successful
func (m *V2RequestBuilder) Companies() *CompaniesRequestBuilder {
	return NewCompaniesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewV2RequestBuilderInternal instantiates a new V2RequestBuilder and sets the default values.
func NewV2RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *V2RequestBuilder {
	m := &V2RequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2", pathParameters),
	}
	return m
}

// NewV2RequestBuilder instantiates a new V2RequestBuilder and sets the default values.
func NewV2RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *V2RequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewV2RequestBuilderInternal(urlParams, requestAdapter)
}

// Contacts the contacts property
// returns a *ContactsRequestBuilder when successful
func (m *V2RequestBuilder) Contacts() *ContactsRequestBuilder {
	return NewContactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ContactsLink the contactsLink property
// returns a *ContactsLinkRequestBuilder when successful
func (m *V2RequestBuilder) ContactsLink() *ContactsLinkRequestBuilder {
	return NewContactsLinkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ContactsSearch the contactsSearch property
// returns a *ContactsSearchRequestBuilder when successful
func (m *V2RequestBuilder) ContactsSearch() *ContactsSearchRequestBuilder {
	return NewContactsSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ContactsUnlink the contactsUnlink property
// returns a *ContactsUnlinkRequestBuilder when successful
func (m *V2RequestBuilder) ContactsUnlink() *ContactsUnlinkRequestBuilder {
	return NewContactsUnlinkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// EmailAddresses the emailAddresses property
// returns a *EmailAddressesRequestBuilder when successful
func (m *V2RequestBuilder) EmailAddresses() *EmailAddressesRequestBuilder {
	return NewEmailAddressesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Emails the emails property
// returns a *EmailsRequestBuilder when successful
func (m *V2RequestBuilder) Emails() *EmailsRequestBuilder {
	return NewEmailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// EmailsBatchAdd the emailsBatchAdd property
// returns a *EmailsBatchAddRequestBuilder when successful
func (m *V2RequestBuilder) EmailsBatchAdd() *EmailsBatchAddRequestBuilder {
	return NewEmailsBatchAddRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// EmailsBatchRemove the emailsBatchRemove property
// returns a *EmailsBatchRemoveRequestBuilder when successful
func (m *V2RequestBuilder) EmailsBatchRemove() *EmailsBatchRemoveRequestBuilder {
	return NewEmailsBatchRemoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// EmailsSend the emailsSend property
// returns a *EmailsSendRequestBuilder when successful
func (m *V2RequestBuilder) EmailsSend() *EmailsSendRequestBuilder {
	return NewEmailsSendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// FunnelIntegration the funnelIntegration property
// returns a *FunnelIntegrationRequestBuilder when successful
func (m *V2RequestBuilder) FunnelIntegration() *FunnelIntegrationRequestBuilder {
	return NewFunnelIntegrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Landingpages the landingpages property
// returns a *LandingpagesRequestBuilder when successful
func (m *V2RequestBuilder) Landingpages() *LandingpagesRequestBuilder {
	return NewLandingpagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Leadsources the leadsources property
// returns a *LeadsourcesRequestBuilder when successful
func (m *V2RequestBuilder) Leadsources() *LeadsourcesRequestBuilder {
	return NewLeadsourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Locales the locales property
// returns a *LocalesRequestBuilder when successful
func (m *V2RequestBuilder) Locales() *LocalesRequestBuilder {
	return NewLocalesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Oauth the oauth property
// returns a *OauthRequestBuilder when successful
func (m *V2RequestBuilder) Oauth() *OauthRequestBuilder {
	return NewOauthRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Orders the orders property
// returns a *OrdersRequestBuilder when successful
func (m *V2RequestBuilder) Orders() *OrdersRequestBuilder {
	return NewOrdersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// PaymentMethodConfigs the paymentMethodConfigs property
// returns a *PaymentMethodConfigsRequestBuilder when successful
func (m *V2RequestBuilder) PaymentMethodConfigs() *PaymentMethodConfigsRequestBuilder {
	return NewPaymentMethodConfigsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Products the products property
// returns a *ProductsRequestBuilder when successful
func (m *V2RequestBuilder) Products() *ProductsRequestBuilder {
	return NewProductsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Referrals the referrals property
// returns a *ReferralsRequestBuilder when successful
func (m *V2RequestBuilder) Referrals() *ReferralsRequestBuilder {
	return NewReferralsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Sales the sales property
// returns a *SalesRequestBuilder when successful
func (m *V2RequestBuilder) Sales() *SalesRequestBuilder {
	return NewSalesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Settings the settings property
// returns a *SettingsRequestBuilder when successful
func (m *V2RequestBuilder) Settings() *SettingsRequestBuilder {
	return NewSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// SubscriptionPlans the subscriptionPlans property
// returns a *SubscriptionPlansRequestBuilder when successful
func (m *V2RequestBuilder) SubscriptionPlans() *SubscriptionPlansRequestBuilder {
	return NewSubscriptionPlansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Subscriptions the subscriptions property
// returns a *SubscriptionsRequestBuilder when successful
func (m *V2RequestBuilder) Subscriptions() *SubscriptionsRequestBuilder {
	return NewSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Tags the tags property
// returns a *TagsRequestBuilder when successful
func (m *V2RequestBuilder) Tags() *TagsRequestBuilder {
	return NewTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Tasks the tasks property
// returns a *TasksRequestBuilder when successful
func (m *V2RequestBuilder) Tasks() *TasksRequestBuilder {
	return NewTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Users the users property
// returns a *UsersRequestBuilder when successful
func (m *V2RequestBuilder) Users() *UsersRequestBuilder {
	return NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
