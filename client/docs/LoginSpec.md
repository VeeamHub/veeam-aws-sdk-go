# LoginSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** | required if grant_type &#x3D; authorization_code | [optional] 
**GrantType** | [**LoginGrantTypes**](LoginGrantTypes.md) |  | 
**MfaToken** | Pointer to **string** |  | [optional] 
**MfaCode** | Pointer to **string** |  | [optional] 

## Methods

### NewLoginSpec

`func NewLoginSpec(grantType LoginGrantTypes, ) *LoginSpec`

NewLoginSpec instantiates a new LoginSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginSpecWithDefaults

`func NewLoginSpecWithDefaults() *LoginSpec`

NewLoginSpecWithDefaults instantiates a new LoginSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *LoginSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LoginSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LoginSpec) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LoginSpec) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *LoginSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoginSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoginSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoginSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRefreshToken

`func (o *LoginSpec) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *LoginSpec) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *LoginSpec) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *LoginSpec) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetCode

`func (o *LoginSpec) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LoginSpec) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LoginSpec) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LoginSpec) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetGrantType

`func (o *LoginSpec) GetGrantType() LoginGrantTypes`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *LoginSpec) GetGrantTypeOk() (*LoginGrantTypes, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *LoginSpec) SetGrantType(v LoginGrantTypes)`

SetGrantType sets GrantType field to given value.


### GetMfaToken

`func (o *LoginSpec) GetMfaToken() string`

GetMfaToken returns the MfaToken field if non-nil, zero value otherwise.

### GetMfaTokenOk

`func (o *LoginSpec) GetMfaTokenOk() (*string, bool)`

GetMfaTokenOk returns a tuple with the MfaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaToken

`func (o *LoginSpec) SetMfaToken(v string)`

SetMfaToken sets MfaToken field to given value.

### HasMfaToken

`func (o *LoginSpec) HasMfaToken() bool`

HasMfaToken returns a boolean if a field has been set.

### GetMfaCode

`func (o *LoginSpec) GetMfaCode() string`

GetMfaCode returns the MfaCode field if non-nil, zero value otherwise.

### GetMfaCodeOk

`func (o *LoginSpec) GetMfaCodeOk() (*string, bool)`

GetMfaCodeOk returns a tuple with the MfaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaCode

`func (o *LoginSpec) SetMfaCode(v string)`

SetMfaCode sets MfaCode field to given value.

### HasMfaCode

`func (o *LoginSpec) HasMfaCode() bool`

HasMfaCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


