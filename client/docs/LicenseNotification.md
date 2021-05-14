# LicenseNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | [**LicenseNotificationTypes**](LicenseNotificationTypes.md) |  | 
**NotificationSeverity** | [**LicenseNotificationSeverities**](LicenseNotificationSeverities.md) |  | 
**Message** | **string** |  | 
**Replies** | Pointer to [**[]LicenseNotificationReply**](LicenseNotificationReply.md) |  | [optional] 

## Methods

### NewLicenseNotification

`func NewLicenseNotification(notificationType LicenseNotificationTypes, notificationSeverity LicenseNotificationSeverities, message string, ) *LicenseNotification`

NewLicenseNotification instantiates a new LicenseNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseNotificationWithDefaults

`func NewLicenseNotificationWithDefaults() *LicenseNotification`

NewLicenseNotificationWithDefaults instantiates a new LicenseNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *LicenseNotification) GetNotificationType() LicenseNotificationTypes`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *LicenseNotification) GetNotificationTypeOk() (*LicenseNotificationTypes, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *LicenseNotification) SetNotificationType(v LicenseNotificationTypes)`

SetNotificationType sets NotificationType field to given value.


### GetNotificationSeverity

`func (o *LicenseNotification) GetNotificationSeverity() LicenseNotificationSeverities`

GetNotificationSeverity returns the NotificationSeverity field if non-nil, zero value otherwise.

### GetNotificationSeverityOk

`func (o *LicenseNotification) GetNotificationSeverityOk() (*LicenseNotificationSeverities, bool)`

GetNotificationSeverityOk returns a tuple with the NotificationSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSeverity

`func (o *LicenseNotification) SetNotificationSeverity(v LicenseNotificationSeverities)`

SetNotificationSeverity sets NotificationSeverity field to given value.


### GetMessage

`func (o *LicenseNotification) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LicenseNotification) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LicenseNotification) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReplies

`func (o *LicenseNotification) GetReplies() []LicenseNotificationReply`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *LicenseNotification) GetRepliesOk() (*[]LicenseNotificationReply, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *LicenseNotification) SetReplies(v []LicenseNotificationReply)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *LicenseNotification) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


