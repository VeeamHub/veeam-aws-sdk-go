# AmazonAccountUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**AccessKeys** | Pointer to [**AmazonAccountAccessKeysUpdateSpec**](AmazonAccountAccessKeysUpdateSpec.md) |  | [optional] 
**IAMRole** | Pointer to [**AmazonAccountIAMRoleUpdateSpec**](AmazonAccountIAMRoleUpdateSpec.md) |  | [optional] 
**IAMRoleFromAnotherAccount** | Pointer to [**AmazonAccountIAMRoleFromAnotherAccountUpdateSpec**](AmazonAccountIAMRoleFromAnotherAccountUpdateSpec.md) |  | [optional] 

## Methods

### NewAmazonAccountUpdateSpec

`func NewAmazonAccountUpdateSpec(name string, ) *AmazonAccountUpdateSpec`

NewAmazonAccountUpdateSpec instantiates a new AmazonAccountUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonAccountUpdateSpecWithDefaults

`func NewAmazonAccountUpdateSpecWithDefaults() *AmazonAccountUpdateSpec`

NewAmazonAccountUpdateSpecWithDefaults instantiates a new AmazonAccountUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AmazonAccountUpdateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmazonAccountUpdateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmazonAccountUpdateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AmazonAccountUpdateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonAccountUpdateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonAccountUpdateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AmazonAccountUpdateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessKeys

`func (o *AmazonAccountUpdateSpec) GetAccessKeys() AmazonAccountAccessKeysUpdateSpec`

GetAccessKeys returns the AccessKeys field if non-nil, zero value otherwise.

### GetAccessKeysOk

`func (o *AmazonAccountUpdateSpec) GetAccessKeysOk() (*AmazonAccountAccessKeysUpdateSpec, bool)`

GetAccessKeysOk returns a tuple with the AccessKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeys

`func (o *AmazonAccountUpdateSpec) SetAccessKeys(v AmazonAccountAccessKeysUpdateSpec)`

SetAccessKeys sets AccessKeys field to given value.

### HasAccessKeys

`func (o *AmazonAccountUpdateSpec) HasAccessKeys() bool`

HasAccessKeys returns a boolean if a field has been set.

### GetIAMRole

`func (o *AmazonAccountUpdateSpec) GetIAMRole() AmazonAccountIAMRoleUpdateSpec`

GetIAMRole returns the IAMRole field if non-nil, zero value otherwise.

### GetIAMRoleOk

`func (o *AmazonAccountUpdateSpec) GetIAMRoleOk() (*AmazonAccountIAMRoleUpdateSpec, bool)`

GetIAMRoleOk returns a tuple with the IAMRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAMRole

`func (o *AmazonAccountUpdateSpec) SetIAMRole(v AmazonAccountIAMRoleUpdateSpec)`

SetIAMRole sets IAMRole field to given value.

### HasIAMRole

`func (o *AmazonAccountUpdateSpec) HasIAMRole() bool`

HasIAMRole returns a boolean if a field has been set.

### GetIAMRoleFromAnotherAccount

`func (o *AmazonAccountUpdateSpec) GetIAMRoleFromAnotherAccount() AmazonAccountIAMRoleFromAnotherAccountUpdateSpec`

GetIAMRoleFromAnotherAccount returns the IAMRoleFromAnotherAccount field if non-nil, zero value otherwise.

### GetIAMRoleFromAnotherAccountOk

`func (o *AmazonAccountUpdateSpec) GetIAMRoleFromAnotherAccountOk() (*AmazonAccountIAMRoleFromAnotherAccountUpdateSpec, bool)`

GetIAMRoleFromAnotherAccountOk returns a tuple with the IAMRoleFromAnotherAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAMRoleFromAnotherAccount

`func (o *AmazonAccountUpdateSpec) SetIAMRoleFromAnotherAccount(v AmazonAccountIAMRoleFromAnotherAccountUpdateSpec)`

SetIAMRoleFromAnotherAccount sets IAMRoleFromAnotherAccount field to given value.

### HasIAMRoleFromAnotherAccount

`func (o *AmazonAccountUpdateSpec) HasIAMRoleFromAnotherAccount() bool`

HasIAMRoleFromAnotherAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


