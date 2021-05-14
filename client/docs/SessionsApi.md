# \SessionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportSessions**](SessionsApi.md#ExportSessions) | **Post** /api/v1/sessions/export | 
[**GetBackupSessionProtectedItems**](SessionsApi.md#GetBackupSessionProtectedItems) | **Get** /api/v1/sessions/{backupSessionId}/protectedItems | 
[**GetRestoreSessionRestoredItems**](SessionsApi.md#GetRestoreSessionRestoredItems) | **Get** /api/v1/sessions/{restoreSessionId}/restoredItems | 
[**GetRetentionSessionDeletedRestorePoints**](SessionsApi.md#GetRetentionSessionDeletedRestorePoints) | **Get** /api/v1/sessions/{retentionSessionId}/deletedItems | 
[**GetSessionById**](SessionsApi.md#GetSessionById) | **Get** /api/v1/sessions/{sessionId} | 
[**GetSessionLog**](SessionsApi.md#GetSessionLog) | **Get** /api/v1/sessions/{sessionId}/log | 
[**GetSessions**](SessionsApi.md#GetSessions) | **Get** /api/v1/sessions | 
[**StopJob**](SessionsApi.md#StopJob) | **Post** /api/v1/sessions/{sessionId}/stop | 



## ExportSessions

> *os.File ExportSessions(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.SessionsApi.ExportSessions(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.ExportSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSessions`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.ExportSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupSessionProtectedItems

> ProtectedItemsPage GetBackupSessionProtectedItems(ctx, backupSessionId).XApiVersion(xApiVersion).SearchPattern(searchPattern).Sort(sort).Offset(offset).Limit(limit).Execute()



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
    backupSessionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    searchPattern := "searchPattern_example" // string |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionsApi.GetBackupSessionProtectedItems(context.Background(), backupSessionId).XApiVersion(xApiVersion).SearchPattern(searchPattern).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetBackupSessionProtectedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackupSessionProtectedItems`: ProtectedItemsPage
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetBackupSessionProtectedItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupSessionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupSessionProtectedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **searchPattern** | **string** |  | 
 **sort** | **[]string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**ProtectedItemsPage**](ProtectedItemsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestoreSessionRestoredItems

> RestoredItemsPage GetRestoreSessionRestoredItems(ctx, restoreSessionId).XApiVersion(xApiVersion).SearchPattern(searchPattern).Sort(sort).Offset(offset).Limit(limit).Execute()



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
    restoreSessionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    searchPattern := "searchPattern_example" // string |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionsApi.GetRestoreSessionRestoredItems(context.Background(), restoreSessionId).XApiVersion(xApiVersion).SearchPattern(searchPattern).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetRestoreSessionRestoredItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestoreSessionRestoredItems`: RestoredItemsPage
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetRestoreSessionRestoredItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restoreSessionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestoreSessionRestoredItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **searchPattern** | **string** |  | 
 **sort** | **[]string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**RestoredItemsPage**](RestoredItemsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetentionSessionDeletedRestorePoints

> RetentionJobDeletedRestorePointsPage GetRetentionSessionDeletedRestorePoints(ctx, retentionSessionId).XApiVersion(xApiVersion).SearchPattern(searchPattern).Sort(sort).Offset(offset).Limit(limit).Execute()



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
    retentionSessionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    searchPattern := "searchPattern_example" // string |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionsApi.GetRetentionSessionDeletedRestorePoints(context.Background(), retentionSessionId).XApiVersion(xApiVersion).SearchPattern(searchPattern).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetRetentionSessionDeletedRestorePoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRetentionSessionDeletedRestorePoints`: RetentionJobDeletedRestorePointsPage
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetRetentionSessionDeletedRestorePoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**retentionSessionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetentionSessionDeletedRestorePointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **searchPattern** | **string** |  | 
 **sort** | **[]string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**RetentionJobDeletedRestorePointsPage**](RetentionJobDeletedRestorePointsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionById

> Session GetSessionById(ctx, sessionId).XApiVersion(xApiVersion).Execute()



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
    sessionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionsApi.GetSessionById(context.Background(), sessionId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetSessionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionById`: Session
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetSessionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionLog

> SessionLog GetSessionLog(ctx, sessionId).XApiVersion(xApiVersion).Execute()



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
    sessionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionsApi.GetSessionLog(context.Background(), sessionId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetSessionLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionLog`: SessionLog
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetSessionLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**SessionLog**](SessionLog.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessions

> SessionsPage GetSessions(ctx).XApiVersion(xApiVersion).Offset(offset).Limit(limit).PolicyId(policyId).Status(status).Types(types).FromUtc(fromUtc).ToUtc(toUtc).Usn(usn).Sort(sort).VmId(vmId).Execute()



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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    policyId := TODO // string |  (optional)
    status := []string{"Status_example"} // []string |  (optional)
    types := []string{"Types_example"} // []string |  (optional)
    fromUtc := time.Now() // time.Time |  (optional)
    toUtc := time.Now() // time.Time |  (optional)
    usn := int64(789) // int64 |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)
    vmId := "vmId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionsApi.GetSessions(context.Background()).XApiVersion(xApiVersion).Offset(offset).Limit(limit).PolicyId(policyId).Status(status).Types(types).FromUtc(fromUtc).ToUtc(toUtc).Usn(usn).Sort(sort).VmId(vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessions`: SessionsPage
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **policyId** | [**string**](string.md) |  | 
 **status** | **[]string** |  | 
 **types** | **[]string** |  | 
 **fromUtc** | **time.Time** |  | 
 **toUtc** | **time.Time** |  | 
 **usn** | **int64** |  | 
 **sort** | **[]string** |  | 
 **vmId** | **string** |  | 

### Return type

[**SessionsPage**](SessionsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopJob

> StopJob(ctx, sessionId).XApiVersion(xApiVersion).Execute()



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
    sessionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionsApi.StopJob(context.Background(), sessionId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.StopJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

