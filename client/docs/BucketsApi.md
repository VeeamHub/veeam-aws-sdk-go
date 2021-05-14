# \BucketsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BucketById**](BucketsApi.md#BucketById) | **Get** /api/v1/cloudInfrastructure/buckets/{bucketId} | 
[**Buckets**](BucketsApi.md#Buckets) | **Get** /api/v1/cloudInfrastructure/buckets | 
[**GetBucketFolder**](BucketsApi.md#GetBucketFolder) | **Get** /api/v1/cloudInfrastructure/buckets/{bucketId}/folders/{folderName} | 
[**GetBucketFolders**](BucketsApi.md#GetBucketFolders) | **Get** /api/v1/cloudInfrastructure/buckets/{bucketId}/folders | 
[**RescanBucket**](BucketsApi.md#RescanBucket) | **Post** /api/v1/cloudInfrastructure/buckets/rescan/{amazonAccountId} | 



## BucketById

> Bucket BucketById(ctx, bucketId).XApiVersion(xApiVersion).Execute()



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
    bucketId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BucketsApi.BucketById(context.Background(), bucketId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.BucketById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketById`: Bucket
    fmt.Fprintf(os.Stdout, "Response from `BucketsApi.BucketById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**Bucket**](Bucket.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Buckets

> BucketsPage Buckets(ctx).XApiVersion(xApiVersion).AmazonAccountId(amazonAccountId).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()



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
    amazonAccountId := TODO // string |  (optional)
    searchPattern := "searchPattern_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BucketsApi.Buckets(context.Background()).XApiVersion(xApiVersion).AmazonAccountId(amazonAccountId).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.Buckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Buckets`: BucketsPage
    fmt.Fprintf(os.Stdout, "Response from `BucketsApi.Buckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **amazonAccountId** | [**string**](string.md) |  | 
 **searchPattern** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**BucketsPage**](BucketsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBucketFolder

> BucketFolder GetBucketFolder(ctx, bucketId, folderName).XApiVersion(xApiVersion).Execute()



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
    bucketId := TODO // string | 
    folderName := "folderName_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BucketsApi.GetBucketFolder(context.Background(), bucketId, folderName).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.GetBucketFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucketFolder`: BucketFolder
    fmt.Fprintf(os.Stdout, "Response from `BucketsApi.GetBucketFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | [**string**](.md) |  | 
**folderName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**BucketFolder**](BucketFolder.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBucketFolders

> BucketFoldersPage GetBucketFolders(ctx, bucketId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).Sort(sort).Execute()



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
    bucketId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BucketsApi.GetBucketFolders(context.Background(), bucketId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.GetBucketFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucketFolders`: BucketFoldersPage
    fmt.Fprintf(os.Stdout, "Response from `BucketsApi.GetBucketFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**BucketFoldersPage**](BucketFoldersPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RescanBucket

> SessionLink RescanBucket(ctx, amazonAccountId).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.BucketsApi.RescanBucket(context.Background(), amazonAccountId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.RescanBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RescanBucket`: SessionLink
    fmt.Fprintf(os.Stdout, "Response from `BucketsApi.RescanBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonAccountId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRescanBucketRequest struct via the builder pattern


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

