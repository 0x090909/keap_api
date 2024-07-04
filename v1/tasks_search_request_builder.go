package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TasksSearchRequestBuilder builds and executes requests for operations under \v1\tasks\search
type TasksSearchRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TasksSearchRequestBuilderGetQueryParameters retrieves Tasks belonging to the authenticated user using the specified search criteria
type TasksSearchRequestBuilderGetQueryParameters struct {
	// Sets completed status of items to return
	Completed *bool `uriparametername:"completed"`
	// Returns tasks for the provided contact id
	Contact_id *int64 `uriparametername:"contact_id"`
	// Returns tasks that have an 'action date' when set to true
	Has_due_date *bool `uriparametername:"has_due_date"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
	// Attribute to order items by
	Order *string `uriparametername:"order"`
	// Date to start searching from ex. `2017-01-01T22:17:59.039Z`
	Since *string `uriparametername:"since"`
	// Date to search to ex. `2017-01-01T22:17:59.039Z`
	Until *string `uriparametername:"until"`
	// Returns tasks for the provided user id
	User_id *int64 `uriparametername:"user_id"`
}

// TasksSearchRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TasksSearchRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *TasksSearchRequestBuilderGetQueryParameters
}

// NewTasksSearchRequestBuilderInternal instantiates a new TasksSearchRequestBuilder and sets the default values.
func NewTasksSearchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksSearchRequestBuilder {
	m := &TasksSearchRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/tasks/search{?completed*,contact_id*,has_due_date*,limit*,offset*,order*,since*,until*,user_id*}", pathParameters),
	}
	return m
}

// NewTasksSearchRequestBuilder instantiates a new TasksSearchRequestBuilder and sets the default values.
func NewTasksSearchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksSearchRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTasksSearchRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves Tasks belonging to the authenticated user using the specified search criteria
// Deprecated: This method is obsolete. Use GetAsSearchGetResponse instead.
// returns a TasksSearchResponseable when successful
// returns a TasksSearch401Error error when the service returns a 401 status code
// returns a TasksSearch403Error error when the service returns a 403 status code
// returns a TasksSearch404Error error when the service returns a 404 status code
func (m *TasksSearchRequestBuilder) Get(ctx context.Context, requestConfiguration *TasksSearchRequestBuilderGetRequestConfiguration) (TasksSearchResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksSearch401ErrorFromDiscriminatorValue,
		"403": CreateTasksSearch403ErrorFromDiscriminatorValue,
		"404": CreateTasksSearch404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksSearchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksSearchResponseable), nil
}

// GetAsSearchGetResponse retrieves Tasks belonging to the authenticated user using the specified search criteria
// returns a TasksSearchGetResponseable when successful
// returns a TasksSearch401Error error when the service returns a 401 status code
// returns a TasksSearch403Error error when the service returns a 403 status code
// returns a TasksSearch404Error error when the service returns a 404 status code
func (m *TasksSearchRequestBuilder) GetAsSearchGetResponse(ctx context.Context, requestConfiguration *TasksSearchRequestBuilderGetRequestConfiguration) (TasksSearchGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasksSearch401ErrorFromDiscriminatorValue,
		"403": CreateTasksSearch403ErrorFromDiscriminatorValue,
		"404": CreateTasksSearch404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksSearchGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksSearchGetResponseable), nil
}

// ToGetRequestInformation retrieves Tasks belonging to the authenticated user using the specified search criteria
// returns a *RequestInformation when successful
func (m *TasksSearchRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TasksSearchRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TasksSearchRequestBuilder when successful
func (m *TasksSearchRequestBuilder) WithUrl(rawUrl string) *TasksSearchRequestBuilder {
	return NewTasksSearchRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
