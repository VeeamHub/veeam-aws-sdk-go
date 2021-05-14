# ReplicaMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceRegionId** | **string** |  | 
**TargetRegionId** | **string** |  | 
**TargetAmazonAccountId** | **string** |  | 
**EncryptionKey** | Pointer to **string** |  | [optional] 

## Methods

### NewReplicaMapping

`func NewReplicaMapping(sourceRegionId string, targetRegionId string, targetAmazonAccountId string, ) *ReplicaMapping`

NewReplicaMapping instantiates a new ReplicaMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaMappingWithDefaults

`func NewReplicaMappingWithDefaults() *ReplicaMapping`

NewReplicaMappingWithDefaults instantiates a new ReplicaMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceRegionId

`func (o *ReplicaMapping) GetSourceRegionId() string`

GetSourceRegionId returns the SourceRegionId field if non-nil, zero value otherwise.

### GetSourceRegionIdOk

`func (o *ReplicaMapping) GetSourceRegionIdOk() (*string, bool)`

GetSourceRegionIdOk returns a tuple with the SourceRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRegionId

`func (o *ReplicaMapping) SetSourceRegionId(v string)`

SetSourceRegionId sets SourceRegionId field to given value.


### GetTargetRegionId

`func (o *ReplicaMapping) GetTargetRegionId() string`

GetTargetRegionId returns the TargetRegionId field if non-nil, zero value otherwise.

### GetTargetRegionIdOk

`func (o *ReplicaMapping) GetTargetRegionIdOk() (*string, bool)`

GetTargetRegionIdOk returns a tuple with the TargetRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRegionId

`func (o *ReplicaMapping) SetTargetRegionId(v string)`

SetTargetRegionId sets TargetRegionId field to given value.


### GetTargetAmazonAccountId

`func (o *ReplicaMapping) GetTargetAmazonAccountId() string`

GetTargetAmazonAccountId returns the TargetAmazonAccountId field if non-nil, zero value otherwise.

### GetTargetAmazonAccountIdOk

`func (o *ReplicaMapping) GetTargetAmazonAccountIdOk() (*string, bool)`

GetTargetAmazonAccountIdOk returns a tuple with the TargetAmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmazonAccountId

`func (o *ReplicaMapping) SetTargetAmazonAccountId(v string)`

SetTargetAmazonAccountId sets TargetAmazonAccountId field to given value.


### GetEncryptionKey

`func (o *ReplicaMapping) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *ReplicaMapping) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *ReplicaMapping) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *ReplicaMapping) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


