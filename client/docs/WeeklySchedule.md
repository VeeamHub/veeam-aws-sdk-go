# WeeklySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | [**[]DaysOfWeek**](DaysOfWeek.md) |  | 

## Methods

### NewWeeklySchedule

`func NewWeeklySchedule(days []DaysOfWeek, ) *WeeklySchedule`

NewWeeklySchedule instantiates a new WeeklySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeeklyScheduleWithDefaults

`func NewWeeklyScheduleWithDefaults() *WeeklySchedule`

NewWeeklyScheduleWithDefaults instantiates a new WeeklySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *WeeklySchedule) GetDays() []DaysOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *WeeklySchedule) GetDaysOk() (*[]DaysOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *WeeklySchedule) SetDays(v []DaysOfWeek)`

SetDays sets Days field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


