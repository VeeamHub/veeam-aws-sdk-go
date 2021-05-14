# \IAMRolesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAmazonAccount**](IAMRolesApi.md#AddAmazonAccount) | **Post** /api/v1/accounts/amazon | 
[**AmazonAccountsGetAll**](IAMRolesApi.md#AmazonAccountsGetAll) | **Get** /api/v1/accounts/amazon | 
[**AmazonAccountsGetOneById**](IAMRolesApi.md#AmazonAccountsGetOneById) | **Get** /api/v1/accounts/amazon/{amazonAccountId} | 
[**DeleteAmazonAccount**](IAMRolesApi.md#DeleteAmazonAccount) | **Delete** /api/v1/accounts/amazon/{amazonAccountId} | 
[**ExportAmazonAccounts**](IAMRolesApi.md#ExportAmazonAccounts) | **Post** /api/v1/accounts/amazon/export | 
[**RescanAmazonAccount**](IAMRolesApi.md#RescanAmazonAccount) | **Post** /api/v1/accounts/amazon/{amazonAccountId}/rescan | 
[**UpdateAmazonAccount**](IAMRolesApi.md#UpdateAmazonAccount) | **Put** /api/v1/accounts/amazon/{amazonAccountId} | 
[**ValidateAmazonAccountName**](IAMRolesApi.md#ValidateAmazonAccountName) | **Post** /api/v1/accounts/amazon/validateName | 
[**ValidateDeleteAmazonAccount**](IAMRolesApi.md#ValidateDeleteAmazonAccount) | **Post** /api/v1/accounts/amazon/{amazonAccountId}/validateDelete | 



## AddAmazonAccount

> AmazonAccount AddAmazonAccount(ctx).XApiVersion(xApiVersion).AmazonAccount(amazonAccount).Execute()



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
    amazonAccount := *openapiclient.NewAmazonAccountCreateSpec("Name_example") // AmazonAccountCreateSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IAMRolesApi.AddAmazonAccount(context.Background()).XApiVersion(xApiVersion).AmazonAccount(amazonAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMRolesApi.AddAmazonAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAmazonAccount`: AmazonAccount
    fmt.Fprintf(os.Stdout, "Response from `IAMRolesApi.AddAmazonAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAmazonAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **amazonAccount** | [**AmazonAccountCreateSpec**](AmazonAccountCreateSpec.md) |  | 

### Return type

[**AmazonAccount**](AmazonAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/csp-report, application/problem+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AmazonAccountsGetAll

> AmazonAccountsPage AmazonAccountsGetAll(ctx).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()



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
    resp, r, err := api_client.IAMRolesApi.AmazonAccountsGetAll(context.Background()).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMRolesApi.AmazonAccountsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AmazonAccountsGetAll`: AmazonAccountsPage
    fmt.Fprintf(os.Stdout, "Response from `IAMRolesApi.AmazonAccountsGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAmazonAccountsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **searchPattern** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**AmazonAccountsPage**](AmazonAccountsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AmazonAccountsGetOneById

> AmazonAccount AmazonAccountsGetOneById(ctx, amazonAccountId).XApiVersion(xApiVersion).Execute()



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
    amazonAccountId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IAMRolesApi.AmazonAccountsGetOneById(context.Background(), amazonAccountId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMRolesApi.AmazonAccountsGetOneById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AmazonAccountsGetOneById`: AmazonAccount
    fmt.Fprintf(os.Stdout, "Response from `IAMRolesApi.AmazonAccountsGetOneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonAccountId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAmazonAccountsGetOneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**AmazonAccount**](AmazonAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAmazonAccount

> DeleteAmazonAccount(ctx, amazonAccountId).XApiVersion(xApiVersion).Execute()



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
    amazonAccountId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IAMRolesApi.DeleteAmazonAccount(context.Background(), amazonAccountId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMRolesApi.DeleteAmazonAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonAccountId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAmazonAccountRequest struct via the builder pattern


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


## ExportAmazonAccounts

> *os.File ExportAmazonAccounts(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.IAMRolesApi.ExportAmazonAccounts(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMRolesApi.ExportAmazonAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAmazonAccounts`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `IAMRolesApi.ExportAmazonAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAmazonAccountsRequest struct via the builder pattern


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


## RescanAmazonAccount

> SessionLink RescanAmazonAccount(ctx, amazonAccountId).XApiVersion(xApiVersion).RescanCloudAccountSpec(rescanCloudAccountSpec).Execute()



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
    amazonAccountId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    rescanCloudAccountSpec := *openapiclient.NewRescanCloudAccountSpec() // RescanCloudAccountSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IAMRolesApi.RescanAmazonAccount(context.Background(), amazonAccountId).XApiVersion(xApiVersion).RescanCloudAccountSpec(rescanCloudAccountSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMRolesApi.RescanAmazonAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RescanAmazonAccount`: SessionLink
    fmt.Fprintf(os.Stdout, "Response from `IAMRolesApi.RescanAmazonAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonAccountId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRescanAmazonAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **rescanCloudAccountSpec** | [**RescanCloudAccountSpec**](RescanCloudAccountSpec.md) |  | 

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


## UpdateAmazonAccount

> AmazonAccount UpdateAmazonAccount(ctx, amazonAccountId).XApiVersion(xApiVersion).AmazonAccount(amazonAccount).Execute()



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
    amazonAccountId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    amazonAccount := *openapiclient.NewAmazonAccountUpdateSpec("Name_example") // AmazonAccountUpdateSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IAMRolesApi.UpdateAmazonAccount(context.Background(), amazonAccountId).XApiVersion(xApiVersion).AmazonAccount(amazonAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMRolesApi.UpdateAmazonAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAmazonAccount`: AmazonAccount
    fmt.Fprintf(os.Stdout, "Response from `IAMRolesApi.UpdateAmazonAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonAccountId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAmazonAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **amazonAccount** | [**AmazonAccountUpdateSpec**](AmazonAccountUpdateSpec.md) |  | 

### Return type

[**AmazonAccount**](AmazonAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/csp-report
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAmazonAccountName

> []ValidationMessage ValidateAmazonAccountName(ctx).XApiVersion(xApiVersion).ValidationSpec(validationSpec).Execute()



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
    validationSpec := *openapiclient.NewAmazonAccountNameValidationSpec("Name_example") // AmazonAccountNameValidationSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IAMRolesApi.ValidateAmazonAccountName(context.Background()).XApiVersion(xApiVersion).ValidationSpec(validationSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMRolesApi.ValidateAmazonAccountName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateAmazonAccountName`: []ValidationMessage
    fmt.Fprintf(os.Stdout, "Response from `IAMRolesApi.ValidateAmazonAccountName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAmazonAccountNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **validationSpec** | [**AmazonAccountNameValidationSpec**](AmazonAccountNameValidationSpec.md) |  | 

### Return type

[**[]ValidationMessage**](ValidationMessage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateDeleteAmazonAccount

> []ValidationMessage ValidateDeleteAmazonAccount(ctx, amazonAccountId).XApiVersion(xApiVersion).Execute()



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
    amazonAccountId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IAMRolesApi.ValidateDeleteAmazonAccount(context.Background(), amazonAccountId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMRolesApi.ValidateDeleteAmazonAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateDeleteAmazonAccount`: []ValidationMessage
    fmt.Fprintf(os.Stdout, "Response from `IAMRolesApi.ValidateDeleteAmazonAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonAccountId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateDeleteAmazonAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**[]ValidationMessage**](ValidationMessage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

