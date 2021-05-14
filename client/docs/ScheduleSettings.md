# ScheduleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyScheduleEnabled** | **bool** |  | 
**DailySchedule** | Pointer to [**DailyScheduleSettings**](DailyScheduleSettings.md) |  | [optional] 
**WeeklyScheduleEnabled** | **bool** |  | 
**WeeklySchedule** | Pointer to [**WeeklyScheduleSettings**](WeeklyScheduleSettings.md) |  | [optional] 
**MonthlyScheduleEnabled** | **bool** |  | 
**MonthlySchedule** | Pointer to [**MonthlyScheduleSettings**](MonthlyScheduleSettings.md) |  | [optional] 
**YearlyScheduleEnabled** | **bool** |  | 
**YearlySchedule** | Pointer to [**YearlyScheduleSettings**](YearlyScheduleSettings.md) |  | [optional] 

## Methods

### NewScheduleSettings

`func NewScheduleSettings(dailyScheduleEnabled bool, weeklyScheduleEnabled bool, monthlyScheduleEnabled bool, yearlyScheduleEnabled bool, ) *ScheduleSettings`

NewScheduleSettings instantiates a new ScheduleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleSettingsWithDefaults

`func NewScheduleSettingsWithDefaults() *ScheduleSettings`

NewScheduleSettingsWithDefaults instantiates a new ScheduleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyScheduleEnabled

`func (o *ScheduleSettings) GetDailyScheduleEnabled() bool`

GetDailyScheduleEnabled returns the DailyScheduleEnabled field if non-nil, zero value otherwise.

### GetDailyScheduleEnabledOk

`func (o *ScheduleSettings) GetDailyScheduleEnabledOk() (*bool, bool)`

GetDailyScheduleEnabledOk returns a tuple with the DailyScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyScheduleEnabled

`func (o *ScheduleSettings) SetDailyScheduleEnabled(v bool)`

SetDailyScheduleEnabled sets DailyScheduleEnabled field to given value.


### GetDailySchedule

`func (o *ScheduleSettings) GetDailySchedule() DailyScheduleSettings`

GetDailySchedule returns the DailySchedule field if non-nil, zero value otherwise.

### GetDailyScheduleOk

`func (o *ScheduleSettings) GetDailyScheduleOk() (*DailyScheduleSettings, bool)`

GetDailyScheduleOk returns a tuple with the DailySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySchedule

`func (o *ScheduleSettings) SetDailySchedule(v DailyScheduleSettings)`

SetDailySchedule sets DailySchedule field to given value.

### HasDailySchedule

`func (o *ScheduleSettings) HasDailySchedule() bool`

HasDailySchedule returns a boolean if a field has been set.

### GetWeeklyScheduleEnabled

`func (o *ScheduleSettings) GetWeeklyScheduleEnabled() bool`

GetWeeklyScheduleEnabled returns the WeeklyScheduleEnabled field if non-nil, zero value otherwise.

### GetWeeklyScheduleEnabledOk

`func (o *ScheduleSettings) GetWeeklyScheduleEnabledOk() (*bool, bool)`

GetWeeklyScheduleEnabledOk returns a tuple with the WeeklyScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyScheduleEnabled

`func (o *ScheduleSettings) SetWeeklyScheduleEnabled(v bool)`

SetWeeklyScheduleEnabled sets WeeklyScheduleEnabled field to given value.


### GetWeeklySchedule

`func (o *ScheduleSettings) GetWeeklySchedule() WeeklyScheduleSettings`

GetWeeklySchedule returns the WeeklySchedule field if non-nil, zero value otherwise.

### GetWeeklyScheduleOk

`func (o *ScheduleSettings) GetWeeklyScheduleOk() (*WeeklyScheduleSettings, bool)`

GetWeeklyScheduleOk returns a tuple with the WeeklySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklySchedule

`func (o *ScheduleSettings) SetWeeklySchedule(v WeeklyScheduleSettings)`

SetWeeklySchedule sets WeeklySchedule field to given value.

### HasWeeklySchedule

`func (o *ScheduleSettings) HasWeeklySchedule() bool`

HasWeeklySchedule returns a boolean if a field has been set.

### GetMonthlyScheduleEnabled

`func (o *ScheduleSettings) GetMonthlyScheduleEnabled() bool`

GetMonthlyScheduleEnabled returns the MonthlyScheduleEnabled field if non-nil, zero value otherwise.

### GetMonthlyScheduleEnabledOk

`func (o *ScheduleSettings) GetMonthlyScheduleEnabledOk() (*bool, bool)`

GetMonthlyScheduleEnabledOk returns a tuple with the MonthlyScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyScheduleEnabled

`func (o *ScheduleSettings) SetMonthlyScheduleEnabled(v bool)`

SetMonthlyScheduleEnabled sets MonthlyScheduleEnabled field to given value.


### GetMonthlySchedule

`func (o *ScheduleSettings) GetMonthlySchedule() MonthlyScheduleSettings`

GetMonthlySchedule returns the MonthlySchedule field if non-nil, zero value otherwise.

### GetMonthlyScheduleOk

`func (o *ScheduleSettings) GetMonthlyScheduleOk() (*MonthlyScheduleSettings, bool)`

GetMonthlyScheduleOk returns a tuple with the MonthlySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlySchedule

`func (o *ScheduleSettings) SetMonthlySchedule(v MonthlyScheduleSettings)`

SetMonthlySchedule sets MonthlySchedule field to given value.

### HasMonthlySchedule

`func (o *ScheduleSettings) HasMonthlySchedule() bool`

HasMonthlySchedule returns a boolean if a field has been set.

### GetYearlyScheduleEnabled

`func (o *ScheduleSettings) GetYearlyScheduleEnabled() bool`

GetYearlyScheduleEnabled returns the YearlyScheduleEnabled field if non-nil, zero value otherwise.

### GetYearlyScheduleEnabledOk

`func (o *ScheduleSettings) GetYearlyScheduleEnabledOk() (*bool, bool)`

GetYearlyScheduleEnabledOk returns a tuple with the YearlyScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyScheduleEnabled

`func (o *ScheduleSettings) SetYearlyScheduleEnabled(v bool)`

SetYearlyScheduleEnabled sets YearlyScheduleEnabled field to given value.


### GetYearlySchedule

`func (o *ScheduleSettings) GetYearlySchedule() YearlyScheduleSettings`

GetYearlySchedule returns the YearlySchedule field if non-nil, zero value otherwise.

### GetYearlyScheduleOk

`func (o *ScheduleSettings) GetYearlyScheduleOk() (*YearlyScheduleSettings, bool)`

GetYearlyScheduleOk returns a tuple with the YearlySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlySchedule

`func (o *ScheduleSettings) SetYearlySchedule(v YearlyScheduleSettings)`

SetYearlySchedule sets YearlySchedule field to given value.

### HasYearlySchedule

`func (o *ScheduleSettings) HasYearlySchedule() bool`

HasYearlySchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


