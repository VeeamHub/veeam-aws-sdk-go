# \CostEstimationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportCostEstimation**](CostEstimationApi.md#ExportCostEstimation) | **Post** /api/v1/costEstimation/export | 
[**GetCostEstimation**](CostEstimationApi.md#GetCostEstimation) | **Post** /api/v1/costEstimation | 



## ExportCostEstimation

> *os.File ExportCostEstimation(ctx).XApiVersion(xApiVersion).PolicySpec(policySpec).Execute()



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
    policySpec := *openapiclient.NewPolicySpec("Name_example", "AmazonAccountId_example", openapiclient.PolicySelectionTypes("AllItems")) // PolicySpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CostEstimationApi.ExportCostEstimation(context.Background()).XApiVersion(xApiVersion).PolicySpec(policySpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CostEstimationApi.ExportCostEstimation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportCostEstimation`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CostEstimationApi.ExportCostEstimation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportCostEstimationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **policySpec** | [**PolicySpec**](PolicySpec.md) |  | 

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


## GetCostEstimation

> CostEstimationsPage GetCostEstimation(ctx).XApiVersion(xApiVersion).PolicySpec(policySpec).VirtualMachineNameFilter(virtualMachineNameFilter).Offset(offset).Limit(limit).Sort(sort).Execute()



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
    policySpec := *openapiclient.NewPolicySpec("Name_example", "AmazonAccountId_example", openapiclient.PolicySelectionTypes("AllItems")) // PolicySpec | 
    virtualMachineNameFilter := "virtualMachineNameFilter_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CostEstimationApi.GetCostEstimation(context.Background()).XApiVersion(xApiVersion).PolicySpec(policySpec).VirtualMachineNameFilter(virtualMachineNameFilter).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CostEstimationApi.GetCostEstimation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCostEstimation`: CostEstimationsPage
    fmt.Fprintf(os.Stdout, "Response from `CostEstimationApi.GetCostEstimation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCostEstimationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **policySpec** | [**PolicySpec**](PolicySpec.md) |  | 
 **virtualMachineNameFilter** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**CostEstimationsPage**](CostEstimationsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

