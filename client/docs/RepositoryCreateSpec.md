# RepositoryCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**AmazonAccountId** | **string** |  | 
**AmazonBucketId** | **string** |  | 
**AmazonStorageFolder** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**Hint** | Pointer to **string** |  | [optional] 
**EnableEncryption** | **bool** |  | 

## Methods

### NewRepositoryCreateSpec

`func NewRepositoryCreateSpec(name string, amazonAccountId string, amazonBucketId string, amazonStorageFolder string, enableEncryption bool, ) *RepositoryCreateSpec`

NewRepositoryCreateSpec instantiates a new RepositoryCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryCreateSpecWithDefaults

`func NewRepositoryCreateSpecWithDefaults() *RepositoryCreateSpec`

NewRepositoryCreateSpecWithDefaults instantiates a new RepositoryCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RepositoryCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryCreateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RepositoryCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepositoryCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepositoryCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RepositoryCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmazonAccountId

`func (o *RepositoryCreateSpec) GetAmazonAccountId() string`

GetAmazonAccountId returns the AmazonAccountId field if non-nil, zero value otherwise.

### GetAmazonAccountIdOk

`func (o *RepositoryCreateSpec) GetAmazonAccountIdOk() (*string, bool)`

GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccountId

`func (o *RepositoryCreateSpec) SetAmazonAccountId(v string)`

SetAmazonAccountId sets AmazonAccountId field to given value.


### GetAmazonBucketId

`func (o *RepositoryCreateSpec) GetAmazonBucketId() string`

GetAmazonBucketId returns the AmazonBucketId field if non-nil, zero value otherwise.

### GetAmazonBucketIdOk

`func (o *RepositoryCreateSpec) GetAmazonBucketIdOk() (*string, bool)`

GetAmazonBucketIdOk returns a tuple with the AmazonBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonBucketId

`func (o *RepositoryCreateSpec) SetAmazonBucketId(v string)`

SetAmazonBucketId sets AmazonBucketId field to given value.


### GetAmazonStorageFolder

`func (o *RepositoryCreateSpec) GetAmazonStorageFolder() string`

GetAmazonStorageFolder returns the AmazonStorageFolder field if non-nil, zero value otherwise.

### GetAmazonStorageFolderOk

`func (o *RepositoryCreateSpec) GetAmazonStorageFolderOk() (*string, bool)`

GetAmazonStorageFolderOk returns a tuple with the AmazonStorageFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonStorageFolder

`func (o *RepositoryCreateSpec) SetAmazonStorageFolder(v string)`

SetAmazonStorageFolder sets AmazonStorageFolder field to given value.


### GetPassword

`func (o *RepositoryCreateSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RepositoryCreateSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RepositoryCreateSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RepositoryCreateSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHint

`func (o *RepositoryCreateSpec) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *RepositoryCreateSpec) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *RepositoryCreateSpec) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *RepositoryCreateSpec) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetEnableEncryption

`func (o *RepositoryCreateSpec) GetEnableEncryption() bool`

GetEnableEncryption returns the EnableEncryption field if non-nil, zero value otherwise.

### GetEnableEncryptionOk

`func (o *RepositoryCreateSpec) GetEnableEncryptionOk() (*bool, bool)`

GetEnableEncryptionOk returns a tuple with the EnableEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEncryption

`func (o *RepositoryCreateSpec) SetEnableEncryption(v bool)`

SetEnableEncryption sets EnableEncryption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


