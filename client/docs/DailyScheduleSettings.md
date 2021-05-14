# DailyScheduleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**DailyScheduleKinds**](DailyScheduleKinds.md) |  | 
**RunsPerHour** | **int32** |  | 
**Days** | Pointer to [**[]DaysOfWeek**](DaysOfWeek.md) |  | [optional] 
**SnapshotOptions** | [**DailySnapshotScheduleSettings**](DailySnapshotScheduleSettings.md) |  | 
**BackupOptions** | Pointer to [**DailyBackupScheduleSettings**](DailyBackupScheduleSettings.md) |  | [optional] 
**ReplicaOptions** | Pointer to [**DailyReplicaScheduleSettings**](DailyReplicaScheduleSettings.md) |  | [optional] 

## Methods

### NewDailyScheduleSettings

`func NewDailyScheduleSettings(kind DailyScheduleKinds, runsPerHour int32, snapshotOptions DailySnapshotScheduleSettings, ) *DailyScheduleSettings`

NewDailyScheduleSettings instantiates a new DailyScheduleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyScheduleSettingsWithDefaults

`func NewDailyScheduleSettingsWithDefaults() *DailyScheduleSettings`

NewDailyScheduleSettingsWithDefaults instantiates a new DailyScheduleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *DailyScheduleSettings) GetKind() DailyScheduleKinds`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DailyScheduleSettings) GetKindOk() (*DailyScheduleKinds, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DailyScheduleSettings) SetKind(v DailyScheduleKinds)`

SetKind sets Kind field to given value.


### GetRunsPerHour

`func (o *DailyScheduleSettings) GetRunsPerHour() int32`

GetRunsPerHour returns the RunsPerHour field if non-nil, zero value otherwise.

### GetRunsPerHourOk

`func (o *DailyScheduleSettings) GetRunsPerHourOk() (*int32, bool)`

GetRunsPerHourOk returns a tuple with the RunsPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunsPerHour

`func (o *DailyScheduleSettings) SetRunsPerHour(v int32)`

SetRunsPerHour sets RunsPerHour field to given value.


### GetDays

`func (o *DailyScheduleSettings) GetDays() []DaysOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *DailyScheduleSettings) GetDaysOk() (*[]DaysOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *DailyScheduleSettings) SetDays(v []DaysOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *DailyScheduleSettings) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetSnapshotOptions

`func (o *DailyScheduleSettings) GetSnapshotOptions() DailySnapshotScheduleSettings`

GetSnapshotOptions returns the SnapshotOptions field if non-nil, zero value otherwise.

### GetSnapshotOptionsOk

`func (o *DailyScheduleSettings) GetSnapshotOptionsOk() (*DailySnapshotScheduleSettings, bool)`

GetSnapshotOptionsOk returns a tuple with the SnapshotOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotOptions

`func (o *DailyScheduleSettings) SetSnapshotOptions(v DailySnapshotScheduleSettings)`

SetSnapshotOptions sets SnapshotOptions field to given value.


### GetBackupOptions

`func (o *DailyScheduleSettings) GetBackupOptions() DailyBackupScheduleSettings`

GetBackupOptions returns the BackupOptions field if non-nil, zero value otherwise.

### GetBackupOptionsOk

`func (o *DailyScheduleSettings) GetBackupOptionsOk() (*DailyBackupScheduleSettings, bool)`

GetBackupOptionsOk returns a tuple with the BackupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOptions

`func (o *DailyScheduleSettings) SetBackupOptions(v DailyBackupScheduleSettings)`

SetBackupOptions sets BackupOptions field to given value.

### HasBackupOptions

`func (o *DailyScheduleSettings) HasBackupOptions() bool`

HasBackupOptions returns a boolean if a field has been set.

### GetReplicaOptions

`func (o *DailyScheduleSettings) GetReplicaOptions() DailyReplicaScheduleSettings`

GetReplicaOptions returns the ReplicaOptions field if non-nil, zero value otherwise.

### GetReplicaOptionsOk

`func (o *DailyScheduleSettings) GetReplicaOptionsOk() (*DailyReplicaScheduleSettings, bool)`

GetReplicaOptionsOk returns a tuple with the ReplicaOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaOptions

`func (o *DailyScheduleSettings) SetReplicaOptions(v DailyReplicaScheduleSettings)`

SetReplicaOptions sets ReplicaOptions field to given value.

### HasReplicaOptions

`func (o *DailyScheduleSettings) HasReplicaOptions() bool`

HasReplicaOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


