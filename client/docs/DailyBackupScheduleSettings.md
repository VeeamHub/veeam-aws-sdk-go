# DailyBackupScheduleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retention** | [**PeriodRetentionOptions**](PeriodRetentionOptions.md) |  | 
**Schedule** | [**DailySchedule**](DailySchedule.md) |  | 

## Methods

### NewDailyBackupScheduleSettings

`func NewDailyBackupScheduleSettings(retention PeriodRetentionOptions, schedule DailySchedule, ) *DailyBackupScheduleSettings`

NewDailyBackupScheduleSettings instantiates a new DailyBackupScheduleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyBackupScheduleSettingsWithDefaults

`func NewDailyBackupScheduleSettingsWithDefaults() *DailyBackupScheduleSettings`

NewDailyBackupScheduleSettingsWithDefaults instantiates a new DailyBackupScheduleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetention

`func (o *DailyBackupScheduleSettings) GetRetention() PeriodRetentionOptions`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *DailyBackupScheduleSettings) GetRetentionOk() (*PeriodRetentionOptions, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *DailyBackupScheduleSettings) SetRetention(v PeriodRetentionOptions)`

SetRetention sets Retention field to given value.


### GetSchedule

`func (o *DailyBackupScheduleSettings) GetSchedule() DailySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *DailyBackupScheduleSettings) GetScheduleOk() (*DailySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *DailyBackupScheduleSettings) SetSchedule(v DailySchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


