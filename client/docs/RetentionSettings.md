# RetentionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LingeringSnapshotRetentionDisabled** | **bool** |  | 
**LingeringSnapshotRetention** | Pointer to [**RetentionSetting**](RetentionSetting.md) |  | [optional] 
**SessionsRetention** | Pointer to [**RetentionSetting**](RetentionSetting.md) |  | [optional] 
**KeepAllSessions** | **bool** |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewRetentionSettings

`func NewRetentionSettings(lingeringSnapshotRetentionDisabled bool, keepAllSessions bool, ) *RetentionSettings`

NewRetentionSettings instantiates a new RetentionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionSettingsWithDefaults

`func NewRetentionSettingsWithDefaults() *RetentionSettings`

NewRetentionSettingsWithDefaults instantiates a new RetentionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLingeringSnapshotRetentionDisabled

`func (o *RetentionSettings) GetLingeringSnapshotRetentionDisabled() bool`

GetLingeringSnapshotRetentionDisabled returns the LingeringSnapshotRetentionDisabled field if non-nil, zero value otherwise.

### GetLingeringSnapshotRetentionDisabledOk

`func (o *RetentionSettings) GetLingeringSnapshotRetentionDisabledOk() (*bool, bool)`

GetLingeringSnapshotRetentionDisabledOk returns a tuple with the LingeringSnapshotRetentionDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLingeringSnapshotRetentionDisabled

`func (o *RetentionSettings) SetLingeringSnapshotRetentionDisabled(v bool)`

SetLingeringSnapshotRetentionDisabled sets LingeringSnapshotRetentionDisabled field to given value.


### GetLingeringSnapshotRetention

`func (o *RetentionSettings) GetLingeringSnapshotRetention() RetentionSetting`

GetLingeringSnapshotRetention returns the LingeringSnapshotRetention field if non-nil, zero value otherwise.

### GetLingeringSnapshotRetentionOk

`func (o *RetentionSettings) GetLingeringSnapshotRetentionOk() (*RetentionSetting, bool)`

GetLingeringSnapshotRetentionOk returns a tuple with the LingeringSnapshotRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLingeringSnapshotRetention

`func (o *RetentionSettings) SetLingeringSnapshotRetention(v RetentionSetting)`

SetLingeringSnapshotRetention sets LingeringSnapshotRetention field to given value.

### HasLingeringSnapshotRetention

`func (o *RetentionSettings) HasLingeringSnapshotRetention() bool`

HasLingeringSnapshotRetention returns a boolean if a field has been set.

### GetSessionsRetention

`func (o *RetentionSettings) GetSessionsRetention() RetentionSetting`

GetSessionsRetention returns the SessionsRetention field if non-nil, zero value otherwise.

### GetSessionsRetentionOk

`func (o *RetentionSettings) GetSessionsRetentionOk() (*RetentionSetting, bool)`

GetSessionsRetentionOk returns a tuple with the SessionsRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsRetention

`func (o *RetentionSettings) SetSessionsRetention(v RetentionSetting)`

SetSessionsRetention sets SessionsRetention field to given value.

### HasSessionsRetention

`func (o *RetentionSettings) HasSessionsRetention() bool`

HasSessionsRetention returns a boolean if a field has been set.

### GetKeepAllSessions

`func (o *RetentionSettings) GetKeepAllSessions() bool`

GetKeepAllSessions returns the KeepAllSessions field if non-nil, zero value otherwise.

### GetKeepAllSessionsOk

`func (o *RetentionSettings) GetKeepAllSessionsOk() (*bool, bool)`

GetKeepAllSessionsOk returns a tuple with the KeepAllSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAllSessions

`func (o *RetentionSettings) SetKeepAllSessions(v bool)`

SetKeepAllSessions sets KeepAllSessions field to given value.


### GetLinks

`func (o *RetentionSettings) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RetentionSettings) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RetentionSettings) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RetentionSettings) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


