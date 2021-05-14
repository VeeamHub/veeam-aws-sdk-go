# \EmailNotificationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmailNotifications**](EmailNotificationsApi.md#GetEmailNotifications) | **Get** /api/v1/settings/emailNotifications | 
[**SaveEmailNotifications**](EmailNotificationsApi.md#SaveEmailNotifications) | **Post** /api/v1/settings/emailNotifications | 
[**SendTestMessage**](EmailNotificationsApi.md#SendTestMessage) | **Post** /api/v1/settings/emailNotifications/sendTestMessage | 



## GetEmailNotifications

> EmailNotifications GetEmailNotifications(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.EmailNotificationsApi.GetEmailNotifications(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailNotificationsApi.GetEmailNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailNotifications`: EmailNotifications
    fmt.Fprintf(os.Stdout, "Response from `EmailNotificationsApi.GetEmailNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**EmailNotifications**](EmailNotifications.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveEmailNotifications

> EmailNotifications SaveEmailNotifications(ctx).XApiVersion(xApiVersion).SmtpSpec(smtpSpec).Execute()



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
    smtpSpec := *openapiclient.NewEmailNotifications(false, "Server_example", int32(123), "From_example", "To_example", "Subject_example") // EmailNotifications | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EmailNotificationsApi.SaveEmailNotifications(context.Background()).XApiVersion(xApiVersion).SmtpSpec(smtpSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailNotificationsApi.SaveEmailNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveEmailNotifications`: EmailNotifications
    fmt.Fprintf(os.Stdout, "Response from `EmailNotificationsApi.SaveEmailNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveEmailNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **smtpSpec** | [**EmailNotifications**](EmailNotifications.md) |  | 

### Return type

[**EmailNotifications**](EmailNotifications.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestMessage

> EmailNotificationsTestResult SendTestMessage(ctx).XApiVersion(xApiVersion).EmailNotifications(emailNotifications).Execute()



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
    emailNotifications := *openapiclient.NewEmailNotifications(false, "Server_example", int32(123), "From_example", "To_example", "Subject_example") // EmailNotifications | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EmailNotificationsApi.SendTestMessage(context.Background()).XApiVersion(xApiVersion).EmailNotifications(emailNotifications).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailNotificationsApi.SendTestMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendTestMessage`: EmailNotificationsTestResult
    fmt.Fprintf(os.Stdout, "Response from `EmailNotificationsApi.SendTestMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendTestMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **emailNotifications** | [**EmailNotifications**](EmailNotifications.md) |  | 

### Return type

[**EmailNotificationsTestResult**](EmailNotificationsTestResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

