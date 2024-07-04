package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FilesWithFileItemRequestBuilder builds and executes requests for operations under \v1\files\{fileId}
type FilesWithFileItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// FilesWithFileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilesWithFileItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// FilesWithFileItemRequestBuilderGetQueryParameters retrieves metadata about a specific file. Optionally returns the base64 encoded file data.
type FilesWithFileItemRequestBuilderGetQueryParameters struct {
	// Comma-delimited list of File properties to include in the response. (Some fields such as `file_data` aren't included, by default.)
	Optional_properties []string `uriparametername:"optional_properties"`
}

// FilesWithFileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilesWithFileItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *FilesWithFileItemRequestBuilderGetQueryParameters
}

// FilesWithFileItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilesWithFileItemRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewFilesWithFileItemRequestBuilderInternal instantiates a new FilesWithFileItemRequestBuilder and sets the default values.
func NewFilesWithFileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *FilesWithFileItemRequestBuilder {
	m := &FilesWithFileItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/files/{fileId}{?optional_properties*}", pathParameters),
	}
	return m
}

// NewFilesWithFileItemRequestBuilder instantiates a new FilesWithFileItemRequestBuilder and sets the default values.
func NewFilesWithFileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *FilesWithFileItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewFilesWithFileItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes the specified file
// returns a FilesItemWithFile401Error error when the service returns a 401 status code
// returns a FilesItemWithFile403Error error when the service returns a 403 status code
// returns a FilesItemWithFile404Error error when the service returns a 404 status code
func (m *FilesWithFileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilesWithFileItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateFilesItemWithFile401ErrorFromDiscriminatorValue,
		"403": CreateFilesItemWithFile403ErrorFromDiscriminatorValue,
		"404": CreateFilesItemWithFile404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves metadata about a specific file. Optionally returns the base64 encoded file data.
// Deprecated: This method is obsolete. Use GetAsWithFileGetResponse instead.
// returns a FilesItemWithFileResponseable when successful
// returns a FilesItemWithFile401Error error when the service returns a 401 status code
// returns a FilesItemWithFile403Error error when the service returns a 403 status code
// returns a FilesItemWithFile404Error error when the service returns a 404 status code
func (m *FilesWithFileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilesWithFileItemRequestBuilderGetRequestConfiguration) (FilesItemWithFileResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateFilesItemWithFile401ErrorFromDiscriminatorValue,
		"403": CreateFilesItemWithFile403ErrorFromDiscriminatorValue,
		"404": CreateFilesItemWithFile404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilesItemWithFileResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(FilesItemWithFileResponseable), nil
}

// GetAsWithFileGetResponse retrieves metadata about a specific file. Optionally returns the base64 encoded file data.
// returns a FilesItemWithFileGetResponseable when successful
// returns a FilesItemWithFile401Error error when the service returns a 401 status code
// returns a FilesItemWithFile403Error error when the service returns a 403 status code
// returns a FilesItemWithFile404Error error when the service returns a 404 status code
func (m *FilesWithFileItemRequestBuilder) GetAsWithFileGetResponse(ctx context.Context, requestConfiguration *FilesWithFileItemRequestBuilderGetRequestConfiguration) (FilesItemWithFileGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateFilesItemWithFile401ErrorFromDiscriminatorValue,
		"403": CreateFilesItemWithFile403ErrorFromDiscriminatorValue,
		"404": CreateFilesItemWithFile404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilesItemWithFileGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(FilesItemWithFileGetResponseable), nil
}

// Put upload a base64 encoded file to replace an existing one. `contact_id` is required only when `file_association` is `CONTACT`.
// Deprecated: This method is obsolete. Use PutAsWithFilePutResponse instead.
// returns a FilesItemWithFileResponseable when successful
// returns a FilesItemWithFile401Error error when the service returns a 401 status code
// returns a FilesItemWithFile403Error error when the service returns a 403 status code
// returns a FilesItemWithFile404Error error when the service returns a 404 status code
func (m *FilesWithFileItemRequestBuilder) Put(ctx context.Context, body FilesItemWithFilePutRequestBodyable, requestConfiguration *FilesWithFileItemRequestBuilderPutRequestConfiguration) (FilesItemWithFileResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateFilesItemWithFile401ErrorFromDiscriminatorValue,
		"403": CreateFilesItemWithFile403ErrorFromDiscriminatorValue,
		"404": CreateFilesItemWithFile404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilesItemWithFileResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(FilesItemWithFileResponseable), nil
}

// PutAsWithFilePutResponse upload a base64 encoded file to replace an existing one. `contact_id` is required only when `file_association` is `CONTACT`.
// returns a FilesItemWithFilePutResponseable when successful
// returns a FilesItemWithFile401Error error when the service returns a 401 status code
// returns a FilesItemWithFile403Error error when the service returns a 403 status code
// returns a FilesItemWithFile404Error error when the service returns a 404 status code
func (m *FilesWithFileItemRequestBuilder) PutAsWithFilePutResponse(ctx context.Context, body FilesItemWithFilePutRequestBodyable, requestConfiguration *FilesWithFileItemRequestBuilderPutRequestConfiguration) (FilesItemWithFilePutResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateFilesItemWithFile401ErrorFromDiscriminatorValue,
		"403": CreateFilesItemWithFile403ErrorFromDiscriminatorValue,
		"404": CreateFilesItemWithFile404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilesItemWithFilePutResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(FilesItemWithFilePutResponseable), nil
}

// ToDeleteRequestInformation deletes the specified file
// returns a *RequestInformation when successful
func (m *FilesWithFileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilesWithFileItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/v1/files/{fileId}", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieves metadata about a specific file. Optionally returns the base64 encoded file data.
// returns a *RequestInformation when successful
func (m *FilesWithFileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilesWithFileItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPutRequestInformation upload a base64 encoded file to replace an existing one. `contact_id` is required only when `file_association` is `CONTACT`.
// returns a *RequestInformation when successful
func (m *FilesWithFileItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body FilesItemWithFilePutRequestBodyable, requestConfiguration *FilesWithFileItemRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, "{+baseurl}/v1/files/{fileId}", m.BaseRequestBuilder.PathParameters)
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
// returns a *FilesWithFileItemRequestBuilder when successful
func (m *FilesWithFileItemRequestBuilder) WithUrl(rawUrl string) *FilesWithFileItemRequestBuilder {
	return NewFilesWithFileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
