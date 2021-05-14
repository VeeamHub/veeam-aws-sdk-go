# ResolveNotificationsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationsToResolve** | [**[]LicenseNotificationTypes**](LicenseNotificationTypes.md) |  | 
**SelectedReplies** | Pointer to [**[]LicenseNotificationReply**](LicenseNotificationReply.md) |  | [optional] 

## Methods

### NewResolveNotificationsSpec

`func NewResolveNotificationsSpec(notificationsToResolve []LicenseNotificationTypes, ) *ResolveNotificationsSpec`

NewResolveNotificationsSpec instantiates a new ResolveNotificationsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResolveNotificationsSpecWithDefaults

`func NewResolveNotificationsSpecWithDefaults() *ResolveNotificationsSpec`

NewResolveNotificationsSpecWithDefaults instantiates a new ResolveNotificationsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationsToResolve

`func (o *ResolveNotificationsSpec) GetNotificationsToResolve() []LicenseNotificationTypes`

GetNotificationsToResolve returns the NotificationsToResolve field if non-nil, zero value otherwise.

### GetNotificationsToResolveOk

`func (o *ResolveNotificationsSpec) GetNotificationsToResolveOk() (*[]LicenseNotificationTypes, bool)`

GetNotificationsToResolveOk returns a tuple with the NotificationsToResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsToResolve

`func (o *ResolveNotificationsSpec) SetNotificationsToResolve(v []LicenseNotificationTypes)`

SetNotificationsToResolve sets NotificationsToResolve field to given value.


### GetSelectedReplies

`func (o *ResolveNotificationsSpec) GetSelectedReplies() []LicenseNotificationReply`

GetSelectedReplies returns the SelectedReplies field if non-nil, zero value otherwise.

### GetSelectedRepliesOk

`func (o *ResolveNotificationsSpec) GetSelectedRepliesOk() (*[]LicenseNotificationReply, bool)`

GetSelectedRepliesOk returns a tuple with the SelectedReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedReplies

`func (o *ResolveNotificationsSpec) SetSelectedReplies(v []LicenseNotificationReply)`

SetSelectedReplies sets SelectedReplies field to given value.

### HasSelectedReplies

`func (o *ResolveNotificationsSpec) HasSelectedReplies() bool`

HasSelectedReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


