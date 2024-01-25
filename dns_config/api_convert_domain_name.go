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

type ConvertDomainNameAPI interface {
	/*
		ConvertDomainNameConvert Convert the object.

		Use this method to convert between Internationalized Domain Name (IDN) and ASCII domain name (Punycode).

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param domainName Input domain name in either of IDN or punycode representations.
		@return ApiConvertDomainNameConvertRequest
	*/
	ConvertDomainNameConvert(ctx context.Context, domainName string) ApiConvertDomainNameConvertRequest

	// ConvertDomainNameConvertExecute executes the request
	//  @return ConfigConvertDomainNameResponse
	ConvertDomainNameConvertExecute(r ApiConvertDomainNameConvertRequest) (*ConfigConvertDomainNameResponse, *http.Response, error)
}

// ConvertDomainNameAPIService ConvertDomainNameAPI service
type ConvertDomainNameAPIService internal.Service

type ApiConvertDomainNameConvertRequest struct {
	ctx        context.Context
	ApiService ConvertDomainNameAPI
	domainName string
}

func (r ApiConvertDomainNameConvertRequest) Execute() (*ConfigConvertDomainNameResponse, *http.Response, error) {
	return r.ApiService.ConvertDomainNameConvertExecute(r)
}

/*
ConvertDomainNameConvert Convert the object.

Use this method to convert between Internationalized Domain Name (IDN) and ASCII domain name (Punycode).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param domainName Input domain name in either of IDN or punycode representations.
 @return ApiConvertDomainNameConvertRequest
*/
func (a *ConvertDomainNameAPIService) ConvertDomainNameConvert(ctx context.Context, domainName string) ApiConvertDomainNameConvertRequest {
	return ApiConvertDomainNameConvertRequest{
		ApiService: a,
		ctx:        ctx,
		domainName: domainName,
	}
}

// Execute executes the request
//  @return ConfigConvertDomainNameResponse
func (a *ConvertDomainNameAPIService) ConvertDomainNameConvertExecute(r ApiConvertDomainNameConvertRequest) (*ConfigConvertDomainNameResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ConfigConvertDomainNameResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ConvertDomainNameAPIService.ConvertDomainNameConvert")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/convert_domain_name/{domain_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"domain_name"+"}", url.PathEscape(internal.ParameterValueToString(r.domainName, "domainName")), -1)

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
