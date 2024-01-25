/*
DDI Keys API

The DDI Keys application is a BloxOne DDI service for managing TSIG keys and GSS-TSIG (Kerberos) keys which are used by other BloxOne DDI applications. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keys

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type GenerateTsigAPI interface {
	/*
		GenerateTsigGenerateTSIG Generate TSIG key with a random secret.

		Use this method to generate a TSIG key with a random secret using the specified algorithm.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGenerateTsigGenerateTSIGRequest
	*/
	GenerateTsigGenerateTSIG(ctx context.Context) ApiGenerateTsigGenerateTSIGRequest

	// GenerateTsigGenerateTSIGExecute executes the request
	//  @return KeysGenerateTSIGResponse
	GenerateTsigGenerateTSIGExecute(r ApiGenerateTsigGenerateTSIGRequest) (*KeysGenerateTSIGResponse, *http.Response, error)
}

// GenerateTsigAPIService GenerateTsigAPI service
type GenerateTsigAPIService internal.Service

type ApiGenerateTsigGenerateTSIGRequest struct {
	ctx        context.Context
	ApiService GenerateTsigAPI
	algorithm  *string
}

// The TSIG key algorithm.  Valid values are: * _hmac_sha256_ * _hmac_sha1_ * _hmac_sha224_ * _hmac_sha384_ * _hmac_sha512_  Defaults to _hmac_sha256_.
func (r ApiGenerateTsigGenerateTSIGRequest) Algorithm(algorithm string) ApiGenerateTsigGenerateTSIGRequest {
	r.algorithm = &algorithm
	return r
}

func (r ApiGenerateTsigGenerateTSIGRequest) Execute() (*KeysGenerateTSIGResponse, *http.Response, error) {
	return r.ApiService.GenerateTsigGenerateTSIGExecute(r)
}

/*
GenerateTsigGenerateTSIG Generate TSIG key with a random secret.

Use this method to generate a TSIG key with a random secret using the specified algorithm.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGenerateTsigGenerateTSIGRequest
*/
func (a *GenerateTsigAPIService) GenerateTsigGenerateTSIG(ctx context.Context) ApiGenerateTsigGenerateTSIGRequest {
	return ApiGenerateTsigGenerateTSIGRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return KeysGenerateTSIGResponse
func (a *GenerateTsigAPIService) GenerateTsigGenerateTSIGExecute(r ApiGenerateTsigGenerateTSIGRequest) (*KeysGenerateTSIGResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *KeysGenerateTSIGResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GenerateTsigAPIService.GenerateTsigGenerateTSIG")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/keys/generate_tsig"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.algorithm != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "algorithm", r.algorithm, "")
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
