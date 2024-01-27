/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

"github.com/infobloxopen/bloxone-go-client/internal"
)


type HaGroupAPI interface {

	/*
	HaGroupCreate Create the HA group.

	Use this method to create an __HAGroup__ object.
The __HAGroup__ object represents on-prem hosts that can serve the same leases for HA.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiHaGroupCreateRequest
	*/
	HaGroupCreate(ctx context.Context) ApiHaGroupCreateRequest

	// HaGroupCreateExecute executes the request
	//  @return IpamsvcCreateHAGroupResponse
	HaGroupCreateExecute(r ApiHaGroupCreateRequest) (*IpamsvcCreateHAGroupResponse, *http.Response, error)

	/*
	HaGroupDelete Delete the HA group.

	Use this method to delete an __HAGroup__ object.
The __HAGroup__ (_dhcp/ha_group_) object represents on-prem hosts that can serve the same leases for HA.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiHaGroupDeleteRequest
	*/
	HaGroupDelete(ctx context.Context, id string) ApiHaGroupDeleteRequest

	// HaGroupDeleteExecute executes the request
	HaGroupDeleteExecute(r ApiHaGroupDeleteRequest) (*http.Response, error)

	/*
	HaGroupList Retrieve HA groups.

	Use this method to retrieve __HAGroup__ objects.
The __HAGroup__ (_dhcp/ha_group_) object represents on-prem hosts that can serve the same leases for HA.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiHaGroupListRequest
	*/
	HaGroupList(ctx context.Context) ApiHaGroupListRequest

	// HaGroupListExecute executes the request
	//  @return IpamsvcListHAGroupResponse
	HaGroupListExecute(r ApiHaGroupListRequest) (*IpamsvcListHAGroupResponse, *http.Response, error)

	/*
	HaGroupRead Retrieve the HA group.

	Use this method to retrieve an __HAGroup__ object.
The __HAGroup__ object represents on-prem hosts that can serve the same leases for HA.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiHaGroupReadRequest
	*/
	HaGroupRead(ctx context.Context, id string) ApiHaGroupReadRequest

	// HaGroupReadExecute executes the request
	//  @return IpamsvcReadHAGroupResponse
	HaGroupReadExecute(r ApiHaGroupReadRequest) (*IpamsvcReadHAGroupResponse, *http.Response, error)

	/*
	HaGroupUpdate Update the HA group.

	Use this method to update an __HAGroup__ object.
The __HAGroup__ object represents on-prem hosts that can serve the same leases for HA.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiHaGroupUpdateRequest
	*/
	HaGroupUpdate(ctx context.Context, id string) ApiHaGroupUpdateRequest

	// HaGroupUpdateExecute executes the request
	//  @return IpamsvcUpdateHAGroupResponse
	HaGroupUpdateExecute(r ApiHaGroupUpdateRequest) (*IpamsvcUpdateHAGroupResponse, *http.Response, error)
}

// HaGroupAPIService HaGroupAPI service
type HaGroupAPIService internal.Service

type ApiHaGroupCreateRequest struct {
	ctx context.Context
	ApiService HaGroupAPI
	body *IpamsvcHAGroup
}

func (r ApiHaGroupCreateRequest) Body(body IpamsvcHAGroup) ApiHaGroupCreateRequest {
	r.body = &body
	return r
}

func (r ApiHaGroupCreateRequest) Execute() (*IpamsvcCreateHAGroupResponse, *http.Response, error) {
	return r.ApiService.HaGroupCreateExecute(r)
}

/*
HaGroupCreate Create the HA group.

Use this method to create an __HAGroup__ object.
The __HAGroup__ object represents on-prem hosts that can serve the same leases for HA.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiHaGroupCreateRequest
*/
func (a *HaGroupAPIService) HaGroupCreate(ctx context.Context) ApiHaGroupCreateRequest {
	return ApiHaGroupCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IpamsvcCreateHAGroupResponse
func (a *HaGroupAPIService) HaGroupCreateExecute(r ApiHaGroupCreateRequest) (*IpamsvcCreateHAGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *IpamsvcCreateHAGroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "HaGroupAPIService.HaGroupCreate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/ha_group"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiHaGroupDeleteRequest struct {
	ctx context.Context
	ApiService HaGroupAPI
	id string
}

func (r ApiHaGroupDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.HaGroupDeleteExecute(r)
}

/*
HaGroupDelete Delete the HA group.

Use this method to delete an __HAGroup__ object.
The __HAGroup__ (_dhcp/ha_group_) object represents on-prem hosts that can serve the same leases for HA.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiHaGroupDeleteRequest
*/
func (a *HaGroupAPIService) HaGroupDelete(ctx context.Context, id string) ApiHaGroupDeleteRequest {
	return ApiHaGroupDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *HaGroupAPIService) HaGroupDeleteExecute(r ApiHaGroupDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "HaGroupAPIService.HaGroupDelete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/ha_group/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiHaGroupListRequest struct {
	ctx context.Context
	ApiService HaGroupAPI
	filter *string
	orderBy *string
	fields *string
	offset *int32
	limit *int32
	pageToken *string
	torderBy *string
	tfilter *string
	collectStats *bool
}

//   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |        
func (r ApiHaGroupListRequest) Filter(filter string) ApiHaGroupListRequest {
	r.filter = &filter
	return r
}

//   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.        
func (r ApiHaGroupListRequest) OrderBy(orderBy string) ApiHaGroupListRequest {
	r.orderBy = &orderBy
	return r
}

//   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.        
func (r ApiHaGroupListRequest) Fields(fields string) ApiHaGroupListRequest {
	r.fields = &fields
	return r
}

//   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.         
func (r ApiHaGroupListRequest) Offset(offset int32) ApiHaGroupListRequest {
	r.offset = &offset
	return r
}

//   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.         
func (r ApiHaGroupListRequest) Limit(limit int32) ApiHaGroupListRequest {
	r.limit = &limit
	return r
}

//   The service-defined string used to identify a page of resources. A null value indicates the first page.         
func (r ApiHaGroupListRequest) PageToken(pageToken string) ApiHaGroupListRequest {
	r.pageToken = &pageToken
	return r
}

// This parameter is used for sorting by tags.
func (r ApiHaGroupListRequest) TorderBy(torderBy string) ApiHaGroupListRequest {
	r.torderBy = &torderBy
	return r
}

// This parameter is used for filtering by tags.
func (r ApiHaGroupListRequest) Tfilter(tfilter string) ApiHaGroupListRequest {
	r.tfilter = &tfilter
	return r
}

// collect_stats gets the HA group stats(state, status, heartbeat) if set to _true_ in the _GET_ _/dhcp/ha_group_ request.
func (r ApiHaGroupListRequest) CollectStats(collectStats bool) ApiHaGroupListRequest {
	r.collectStats = &collectStats
	return r
}

func (r ApiHaGroupListRequest) Execute() (*IpamsvcListHAGroupResponse, *http.Response, error) {
	return r.ApiService.HaGroupListExecute(r)
}

/*
HaGroupList Retrieve HA groups.

Use this method to retrieve __HAGroup__ objects.
The __HAGroup__ (_dhcp/ha_group_) object represents on-prem hosts that can serve the same leases for HA.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiHaGroupListRequest
*/
func (a *HaGroupAPIService) HaGroupList(ctx context.Context) ApiHaGroupListRequest {
	return ApiHaGroupListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IpamsvcListHAGroupResponse
func (a *HaGroupAPIService) HaGroupListExecute(r ApiHaGroupListRequest) (*IpamsvcListHAGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *IpamsvcListHAGroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "HaGroupAPIService.HaGroupList")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/ha_group"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_filter", r.filter, "")
	}
	if r.orderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_order_by", r.orderBy, "")
	}
	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.offset != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_offset", r.offset, "")
	}
	if r.limit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_limit", r.limit, "")
	}
	if r.pageToken != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_token", r.pageToken, "")
	}
	if r.torderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_torder_by", r.torderBy, "")
	}
	if r.tfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_tfilter", r.tfilter, "")
	}
	if r.collectStats != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "collect_stats", r.collectStats, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiHaGroupReadRequest struct {
	ctx context.Context
	ApiService HaGroupAPI
	id string
	fields *string
	collectStats *bool
}

//   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.        
func (r ApiHaGroupReadRequest) Fields(fields string) ApiHaGroupReadRequest {
	r.fields = &fields
	return r
}

// collect_stats gets the HA group stats(state, status, heartbeat) if set to _true_ in the _GET_ _/dhcp/ha_group_ request.
func (r ApiHaGroupReadRequest) CollectStats(collectStats bool) ApiHaGroupReadRequest {
	r.collectStats = &collectStats
	return r
}

func (r ApiHaGroupReadRequest) Execute() (*IpamsvcReadHAGroupResponse, *http.Response, error) {
	return r.ApiService.HaGroupReadExecute(r)
}

/*
HaGroupRead Retrieve the HA group.

Use this method to retrieve an __HAGroup__ object.
The __HAGroup__ object represents on-prem hosts that can serve the same leases for HA.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiHaGroupReadRequest
*/
func (a *HaGroupAPIService) HaGroupRead(ctx context.Context, id string) ApiHaGroupReadRequest {
	return ApiHaGroupReadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return IpamsvcReadHAGroupResponse
func (a *HaGroupAPIService) HaGroupReadExecute(r ApiHaGroupReadRequest) (*IpamsvcReadHAGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *IpamsvcReadHAGroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "HaGroupAPIService.HaGroupRead")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/ha_group/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.collectStats != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "collect_stats", r.collectStats, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiHaGroupUpdateRequest struct {
	ctx context.Context
	ApiService HaGroupAPI
	id string
	body *IpamsvcHAGroup
}

func (r ApiHaGroupUpdateRequest) Body(body IpamsvcHAGroup) ApiHaGroupUpdateRequest {
	r.body = &body
	return r
}

func (r ApiHaGroupUpdateRequest) Execute() (*IpamsvcUpdateHAGroupResponse, *http.Response, error) {
	return r.ApiService.HaGroupUpdateExecute(r)
}

/*
HaGroupUpdate Update the HA group.

Use this method to update an __HAGroup__ object.
The __HAGroup__ object represents on-prem hosts that can serve the same leases for HA.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiHaGroupUpdateRequest
*/
func (a *HaGroupAPIService) HaGroupUpdate(ctx context.Context, id string) ApiHaGroupUpdateRequest {
	return ApiHaGroupUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return IpamsvcUpdateHAGroupResponse
func (a *HaGroupAPIService) HaGroupUpdateExecute(r ApiHaGroupUpdateRequest) (*IpamsvcUpdateHAGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *IpamsvcUpdateHAGroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "HaGroupAPIService.HaGroupUpdate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/ha_group/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
