/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.   

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

"github.com/infobloxopen/bloxone-go-client/internal"
)


type ForwardNsgAPI interface {

	/*
	ForwardNsgCreate Create the ForwardNSG object.

	Use this method to create a ForwardNSG object.
The _dns/forward_nsg_ object represents a Forward DNS Server Group for forward zones.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiForwardNsgCreateRequest
	*/
	ForwardNsgCreate(ctx context.Context) ApiForwardNsgCreateRequest

	// ForwardNsgCreateExecute executes the request
	//  @return ConfigCreateForwardNSGResponse
	ForwardNsgCreateExecute(r ApiForwardNsgCreateRequest) (*ConfigCreateForwardNSGResponse, *http.Response, error)

	/*
	ForwardNsgDelete Move the ForwardNSG object to Recyclebin.

	Use this method to move a ForwardNSG object to Recyclebin.
The _dns/forward_nsg_ object represents a Forward DNS Server Group for forward zones.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiForwardNsgDeleteRequest
	*/
	ForwardNsgDelete(ctx context.Context, id string) ApiForwardNsgDeleteRequest

	// ForwardNsgDeleteExecute executes the request
	ForwardNsgDeleteExecute(r ApiForwardNsgDeleteRequest) (*http.Response, error)

	/*
	ForwardNsgList List ForwardNSG objects.

	Use this method to list ForwardNSG objects.
The _dns/forward_nsg_ object represents a Forward DNS Server Group for forward zones.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiForwardNsgListRequest
	*/
	ForwardNsgList(ctx context.Context) ApiForwardNsgListRequest

	// ForwardNsgListExecute executes the request
	//  @return ConfigListForwardNSGResponse
	ForwardNsgListExecute(r ApiForwardNsgListRequest) (*ConfigListForwardNSGResponse, *http.Response, error)

	/*
	ForwardNsgRead Read the ForwardNSG object.

	Use this method to read a ForwardNSG object.
The _dns/forward_nsg_ object represents a Forward DNS Server Group for forward zones.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiForwardNsgReadRequest
	*/
	ForwardNsgRead(ctx context.Context, id string) ApiForwardNsgReadRequest

	// ForwardNsgReadExecute executes the request
	//  @return ConfigReadForwardNSGResponse
	ForwardNsgReadExecute(r ApiForwardNsgReadRequest) (*ConfigReadForwardNSGResponse, *http.Response, error)

	/*
	ForwardNsgUpdate Update the ForwardNSG object.

	Use this method to update a ForwardNSG object.
The _dns/forward_nsg_ object represents a Forward DNS Server Group for forward zones.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiForwardNsgUpdateRequest
	*/
	ForwardNsgUpdate(ctx context.Context, id string) ApiForwardNsgUpdateRequest

	// ForwardNsgUpdateExecute executes the request
	//  @return ConfigUpdateForwardNSGResponse
	ForwardNsgUpdateExecute(r ApiForwardNsgUpdateRequest) (*ConfigUpdateForwardNSGResponse, *http.Response, error)
}

// ForwardNsgAPIService ForwardNsgAPI service
type ForwardNsgAPIService internal.Service

type ApiForwardNsgCreateRequest struct {
	ctx context.Context
	ApiService ForwardNsgAPI
	body *ConfigForwardNSG
}

func (r ApiForwardNsgCreateRequest) Body(body ConfigForwardNSG) ApiForwardNsgCreateRequest {
	r.body = &body
	return r
}

func (r ApiForwardNsgCreateRequest) Execute() (*ConfigCreateForwardNSGResponse, *http.Response, error) {
	return r.ApiService.ForwardNsgCreateExecute(r)
}

/*
ForwardNsgCreate Create the ForwardNSG object.

Use this method to create a ForwardNSG object.
The _dns/forward_nsg_ object represents a Forward DNS Server Group for forward zones.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiForwardNsgCreateRequest
*/
func (a *ForwardNsgAPIService) ForwardNsgCreate(ctx context.Context) ApiForwardNsgCreateRequest {
	return ApiForwardNsgCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConfigCreateForwardNSGResponse
func (a *ForwardNsgAPIService) ForwardNsgCreateExecute(r ApiForwardNsgCreateRequest) (*ConfigCreateForwardNSGResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *ConfigCreateForwardNSGResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ForwardNsgAPIService.ForwardNsgCreate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/forward_nsg"

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

type ApiForwardNsgDeleteRequest struct {
	ctx context.Context
	ApiService ForwardNsgAPI
	id string
}

func (r ApiForwardNsgDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ForwardNsgDeleteExecute(r)
}

/*
ForwardNsgDelete Move the ForwardNSG object to Recyclebin.

Use this method to move a ForwardNSG object to Recyclebin.
The _dns/forward_nsg_ object represents a Forward DNS Server Group for forward zones.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiForwardNsgDeleteRequest
*/
func (a *ForwardNsgAPIService) ForwardNsgDelete(ctx context.Context, id string) ApiForwardNsgDeleteRequest {
	return ApiForwardNsgDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ForwardNsgAPIService) ForwardNsgDeleteExecute(r ApiForwardNsgDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ForwardNsgAPIService.ForwardNsgDelete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/forward_nsg/{id}"
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

type ApiForwardNsgListRequest struct {
	ctx context.Context
	ApiService ForwardNsgAPI
	fields *string
	filter *string
	offset *int32
	limit *int32
	pageToken *string
	orderBy *string
	tfilter *string
	torderBy *string
}

//   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.        
func (r ApiForwardNsgListRequest) Fields(fields string) ApiForwardNsgListRequest {
	r.fields = &fields
	return r
}

//   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |        
func (r ApiForwardNsgListRequest) Filter(filter string) ApiForwardNsgListRequest {
	r.filter = &filter
	return r
}

//   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.         
func (r ApiForwardNsgListRequest) Offset(offset int32) ApiForwardNsgListRequest {
	r.offset = &offset
	return r
}

//   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.         
func (r ApiForwardNsgListRequest) Limit(limit int32) ApiForwardNsgListRequest {
	r.limit = &limit
	return r
}

//   The service-defined string used to identify a page of resources. A null value indicates the first page.         
func (r ApiForwardNsgListRequest) PageToken(pageToken string) ApiForwardNsgListRequest {
	r.pageToken = &pageToken
	return r
}

//   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.        
func (r ApiForwardNsgListRequest) OrderBy(orderBy string) ApiForwardNsgListRequest {
	r.orderBy = &orderBy
	return r
}

// This parameter is used for filtering by tags.
func (r ApiForwardNsgListRequest) Tfilter(tfilter string) ApiForwardNsgListRequest {
	r.tfilter = &tfilter
	return r
}

// This parameter is used for sorting by tags.
func (r ApiForwardNsgListRequest) TorderBy(torderBy string) ApiForwardNsgListRequest {
	r.torderBy = &torderBy
	return r
}

func (r ApiForwardNsgListRequest) Execute() (*ConfigListForwardNSGResponse, *http.Response, error) {
	return r.ApiService.ForwardNsgListExecute(r)
}

/*
ForwardNsgList List ForwardNSG objects.

Use this method to list ForwardNSG objects.
The _dns/forward_nsg_ object represents a Forward DNS Server Group for forward zones.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiForwardNsgListRequest
*/
func (a *ForwardNsgAPIService) ForwardNsgList(ctx context.Context) ApiForwardNsgListRequest {
	return ApiForwardNsgListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConfigListForwardNSGResponse
func (a *ForwardNsgAPIService) ForwardNsgListExecute(r ApiForwardNsgListRequest) (*ConfigListForwardNSGResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *ConfigListForwardNSGResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ForwardNsgAPIService.ForwardNsgList")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/forward_nsg"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.filter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_filter", r.filter, "")
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
	if r.orderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_order_by", r.orderBy, "")
	}
	if r.tfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_tfilter", r.tfilter, "")
	}
	if r.torderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_torder_by", r.torderBy, "")
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

type ApiForwardNsgReadRequest struct {
	ctx context.Context
	ApiService ForwardNsgAPI
	id string
	fields *string
}

//   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.        
func (r ApiForwardNsgReadRequest) Fields(fields string) ApiForwardNsgReadRequest {
	r.fields = &fields
	return r
}

func (r ApiForwardNsgReadRequest) Execute() (*ConfigReadForwardNSGResponse, *http.Response, error) {
	return r.ApiService.ForwardNsgReadExecute(r)
}

/*
ForwardNsgRead Read the ForwardNSG object.

Use this method to read a ForwardNSG object.
The _dns/forward_nsg_ object represents a Forward DNS Server Group for forward zones.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiForwardNsgReadRequest
*/
func (a *ForwardNsgAPIService) ForwardNsgRead(ctx context.Context, id string) ApiForwardNsgReadRequest {
	return ApiForwardNsgReadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ConfigReadForwardNSGResponse
func (a *ForwardNsgAPIService) ForwardNsgReadExecute(r ApiForwardNsgReadRequest) (*ConfigReadForwardNSGResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *ConfigReadForwardNSGResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ForwardNsgAPIService.ForwardNsgRead")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/forward_nsg/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
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

type ApiForwardNsgUpdateRequest struct {
	ctx context.Context
	ApiService ForwardNsgAPI
	id string
	body *ConfigForwardNSG
}

func (r ApiForwardNsgUpdateRequest) Body(body ConfigForwardNSG) ApiForwardNsgUpdateRequest {
	r.body = &body
	return r
}

func (r ApiForwardNsgUpdateRequest) Execute() (*ConfigUpdateForwardNSGResponse, *http.Response, error) {
	return r.ApiService.ForwardNsgUpdateExecute(r)
}

/*
ForwardNsgUpdate Update the ForwardNSG object.

Use this method to update a ForwardNSG object.
The _dns/forward_nsg_ object represents a Forward DNS Server Group for forward zones.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiForwardNsgUpdateRequest
*/
func (a *ForwardNsgAPIService) ForwardNsgUpdate(ctx context.Context, id string) ApiForwardNsgUpdateRequest {
	return ApiForwardNsgUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ConfigUpdateForwardNSGResponse
func (a *ForwardNsgAPIService) ForwardNsgUpdateExecute(r ApiForwardNsgUpdateRequest) (*ConfigUpdateForwardNSGResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *ConfigUpdateForwardNSGResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ForwardNsgAPIService.ForwardNsgUpdate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/forward_nsg/{id}"
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
