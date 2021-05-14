# MonthlyReplicaScheduleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retention** | [**PointsRetentionOptions**](PointsRetentionOptions.md) |  | 
**Schedule** | [**MonthlySchedule**](MonthlySchedule.md) |  | 

## Methods

### NewMonthlyReplicaScheduleSettings

`func NewMonthlyReplicaScheduleSettings(retention PointsRetentionOptions, schedule MonthlySchedule, ) *MonthlyReplicaScheduleSettings`

NewMonthlyReplicaScheduleSettings instantiates a new MonthlyReplicaScheduleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthlyReplicaScheduleSettingsWithDefaults

`func NewMonthlyReplicaScheduleSettingsWithDefaults() *MonthlyReplicaScheduleSettings`

NewMonthlyReplicaScheduleSettingsWithDefaults instantiates a new MonthlyReplicaScheduleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetention

`func (o *MonthlyReplicaScheduleSettings) GetRetention() PointsRetentionOptions`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *MonthlyReplicaScheduleSettings) GetRetentionOk() (*PointsRetentionOptions, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *MonthlyReplicaScheduleSettings) SetRetention(v PointsRetentionOptions)`

SetRetention sets Retention field to given value.


### GetSchedule

`func (o *MonthlyReplicaScheduleSettings) GetSchedule() MonthlySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MonthlyReplicaScheduleSettings) GetScheduleOk() (*MonthlySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MonthlyReplicaScheduleSettings) SetSchedule(v MonthlySchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


