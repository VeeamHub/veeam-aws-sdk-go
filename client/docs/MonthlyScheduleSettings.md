# MonthlyScheduleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeLocal** | **time.Time** |  | 
**DayNumberInMonth** | [**DayNumbersInMonth**](DayNumbersInMonth.md) |  | 
**DayOfWeek** | Pointer to [**DaysOfWeek**](DaysOfWeek.md) |  | [optional] 
**DayOfMonth** | Pointer to **int32** |  | [optional] 
**SnapshotOptions** | [**MonthlySnapshotScheduleSettings**](MonthlySnapshotScheduleSettings.md) |  | 
**BackupOptions** | Pointer to [**MonthlyBackupScheduleSettings**](MonthlyBackupScheduleSettings.md) |  | [optional] 
**ReplicaOptions** | Pointer to [**MonthlyReplicaScheduleSettings**](MonthlyReplicaScheduleSettings.md) |  | [optional] 

## Methods

### NewMonthlyScheduleSettings

`func NewMonthlyScheduleSettings(timeLocal time.Time, dayNumberInMonth DayNumbersInMonth, snapshotOptions MonthlySnapshotScheduleSettings, ) *MonthlyScheduleSettings`

NewMonthlyScheduleSettings instantiates a new MonthlyScheduleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthlyScheduleSettingsWithDefaults

`func NewMonthlyScheduleSettingsWithDefaults() *MonthlyScheduleSettings`

NewMonthlyScheduleSettingsWithDefaults instantiates a new MonthlyScheduleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeLocal

`func (o *MonthlyScheduleSettings) GetTimeLocal() time.Time`

GetTimeLocal returns the TimeLocal field if non-nil, zero value otherwise.

### GetTimeLocalOk

`func (o *MonthlyScheduleSettings) GetTimeLocalOk() (*time.Time, bool)`

GetTimeLocalOk returns a tuple with the TimeLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLocal

`func (o *MonthlyScheduleSettings) SetTimeLocal(v time.Time)`

SetTimeLocal sets TimeLocal field to given value.


### GetDayNumberInMonth

`func (o *MonthlyScheduleSettings) GetDayNumberInMonth() DayNumbersInMonth`

GetDayNumberInMonth returns the DayNumberInMonth field if non-nil, zero value otherwise.

### GetDayNumberInMonthOk

`func (o *MonthlyScheduleSettings) GetDayNumberInMonthOk() (*DayNumbersInMonth, bool)`

GetDayNumberInMonthOk returns a tuple with the DayNumberInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayNumberInMonth

`func (o *MonthlyScheduleSettings) SetDayNumberInMonth(v DayNumbersInMonth)`

SetDayNumberInMonth sets DayNumberInMonth field to given value.


### GetDayOfWeek

`func (o *MonthlyScheduleSettings) GetDayOfWeek() DaysOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *MonthlyScheduleSettings) GetDayOfWeekOk() (*DaysOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *MonthlyScheduleSettings) SetDayOfWeek(v DaysOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *MonthlyScheduleSettings) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *MonthlyScheduleSettings) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *MonthlyScheduleSettings) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *MonthlyScheduleSettings) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *MonthlyScheduleSettings) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetSnapshotOptions

`func (o *MonthlyScheduleSettings) GetSnapshotOptions() MonthlySnapshotScheduleSettings`

GetSnapshotOptions returns the SnapshotOptions field if non-nil, zero value otherwise.

### GetSnapshotOptionsOk

`func (o *MonthlyScheduleSettings) GetSnapshotOptionsOk() (*MonthlySnapshotScheduleSettings, bool)`

GetSnapshotOptionsOk returns a tuple with the SnapshotOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotOptions

`func (o *MonthlyScheduleSettings) SetSnapshotOptions(v MonthlySnapshotScheduleSettings)`

SetSnapshotOptions sets SnapshotOptions field to given value.


### GetBackupOptions

`func (o *MonthlyScheduleSettings) GetBackupOptions() MonthlyBackupScheduleSettings`

GetBackupOptions returns the BackupOptions field if non-nil, zero value otherwise.

### GetBackupOptionsOk

`func (o *MonthlyScheduleSettings) GetBackupOptionsOk() (*MonthlyBackupScheduleSettings, bool)`

GetBackupOptionsOk returns a tuple with the BackupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOptions

`func (o *MonthlyScheduleSettings) SetBackupOptions(v MonthlyBackupScheduleSettings)`

SetBackupOptions sets BackupOptions field to given value.

### HasBackupOptions

`func (o *MonthlyScheduleSettings) HasBackupOptions() bool`

HasBackupOptions returns a boolean if a field has been set.

### GetReplicaOptions

`func (o *MonthlyScheduleSettings) GetReplicaOptions() MonthlyReplicaScheduleSettings`

GetReplicaOptions returns the ReplicaOptions field if non-nil, zero value otherwise.

### GetReplicaOptionsOk

`func (o *MonthlyScheduleSettings) GetReplicaOptionsOk() (*MonthlyReplicaScheduleSettings, bool)`

GetReplicaOptionsOk returns a tuple with the ReplicaOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaOptions

`func (o *MonthlyScheduleSettings) SetReplicaOptions(v MonthlyReplicaScheduleSettings)`

SetReplicaOptions sets ReplicaOptions field to given value.

### HasReplicaOptions

`func (o *MonthlyScheduleSettings) HasReplicaOptions() bool`

HasReplicaOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


