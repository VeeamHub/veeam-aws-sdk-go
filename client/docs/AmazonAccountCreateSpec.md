# AmazonAccountCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**AccessKeys** | Pointer to [**AmazonAccountAccessKeysCreateSpec**](AmazonAccountAccessKeysCreateSpec.md) |  | [optional] 
**IAMRole** | Pointer to [**AmazonAccountIAMRoleCreateSpec**](AmazonAccountIAMRoleCreateSpec.md) |  | [optional] 
**IAMRoleFromAnotherAccount** | Pointer to [**AmazonAccountIAMRoleFromAnotherAccountCreateSpec**](AmazonAccountIAMRoleFromAnotherAccountCreateSpec.md) |  | [optional] 

## Methods

### NewAmazonAccountCreateSpec

`func NewAmazonAccountCreateSpec(name string, ) *AmazonAccountCreateSpec`

NewAmazonAccountCreateSpec instantiates a new AmazonAccountCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonAccountCreateSpecWithDefaults

`func NewAmazonAccountCreateSpecWithDefaults() *AmazonAccountCreateSpec`

NewAmazonAccountCreateSpecWithDefaults instantiates a new AmazonAccountCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AmazonAccountCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmazonAccountCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmazonAccountCreateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AmazonAccountCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonAccountCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonAccountCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AmazonAccountCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessKeys

`func (o *AmazonAccountCreateSpec) GetAccessKeys() AmazonAccountAccessKeysCreateSpec`

GetAccessKeys returns the AccessKeys field if non-nil, zero value otherwise.

### GetAccessKeysOk

`func (o *AmazonAccountCreateSpec) GetAccessKeysOk() (*AmazonAccountAccessKeysCreateSpec, bool)`

GetAccessKeysOk returns a tuple with the AccessKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeys

`func (o *AmazonAccountCreateSpec) SetAccessKeys(v AmazonAccountAccessKeysCreateSpec)`

SetAccessKeys sets AccessKeys field to given value.

### HasAccessKeys

`func (o *AmazonAccountCreateSpec) HasAccessKeys() bool`

HasAccessKeys returns a boolean if a field has been set.

### GetIAMRole

`func (o *AmazonAccountCreateSpec) GetIAMRole() AmazonAccountIAMRoleCreateSpec`

GetIAMRole returns the IAMRole field if non-nil, zero value otherwise.

### GetIAMRoleOk

`func (o *AmazonAccountCreateSpec) GetIAMRoleOk() (*AmazonAccountIAMRoleCreateSpec, bool)`

GetIAMRoleOk returns a tuple with the IAMRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAMRole

`func (o *AmazonAccountCreateSpec) SetIAMRole(v AmazonAccountIAMRoleCreateSpec)`

SetIAMRole sets IAMRole field to given value.

### HasIAMRole

`func (o *AmazonAccountCreateSpec) HasIAMRole() bool`

HasIAMRole returns a boolean if a field has been set.

### GetIAMRoleFromAnotherAccount

`func (o *AmazonAccountCreateSpec) GetIAMRoleFromAnotherAccount() AmazonAccountIAMRoleFromAnotherAccountCreateSpec`

GetIAMRoleFromAnotherAccount returns the IAMRoleFromAnotherAccount field if non-nil, zero value otherwise.

### GetIAMRoleFromAnotherAccountOk

`func (o *AmazonAccountCreateSpec) GetIAMRoleFromAnotherAccountOk() (*AmazonAccountIAMRoleFromAnotherAccountCreateSpec, bool)`

GetIAMRoleFromAnotherAccountOk returns a tuple with the IAMRoleFromAnotherAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAMRoleFromAnotherAccount

`func (o *AmazonAccountCreateSpec) SetIAMRoleFromAnotherAccount(v AmazonAccountIAMRoleFromAnotherAccountCreateSpec)`

SetIAMRoleFromAnotherAccount sets IAMRoleFromAnotherAccount field to given value.

### HasIAMRoleFromAnotherAccount

`func (o *AmazonAccountCreateSpec) HasIAMRoleFromAnotherAccount() bool`

HasIAMRoleFromAnotherAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


