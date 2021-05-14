# WeeklyBackupScheduleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retention** | [**PeriodRetentionOptions**](PeriodRetentionOptions.md) |  | 
**Schedule** | [**WeeklySchedule**](WeeklySchedule.md) |  | 

## Methods

### NewWeeklyBackupScheduleSettings

`func NewWeeklyBackupScheduleSettings(retention PeriodRetentionOptions, schedule WeeklySchedule, ) *WeeklyBackupScheduleSettings`

NewWeeklyBackupScheduleSettings instantiates a new WeeklyBackupScheduleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeeklyBackupScheduleSettingsWithDefaults

`func NewWeeklyBackupScheduleSettingsWithDefaults() *WeeklyBackupScheduleSettings`

NewWeeklyBackupScheduleSettingsWithDefaults instantiates a new WeeklyBackupScheduleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetention

`func (o *WeeklyBackupScheduleSettings) GetRetention() PeriodRetentionOptions`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *WeeklyBackupScheduleSettings) GetRetentionOk() (*PeriodRetentionOptions, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *WeeklyBackupScheduleSettings) SetRetention(v PeriodRetentionOptions)`

SetRetention sets Retention field to given value.


### GetSchedule

`func (o *WeeklyBackupScheduleSettings) GetSchedule() WeeklySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *WeeklyBackupScheduleSettings) GetScheduleOk() (*WeeklySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *WeeklyBackupScheduleSettings) SetSchedule(v WeeklySchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


