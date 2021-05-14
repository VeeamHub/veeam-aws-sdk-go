# RESTLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**Issued** | Pointer to **time.Time** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**MfaEnabled** | Pointer to **bool** |  | [optional] 
**MfaToken** | Pointer to **string** |  | [optional] 

## Methods

### NewRESTLogin

`func NewRESTLogin() *RESTLogin`

NewRESTLogin instantiates a new RESTLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRESTLoginWithDefaults

`func NewRESTLoginWithDefaults() *RESTLogin`

NewRESTLoginWithDefaults instantiates a new RESTLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *RESTLogin) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *RESTLogin) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *RESTLogin) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *RESTLogin) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetTokenType

`func (o *RESTLogin) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *RESTLogin) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *RESTLogin) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *RESTLogin) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetRefreshToken

`func (o *RESTLogin) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *RESTLogin) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *RESTLogin) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *RESTLogin) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *RESTLogin) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *RESTLogin) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *RESTLogin) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *RESTLogin) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetIssued

`func (o *RESTLogin) GetIssued() time.Time`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *RESTLogin) GetIssuedOk() (*time.Time, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *RESTLogin) SetIssued(v time.Time)`

SetIssued sets Issued field to given value.

### HasIssued

`func (o *RESTLogin) HasIssued() bool`

HasIssued returns a boolean if a field has been set.

### GetExpires

`func (o *RESTLogin) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *RESTLogin) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *RESTLogin) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *RESTLogin) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetUsername

`func (o *RESTLogin) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RESTLogin) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RESTLogin) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RESTLogin) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRoleName

`func (o *RESTLogin) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *RESTLogin) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *RESTLogin) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *RESTLogin) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetMfaEnabled

`func (o *RESTLogin) GetMfaEnabled() bool`

GetMfaEnabled returns the MfaEnabled field if non-nil, zero value otherwise.

### GetMfaEnabledOk

`func (o *RESTLogin) GetMfaEnabledOk() (*bool, bool)`

GetMfaEnabledOk returns a tuple with the MfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnabled

`func (o *RESTLogin) SetMfaEnabled(v bool)`

SetMfaEnabled sets MfaEnabled field to given value.

### HasMfaEnabled

`func (o *RESTLogin) HasMfaEnabled() bool`

HasMfaEnabled returns a boolean if a field has been set.

### GetMfaToken

`func (o *RESTLogin) GetMfaToken() string`

GetMfaToken returns the MfaToken field if non-nil, zero value otherwise.

### GetMfaTokenOk

`func (o *RESTLogin) GetMfaTokenOk() (*string, bool)`

GetMfaTokenOk returns a tuple with the MfaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaToken

`func (o *RESTLogin) SetMfaToken(v string)`

SetMfaToken sets MfaToken field to given value.

### HasMfaToken

`func (o *RESTLogin) HasMfaToken() bool`

HasMfaToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


