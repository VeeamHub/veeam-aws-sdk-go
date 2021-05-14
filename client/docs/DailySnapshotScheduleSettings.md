# DailySnapshotScheduleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retention** | [**PointsRetentionOptions**](PointsRetentionOptions.md) |  | 
**Schedule** | [**DailySchedule**](DailySchedule.md) |  | 

## Methods

### NewDailySnapshotScheduleSettings

`func NewDailySnapshotScheduleSettings(retention PointsRetentionOptions, schedule DailySchedule, ) *DailySnapshotScheduleSettings`

NewDailySnapshotScheduleSettings instantiates a new DailySnapshotScheduleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailySnapshotScheduleSettingsWithDefaults

`func NewDailySnapshotScheduleSettingsWithDefaults() *DailySnapshotScheduleSettings`

NewDailySnapshotScheduleSettingsWithDefaults instantiates a new DailySnapshotScheduleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetention

`func (o *DailySnapshotScheduleSettings) GetRetention() PointsRetentionOptions`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *DailySnapshotScheduleSettings) GetRetentionOk() (*PointsRetentionOptions, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *DailySnapshotScheduleSettings) SetRetention(v PointsRetentionOptions)`

SetRetention sets Retention field to given value.


### GetSchedule

`func (o *DailySnapshotScheduleSettings) GetSchedule() DailySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *DailySnapshotScheduleSettings) GetScheduleOk() (*DailySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *DailySnapshotScheduleSettings) SetSchedule(v DailySchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


