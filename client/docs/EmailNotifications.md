# EmailNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Server** | **string** |  | 
**Port** | **int32** |  | 
**UseSsl** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**UseCredentials** | Pointer to **bool** |  | [optional] 
**CredentialsId** | Pointer to **string** |  | [optional] 
**From** | **string** |  | 
**To** | **string** |  | 
**Subject** | **string** |  | 
**OnSuccess** | Pointer to **bool** |  | [optional] 
**OnWarning** | Pointer to **bool** |  | [optional] 
**OnFailure** | Pointer to **bool** |  | [optional] 
**EnableScheduledNotification** | Pointer to **bool** |  | [optional] 
**NotifyTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEmailNotifications

`func NewEmailNotifications(enabled bool, server string, port int32, from string, to string, subject string, ) *EmailNotifications`

NewEmailNotifications instantiates a new EmailNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailNotificationsWithDefaults

`func NewEmailNotificationsWithDefaults() *EmailNotifications`

NewEmailNotificationsWithDefaults instantiates a new EmailNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *EmailNotifications) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EmailNotifications) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EmailNotifications) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetServer

`func (o *EmailNotifications) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *EmailNotifications) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *EmailNotifications) SetServer(v string)`

SetServer sets Server field to given value.


### GetPort

`func (o *EmailNotifications) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EmailNotifications) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EmailNotifications) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUseSsl

`func (o *EmailNotifications) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *EmailNotifications) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *EmailNotifications) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *EmailNotifications) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetTimeout

`func (o *EmailNotifications) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EmailNotifications) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EmailNotifications) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EmailNotifications) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUseCredentials

`func (o *EmailNotifications) GetUseCredentials() bool`

GetUseCredentials returns the UseCredentials field if non-nil, zero value otherwise.

### GetUseCredentialsOk

`func (o *EmailNotifications) GetUseCredentialsOk() (*bool, bool)`

GetUseCredentialsOk returns a tuple with the UseCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCredentials

`func (o *EmailNotifications) SetUseCredentials(v bool)`

SetUseCredentials sets UseCredentials field to given value.

### HasUseCredentials

`func (o *EmailNotifications) HasUseCredentials() bool`

HasUseCredentials returns a boolean if a field has been set.

### GetCredentialsId

`func (o *EmailNotifications) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *EmailNotifications) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *EmailNotifications) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.

### HasCredentialsId

`func (o *EmailNotifications) HasCredentialsId() bool`

HasCredentialsId returns a boolean if a field has been set.

### GetFrom

`func (o *EmailNotifications) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailNotifications) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailNotifications) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *EmailNotifications) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailNotifications) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailNotifications) SetTo(v string)`

SetTo sets To field to given value.


### GetSubject

`func (o *EmailNotifications) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailNotifications) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailNotifications) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetOnSuccess

`func (o *EmailNotifications) GetOnSuccess() bool`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *EmailNotifications) GetOnSuccessOk() (*bool, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *EmailNotifications) SetOnSuccess(v bool)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *EmailNotifications) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.

### GetOnWarning

`func (o *EmailNotifications) GetOnWarning() bool`

GetOnWarning returns the OnWarning field if non-nil, zero value otherwise.

### GetOnWarningOk

`func (o *EmailNotifications) GetOnWarningOk() (*bool, bool)`

GetOnWarningOk returns a tuple with the OnWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnWarning

`func (o *EmailNotifications) SetOnWarning(v bool)`

SetOnWarning sets OnWarning field to given value.

### HasOnWarning

`func (o *EmailNotifications) HasOnWarning() bool`

HasOnWarning returns a boolean if a field has been set.

### GetOnFailure

`func (o *EmailNotifications) GetOnFailure() bool`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *EmailNotifications) GetOnFailureOk() (*bool, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *EmailNotifications) SetOnFailure(v bool)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *EmailNotifications) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetEnableScheduledNotification

`func (o *EmailNotifications) GetEnableScheduledNotification() bool`

GetEnableScheduledNotification returns the EnableScheduledNotification field if non-nil, zero value otherwise.

### GetEnableScheduledNotificationOk

`func (o *EmailNotifications) GetEnableScheduledNotificationOk() (*bool, bool)`

GetEnableScheduledNotificationOk returns a tuple with the EnableScheduledNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableScheduledNotification

`func (o *EmailNotifications) SetEnableScheduledNotification(v bool)`

SetEnableScheduledNotification sets EnableScheduledNotification field to given value.

### HasEnableScheduledNotification

`func (o *EmailNotifications) HasEnableScheduledNotification() bool`

HasEnableScheduledNotification returns a boolean if a field has been set.

### GetNotifyTime

`func (o *EmailNotifications) GetNotifyTime() time.Time`

GetNotifyTime returns the NotifyTime field if non-nil, zero value otherwise.

### GetNotifyTimeOk

`func (o *EmailNotifications) GetNotifyTimeOk() (*time.Time, bool)`

GetNotifyTimeOk returns a tuple with the NotifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTime

`func (o *EmailNotifications) SetNotifyTime(v time.Time)`

SetNotifyTime sets NotifyTime field to given value.

### HasNotifyTime

`func (o *EmailNotifications) HasNotifyTime() bool`

HasNotifyTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


