# UserUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**OldPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewUserUpdateSpec

`func NewUserUpdateSpec() *UserUpdateSpec`

NewUserUpdateSpec instantiates a new UserUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateSpecWithDefaults

`func NewUserUpdateSpecWithDefaults() *UserUpdateSpec`

NewUserUpdateSpecWithDefaults instantiates a new UserUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UserUpdateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserUpdateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserUpdateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserUpdateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPassword

`func (o *UserUpdateSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserUpdateSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserUpdateSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserUpdateSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetOldPassword

`func (o *UserUpdateSpec) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *UserUpdateSpec) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *UserUpdateSpec) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *UserUpdateSpec) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


