# \TagsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Tag**](TagsApi.md#Tag) | **Get** /api/v1/cloudInfrastructure/tags/{tagId} | 
[**Tags**](TagsApi.md#Tags) | **Get** /api/v1/cloudInfrastructure/tags | 



## Tag

> Tag Tag(ctx, tagId).XApiVersion(xApiVersion).Execute()



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
    tagId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.Tag(context.Background(), tagId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.Tag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Tag`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.Tag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**Tag**](Tag.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Tags

> TagsPage Tags(ctx).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()



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
    resp, r, err := api_client.TagsApi.Tags(context.Background()).XApiVersion(xApiVersion).SearchPattern(searchPattern).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.Tags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Tags`: TagsPage
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.Tags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **searchPattern** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**TagsPage**](TagsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

