# LicenseNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]LicenseNotification**](LicenseNotification.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewLicenseNotifications

`func NewLicenseNotifications(results []LicenseNotification, ) *LicenseNotifications`

NewLicenseNotifications instantiates a new LicenseNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseNotificationsWithDefaults

`func NewLicenseNotificationsWithDefaults() *LicenseNotifications`

NewLicenseNotificationsWithDefaults instantiates a new LicenseNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *LicenseNotifications) GetResults() []LicenseNotification`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *LicenseNotifications) GetResultsOk() (*[]LicenseNotification, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *LicenseNotifications) SetResults(v []LicenseNotification)`

SetResults sets Results field to given value.


### GetLinks

`func (o *LicenseNotifications) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LicenseNotifications) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LicenseNotifications) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LicenseNotifications) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


