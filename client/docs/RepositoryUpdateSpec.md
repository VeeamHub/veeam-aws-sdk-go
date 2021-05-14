# RepositoryUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**AmazonAccountId** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**Hint** | Pointer to **string** |  | [optional] 
**EnableEncryption** | **bool** |  | 

## Methods

### NewRepositoryUpdateSpec

`func NewRepositoryUpdateSpec(name string, amazonAccountId string, enableEncryption bool, ) *RepositoryUpdateSpec`

NewRepositoryUpdateSpec instantiates a new RepositoryUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryUpdateSpecWithDefaults

`func NewRepositoryUpdateSpecWithDefaults() *RepositoryUpdateSpec`

NewRepositoryUpdateSpecWithDefaults instantiates a new RepositoryUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RepositoryUpdateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryUpdateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryUpdateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RepositoryUpdateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepositoryUpdateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepositoryUpdateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RepositoryUpdateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmazonAccountId

`func (o *RepositoryUpdateSpec) GetAmazonAccountId() string`

GetAmazonAccountId returns the AmazonAccountId field if non-nil, zero value otherwise.

### GetAmazonAccountIdOk

`func (o *RepositoryUpdateSpec) GetAmazonAccountIdOk() (*string, bool)`

GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccountId

`func (o *RepositoryUpdateSpec) SetAmazonAccountId(v string)`

SetAmazonAccountId sets AmazonAccountId field to given value.


### GetPassword

`func (o *RepositoryUpdateSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RepositoryUpdateSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RepositoryUpdateSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RepositoryUpdateSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHint

`func (o *RepositoryUpdateSpec) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *RepositoryUpdateSpec) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *RepositoryUpdateSpec) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *RepositoryUpdateSpec) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetEnableEncryption

`func (o *RepositoryUpdateSpec) GetEnableEncryption() bool`

GetEnableEncryption returns the EnableEncryption field if non-nil, zero value otherwise.

### GetEnableEncryptionOk

`func (o *RepositoryUpdateSpec) GetEnableEncryptionOk() (*bool, bool)`

GetEnableEncryptionOk returns a tuple with the EnableEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEncryption

`func (o *RepositoryUpdateSpec) SetEnableEncryption(v bool)`

SetEnableEncryption sets EnableEncryption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


