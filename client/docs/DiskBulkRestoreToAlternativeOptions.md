# DiskBulkRestoreToAlternativeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneId** | **string** |  | 
**PreserveEncryptionForVolumes** | **bool** |  | 
**EncryptionKeyId** | Pointer to **string** |  | [optional] 

## Methods

### NewDiskBulkRestoreToAlternativeOptions

`func NewDiskBulkRestoreToAlternativeOptions(availabilityZoneId string, preserveEncryptionForVolumes bool, ) *DiskBulkRestoreToAlternativeOptions`

NewDiskBulkRestoreToAlternativeOptions instantiates a new DiskBulkRestoreToAlternativeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBulkRestoreToAlternativeOptionsWithDefaults

`func NewDiskBulkRestoreToAlternativeOptionsWithDefaults() *DiskBulkRestoreToAlternativeOptions`

NewDiskBulkRestoreToAlternativeOptionsWithDefaults instantiates a new DiskBulkRestoreToAlternativeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneId

`func (o *DiskBulkRestoreToAlternativeOptions) GetAvailabilityZoneId() string`

GetAvailabilityZoneId returns the AvailabilityZoneId field if non-nil, zero value otherwise.

### GetAvailabilityZoneIdOk

`func (o *DiskBulkRestoreToAlternativeOptions) GetAvailabilityZoneIdOk() (*string, bool)`

GetAvailabilityZoneIdOk returns a tuple with the AvailabilityZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneId

`func (o *DiskBulkRestoreToAlternativeOptions) SetAvailabilityZoneId(v string)`

SetAvailabilityZoneId sets AvailabilityZoneId field to given value.


### GetPreserveEncryptionForVolumes

`func (o *DiskBulkRestoreToAlternativeOptions) GetPreserveEncryptionForVolumes() bool`

GetPreserveEncryptionForVolumes returns the PreserveEncryptionForVolumes field if non-nil, zero value otherwise.

### GetPreserveEncryptionForVolumesOk

`func (o *DiskBulkRestoreToAlternativeOptions) GetPreserveEncryptionForVolumesOk() (*bool, bool)`

GetPreserveEncryptionForVolumesOk returns a tuple with the PreserveEncryptionForVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveEncryptionForVolumes

`func (o *DiskBulkRestoreToAlternativeOptions) SetPreserveEncryptionForVolumes(v bool)`

SetPreserveEncryptionForVolumes sets PreserveEncryptionForVolumes field to given value.


### GetEncryptionKeyId

`func (o *DiskBulkRestoreToAlternativeOptions) GetEncryptionKeyId() string`

GetEncryptionKeyId returns the EncryptionKeyId field if non-nil, zero value otherwise.

### GetEncryptionKeyIdOk

`func (o *DiskBulkRestoreToAlternativeOptions) GetEncryptionKeyIdOk() (*string, bool)`

GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyId

`func (o *DiskBulkRestoreToAlternativeOptions) SetEncryptionKeyId(v string)`

SetEncryptionKeyId sets EncryptionKeyId field to given value.

### HasEncryptionKeyId

`func (o *DiskBulkRestoreToAlternativeOptions) HasEncryptionKeyId() bool`

HasEncryptionKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


