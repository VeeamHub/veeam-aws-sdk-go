# DateTimeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | [**TimeZone**](TimeZone.md) |  | 
**UseCustomSettings** | **bool** |  | 

## Methods

### NewDateTimeSettings

`func NewDateTimeSettings(timezone TimeZone, useCustomSettings bool, ) *DateTimeSettings`

NewDateTimeSettings instantiates a new DateTimeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeSettingsWithDefaults

`func NewDateTimeSettingsWithDefaults() *DateTimeSettings`

NewDateTimeSettingsWithDefaults instantiates a new DateTimeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *DateTimeSettings) GetTimezone() TimeZone`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *DateTimeSettings) GetTimezoneOk() (*TimeZone, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *DateTimeSettings) SetTimezone(v TimeZone)`

SetTimezone sets Timezone field to given value.


### GetUseCustomSettings

`func (o *DateTimeSettings) GetUseCustomSettings() bool`

GetUseCustomSettings returns the UseCustomSettings field if non-nil, zero value otherwise.

### GetUseCustomSettingsOk

`func (o *DateTimeSettings) GetUseCustomSettingsOk() (*bool, bool)`

GetUseCustomSettingsOk returns a tuple with the UseCustomSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomSettings

`func (o *DateTimeSettings) SetUseCustomSettings(v bool)`

SetUseCustomSettings sets UseCustomSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


