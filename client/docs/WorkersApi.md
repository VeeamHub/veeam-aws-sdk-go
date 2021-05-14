# \WorkersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegionSettings**](WorkersApi.md#CreateRegionSettings) | **Post** /api/v1/settings/regions | 
[**DeleteRegionSettings**](WorkersApi.md#DeleteRegionSettings) | **Delete** /api/v1/settings/regions/{regionId} | 
[**GetAllRegionSettings**](WorkersApi.md#GetAllRegionSettings) | **Get** /api/v1/settings/regions | 
[**GetRegionSettings**](WorkersApi.md#GetRegionSettings) | **Get** /api/v1/settings/regions/{regionId} | 
[**GetServiceAccountSettings**](WorkersApi.md#GetServiceAccountSettings) | **Get** /api/v1/settings/serviceAccount | 
[**UpdateRegionSettings**](WorkersApi.md#UpdateRegionSettings) | **Put** /api/v1/settings/regions/{regionId} | 
[**UpdateServiceAccountSettings**](WorkersApi.md#UpdateServiceAccountSettings) | **Put** /api/v1/settings/serviceAccount | 



## CreateRegionSettings

> RegionSettings CreateRegionSettings(ctx).XApiVersion(xApiVersion).RegionConfigurationSpec(regionConfigurationSpec).Execute()



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
    regionConfigurationSpec := *openapiclient.NewRegionSettingsSpec("AvailableZoneId_example", "CloudNetworkId_example", "CloudSubnetworkId_example", "CloudSecurityGroupId_example") // RegionSettingsSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkersApi.CreateRegionSettings(context.Background()).XApiVersion(xApiVersion).RegionConfigurationSpec(regionConfigurationSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersApi.CreateRegionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegionSettings`: RegionSettings
    fmt.Fprintf(os.Stdout, "Response from `WorkersApi.CreateRegionSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **regionConfigurationSpec** | [**RegionSettingsSpec**](RegionSettingsSpec.md) |  | 

### Return type

[**RegionSettings**](RegionSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegionSettings

> map[string]interface{} DeleteRegionSettings(ctx, regionId).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.WorkersApi.DeleteRegionSettings(context.Background(), regionId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersApi.DeleteRegionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRegionSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WorkersApi.DeleteRegionSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRegionSettings

> RegionSettingsPage GetAllRegionSettings(ctx).XApiVersion(xApiVersion).Offset(offset).Limit(limit).Execute()



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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkersApi.GetAllRegionSettings(context.Background()).XApiVersion(xApiVersion).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersApi.GetAllRegionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRegionSettings`: RegionSettingsPage
    fmt.Fprintf(os.Stdout, "Response from `WorkersApi.GetAllRegionSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRegionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**RegionSettingsPage**](RegionSettingsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegionSettings

> RegionSettings GetRegionSettings(ctx, regionId).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.WorkersApi.GetRegionSettings(context.Background(), regionId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersApi.GetRegionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegionSettings`: RegionSettings
    fmt.Fprintf(os.Stdout, "Response from `WorkersApi.GetRegionSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**RegionSettings**](RegionSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccountSettings

> ServiceAccountSettings GetServiceAccountSettings(ctx).XApiVersion(xApiVersion).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkersApi.GetServiceAccountSettings(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersApi.GetServiceAccountSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAccountSettings`: ServiceAccountSettings
    fmt.Fprintf(os.Stdout, "Response from `WorkersApi.GetServiceAccountSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**ServiceAccountSettings**](ServiceAccountSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRegionSettings

> RegionSettings UpdateRegionSettings(ctx, regionId).XApiVersion(xApiVersion).LocationConfigurationSpec(locationConfigurationSpec).Execute()



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
    locationConfigurationSpec := *openapiclient.NewRegionSettings("RegionId_example", "CloudNetworkId_example", "CloudSubnetworkId_example", "CloudSecurityGroupId_example") // RegionSettings | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkersApi.UpdateRegionSettings(context.Background(), regionId).XApiVersion(xApiVersion).LocationConfigurationSpec(locationConfigurationSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersApi.UpdateRegionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRegionSettings`: RegionSettings
    fmt.Fprintf(os.Stdout, "Response from `WorkersApi.UpdateRegionSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **locationConfigurationSpec** | [**RegionSettings**](RegionSettings.md) |  | 

### Return type

[**RegionSettings**](RegionSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceAccountSettings

> ServiceAccountSettings UpdateServiceAccountSettings(ctx).XApiVersion(xApiVersion).ConfigurationSpec(configurationSpec).Execute()



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
    configurationSpec := *openapiclient.NewServiceAccountSettings("ServiceAmazonAccountId_example") // ServiceAccountSettings | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkersApi.UpdateServiceAccountSettings(context.Background()).XApiVersion(xApiVersion).ConfigurationSpec(configurationSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkersApi.UpdateServiceAccountSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceAccountSettings`: ServiceAccountSettings
    fmt.Fprintf(os.Stdout, "Response from `WorkersApi.UpdateServiceAccountSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAccountSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **configurationSpec** | [**ServiceAccountSettings**](ServiceAccountSettings.md) |  | 

### Return type

[**ServiceAccountSettings**](ServiceAccountSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

