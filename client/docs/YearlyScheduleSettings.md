# YearlyScheduleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeLocal** | **time.Time** |  | 
**DayNumberInMonth** | [**DayNumbersInMonth**](DayNumbersInMonth.md) |  | 
**DayOfWeek** | Pointer to [**DaysOfWeek**](DaysOfWeek.md) |  | [optional] 
**DayOfMonth** | Pointer to **int32** |  | [optional] 
**Month** | [**Months**](Months.md) |  | 
**Retention** | [**YearlyRetentionOptions**](YearlyRetentionOptions.md) |  | 

## Methods

### NewYearlyScheduleSettings

`func NewYearlyScheduleSettings(timeLocal time.Time, dayNumberInMonth DayNumbersInMonth, month Months, retention YearlyRetentionOptions, ) *YearlyScheduleSettings`

NewYearlyScheduleSettings instantiates a new YearlyScheduleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYearlyScheduleSettingsWithDefaults

`func NewYearlyScheduleSettingsWithDefaults() *YearlyScheduleSettings`

NewYearlyScheduleSettingsWithDefaults instantiates a new YearlyScheduleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeLocal

`func (o *YearlyScheduleSettings) GetTimeLocal() time.Time`

GetTimeLocal returns the TimeLocal field if non-nil, zero value otherwise.

### GetTimeLocalOk

`func (o *YearlyScheduleSettings) GetTimeLocalOk() (*time.Time, bool)`

GetTimeLocalOk returns a tuple with the TimeLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLocal

`func (o *YearlyScheduleSettings) SetTimeLocal(v time.Time)`

SetTimeLocal sets TimeLocal field to given value.


### GetDayNumberInMonth

`func (o *YearlyScheduleSettings) GetDayNumberInMonth() DayNumbersInMonth`

GetDayNumberInMonth returns the DayNumberInMonth field if non-nil, zero value otherwise.

### GetDayNumberInMonthOk

`func (o *YearlyScheduleSettings) GetDayNumberInMonthOk() (*DayNumbersInMonth, bool)`

GetDayNumberInMonthOk returns a tuple with the DayNumberInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayNumberInMonth

`func (o *YearlyScheduleSettings) SetDayNumberInMonth(v DayNumbersInMonth)`

SetDayNumberInMonth sets DayNumberInMonth field to given value.


### GetDayOfWeek

`func (o *YearlyScheduleSettings) GetDayOfWeek() DaysOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *YearlyScheduleSettings) GetDayOfWeekOk() (*DaysOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *YearlyScheduleSettings) SetDayOfWeek(v DaysOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *YearlyScheduleSettings) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *YearlyScheduleSettings) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *YearlyScheduleSettings) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *YearlyScheduleSettings) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *YearlyScheduleSettings) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetMonth

`func (o *YearlyScheduleSettings) GetMonth() Months`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *YearlyScheduleSettings) GetMonthOk() (*Months, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *YearlyScheduleSettings) SetMonth(v Months)`

SetMonth sets Month field to given value.


### GetRetention

`func (o *YearlyScheduleSettings) GetRetention() YearlyRetentionOptions`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *YearlyScheduleSettings) GetRetentionOk() (*YearlyRetentionOptions, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *YearlyScheduleSettings) SetRetention(v YearlyRetentionOptions)`

SetRetention sets Retention field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


