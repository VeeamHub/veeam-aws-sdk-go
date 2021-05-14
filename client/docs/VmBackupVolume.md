# VmBackupVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskId** | **string** |  | 
**Type** | [**Ec2DiskTypes**](Ec2DiskTypes.md) |  | 
**BackupSizeBytes** | **int64** |  | 
**Iops** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Device** | Pointer to **string** |  | [optional] 
**EncryptionKey** | Pointer to **string** |  | [optional] 
**EncryptionKeyName** | Pointer to **string** |  | [optional] 

## Methods

### NewVmBackupVolume

`func NewVmBackupVolume(diskId string, type_ Ec2DiskTypes, backupSizeBytes int64, ) *VmBackupVolume`

NewVmBackupVolume instantiates a new VmBackupVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmBackupVolumeWithDefaults

`func NewVmBackupVolumeWithDefaults() *VmBackupVolume`

NewVmBackupVolumeWithDefaults instantiates a new VmBackupVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskId

`func (o *VmBackupVolume) GetDiskId() string`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *VmBackupVolume) GetDiskIdOk() (*string, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *VmBackupVolume) SetDiskId(v string)`

SetDiskId sets DiskId field to given value.


### GetType

`func (o *VmBackupVolume) GetType() Ec2DiskTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VmBackupVolume) GetTypeOk() (*Ec2DiskTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VmBackupVolume) SetType(v Ec2DiskTypes)`

SetType sets Type field to given value.


### GetBackupSizeBytes

`func (o *VmBackupVolume) GetBackupSizeBytes() int64`

GetBackupSizeBytes returns the BackupSizeBytes field if non-nil, zero value otherwise.

### GetBackupSizeBytesOk

`func (o *VmBackupVolume) GetBackupSizeBytesOk() (*int64, bool)`

GetBackupSizeBytesOk returns a tuple with the BackupSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSizeBytes

`func (o *VmBackupVolume) SetBackupSizeBytes(v int64)`

SetBackupSizeBytes sets BackupSizeBytes field to given value.


### GetIops

`func (o *VmBackupVolume) GetIops() int32`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *VmBackupVolume) GetIopsOk() (*int32, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *VmBackupVolume) SetIops(v int32)`

SetIops sets Iops field to given value.

### HasIops

`func (o *VmBackupVolume) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetName

`func (o *VmBackupVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmBackupVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmBackupVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmBackupVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDevice

`func (o *VmBackupVolume) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VmBackupVolume) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VmBackupVolume) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *VmBackupVolume) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *VmBackupVolume) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *VmBackupVolume) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *VmBackupVolume) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *VmBackupVolume) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetEncryptionKeyName

`func (o *VmBackupVolume) GetEncryptionKeyName() string`

GetEncryptionKeyName returns the EncryptionKeyName field if non-nil, zero value otherwise.

### GetEncryptionKeyNameOk

`func (o *VmBackupVolume) GetEncryptionKeyNameOk() (*string, bool)`

GetEncryptionKeyNameOk returns a tuple with the EncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyName

`func (o *VmBackupVolume) SetEncryptionKeyName(v string)`

SetEncryptionKeyName sets EncryptionKeyName field to given value.

### HasEncryptionKeyName

`func (o *VmBackupVolume) HasEncryptionKeyName() bool`

HasEncryptionKeyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


