# LicenseNotificationReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplyId** | **string** |  | 
**NotificationType** | [**LicenseNotificationTypes**](LicenseNotificationTypes.md) |  | 
**ReplyText** | **string** |  | 

## Methods

### NewLicenseNotificationReply

`func NewLicenseNotificationReply(replyId string, notificationType LicenseNotificationTypes, replyText string, ) *LicenseNotificationReply`

NewLicenseNotificationReply instantiates a new LicenseNotificationReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseNotificationReplyWithDefaults

`func NewLicenseNotificationReplyWithDefaults() *LicenseNotificationReply`

NewLicenseNotificationReplyWithDefaults instantiates a new LicenseNotificationReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplyId

`func (o *LicenseNotificationReply) GetReplyId() string`

GetReplyId returns the ReplyId field if non-nil, zero value otherwise.

### GetReplyIdOk

`func (o *LicenseNotificationReply) GetReplyIdOk() (*string, bool)`

GetReplyIdOk returns a tuple with the ReplyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyId

`func (o *LicenseNotificationReply) SetReplyId(v string)`

SetReplyId sets ReplyId field to given value.


### GetNotificationType

`func (o *LicenseNotificationReply) GetNotificationType() LicenseNotificationTypes`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *LicenseNotificationReply) GetNotificationTypeOk() (*LicenseNotificationTypes, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *LicenseNotificationReply) SetNotificationType(v LicenseNotificationTypes)`

SetNotificationType sets NotificationType field to given value.


### GetReplyText

`func (o *LicenseNotificationReply) GetReplyText() string`

GetReplyText returns the ReplyText field if non-nil, zero value otherwise.

### GetReplyTextOk

`func (o *LicenseNotificationReply) GetReplyTextOk() (*string, bool)`

GetReplyTextOk returns a tuple with the ReplyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyText

`func (o *LicenseNotificationReply) SetReplyText(v string)`

SetReplyText sets ReplyText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


