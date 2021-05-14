# \InstancesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVirtualMachineBackups**](InstancesApi.md#DeleteVirtualMachineBackups) | **Post** /api/v1/virtualMachines/{virtualMachineId}/deleteBackups | 
[**DeleteVirtualMachineSnapshots**](InstancesApi.md#DeleteVirtualMachineSnapshots) | **Post** /api/v1/virtualMachines/{virtualMachineId}/deleteSnapshots | 
[**ExportVirtualMachines**](InstancesApi.md#ExportVirtualMachines) | **Post** /api/v1/virtualMachines/export | 
[**TakeSnapshot**](InstancesApi.md#TakeSnapshot) | **Post** /api/v1/virtualMachines/{virtualMachineId}/takeSnapshot | 
[**VirtualMachineById**](InstancesApi.md#VirtualMachineById) | **Get** /api/v1/virtualMachines/{virtualMachineId} | 
[**VirtualMachines**](InstancesApi.md#VirtualMachines) | **Get** /api/v1/virtualMachines | 



## DeleteVirtualMachineBackups

> SessionLink DeleteVirtualMachineBackups(ctx, virtualMachineId).XApiVersion(xApiVersion).Execute()



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
    virtualMachineId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.DeleteVirtualMachineBackups(context.Background(), virtualMachineId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.DeleteVirtualMachineBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualMachineBackups`: SessionLink
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.DeleteVirtualMachineBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualMachineBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**SessionLink**](SessionLink.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualMachineSnapshots

> SessionLink DeleteVirtualMachineSnapshots(ctx, virtualMachineId).XApiVersion(xApiVersion).Execute()



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
    virtualMachineId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.DeleteVirtualMachineSnapshots(context.Background(), virtualMachineId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.DeleteVirtualMachineSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualMachineSnapshots`: SessionLink
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.DeleteVirtualMachineSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualMachineSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**SessionLink**](SessionLink.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportVirtualMachines

> *os.File ExportVirtualMachines(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.InstancesApi.ExportVirtualMachines(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ExportVirtualMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportVirtualMachines`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ExportVirtualMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportVirtualMachinesRequest struct via the builder pattern


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


## TakeSnapshot

> SessionLink TakeSnapshot(ctx, virtualMachineId).XApiVersion(xApiVersion).SnapshotSpec(snapshotSpec).Execute()



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
    virtualMachineId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    snapshotSpec := *openapiclient.NewVirtualMachineSnapshotSpec("AmazonAccountId_example") // VirtualMachineSnapshotSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.TakeSnapshot(context.Background(), virtualMachineId).XApiVersion(xApiVersion).SnapshotSpec(snapshotSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.TakeSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TakeSnapshot`: SessionLink
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.TakeSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTakeSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **snapshotSpec** | [**VirtualMachineSnapshotSpec**](VirtualMachineSnapshotSpec.md) |  | 

### Return type

[**SessionLink**](SessionLink.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirtualMachineById

> VirtualMachine VirtualMachineById(ctx, virtualMachineId).XApiVersion(xApiVersion).Execute()



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
    virtualMachineId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.VirtualMachineById(context.Background(), virtualMachineId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.VirtualMachineById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirtualMachineById`: VirtualMachine
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.VirtualMachineById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirtualMachineByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**VirtualMachine**](VirtualMachine.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirtualMachines

> VirtualMachinesPage VirtualMachines(ctx).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).ProtectionStatus(protectionStatus).BuckupDestination(buckupDestination).BackupState(backupState).Execute()



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
    protectionStatus := []string{"ProtectionStatus_example"} // []string |  (optional)
    buckupDestination := []string{"BuckupDestination_example"} // []string |  (optional)
    backupState := "backupState_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.VirtualMachines(context.Background()).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).ProtectionStatus(protectionStatus).BuckupDestination(buckupDestination).BackupState(backupState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.VirtualMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirtualMachines`: VirtualMachinesPage
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.VirtualMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **searchPattern** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 
 **protectionStatus** | **[]string** |  | 
 **buckupDestination** | **[]string** |  | 
 **backupState** | **string** |  | 

### Return type

[**VirtualMachinesPage**](VirtualMachinesPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

