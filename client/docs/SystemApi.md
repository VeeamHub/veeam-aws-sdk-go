# \SystemApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemDateTimeInfo**](SystemApi.md#GetSystemDateTimeInfo) | **Get** /api/v1/system/time | 
[**GetSystemLogs**](SystemApi.md#GetSystemLogs) | **Get** /api/v1/system/logs | 



## GetSystemDateTimeInfo

> DateTimeInfo GetSystemDateTimeInfo(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.SystemApi.GetSystemDateTimeInfo(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSystemDateTimeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemDateTimeInfo`: DateTimeInfo
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetSystemDateTimeInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemDateTimeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**DateTimeInfo**](DateTimeInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemLogs

> *os.File GetSystemLogs(ctx).XApiVersion(xApiVersion).ExportLogsType(exportLogsType).Days(days).FromDateUtc(fromDateUtc).ToDateUtc(toDateUtc).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    exportLogsType := "exportLogsType_example" // string |  (optional)
    days := int32(56) // int32 |  (optional)
    fromDateUtc := time.Now() // time.Time |  (optional)
    toDateUtc := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemApi.GetSystemLogs(context.Background()).XApiVersion(xApiVersion).ExportLogsType(exportLogsType).Days(days).FromDateUtc(fromDateUtc).ToDateUtc(toDateUtc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSystemLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemLogs`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetSystemLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **exportLogsType** | **string** |  | 
 **days** | **int32** |  | 
 **fromDateUtc** | **time.Time** |  | 
 **toDateUtc** | **time.Time** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

