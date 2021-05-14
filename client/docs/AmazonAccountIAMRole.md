# AmazonAccountIAMRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentAmazonAccountId** | Pointer to **string** |  | [optional] 
**RoleName** | **string** |  | 
**IsDefault** | **bool** |  | 

## Methods

### NewAmazonAccountIAMRole

`func NewAmazonAccountIAMRole(roleName string, isDefault bool, ) *AmazonAccountIAMRole`

NewAmazonAccountIAMRole instantiates a new AmazonAccountIAMRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonAccountIAMRoleWithDefaults

`func NewAmazonAccountIAMRoleWithDefaults() *AmazonAccountIAMRole`

NewAmazonAccountIAMRoleWithDefaults instantiates a new AmazonAccountIAMRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentAmazonAccountId

`func (o *AmazonAccountIAMRole) GetParentAmazonAccountId() string`

GetParentAmazonAccountId returns the ParentAmazonAccountId field if non-nil, zero value otherwise.

### GetParentAmazonAccountIdOk

`func (o *AmazonAccountIAMRole) GetParentAmazonAccountIdOk() (*string, bool)`

GetParentAmazonAccountIdOk returns a tuple with the ParentAmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAmazonAccountId

`func (o *AmazonAccountIAMRole) SetParentAmazonAccountId(v string)`

SetParentAmazonAccountId sets ParentAmazonAccountId field to given value.

### HasParentAmazonAccountId

`func (o *AmazonAccountIAMRole) HasParentAmazonAccountId() bool`

HasParentAmazonAccountId returns a boolean if a field has been set.

### GetRoleName

`func (o *AmazonAccountIAMRole) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AmazonAccountIAMRole) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AmazonAccountIAMRole) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetIsDefault

`func (o *AmazonAccountIAMRole) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AmazonAccountIAMRole) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AmazonAccountIAMRole) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


