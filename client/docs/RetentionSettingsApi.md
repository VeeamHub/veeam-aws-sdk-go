# \RetentionSettingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRetentionSettings**](RetentionSettingsApi.md#GetRetentionSettings) | **Get** /api/v1/settings/retentionSettings | 
[**SetDefaultRetentionSettings**](RetentionSettingsApi.md#SetDefaultRetentionSettings) | **Post** /api/v1/settings/retentionSettings/reset | 
[**UpdateRetentionSettings**](RetentionSettingsApi.md#UpdateRetentionSettings) | **Post** /api/v1/settings/retentionSettings | 



## GetRetentionSettings

> RetentionSettings GetRetentionSettings(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.RetentionSettingsApi.GetRetentionSettings(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionSettingsApi.GetRetentionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRetentionSettings`: RetentionSettings
    fmt.Fprintf(os.Stdout, "Response from `RetentionSettingsApi.GetRetentionSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRetentionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**RetentionSettings**](RetentionSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultRetentionSettings

> RetentionSettings SetDefaultRetentionSettings(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.RetentionSettingsApi.SetDefaultRetentionSettings(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionSettingsApi.SetDefaultRetentionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDefaultRetentionSettings`: RetentionSettings
    fmt.Fprintf(os.Stdout, "Response from `RetentionSettingsApi.SetDefaultRetentionSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultRetentionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**RetentionSettings**](RetentionSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRetentionSettings

> RetentionSettings UpdateRetentionSettings(ctx).XApiVersion(xApiVersion).RetentionSpec(retentionSpec).Execute()



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
    retentionSpec := *openapiclient.NewRetentionSettings(false, false) // RetentionSettings | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RetentionSettingsApi.UpdateRetentionSettings(context.Background()).XApiVersion(xApiVersion).RetentionSpec(retentionSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionSettingsApi.UpdateRetentionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRetentionSettings`: RetentionSettings
    fmt.Fprintf(os.Stdout, "Response from `RetentionSettingsApi.UpdateRetentionSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRetentionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **retentionSpec** | [**RetentionSettings**](RetentionSettings.md) |  | 

### Return type

[**RetentionSettings**](RetentionSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

