# WeeklyScheduleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeLocal** | **time.Time** |  | 
**SnapshotOptions** | [**WeeklySnapshotScheduleSettings**](WeeklySnapshotScheduleSettings.md) |  | 
**BackupOptions** | Pointer to [**WeeklyBackupScheduleSettings**](WeeklyBackupScheduleSettings.md) |  | [optional] 
**ReplicaOptions** | Pointer to [**WeeklyReplicaScheduleSettings**](WeeklyReplicaScheduleSettings.md) |  | [optional] 

## Methods

### NewWeeklyScheduleSettings

`func NewWeeklyScheduleSettings(timeLocal time.Time, snapshotOptions WeeklySnapshotScheduleSettings, ) *WeeklyScheduleSettings`

NewWeeklyScheduleSettings instantiates a new WeeklyScheduleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeeklyScheduleSettingsWithDefaults

`func NewWeeklyScheduleSettingsWithDefaults() *WeeklyScheduleSettings`

NewWeeklyScheduleSettingsWithDefaults instantiates a new WeeklyScheduleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeLocal

`func (o *WeeklyScheduleSettings) GetTimeLocal() time.Time`

GetTimeLocal returns the TimeLocal field if non-nil, zero value otherwise.

### GetTimeLocalOk

`func (o *WeeklyScheduleSettings) GetTimeLocalOk() (*time.Time, bool)`

GetTimeLocalOk returns a tuple with the TimeLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLocal

`func (o *WeeklyScheduleSettings) SetTimeLocal(v time.Time)`

SetTimeLocal sets TimeLocal field to given value.


### GetSnapshotOptions

`func (o *WeeklyScheduleSettings) GetSnapshotOptions() WeeklySnapshotScheduleSettings`

GetSnapshotOptions returns the SnapshotOptions field if non-nil, zero value otherwise.

### GetSnapshotOptionsOk

`func (o *WeeklyScheduleSettings) GetSnapshotOptionsOk() (*WeeklySnapshotScheduleSettings, bool)`

GetSnapshotOptionsOk returns a tuple with the SnapshotOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotOptions

`func (o *WeeklyScheduleSettings) SetSnapshotOptions(v WeeklySnapshotScheduleSettings)`

SetSnapshotOptions sets SnapshotOptions field to given value.


### GetBackupOptions

`func (o *WeeklyScheduleSettings) GetBackupOptions() WeeklyBackupScheduleSettings`

GetBackupOptions returns the BackupOptions field if non-nil, zero value otherwise.

### GetBackupOptionsOk

`func (o *WeeklyScheduleSettings) GetBackupOptionsOk() (*WeeklyBackupScheduleSettings, bool)`

GetBackupOptionsOk returns a tuple with the BackupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOptions

`func (o *WeeklyScheduleSettings) SetBackupOptions(v WeeklyBackupScheduleSettings)`

SetBackupOptions sets BackupOptions field to given value.

### HasBackupOptions

`func (o *WeeklyScheduleSettings) HasBackupOptions() bool`

HasBackupOptions returns a boolean if a field has been set.

### GetReplicaOptions

`func (o *WeeklyScheduleSettings) GetReplicaOptions() WeeklyReplicaScheduleSettings`

GetReplicaOptions returns the ReplicaOptions field if non-nil, zero value otherwise.

### GetReplicaOptionsOk

`func (o *WeeklyScheduleSettings) GetReplicaOptionsOk() (*WeeklyReplicaScheduleSettings, bool)`

GetReplicaOptionsOk returns a tuple with the ReplicaOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaOptions

`func (o *WeeklyScheduleSettings) SetReplicaOptions(v WeeklyReplicaScheduleSettings)`

SetReplicaOptions sets ReplicaOptions field to given value.

### HasReplicaOptions

`func (o *WeeklyScheduleSettings) HasReplicaOptions() bool`

HasReplicaOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


