# \PoliciesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicy**](PoliciesApi.md#CreatePolicy) | **Post** /api/v1/policies | 
[**DeletePolicy**](PoliciesApi.md#DeletePolicy) | **Delete** /api/v1/policies/{policyId} | 
[**DisablePolicy**](PoliciesApi.md#DisablePolicy) | **Post** /api/v1/policies/{policyId}/disable | 
[**EnablePolicy**](PoliciesApi.md#EnablePolicy) | **Post** /api/v1/policies/{policyId}/enable | 
[**ExportPolicies**](PoliciesApi.md#ExportPolicies) | **Post** /api/v1/policies/export | 
[**ExportPolicy**](PoliciesApi.md#ExportPolicy) | **Get** /api/v1/policies/{policyId}/export | 
[**GetPolicyCost**](PoliciesApi.md#GetPolicyCost) | **Get** /api/v1/policies/{policyId}/cost | 
[**ImportPolicy**](PoliciesApi.md#ImportPolicy) | **Post** /api/v1/policies/import | 
[**PoliciesGetAll**](PoliciesApi.md#PoliciesGetAll) | **Get** /api/v1/policies | 
[**PoliciesGetExcludedItemsById**](PoliciesApi.md#PoliciesGetExcludedItemsById) | **Get** /api/v1/policies/{policyId}/excludedItems | 
[**PoliciesGetOneById**](PoliciesApi.md#PoliciesGetOneById) | **Get** /api/v1/policies/{policyId} | 
[**PoliciesGetRegionsById**](PoliciesApi.md#PoliciesGetRegionsById) | **Get** /api/v1/policies/{policyId}/regions | 
[**PoliciesGetSelectedItemsById**](PoliciesApi.md#PoliciesGetSelectedItemsById) | **Get** /api/v1/policies/{policyId}/selectedItems | 
[**PolicyOrderPriorityOfPolicies**](PoliciesApi.md#PolicyOrderPriorityOfPolicies) | **Post** /api/v1/policies/orderPriorities | 
[**StartPolicy**](PoliciesApi.md#StartPolicy) | **Post** /api/v1/policies/{policyId}/start | 
[**StopPolicy**](PoliciesApi.md#StopPolicy) | **Post** /api/v1/policies/{policyId}/stop | 
[**UpdatePolicy**](PoliciesApi.md#UpdatePolicy) | **Put** /api/v1/policies/{policyId} | 
[**Validate**](PoliciesApi.md#Validate) | **Post** /api/v1/policies/validate | 



## CreatePolicy

> Policy CreatePolicy(ctx).XApiVersion(xApiVersion).Policy(policy).Execute()



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
    policy := *openapiclient.NewPolicySpec("Name_example", "AmazonAccountId_example", openapiclient.PolicySelectionTypes("AllItems")) // PolicySpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.CreatePolicy(context.Background()).XApiVersion(xApiVersion).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.CreatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicy`: Policy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **policy** | [**PolicySpec**](PolicySpec.md) |  | 

### Return type

[**Policy**](Policy.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/csp-report
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> DeletePolicy(ctx, policyId).XApiVersion(xApiVersion).Execute()



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
    policyId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.DeletePolicy(context.Background(), policyId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.DeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


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


## DisablePolicy

> DisablePolicy(ctx, policyId).XApiVersion(xApiVersion).Execute()



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
    policyId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.DisablePolicy(context.Background(), policyId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.DisablePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisablePolicyRequest struct via the builder pattern


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


## EnablePolicy

> EnablePolicy(ctx, policyId).XApiVersion(xApiVersion).Execute()



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
    policyId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.EnablePolicy(context.Background(), policyId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.EnablePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnablePolicyRequest struct via the builder pattern


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


## ExportPolicies

> *os.File ExportPolicies(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.PoliciesApi.ExportPolicies(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.ExportPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPolicies`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.ExportPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportPoliciesRequest struct via the builder pattern


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


## ExportPolicy

> PolicyExportModel ExportPolicy(ctx, policyId).XApiVersion(xApiVersion).Execute()



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
    policyId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.ExportPolicy(context.Background(), policyId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.ExportPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPolicy`: PolicyExportModel
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.ExportPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**PolicyExportModel**](PolicyExportModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyCost

> CostEstimationsPage GetPolicyCost(ctx, policyId).XApiVersion(xApiVersion).Execute()



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
    policyId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.GetPolicyCost(context.Background(), policyId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.GetPolicyCost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyCost`: CostEstimationsPage
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.GetPolicyCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

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


## ImportPolicy

> SessionLink ImportPolicy(ctx).XApiVersion(xApiVersion).PolicySpec(policySpec).Execute()



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
    policySpec := *openapiclient.NewPolicyExportModel([]string{"Regions_example"}, "Name_example", "AmazonAccountName_example", false, openapiclient.PolicySelectionTypes("AllItems")) // PolicyExportModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.ImportPolicy(context.Background()).XApiVersion(xApiVersion).PolicySpec(policySpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.ImportPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportPolicy`: SessionLink
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.ImportPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **policySpec** | [**PolicyExportModel**](PolicyExportModel.md) |  | 

### Return type

[**SessionLink**](SessionLink.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGetAll

> PoliciesPage PoliciesGetAll(ctx).XApiVersion(xApiVersion).SearchPattern(searchPattern).VirtualMachineId(virtualMachineId).LastPolicySessionStatus(lastPolicySessionStatus).PolicyStatus(policyStatus).Sort(sort).Usn(usn).RepositoryId(repositoryId).Offset(offset).Limit(limit).Execute()



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
    virtualMachineId := TODO // string |  (optional)
    lastPolicySessionStatus := []string{"LastPolicySessionStatus_example"} // []string |  (optional)
    policyStatus := []string{"PolicyStatus_example"} // []string |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)
    usn := int64(789) // int64 |  (optional)
    repositoryId := TODO // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesGetAll(context.Background()).XApiVersion(xApiVersion).SearchPattern(searchPattern).VirtualMachineId(virtualMachineId).LastPolicySessionStatus(lastPolicySessionStatus).PolicyStatus(policyStatus).Sort(sort).Usn(usn).RepositoryId(repositoryId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetAll`: PoliciesPage
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **searchPattern** | **string** |  | 
 **virtualMachineId** | [**string**](string.md) |  | 
 **lastPolicySessionStatus** | **[]string** |  | 
 **policyStatus** | **[]string** |  | 
 **sort** | **[]string** |  | 
 **usn** | **int64** |  | 
 **repositoryId** | [**string**](string.md) |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**PoliciesPage**](PoliciesPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGetExcludedItemsById

> PolicyItemsPage PoliciesGetExcludedItemsById(ctx, policyId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).SearchPattern(searchPattern).Sort(sort).Execute()



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
    policyId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    searchPattern := "searchPattern_example" // string |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesGetExcludedItemsById(context.Background(), policyId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).SearchPattern(searchPattern).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesGetExcludedItemsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetExcludedItemsById`: PolicyItemsPage
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesGetExcludedItemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetExcludedItemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **searchPattern** | **string** |  | 
 **sort** | **[]string** |  | 

### Return type

[**PolicyItemsPage**](PolicyItemsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGetOneById

> Policy PoliciesGetOneById(ctx, policyId).XApiVersion(xApiVersion).Execute()



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
    policyId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesGetOneById(context.Background(), policyId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesGetOneById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetOneById`: Policy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesGetOneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetOneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**Policy**](Policy.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGetRegionsById

> RegionsPage PoliciesGetRegionsById(ctx, policyId).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()



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
    policyId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    searchPattern := "searchPattern_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesGetRegionsById(context.Background(), policyId).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesGetRegionsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetRegionsById`: RegionsPage
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesGetRegionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetRegionsByIdRequest struct via the builder pattern


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


## PoliciesGetSelectedItemsById

> PolicyItemsPage PoliciesGetSelectedItemsById(ctx, policyId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).SearchPattern(searchPattern).Sort(sort).Execute()



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
    policyId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    searchPattern := "searchPattern_example" // string |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesGetSelectedItemsById(context.Background(), policyId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).SearchPattern(searchPattern).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesGetSelectedItemsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetSelectedItemsById`: PolicyItemsPage
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesGetSelectedItemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetSelectedItemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **searchPattern** | **string** |  | 
 **sort** | **[]string** |  | 

### Return type

[**PolicyItemsPage**](PolicyItemsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyOrderPriorityOfPolicies

> map[string]interface{} PolicyOrderPriorityOfPolicies(ctx).XApiVersion(xApiVersion).PoliciesPrioritySpec(policiesPrioritySpec).Execute()



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
    policiesPrioritySpec := *openapiclient.NewPoliciesPrioritySpec([]openapiclient.PolicyPrioritySpec{*openapiclient.NewPolicyPrioritySpec("PolicyId_example", int32(123))}) // PoliciesPrioritySpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PolicyOrderPriorityOfPolicies(context.Background()).XApiVersion(xApiVersion).PoliciesPrioritySpec(policiesPrioritySpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PolicyOrderPriorityOfPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyOrderPriorityOfPolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PolicyOrderPriorityOfPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPolicyOrderPriorityOfPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **policiesPrioritySpec** | [**PoliciesPrioritySpec**](PoliciesPrioritySpec.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartPolicy

> SessionLink StartPolicy(ctx, policyId).XApiVersion(xApiVersion).Execute()



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
    policyId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.StartPolicy(context.Background(), policyId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.StartPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartPolicy`: SessionLink
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.StartPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartPolicyRequest struct via the builder pattern


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


## StopPolicy

> StopPolicy(ctx, policyId).XApiVersion(xApiVersion).Execute()



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
    policyId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.StopPolicy(context.Background(), policyId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.StopPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopPolicyRequest struct via the builder pattern


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


## UpdatePolicy

> Policy UpdatePolicy(ctx, policyId).XApiVersion(xApiVersion).Policy(policy).Execute()



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
    policyId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    policy := *openapiclient.NewPolicy("Id_example", []string{"RegionIds_example"}, "Name_example", "Description_example", int64(123), "AmazonAccountId_example", false, openapiclient.PolicySelectionTypes("AllItems")) // Policy |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.UpdatePolicy(context.Background(), policyId).XApiVersion(xApiVersion).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.UpdatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicy`: Policy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.UpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **policy** | [**Policy**](Policy.md) |  | 

### Return type

[**Policy**](Policy.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/csp-report
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Validate

> []ValidationMessage Validate(ctx).XApiVersion(xApiVersion).ValidationSpec(validationSpec).Execute()



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
    validationSpec := *openapiclient.NewPolicySpec("Name_example", "AmazonAccountId_example", openapiclient.PolicySelectionTypes("AllItems")) // PolicySpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.Validate(context.Background()).XApiVersion(xApiVersion).ValidationSpec(validationSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.Validate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Validate`: []ValidationMessage
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.Validate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **validationSpec** | [**PolicySpec**](PolicySpec.md) |  | 

### Return type

[**[]ValidationMessage**](ValidationMessage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

