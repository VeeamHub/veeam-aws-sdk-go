# \AmazonConnectionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrowseAvailabilityZones**](AmazonConnectionsApi.md#BrowseAvailabilityZones) | **Get** /api/v1/amazonConnections/{amazonConnectionId}/availabilityZones | 
[**BrowseBucketFolders**](AmazonConnectionsApi.md#BrowseBucketFolders) | **Get** /api/v1/amazonConnections/{amazonConnectionId}/browseBucketFolders/{bucketId} | 
[**BrowseCloudNetworks**](AmazonConnectionsApi.md#BrowseCloudNetworks) | **Get** /api/v1/amazonConnections/{amazonConnectionId}/cloudNetworks | 
[**BrowseCloudSecurityGroups**](AmazonConnectionsApi.md#BrowseCloudSecurityGroups) | **Get** /api/v1/amazonConnections/{amazonConnectionId}/cloudSecurityGroups | 
[**BrowseCloudSubnetworks**](AmazonConnectionsApi.md#BrowseCloudSubnetworks) | **Get** /api/v1/amazonConnections/{amazonConnectionId}/cloudSubnetworks | 
[**BrowseEncryptionKeys**](AmazonConnectionsApi.md#BrowseEncryptionKeys) | **Get** /api/v1/amazonConnections/{amazonConnectionId}/cloudEncryptionKeys | 
[**CreateAmazonConnection**](AmazonConnectionsApi.md#CreateAmazonConnection) | **Post** /api/v1/amazonConnections | 
[**DeleteAmazonConnection**](AmazonConnectionsApi.md#DeleteAmazonConnection) | **Delete** /api/v1/amazonConnections/{amazonConnectionId} | 
[**GetAllAmazonConnections**](AmazonConnectionsApi.md#GetAllAmazonConnections) | **Get** /api/v1/amazonConnections | 
[**GetAmazonConnectionById**](AmazonConnectionsApi.md#GetAmazonConnectionById) | **Get** /api/v1/amazonConnections/{amazonConnectionId} | 
[**ValidateS3Endpoint**](AmazonConnectionsApi.md#ValidateS3Endpoint) | **Get** /api/v1/amazonConnections/{amazonConnectionId}/validateS3endpoint | 



## BrowseAvailabilityZones

> RegionsPage BrowseAvailabilityZones(ctx, amazonConnectionId).XApiVersion(xApiVersion).Execute()



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
    amazonConnectionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AmazonConnectionsApi.BrowseAvailabilityZones(context.Background(), amazonConnectionId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmazonConnectionsApi.BrowseAvailabilityZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrowseAvailabilityZones`: RegionsPage
    fmt.Fprintf(os.Stdout, "Response from `AmazonConnectionsApi.BrowseAvailabilityZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonConnectionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrowseAvailabilityZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

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


## BrowseBucketFolders

> BucketFoldersPage BrowseBucketFolders(ctx, amazonConnectionId, bucketId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).Sort(sort).Execute()



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
    amazonConnectionId := TODO // string | 
    bucketId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sort := []string{"Sort_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AmazonConnectionsApi.BrowseBucketFolders(context.Background(), amazonConnectionId, bucketId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmazonConnectionsApi.BrowseBucketFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrowseBucketFolders`: BucketFoldersPage
    fmt.Fprintf(os.Stdout, "Response from `AmazonConnectionsApi.BrowseBucketFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonConnectionId** | [**string**](.md) |  | 
**bucketId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrowseBucketFoldersRequest struct via the builder pattern


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


## BrowseCloudNetworks

> CloudNetworksPage BrowseCloudNetworks(ctx, amazonConnectionId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).Execute()



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
    amazonConnectionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AmazonConnectionsApi.BrowseCloudNetworks(context.Background(), amazonConnectionId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmazonConnectionsApi.BrowseCloudNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrowseCloudNetworks`: CloudNetworksPage
    fmt.Fprintf(os.Stdout, "Response from `AmazonConnectionsApi.BrowseCloudNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonConnectionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrowseCloudNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CloudNetworksPage**](CloudNetworksPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrowseCloudSecurityGroups

> CloudSecurityGroupsPage BrowseCloudSecurityGroups(ctx, amazonConnectionId).XApiVersion(xApiVersion).CloudNetworkId(cloudNetworkId).Offset(offset).Limit(limit).Execute()



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
    amazonConnectionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    cloudNetworkId := "cloudNetworkId_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AmazonConnectionsApi.BrowseCloudSecurityGroups(context.Background(), amazonConnectionId).XApiVersion(xApiVersion).CloudNetworkId(cloudNetworkId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmazonConnectionsApi.BrowseCloudSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrowseCloudSecurityGroups`: CloudSecurityGroupsPage
    fmt.Fprintf(os.Stdout, "Response from `AmazonConnectionsApi.BrowseCloudSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonConnectionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrowseCloudSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **cloudNetworkId** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CloudSecurityGroupsPage**](CloudSecurityGroupsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrowseCloudSubnetworks

> CloudSubnetworksPage BrowseCloudSubnetworks(ctx, amazonConnectionId).XApiVersion(xApiVersion).CloudNetworkId(cloudNetworkId).AvailabilityZone(availabilityZone).Offset(offset).Limit(limit).Execute()



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
    amazonConnectionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    cloudNetworkId := "cloudNetworkId_example" // string |  (optional)
    availabilityZone := TODO // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AmazonConnectionsApi.BrowseCloudSubnetworks(context.Background(), amazonConnectionId).XApiVersion(xApiVersion).CloudNetworkId(cloudNetworkId).AvailabilityZone(availabilityZone).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmazonConnectionsApi.BrowseCloudSubnetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrowseCloudSubnetworks`: CloudSubnetworksPage
    fmt.Fprintf(os.Stdout, "Response from `AmazonConnectionsApi.BrowseCloudSubnetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonConnectionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrowseCloudSubnetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **cloudNetworkId** | **string** |  | 
 **availabilityZone** | [**string**](string.md) |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CloudSubnetworksPage**](CloudSubnetworksPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrowseEncryptionKeys

> CloudEncryptionKeysPage BrowseEncryptionKeys(ctx, amazonConnectionId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).Execute()



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
    amazonConnectionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AmazonConnectionsApi.BrowseEncryptionKeys(context.Background(), amazonConnectionId).XApiVersion(xApiVersion).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmazonConnectionsApi.BrowseEncryptionKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrowseEncryptionKeys`: CloudEncryptionKeysPage
    fmt.Fprintf(os.Stdout, "Response from `AmazonConnectionsApi.BrowseEncryptionKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonConnectionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrowseEncryptionKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CloudEncryptionKeysPage**](CloudEncryptionKeysPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAmazonConnection

> AmazonConnection CreateAmazonConnection(ctx).XApiVersion(xApiVersion).AmazonConnectionSpec(amazonConnectionSpec).Execute()



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
    amazonConnectionSpec := *openapiclient.NewAmazonConnectionSpec("RegionId_example") // AmazonConnectionSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AmazonConnectionsApi.CreateAmazonConnection(context.Background()).XApiVersion(xApiVersion).AmazonConnectionSpec(amazonConnectionSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmazonConnectionsApi.CreateAmazonConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAmazonConnection`: AmazonConnection
    fmt.Fprintf(os.Stdout, "Response from `AmazonConnectionsApi.CreateAmazonConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAmazonConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **amazonConnectionSpec** | [**AmazonConnectionSpec**](AmazonConnectionSpec.md) |  | 

### Return type

[**AmazonConnection**](AmazonConnection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAmazonConnection

> DeleteAmazonConnection(ctx, amazonConnectionId).XApiVersion(xApiVersion).Execute()



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
    amazonConnectionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AmazonConnectionsApi.DeleteAmazonConnection(context.Background(), amazonConnectionId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmazonConnectionsApi.DeleteAmazonConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonConnectionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAmazonConnectionRequest struct via the builder pattern


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


## GetAllAmazonConnections

> AmazonConnectionsPage GetAllAmazonConnections(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.AmazonConnectionsApi.GetAllAmazonConnections(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmazonConnectionsApi.GetAllAmazonConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllAmazonConnections`: AmazonConnectionsPage
    fmt.Fprintf(os.Stdout, "Response from `AmazonConnectionsApi.GetAllAmazonConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAmazonConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**AmazonConnectionsPage**](AmazonConnectionsPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmazonConnectionById

> AmazonConnection GetAmazonConnectionById(ctx, amazonConnectionId).XApiVersion(xApiVersion).Execute()



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
    amazonConnectionId := TODO // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AmazonConnectionsApi.GetAmazonConnectionById(context.Background(), amazonConnectionId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmazonConnectionsApi.GetAmazonConnectionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmazonConnectionById`: AmazonConnection
    fmt.Fprintf(os.Stdout, "Response from `AmazonConnectionsApi.GetAmazonConnectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonConnectionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAmazonConnectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**AmazonConnection**](AmazonConnection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateS3Endpoint

> CloudSubnetS3EndpointValidationResult ValidateS3Endpoint(ctx, amazonConnectionId).SubnetId(subnetId).XApiVersion(xApiVersion).Execute()



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
    amazonConnectionId := TODO // string | 
    subnetId := "subnetId_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AmazonConnectionsApi.ValidateS3Endpoint(context.Background(), amazonConnectionId).SubnetId(subnetId).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmazonConnectionsApi.ValidateS3Endpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateS3Endpoint`: CloudSubnetS3EndpointValidationResult
    fmt.Fprintf(os.Stdout, "Response from `AmazonConnectionsApi.ValidateS3Endpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amazonConnectionId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateS3EndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subnetId** | **string** |  | 
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**CloudSubnetS3EndpointValidationResult**](CloudSubnetS3EndpointValidationResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

