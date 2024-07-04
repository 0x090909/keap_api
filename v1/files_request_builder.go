package v1

import (
	"context"
	ia794ecc2b2117ec37157fd20d3d685755a11061c3c8927401ea98be3c3252800 "github.com/0x090909/keap_api/v1/files"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// FilesRequestBuilder builds and executes requests for operations under \v1\files
type FilesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// FilesRequestBuilderGetQueryParameters retrieves a list of all files
type FilesRequestBuilderGetQueryParameters struct {
	// Filter based on Contact Id, if user has permission to see Contact files.
	Contact_id *int64 `uriparametername:"contact_id"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Filter files based on name, with '*' preceding or following to indicate LIKE queries.
	Name *string `uriparametername:"name"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
	// Filter based on the permission of files (USER or COMPANY), defaults to BOTH.
	// Deprecated: This property is deprecated, use PermissionAsGetPermissionQueryParameterType instead
	Permission *string `uriparametername:"permission"`
	// Filter based on the permission of files (USER or COMPANY), defaults to BOTH.
	PermissionAsGetPermissionQueryParameterType *ia794ecc2b2117ec37157fd20d3d685755a11061c3c8927401ea98be3c3252800.GetPermissionQueryParameterType `uriparametername:"permission"`
	// Filter based on the type of file.
	// Deprecated: This property is deprecated, use TypeAsGetTypeQueryParameterType instead
	Type *string `uriparametername:"type"`
	// Filter based on the type of file.
	TypeAsGetTypeQueryParameterType *ia794ecc2b2117ec37157fd20d3d685755a11061c3c8927401ea98be3c3252800.GetTypeQueryParameterType `uriparametername:"type"`
	// Include public or private files in response (PUBLIC or PRIVATE), defaults to BOTH.
	// Deprecated: This property is deprecated, use ViewableAsGetViewableQueryParameterType instead
	Viewable *string `uriparametername:"viewable"`
	// Include public or private files in response (PUBLIC or PRIVATE), defaults to BOTH.
	ViewableAsGetViewableQueryParameterType *ia794ecc2b2117ec37157fd20d3d685755a11061c3c8927401ea98be3c3252800.GetViewableQueryParameterType `uriparametername:"viewable"`
}

// FilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *FilesRequestBuilderGetQueryParameters
}

// FilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByFileId gets an item from the github.com/0x090909/keap_api.v1.files.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *FilesWithFileItemRequestBuilder when successful
func (m *FilesRequestBuilder) ByFileId(fileId string) *FilesWithFileItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if fileId != "" {
		urlTplParams["fileId"] = fileId
	}
	return NewFilesWithFileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByFileIdInt64 gets an item from the github.com/0x090909/keap_api.v1.files.item collection
// returns a *FilesWithFileItemRequestBuilder when successful
func (m *FilesRequestBuilder) ByFileIdInt64(fileId int64) *FilesWithFileItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["fileId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(fileId, 10)
	return NewFilesWithFileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewFilesRequestBuilderInternal instantiates a new FilesRequestBuilder and sets the default values.
func NewFilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *FilesRequestBuilder {
	m := &FilesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/files{?contact_id*,limit*,name*,offset*,permission*,type*,viewable*}", pathParameters),
	}
	return m
}

// NewFilesRequestBuilder instantiates a new FilesRequestBuilder and sets the default values.
func NewFilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *FilesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewFilesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all files
// Deprecated: This method is obsolete. Use GetAsFilesGetResponse instead.
// returns a FilesResponseable when successful
// returns a Files401Error error when the service returns a 401 status code
// returns a Files403Error error when the service returns a 403 status code
// returns a Files404Error error when the service returns a 404 status code
func (m *FilesRequestBuilder) Get(ctx context.Context, requestConfiguration *FilesRequestBuilderGetRequestConfiguration) (FilesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateFiles401ErrorFromDiscriminatorValue,
		"403": CreateFiles403ErrorFromDiscriminatorValue,
		"404": CreateFiles404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(FilesResponseable), nil
}

// GetAsFilesGetResponse retrieves a list of all files
// returns a FilesGetResponseable when successful
// returns a Files401Error error when the service returns a 401 status code
// returns a Files403Error error when the service returns a 403 status code
// returns a Files404Error error when the service returns a 404 status code
func (m *FilesRequestBuilder) GetAsFilesGetResponse(ctx context.Context, requestConfiguration *FilesRequestBuilderGetRequestConfiguration) (FilesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateFiles401ErrorFromDiscriminatorValue,
		"403": CreateFiles403ErrorFromDiscriminatorValue,
		"404": CreateFiles404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(FilesGetResponseable), nil
}

// Post upload a base64 encoded file. `contact_id` is required only when `file_association` is `CONTACT`.
// Deprecated: This method is obsolete. Use PostAsFilesPostResponse instead.
// returns a FilesResponseable when successful
// returns a Files401Error error when the service returns a 401 status code
// returns a Files403Error error when the service returns a 403 status code
func (m *FilesRequestBuilder) Post(ctx context.Context, body FilesPostRequestBodyable, requestConfiguration *FilesRequestBuilderPostRequestConfiguration) (FilesResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateFiles401ErrorFromDiscriminatorValue,
		"403": CreateFiles403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(FilesResponseable), nil
}

// PostAsFilesPostResponse upload a base64 encoded file. `contact_id` is required only when `file_association` is `CONTACT`.
// returns a FilesPostResponseable when successful
// returns a Files401Error error when the service returns a 401 status code
// returns a Files403Error error when the service returns a 403 status code
func (m *FilesRequestBuilder) PostAsFilesPostResponse(ctx context.Context, body FilesPostRequestBodyable, requestConfiguration *FilesRequestBuilderPostRequestConfiguration) (FilesPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateFiles401ErrorFromDiscriminatorValue,
		"403": CreateFiles403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilesPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(FilesPostResponseable), nil
}

// ToGetRequestInformation retrieves a list of all files
// returns a *RequestInformation when successful
func (m *FilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation upload a base64 encoded file. `contact_id` is required only when `file_association` is `CONTACT`.
// returns a *RequestInformation when successful
func (m *FilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilesPostRequestBodyable, requestConfiguration *FilesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/files", m.BaseRequestBuilder.PathParameters)
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
// returns a *FilesRequestBuilder when successful
func (m *FilesRequestBuilder) WithUrl(rawUrl string) *FilesRequestBuilder {
	return NewFilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
