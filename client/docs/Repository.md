# Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**AmazonAccountId** | **string** |  | 
**AmazonStorageFolder** | **string** |  | 
**AmazonBucketId** | Pointer to **string** |  | [optional] 
**Hint** | Pointer to **string** |  | [optional] 
**EnableEncryption** | **bool** |  | 
**Embedded** | Pointer to [**RepositoryEmbedded**](RepositoryEmbedded.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewRepository

`func NewRepository(id string, name string, description string, amazonAccountId string, amazonStorageFolder string, enableEncryption bool, ) *Repository`

NewRepository instantiates a new Repository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWithDefaults

`func NewRepositoryWithDefaults() *Repository`

NewRepositoryWithDefaults instantiates a new Repository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Repository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Repository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Repository) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Repository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Repository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Repository) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Repository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Repository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Repository) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAmazonAccountId

`func (o *Repository) GetAmazonAccountId() string`

GetAmazonAccountId returns the AmazonAccountId field if non-nil, zero value otherwise.

### GetAmazonAccountIdOk

`func (o *Repository) GetAmazonAccountIdOk() (*string, bool)`

GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccountId

`func (o *Repository) SetAmazonAccountId(v string)`

SetAmazonAccountId sets AmazonAccountId field to given value.


### GetAmazonStorageFolder

`func (o *Repository) GetAmazonStorageFolder() string`

GetAmazonStorageFolder returns the AmazonStorageFolder field if non-nil, zero value otherwise.

### GetAmazonStorageFolderOk

`func (o *Repository) GetAmazonStorageFolderOk() (*string, bool)`

GetAmazonStorageFolderOk returns a tuple with the AmazonStorageFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonStorageFolder

`func (o *Repository) SetAmazonStorageFolder(v string)`

SetAmazonStorageFolder sets AmazonStorageFolder field to given value.


### GetAmazonBucketId

`func (o *Repository) GetAmazonBucketId() string`

GetAmazonBucketId returns the AmazonBucketId field if non-nil, zero value otherwise.

### GetAmazonBucketIdOk

`func (o *Repository) GetAmazonBucketIdOk() (*string, bool)`

GetAmazonBucketIdOk returns a tuple with the AmazonBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonBucketId

`func (o *Repository) SetAmazonBucketId(v string)`

SetAmazonBucketId sets AmazonBucketId field to given value.

### HasAmazonBucketId

`func (o *Repository) HasAmazonBucketId() bool`

HasAmazonBucketId returns a boolean if a field has been set.

### GetHint

`func (o *Repository) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *Repository) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *Repository) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *Repository) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetEnableEncryption

`func (o *Repository) GetEnableEncryption() bool`

GetEnableEncryption returns the EnableEncryption field if non-nil, zero value otherwise.

### GetEnableEncryptionOk

`func (o *Repository) GetEnableEncryptionOk() (*bool, bool)`

GetEnableEncryptionOk returns a tuple with the EnableEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEncryption

`func (o *Repository) SetEnableEncryption(v bool)`

SetEnableEncryption sets EnableEncryption field to given value.


### GetEmbedded

`func (o *Repository) GetEmbedded() RepositoryEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *Repository) GetEmbeddedOk() (*RepositoryEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *Repository) SetEmbedded(v RepositoryEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *Repository) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *Repository) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Repository) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Repository) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Repository) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


