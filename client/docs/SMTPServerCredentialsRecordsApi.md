# \SMTPServerCredentialsRecordsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddStandardAccount**](SMTPServerCredentialsRecordsApi.md#AddStandardAccount) | **Post** /api/v1/accounts/standard | 
[**DeleteStandardAccount**](SMTPServerCredentialsRecordsApi.md#DeleteStandardAccount) | **Delete** /api/v1/accounts/standard/{standardAccountId} | 
[**ExportStandardAccounts**](SMTPServerCredentialsRecordsApi.md#ExportStandardAccounts) | **Post** /api/v1/accounts/standard/export | 
[**StandardAccountsGetAll**](SMTPServerCredentialsRecordsApi.md#StandardAccountsGetAll) | **Get** /api/v1/accounts/standard | 
[**StandardAccountsGetOneById**](SMTPServerCredentialsRecordsApi.md#StandardAccountsGetOneById) | **Get** /api/v1/accounts/standard/{standardAccountId} | 
[**UpdateStandardAccount**](SMTPServerCredentialsRecordsApi.md#UpdateStandardAccount) | **Put** /api/v1/accounts/standard/{standardAccountId} | 
[**ValidateStandardAccountName**](SMTPServerCredentialsRecordsApi.md#ValidateStandardAccountName) | **Post** /api/v1/accounts/standard/validateName | 



## AddStandardAccount

> StandardAccount AddStandardAccount(ctx).XApiVersion(xApiVersion).StandardAccount(standardAccount).Execute()



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
    standardAccount := *openapiclient.NewStandardAccountCreateSpec("Name_example", "Username_example") // StandardAccountCreateSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SMTPServerCredentialsRecordsApi.AddStandardAccount(context.Background()).XApiVersion(xApiVersion).StandardAccount(standardAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPServerCredentialsRecordsApi.AddStandardAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddStandardAccount`: StandardAccount
    fmt.Fprintf(os.Stdout, "Response from `SMTPServerCredentialsRecordsApi.AddStandardAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddStandardAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **standardAccount** | [**StandardAccountCreateSpec**](StandardAccountCreateSpec.md) |  | 

### Return type

[**StandardAccount**](StandardAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/csp-report
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStandardAccount

> DeleteStandardAccount(ctx, standardAccountId).XApiVersion(xApiVersion).Execute()



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
    standardAccountId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SMTPServerCredentialsRecordsApi.DeleteStandardAccount(context.Background(), standardAccountId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPServerCredentialsRecordsApi.DeleteStandardAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**standardAccountId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStandardAccountRequest struct via the builder pattern


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


## ExportStandardAccounts

> *os.File ExportStandardAccounts(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.SMTPServerCredentialsRecordsApi.ExportStandardAccounts(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPServerCredentialsRecordsApi.ExportStandardAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportStandardAccounts`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SMTPServerCredentialsRecordsApi.ExportStandardAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportStandardAccountsRequest struct via the builder pattern


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


## StandardAccountsGetAll

> StandardAccountsPage StandardAccountsGetAll(ctx).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()



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
    resp, r, err := api_client.SMTPServerCredentialsRecordsApi.StandardAccountsGetAll(context.Background()).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPServerCredentialsRecordsApi.StandardAccountsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandardAccountsGetAll`: StandardAccountsPage
    fmt.Fprintf(os.Stdout, "Response from `SMTPServerCredentialsRecordsApi.StandardAccountsGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandardAccountsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **searchPattern** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**StandardAccountsPage**](StandardAccountsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandardAccountsGetOneById

> StandardAccount StandardAccountsGetOneById(ctx, standardAccountId).XApiVersion(xApiVersion).Execute()



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
    standardAccountId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SMTPServerCredentialsRecordsApi.StandardAccountsGetOneById(context.Background(), standardAccountId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPServerCredentialsRecordsApi.StandardAccountsGetOneById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandardAccountsGetOneById`: StandardAccount
    fmt.Fprintf(os.Stdout, "Response from `SMTPServerCredentialsRecordsApi.StandardAccountsGetOneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**standardAccountId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardAccountsGetOneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**StandardAccount**](StandardAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStandardAccount

> StandardAccount UpdateStandardAccount(ctx, standardAccountId).XApiVersion(xApiVersion).StandardAccount(standardAccount).Execute()



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
    standardAccountId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    standardAccount := *openapiclient.NewStandardAccountUpdateSpec("Name_example") // StandardAccountUpdateSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SMTPServerCredentialsRecordsApi.UpdateStandardAccount(context.Background(), standardAccountId).XApiVersion(xApiVersion).StandardAccount(standardAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPServerCredentialsRecordsApi.UpdateStandardAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStandardAccount`: StandardAccount
    fmt.Fprintf(os.Stdout, "Response from `SMTPServerCredentialsRecordsApi.UpdateStandardAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**standardAccountId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStandardAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **standardAccount** | [**StandardAccountUpdateSpec**](StandardAccountUpdateSpec.md) |  | 

### Return type

[**StandardAccount**](StandardAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/csp-report
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateStandardAccountName

> []ValidationMessage ValidateStandardAccountName(ctx).XApiVersion(xApiVersion).ValidationSpec(validationSpec).Execute()



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
    validationSpec := *openapiclient.NewStandardAccountNameValidationSpec("Name_example") // StandardAccountNameValidationSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SMTPServerCredentialsRecordsApi.ValidateStandardAccountName(context.Background()).XApiVersion(xApiVersion).ValidationSpec(validationSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPServerCredentialsRecordsApi.ValidateStandardAccountName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateStandardAccountName`: []ValidationMessage
    fmt.Fprintf(os.Stdout, "Response from `SMTPServerCredentialsRecordsApi.ValidateStandardAccountName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateStandardAccountNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **validationSpec** | [**StandardAccountNameValidationSpec**](StandardAccountNameValidationSpec.md) |  | 

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

