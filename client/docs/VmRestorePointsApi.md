# \VmRestorePointsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestoreDisk**](VmRestorePointsApi.md#RestoreDisk) | **Post** /api/v1/vmRestorePoints/{vmRestorePointId}/restoreDisk | 
[**RestorePointsGetAll**](VmRestorePointsApi.md#RestorePointsGetAll) | **Get** /api/v1/vmRestorePoints | 
[**RestorePointsGetOneById**](VmRestorePointsApi.md#RestorePointsGetOneById) | **Get** /api/v1/vmRestorePoints/{vmRestorePointId} | 
[**RestoreVm**](VmRestorePointsApi.md#RestoreVm) | **Post** /api/v1/vmRestorePoints/{vmRestorePointId}/restoreVm | 
[**StartFlr**](VmRestorePointsApi.md#StartFlr) | **Post** /api/v1/vmRestorePoints/{vmRestorePointId}/restoreFile | 
[**ValidateRestoreDisk**](VmRestorePointsApi.md#ValidateRestoreDisk) | **Post** /api/v1/vmRestorePoints/{vmRestorePointId}/validateRestoreDisk | 
[**ValidateRestoreVm**](VmRestorePointsApi.md#ValidateRestoreVm) | **Post** /api/v1/vmRestorePoints/{vmRestorePointId}/validateRestoreVm | 
[**VmRestorePointsGetDefaultNetworkSettings**](VmRestorePointsApi.md#VmRestorePointsGetDefaultNetworkSettings) | **Get** /api/v1/vmRestorePoints/{vmRestorePointId}/defaultNetworkSettings | 



## RestoreDisk

> SessionLink RestoreDisk(ctx, vmRestorePointId).XApiVersion(xApiVersion).Options(options).Execute()



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
    vmRestorePointId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    options := *openapiclient.NewDiskBulkRestoreOptions([]openapiclient.DiskRestoreOptions{*openapiclient.NewDiskRestoreOptions("DiskId_example")}) // DiskBulkRestoreOptions |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmRestorePointsApi.RestoreDisk(context.Background(), vmRestorePointId).XApiVersion(xApiVersion).Options(options).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmRestorePointsApi.RestoreDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreDisk`: SessionLink
    fmt.Fprintf(os.Stdout, "Response from `VmRestorePointsApi.RestoreDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmRestorePointId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **options** | [**DiskBulkRestoreOptions**](DiskBulkRestoreOptions.md) |  | 

### Return type

[**SessionLink**](SessionLink.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/csp-report
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestorePointsGetAll

> VmRestorePointsPage RestorePointsGetAll(ctx).XApiVersion(xApiVersion).VirtualMachineId(virtualMachineId).Offset(offset).Limit(limit).Sort(sort).Execute()



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
    virtualMachineId := TODO // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmRestorePointsApi.RestorePointsGetAll(context.Background()).XApiVersion(xApiVersion).VirtualMachineId(virtualMachineId).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmRestorePointsApi.RestorePointsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestorePointsGetAll`: VmRestorePointsPage
    fmt.Fprintf(os.Stdout, "Response from `VmRestorePointsApi.RestorePointsGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestorePointsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **virtualMachineId** | [**string**](string.md) |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**VmRestorePointsPage**](VmRestorePointsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestorePointsGetOneById

> VmRestorePoint RestorePointsGetOneById(ctx, vmRestorePointId).XApiVersion(xApiVersion).Execute()



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
    vmRestorePointId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmRestorePointsApi.RestorePointsGetOneById(context.Background(), vmRestorePointId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmRestorePointsApi.RestorePointsGetOneById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestorePointsGetOneById`: VmRestorePoint
    fmt.Fprintf(os.Stdout, "Response from `VmRestorePointsApi.RestorePointsGetOneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmRestorePointId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestorePointsGetOneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**VmRestorePoint**](VmRestorePoint.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreVm

> SessionLink RestoreVm(ctx, vmRestorePointId).XApiVersion(xApiVersion).VmRestoreOptions(vmRestoreOptions).Execute()



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
    vmRestorePointId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    vmRestoreOptions := *openapiclient.NewVirtualMachineRestoreOptions() // VirtualMachineRestoreOptions |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmRestorePointsApi.RestoreVm(context.Background(), vmRestorePointId).XApiVersion(xApiVersion).VmRestoreOptions(vmRestoreOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmRestorePointsApi.RestoreVm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreVm`: SessionLink
    fmt.Fprintf(os.Stdout, "Response from `VmRestorePointsApi.RestoreVm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmRestorePointId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreVmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **vmRestoreOptions** | [**VirtualMachineRestoreOptions**](VirtualMachineRestoreOptions.md) |  | 

### Return type

[**SessionLink**](SessionLink.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/csp-report
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartFlr

> StartFlrResponse StartFlr(ctx, vmRestorePointId).FlrSpec(flrSpec).Execute()



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
    vmRestorePointId := TODO // string | 
    flrSpec := *openapiclient.NewFlrOptions() // FlrOptions |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmRestorePointsApi.StartFlr(context.Background(), vmRestorePointId).FlrSpec(flrSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmRestorePointsApi.StartFlr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartFlr`: StartFlrResponse
    fmt.Fprintf(os.Stdout, "Response from `VmRestorePointsApi.StartFlr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmRestorePointId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartFlrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flrSpec** | [**FlrOptions**](FlrOptions.md) |  | 

### Return type

[**StartFlrResponse**](StartFlrResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateRestoreDisk

> []ValidationMessage ValidateRestoreDisk(ctx, vmRestorePointId).XApiVersion(xApiVersion).ValidationConfig(validationConfig).Execute()



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
    vmRestorePointId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    validationConfig := *openapiclient.NewDiskBulkRestoreOptions([]openapiclient.DiskRestoreOptions{*openapiclient.NewDiskRestoreOptions("DiskId_example")}) // DiskBulkRestoreOptions |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmRestorePointsApi.ValidateRestoreDisk(context.Background(), vmRestorePointId).XApiVersion(xApiVersion).ValidationConfig(validationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmRestorePointsApi.ValidateRestoreDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateRestoreDisk`: []ValidationMessage
    fmt.Fprintf(os.Stdout, "Response from `VmRestorePointsApi.ValidateRestoreDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmRestorePointId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateRestoreDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **validationConfig** | [**DiskBulkRestoreOptions**](DiskBulkRestoreOptions.md) |  | 

### Return type

[**[]ValidationMessage**](ValidationMessage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/csp-report
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateRestoreVm

> []ValidationMessage ValidateRestoreVm(ctx, vmRestorePointId).XApiVersion(xApiVersion).ValidationConfig(validationConfig).Execute()



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
    vmRestorePointId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    validationConfig := *openapiclient.NewVirtualMachineRestoreOptions() // VirtualMachineRestoreOptions |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmRestorePointsApi.ValidateRestoreVm(context.Background(), vmRestorePointId).XApiVersion(xApiVersion).ValidationConfig(validationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmRestorePointsApi.ValidateRestoreVm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateRestoreVm`: []ValidationMessage
    fmt.Fprintf(os.Stdout, "Response from `VmRestorePointsApi.ValidateRestoreVm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmRestorePointId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateRestoreVmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **validationConfig** | [**VirtualMachineRestoreOptions**](VirtualMachineRestoreOptions.md) |  | 

### Return type

[**[]ValidationMessage**](ValidationMessage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/csp-report
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmRestorePointsGetDefaultNetworkSettings

> VmRestorePointDefaultNetworkSettings VmRestorePointsGetDefaultNetworkSettings(ctx, vmRestorePointId).XApiVersion(xApiVersion).Execute()



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
    vmRestorePointId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VmRestorePointsApi.VmRestorePointsGetDefaultNetworkSettings(context.Background(), vmRestorePointId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmRestorePointsApi.VmRestorePointsGetDefaultNetworkSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmRestorePointsGetDefaultNetworkSettings`: VmRestorePointDefaultNetworkSettings
    fmt.Fprintf(os.Stdout, "Response from `VmRestorePointsApi.VmRestorePointsGetDefaultNetworkSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmRestorePointId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmRestorePointsGetDefaultNetworkSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**VmRestorePointDefaultNetworkSettings**](VmRestorePointDefaultNetworkSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

