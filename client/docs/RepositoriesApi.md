# \RepositoriesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRepository**](RepositoriesApi.md#CreateRepository) | **Post** /api/v1/repositories | 
[**DeleteRepository**](RepositoriesApi.md#DeleteRepository) | **Delete** /api/v1/repositories/{repositoryId} | 
[**ExportRepositories**](RepositoriesApi.md#ExportRepositories) | **Post** /api/v1/repositories/export | 
[**RepositoriesGetAll**](RepositoriesApi.md#RepositoriesGetAll) | **Get** /api/v1/repositories | 
[**RepositoriesGetOneById**](RepositoriesApi.md#RepositoriesGetOneById) | **Get** /api/v1/repositories/{repositoryId} | 
[**RepositoryPasswordIsCorrect**](RepositoriesApi.md#RepositoryPasswordIsCorrect) | **Post** /api/v1/repositories/{repositoryId}/checkPassword | 
[**RepositoryValidateSettings**](RepositoriesApi.md#RepositoryValidateSettings) | **Post** /api/v1/repositories/settings/validate | 
[**RescanRepository**](RepositoriesApi.md#RescanRepository) | **Post** /api/v1/repositories/{repositoryId}/rescan | 
[**UpdateRepository**](RepositoriesApi.md#UpdateRepository) | **Put** /api/v1/repositories/{repositoryId} | 



## CreateRepository

> Repository CreateRepository(ctx).XApiVersion(xApiVersion).Repository(repository).Execute()



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
    repository := *openapiclient.NewRepositoryCreateSpec("Name_example", "AmazonAccountId_example", "AmazonBucketId_example", "AmazonStorageFolder_example", false) // RepositoryCreateSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.CreateRepository(context.Background()).XApiVersion(xApiVersion).Repository(repository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.CreateRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRepository`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.CreateRepository`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **repository** | [**RepositoryCreateSpec**](RepositoryCreateSpec.md) |  | 

### Return type

[**Repository**](Repository.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/csp-report
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRepository

> DeleteRepository(ctx, repositoryId).XApiVersion(xApiVersion).Execute()



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
    repositoryId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.DeleteRepository(context.Background(), repositoryId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.DeleteRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryRequest struct via the builder pattern


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


## ExportRepositories

> *os.File ExportRepositories(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.RepositoriesApi.ExportRepositories(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.ExportRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportRepositories`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.ExportRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportRepositoriesRequest struct via the builder pattern


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


## RepositoriesGetAll

> RepositoriesPage RepositoriesGetAll(ctx).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Sort(sort).Limit(limit).Execute()



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
    sort := []string{"Sort_example"} // []string |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesGetAll(context.Background()).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Sort(sort).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesGetAll`: RepositoriesPage
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **searchPattern** | **string** |  | 
 **offset** | **int32** |  | 
 **sort** | **[]string** |  | 
 **limit** | **int32** |  | 

### Return type

[**RepositoriesPage**](RepositoriesPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesGetOneById

> Repository RepositoriesGetOneById(ctx, repositoryId).XApiVersion(xApiVersion).Execute()



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
    repositoryId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoriesGetOneById(context.Background(), repositoryId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoriesGetOneById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesGetOneById`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoriesGetOneById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesGetOneByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**Repository**](Repository.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryPasswordIsCorrect

> RepositoryPasswordValidationResult RepositoryPasswordIsCorrect(ctx, repositoryId).XApiVersion(xApiVersion).RepositoryPassword(repositoryPassword).Execute()



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
    repositoryId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    repositoryPassword := *openapiclient.NewRepositoryPasswordValidationSpec("Password_example") // RepositoryPasswordValidationSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoryPasswordIsCorrect(context.Background(), repositoryId).XApiVersion(xApiVersion).RepositoryPassword(repositoryPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoryPasswordIsCorrect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryPasswordIsCorrect`: RepositoryPasswordValidationResult
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoryPasswordIsCorrect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryPasswordIsCorrectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **repositoryPassword** | [**RepositoryPasswordValidationSpec**](RepositoryPasswordValidationSpec.md) |  | 

### Return type

[**RepositoryPasswordValidationResult**](RepositoryPasswordValidationResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryValidateSettings

> []ValidationMessage RepositoryValidateSettings(ctx).XApiVersion(xApiVersion).Repository(repository).Execute()



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
    repository := *openapiclient.NewRepositoryCreateSpec("Name_example", "AmazonAccountId_example", "AmazonBucketId_example", "AmazonStorageFolder_example", false) // RepositoryCreateSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RepositoryValidateSettings(context.Background()).XApiVersion(xApiVersion).Repository(repository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RepositoryValidateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoryValidateSettings`: []ValidationMessage
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RepositoryValidateSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryValidateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **repository** | [**RepositoryCreateSpec**](RepositoryCreateSpec.md) |  | 

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


## RescanRepository

> SessionLink RescanRepository(ctx, repositoryId).XApiVersion(xApiVersion).Execute()



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
    repositoryId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.RescanRepository(context.Background(), repositoryId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RescanRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RescanRepository`: SessionLink
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RescanRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRescanRepositoryRequest struct via the builder pattern


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


## UpdateRepository

> Repository UpdateRepository(ctx, repositoryId).XApiVersion(xApiVersion).Repository(repository).Execute()



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
    repositoryId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    repository := *openapiclient.NewRepositoryUpdateSpec("Name_example", "AmazonAccountId_example", false) // RepositoryUpdateSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.UpdateRepository(context.Background(), repositoryId).XApiVersion(xApiVersion).Repository(repository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.UpdateRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRepository`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.UpdateRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **repository** | [**RepositoryUpdateSpec**](RepositoryUpdateSpec.md) |  | 

### Return type

[**Repository**](Repository.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/csp-report
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

