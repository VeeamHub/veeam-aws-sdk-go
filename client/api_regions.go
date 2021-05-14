/*
 * Veeam Backup for AWS public API 1.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0-rev0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// RegionsApiService RegionsApi service
type RegionsApiService service

type ApiGetAllRegionsRequest struct {
	ctx _context.Context
	ApiService *RegionsApiService
	xApiVersion *string
	searchPattern *string
	offset *int32
	limit *int32
	sort *[]string
}

func (r ApiGetAllRegionsRequest) XApiVersion(xApiVersion string) ApiGetAllRegionsRequest {
	r.xApiVersion = &xApiVersion
	return r
}
func (r ApiGetAllRegionsRequest) SearchPattern(searchPattern string) ApiGetAllRegionsRequest {
	r.searchPattern = &searchPattern
	return r
}
func (r ApiGetAllRegionsRequest) Offset(offset int32) ApiGetAllRegionsRequest {
	r.offset = &offset
	return r
}
func (r ApiGetAllRegionsRequest) Limit(limit int32) ApiGetAllRegionsRequest {
	r.limit = &limit
	return r
}
func (r ApiGetAllRegionsRequest) Sort(sort []string) ApiGetAllRegionsRequest {
	r.sort = &sort
	return r
}

func (r ApiGetAllRegionsRequest) Execute() (RegionsPage, *_nethttp.Response, error) {
	return r.ApiService.GetAllRegionsExecute(r)
}

/*
 * GetAllRegions Method for GetAllRegions
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetAllRegionsRequest
 */
func (a *RegionsApiService) GetAllRegions(ctx _context.Context) ApiGetAllRegionsRequest {
	return ApiGetAllRegionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return RegionsPage
 */
func (a *RegionsApiService) GetAllRegionsExecute(r ApiGetAllRegionsRequest) (RegionsPage, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RegionsPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegionsApiService.GetAllRegions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/cloudInfrastructure/regions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}

	if r.searchPattern != nil {
		localVarQueryParams.Add("SearchPattern", parameterToString(*r.searchPattern, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("Offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("Limit", parameterToString(*r.limit, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("Sort", parameterToString(*r.sort, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-api-version"] = parameterToString(*r.xApiVersion, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetInnerRegionsRequest struct {
	ctx _context.Context
	ApiService *RegionsApiService
	regionId string
	xApiVersion *string
	searchPattern *string
	offset *int32
	limit *int32
	sort *[]string
}

func (r ApiGetInnerRegionsRequest) XApiVersion(xApiVersion string) ApiGetInnerRegionsRequest {
	r.xApiVersion = &xApiVersion
	return r
}
func (r ApiGetInnerRegionsRequest) SearchPattern(searchPattern string) ApiGetInnerRegionsRequest {
	r.searchPattern = &searchPattern
	return r
}
func (r ApiGetInnerRegionsRequest) Offset(offset int32) ApiGetInnerRegionsRequest {
	r.offset = &offset
	return r
}
func (r ApiGetInnerRegionsRequest) Limit(limit int32) ApiGetInnerRegionsRequest {
	r.limit = &limit
	return r
}
func (r ApiGetInnerRegionsRequest) Sort(sort []string) ApiGetInnerRegionsRequest {
	r.sort = &sort
	return r
}

func (r ApiGetInnerRegionsRequest) Execute() (RegionsPage, *_nethttp.Response, error) {
	return r.ApiService.GetInnerRegionsExecute(r)
}

/*
 * GetInnerRegions Method for GetInnerRegions
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param regionId
 * @return ApiGetInnerRegionsRequest
 */
func (a *RegionsApiService) GetInnerRegions(ctx _context.Context, regionId string) ApiGetInnerRegionsRequest {
	return ApiGetInnerRegionsRequest{
		ApiService: a,
		ctx: ctx,
		regionId: regionId,
	}
}

/*
 * Execute executes the request
 * @return RegionsPage
 */
func (a *RegionsApiService) GetInnerRegionsExecute(r ApiGetInnerRegionsRequest) (RegionsPage, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RegionsPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegionsApiService.GetInnerRegions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/cloudInfrastructure/regions/{regionId}/zones"
	localVarPath = strings.Replace(localVarPath, "{"+"regionId"+"}", _neturl.PathEscape(parameterToString(r.regionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}

	if r.searchPattern != nil {
		localVarQueryParams.Add("SearchPattern", parameterToString(*r.searchPattern, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("Offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("Limit", parameterToString(*r.limit, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("Sort", parameterToString(*r.sort, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-api-version"] = parameterToString(*r.xApiVersion, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetInstancesTypesByRegionRequest struct {
	ctx _context.Context
	ApiService *RegionsApiService
	regionId string
	xApiVersion *string
	offset *int32
	limit *int32
	isArm64 *bool
	virtualizationTypes *[]string
}

func (r ApiGetInstancesTypesByRegionRequest) XApiVersion(xApiVersion string) ApiGetInstancesTypesByRegionRequest {
	r.xApiVersion = &xApiVersion
	return r
}
func (r ApiGetInstancesTypesByRegionRequest) Offset(offset int32) ApiGetInstancesTypesByRegionRequest {
	r.offset = &offset
	return r
}
func (r ApiGetInstancesTypesByRegionRequest) Limit(limit int32) ApiGetInstancesTypesByRegionRequest {
	r.limit = &limit
	return r
}
func (r ApiGetInstancesTypesByRegionRequest) IsArm64(isArm64 bool) ApiGetInstancesTypesByRegionRequest {
	r.isArm64 = &isArm64
	return r
}
func (r ApiGetInstancesTypesByRegionRequest) VirtualizationTypes(virtualizationTypes []string) ApiGetInstancesTypesByRegionRequest {
	r.virtualizationTypes = &virtualizationTypes
	return r
}

func (r ApiGetInstancesTypesByRegionRequest) Execute() (AmazonInstanceTypesByRegionPage, *_nethttp.Response, error) {
	return r.ApiService.GetInstancesTypesByRegionExecute(r)
}

/*
 * GetInstancesTypesByRegion Method for GetInstancesTypesByRegion
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param regionId
 * @return ApiGetInstancesTypesByRegionRequest
 */
func (a *RegionsApiService) GetInstancesTypesByRegion(ctx _context.Context, regionId string) ApiGetInstancesTypesByRegionRequest {
	return ApiGetInstancesTypesByRegionRequest{
		ApiService: a,
		ctx: ctx,
		regionId: regionId,
	}
}

/*
 * Execute executes the request
 * @return AmazonInstanceTypesByRegionPage
 */
func (a *RegionsApiService) GetInstancesTypesByRegionExecute(r ApiGetInstancesTypesByRegionRequest) (AmazonInstanceTypesByRegionPage, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AmazonInstanceTypesByRegionPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegionsApiService.GetInstancesTypesByRegion")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/cloudInfrastructure/regions/{regionId}/instancesTypes"
	localVarPath = strings.Replace(localVarPath, "{"+"regionId"+"}", _neturl.PathEscape(parameterToString(r.regionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}

	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.isArm64 != nil {
		localVarQueryParams.Add("isArm64", parameterToString(*r.isArm64, ""))
	}
	if r.virtualizationTypes != nil {
		localVarQueryParams.Add("virtualizationTypes", parameterToString(*r.virtualizationTypes, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-api-version"] = parameterToString(*r.xApiVersion, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetRegionByIdRequest struct {
	ctx _context.Context
	ApiService *RegionsApiService
	regionId string
	xApiVersion *string
}

func (r ApiGetRegionByIdRequest) XApiVersion(xApiVersion string) ApiGetRegionByIdRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiGetRegionByIdRequest) Execute() (Region, *_nethttp.Response, error) {
	return r.ApiService.GetRegionByIdExecute(r)
}

/*
 * GetRegionById Method for GetRegionById
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param regionId
 * @return ApiGetRegionByIdRequest
 */
func (a *RegionsApiService) GetRegionById(ctx _context.Context, regionId string) ApiGetRegionByIdRequest {
	return ApiGetRegionByIdRequest{
		ApiService: a,
		ctx: ctx,
		regionId: regionId,
	}
}

/*
 * Execute executes the request
 * @return Region
 */
func (a *RegionsApiService) GetRegionByIdExecute(r ApiGetRegionByIdRequest) (Region, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Region
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegionsApiService.GetRegionById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/cloudInfrastructure/regions/{regionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"regionId"+"}", _neturl.PathEscape(parameterToString(r.regionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-api-version"] = parameterToString(*r.xApiVersion, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
