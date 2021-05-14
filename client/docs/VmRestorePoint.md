# VmRestorePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**BackupId** | **string** |  | 
**JobType** | [**JobTypes**](JobTypes.md) |  | 
**PointInTime** | **time.Time** |  | 
**BackupSizeBytes** | **int64** |  | 
**Volumes** | [**[]VmBackupVolume**](VmBackupVolume.md) |  | 
**RegionId** | Pointer to **string** |  | [optional] 
**AvailabilityZoneId** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewVmRestorePoint

`func NewVmRestorePoint(id string, backupId string, jobType JobTypes, pointInTime time.Time, backupSizeBytes int64, volumes []VmBackupVolume, ) *VmRestorePoint`

NewVmRestorePoint instantiates a new VmRestorePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRestorePointWithDefaults

`func NewVmRestorePointWithDefaults() *VmRestorePoint`

NewVmRestorePointWithDefaults instantiates a new VmRestorePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmRestorePoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmRestorePoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmRestorePoint) SetId(v string)`

SetId sets Id field to given value.


### GetBackupId

`func (o *VmRestorePoint) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *VmRestorePoint) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *VmRestorePoint) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.


### GetJobType

`func (o *VmRestorePoint) GetJobType() JobTypes`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *VmRestorePoint) GetJobTypeOk() (*JobTypes, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *VmRestorePoint) SetJobType(v JobTypes)`

SetJobType sets JobType field to given value.


### GetPointInTime

`func (o *VmRestorePoint) GetPointInTime() time.Time`

GetPointInTime returns the PointInTime field if non-nil, zero value otherwise.

### GetPointInTimeOk

`func (o *VmRestorePoint) GetPointInTimeOk() (*time.Time, bool)`

GetPointInTimeOk returns a tuple with the PointInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTime

`func (o *VmRestorePoint) SetPointInTime(v time.Time)`

SetPointInTime sets PointInTime field to given value.


### GetBackupSizeBytes

`func (o *VmRestorePoint) GetBackupSizeBytes() int64`

GetBackupSizeBytes returns the BackupSizeBytes field if non-nil, zero value otherwise.

### GetBackupSizeBytesOk

`func (o *VmRestorePoint) GetBackupSizeBytesOk() (*int64, bool)`

GetBackupSizeBytesOk returns a tuple with the BackupSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSizeBytes

`func (o *VmRestorePoint) SetBackupSizeBytes(v int64)`

SetBackupSizeBytes sets BackupSizeBytes field to given value.


### GetVolumes

`func (o *VmRestorePoint) GetVolumes() []VmBackupVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *VmRestorePoint) GetVolumesOk() (*[]VmBackupVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *VmRestorePoint) SetVolumes(v []VmBackupVolume)`

SetVolumes sets Volumes field to given value.


### GetRegionId

`func (o *VmRestorePoint) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *VmRestorePoint) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *VmRestorePoint) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *VmRestorePoint) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetAvailabilityZoneId

`func (o *VmRestorePoint) GetAvailabilityZoneId() string`

GetAvailabilityZoneId returns the AvailabilityZoneId field if non-nil, zero value otherwise.

### GetAvailabilityZoneIdOk

`func (o *VmRestorePoint) GetAvailabilityZoneIdOk() (*string, bool)`

GetAvailabilityZoneIdOk returns a tuple with the AvailabilityZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneId

`func (o *VmRestorePoint) SetAvailabilityZoneId(v string)`

SetAvailabilityZoneId sets AvailabilityZoneId field to given value.

### HasAvailabilityZoneId

`func (o *VmRestorePoint) HasAvailabilityZoneId() bool`

HasAvailabilityZoneId returns a boolean if a field has been set.

### GetLinks

`func (o *VmRestorePoint) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VmRestorePoint) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VmRestorePoint) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VmRestorePoint) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


