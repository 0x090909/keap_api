package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TasksRequestBuilder builds and executes requests for operations under \v1\tasks
type TasksRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TasksRequestBuilderGetQueryParameters retrieves a list of all tasks using the specified search criteria
type TasksRequestBuilderGetQueryParameters struct {
	// Sets completed status of items to return
	Completed *bool `uriparametername:"completed"`
	// contact_id
	Contact_id *int64 `uriparametername:"contact_id"`
	// has_due_date
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
	// user_id
	User_id *int64 `uriparametername:"user_id"`
}

// TasksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TasksRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *TasksRequestBuilderGetQueryParameters
}

// TasksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TasksRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByTaskId gets an item from the github.com/0x090909/keap_api.v1.tasks.item collection
// returns a *TasksWithTaskItemRequestBuilder when successful
func (m *TasksRequestBuilder) ByTaskId(taskId string) *TasksWithTaskItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if taskId != "" {
		urlTplParams["taskId"] = taskId
	}
	return NewTasksWithTaskItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewTasksRequestBuilderInternal instantiates a new TasksRequestBuilder and sets the default values.
func NewTasksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksRequestBuilder {
	m := &TasksRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/tasks{?completed*,contact_id*,has_due_date*,limit*,offset*,order*,since*,until*,user_id*}", pathParameters),
	}
	return m
}

// NewTasksRequestBuilder instantiates a new TasksRequestBuilder and sets the default values.
func NewTasksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TasksRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTasksRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all tasks using the specified search criteria
// Deprecated: This method is obsolete. Use GetAsTasksGetResponse instead.
// returns a TasksResponseable when successful
// returns a Tasks401Error error when the service returns a 401 status code
// returns a Tasks403Error error when the service returns a 403 status code
// returns a Tasks404Error error when the service returns a 404 status code
func (m *TasksRequestBuilder) Get(ctx context.Context, requestConfiguration *TasksRequestBuilderGetRequestConfiguration) (TasksResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasks401ErrorFromDiscriminatorValue,
		"403": CreateTasks403ErrorFromDiscriminatorValue,
		"404": CreateTasks404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksResponseable), nil
}

// GetAsTasksGetResponse retrieves a list of all tasks using the specified search criteria
// returns a TasksGetResponseable when successful
// returns a Tasks401Error error when the service returns a 401 status code
// returns a Tasks403Error error when the service returns a 403 status code
// returns a Tasks404Error error when the service returns a 404 status code
func (m *TasksRequestBuilder) GetAsTasksGetResponse(ctx context.Context, requestConfiguration *TasksRequestBuilderGetRequestConfiguration) (TasksGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasks401ErrorFromDiscriminatorValue,
		"403": CreateTasks403ErrorFromDiscriminatorValue,
		"404": CreateTasks404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksGetResponseable), nil
}

// Model the model property
// returns a *TasksModelRequestBuilder when successful
func (m *TasksRequestBuilder) Model() *TasksModelRequestBuilder {
	return NewTasksModelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Post creates a new task as the authenticated user.  *Note:* Contact must contain at least one item in the fields `title` and `due_date`.  All other attributes are optional.  This endpoint does not currently support setting Custom Field values.
// Deprecated: This method is obsolete. Use PostAsTasksPostResponse instead.
// returns a TasksResponseable when successful
// returns a Tasks401Error error when the service returns a 401 status code
// returns a Tasks403Error error when the service returns a 403 status code
func (m *TasksRequestBuilder) Post(ctx context.Context, body TasksPostRequestBodyable, requestConfiguration *TasksRequestBuilderPostRequestConfiguration) (TasksResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasks401ErrorFromDiscriminatorValue,
		"403": CreateTasks403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksResponseable), nil
}

// PostAsTasksPostResponse creates a new task as the authenticated user.  *Note:* Contact must contain at least one item in the fields `title` and `due_date`.  All other attributes are optional.  This endpoint does not currently support setting Custom Field values.
// returns a TasksPostResponseable when successful
// returns a Tasks401Error error when the service returns a 401 status code
// returns a Tasks403Error error when the service returns a 403 status code
func (m *TasksRequestBuilder) PostAsTasksPostResponse(ctx context.Context, body TasksPostRequestBodyable, requestConfiguration *TasksRequestBuilderPostRequestConfiguration) (TasksPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTasks401ErrorFromDiscriminatorValue,
		"403": CreateTasks403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTasksPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TasksPostResponseable), nil
}

// Search the search property
// returns a *TasksSearchRequestBuilder when successful
func (m *TasksRequestBuilder) Search() *TasksSearchRequestBuilder {
	return NewTasksSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToGetRequestInformation retrieves a list of all tasks using the specified search criteria
// returns a *RequestInformation when successful
func (m *TasksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TasksRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation creates a new task as the authenticated user.  *Note:* Contact must contain at least one item in the fields `title` and `due_date`.  All other attributes are optional.  This endpoint does not currently support setting Custom Field values.
// returns a *RequestInformation when successful
func (m *TasksRequestBuilder) ToPostRequestInformation(ctx context.Context, body TasksPostRequestBodyable, requestConfiguration *TasksRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/tasks", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TasksRequestBuilder when successful
func (m *TasksRequestBuilder) WithUrl(rawUrl string) *TasksRequestBuilder {
	return NewTasksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
