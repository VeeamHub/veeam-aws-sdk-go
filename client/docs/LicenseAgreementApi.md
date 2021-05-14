# \LicenseAgreementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptLicenseAgreement**](LicenseAgreementApi.md#AcceptLicenseAgreement) | **Post** /api/v1/licenseAgreement/accept | 
[**DownloadLicenseAgreementFile**](LicenseAgreementApi.md#DownloadLicenseAgreementFile) | **Get** /api/v1/licenseAgreement/download | 
[**GetCurrentLicenseAgreementAccepted**](LicenseAgreementApi.md#GetCurrentLicenseAgreementAccepted) | **Get** /api/v1/licenseAgreement/accepted | 
[**GetLicenseAgreement**](LicenseAgreementApi.md#GetLicenseAgreement) | **Get** /api/v1/licenseAgreement | 



## AcceptLicenseAgreement

> AcceptLicenseAgreement(ctx).EulaChecksum(eulaChecksum).ThirdPartyLicenseAgreementChecksum(thirdPartyLicenseAgreementChecksum).XApiVersion(xApiVersion).Execute()



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
    eulaChecksum := "eulaChecksum_example" // string | 
    thirdPartyLicenseAgreementChecksum := "thirdPartyLicenseAgreementChecksum_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseAgreementApi.AcceptLicenseAgreement(context.Background()).EulaChecksum(eulaChecksum).ThirdPartyLicenseAgreementChecksum(thirdPartyLicenseAgreementChecksum).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseAgreementApi.AcceptLicenseAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcceptLicenseAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eulaChecksum** | **string** |  | 
 **thirdPartyLicenseAgreementChecksum** | **string** |  | 
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


## DownloadLicenseAgreementFile

> *os.File DownloadLicenseAgreementFile(ctx).LicenseAgreementType(licenseAgreementType).XApiVersion(xApiVersion).Execute()



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
    licenseAgreementType := "licenseAgreementType_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (default to "1.0-rev0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseAgreementApi.DownloadLicenseAgreementFile(context.Background()).LicenseAgreementType(licenseAgreementType).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseAgreementApi.DownloadLicenseAgreementFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadLicenseAgreementFile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `LicenseAgreementApi.DownloadLicenseAgreementFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLicenseAgreementFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licenseAgreementType** | **string** |  | 
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentLicenseAgreementAccepted

> LicenseAgreementAcceptResult GetCurrentLicenseAgreementAccepted(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.LicenseAgreementApi.GetCurrentLicenseAgreementAccepted(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseAgreementApi.GetCurrentLicenseAgreementAccepted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentLicenseAgreementAccepted`: LicenseAgreementAcceptResult
    fmt.Fprintf(os.Stdout, "Response from `LicenseAgreementApi.GetCurrentLicenseAgreementAccepted`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentLicenseAgreementAcceptedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**LicenseAgreementAcceptResult**](LicenseAgreementAcceptResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseAgreement

> LicenseAgreements GetLicenseAgreement(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.LicenseAgreementApi.GetLicenseAgreement(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseAgreementApi.GetLicenseAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseAgreement`: LicenseAgreements
    fmt.Fprintf(os.Stdout, "Response from `LicenseAgreementApi.GetLicenseAgreement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**LicenseAgreements**](LicenseAgreements.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

