# \TokenApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authenticate**](TokenApi.md#Authenticate) | **Post** /api/v1/token | 
[**CreateAuthorizationCode**](TokenApi.md#CreateAuthorizationCode) | **Post** /api/v1/token/authorizationCode | 
[**SignOut**](TokenApi.md#SignOut) | **Delete** /api/v1/token | 



## Authenticate

> RESTLogin Authenticate(ctx).XApiVersion(xApiVersion).GrantType(grantType).Username(username).Password(password).RefreshToken(refreshToken).Code(code).MfaToken(mfaToken).MfaCode(mfaCode).Execute()



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
    grantType := "grantType_example" // string | Authentication grant type
    username := "username_example" // string | User name (required for the 'password' grant type) (optional)
    password := "password_example" // string | Password (required for the 'password' grant type) (optional)
    refreshToken := "refreshToken_example" // string | Refresh token (required for the 'refresh_token' grant type) (optional)
    code := "code_example" // string | Authorization code (required for the 'authorization_code' grant type) (optional)
    mfaToken := "mfaToken_example" // string | MFA token (required for the 'mfa' grant type) (optional)
    mfaCode := "mfaCode_example" // string | Verification code (required for the 'mfa' grant type) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenApi.Authenticate(context.Background()).XApiVersion(xApiVersion).GrantType(grantType).Username(username).Password(password).RefreshToken(refreshToken).Code(code).MfaToken(mfaToken).MfaCode(mfaCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.Authenticate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Authenticate`: RESTLogin
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.Authenticate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]
 **grantType** | **string** | Authentication grant type | 
 **username** | **string** | User name (required for the &#39;password&#39; grant type) | 
 **password** | **string** | Password (required for the &#39;password&#39; grant type) | 
 **refreshToken** | **string** | Refresh token (required for the &#39;refresh_token&#39; grant type) | 
 **code** | **string** | Authorization code (required for the &#39;authorization_code&#39; grant type) | 
 **mfaToken** | **string** | MFA token (required for the &#39;mfa&#39; grant type) | 
 **mfaCode** | **string** | Verification code (required for the &#39;mfa&#39; grant type) | 

### Return type

[**RESTLogin**](RESTLogin.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthorizationCode

> AuthorizationCode CreateAuthorizationCode(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.TokenApi.CreateAuthorizationCode(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.CreateAuthorizationCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationCode`: AuthorizationCode
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.CreateAuthorizationCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | [default to &quot;1.0-rev0&quot;]

### Return type

[**AuthorizationCode**](AuthorizationCode.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignOut

> SignOut(ctx).XApiVersion(xApiVersion).Execute()



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
    resp, r, err := api_client.TokenApi.SignOut(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.SignOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignOutRequest struct via the builder pattern


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

