# \RegionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllRegions**](RegionsApi.md#GetAllRegions) | **Get** /api/v1/cloudInfrastructure/regions | 
[**GetInnerRegions**](RegionsApi.md#GetInnerRegions) | **Get** /api/v1/cloudInfrastructure/regions/{regionId}/zones | 
[**GetInstancesTypesByRegion**](RegionsApi.md#GetInstancesTypesByRegion) | **Get** /api/v1/cloudInfrastructure/regions/{regionId}/instancesTypes | 
[**GetRegionById**](RegionsApi.md#GetRegionById) | **Get** /api/v1/cloudInfrastructure/regions/{regionId} | 



## GetAllRegions

> RegionsPage GetAllRegions(ctx).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    searchPattern := "searchPattern_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionsApi.GetAllRegions(context.Background()).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsApi.GetAllRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRegions`: RegionsPage
    fmt.Fprintf(os.Stdout, "Response from `RegionsApi.GetAllRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **searchPattern** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**RegionsPage**](RegionsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInnerRegions

> RegionsPage GetInnerRegions(ctx, regionId).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    regionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    searchPattern := "searchPattern_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionsApi.GetInnerRegions(context.Background(), regionId).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsApi.GetInnerRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInnerRegions`: RegionsPage
    fmt.Fprintf(os.Stdout, "Response from `RegionsApi.GetInnerRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInnerRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **searchPattern** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**RegionsPage**](RegionsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstancesTypesByRegion

> AmazonInstanceTypesByRegionPage GetInstancesTypesByRegion(ctx, regionId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).IsArm64(isArm64).VirtualizationTypes(virtualizationTypes).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    regionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    isArm64 := true // bool |  (optional)
    virtualizationTypes := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionsApi.GetInstancesTypesByRegion(context.Background(), regionId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).IsArm64(isArm64).VirtualizationTypes(virtualizationTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsApi.GetInstancesTypesByRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstancesTypesByRegion`: AmazonInstanceTypesByRegionPage
    fmt.Fprintf(os.Stdout, "Response from `RegionsApi.GetInstancesTypesByRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancesTypesByRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **isArm64** | **bool** |  | 
 **virtualizationTypes** | **[]string** |  | 

### Return type

[**AmazonInstanceTypesByRegionPage**](AmazonInstanceTypesByRegionPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegionById

> Region GetRegionById(ctx, regionId).XApiVersion(xApiVersion).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    regionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionsApi.GetRegionById(context.Background(), regionId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsApi.GetRegionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegionById`: Region
    fmt.Fprintf(os.Stdout, "Response from `RegionsApi.GetRegionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**Region**](Region.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

