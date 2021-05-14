# ServiceAccountSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAmazonAccountId** | **string** |  | 
**ServiceAmazonAccountName** | Pointer to **string** |  | [optional] 
**ServiceAmazonAccountRegionType** | Pointer to [**RegionTypes**](RegionTypes.md) |  | [optional] 
**ServiceRegion** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewServiceAccountSettings

`func NewServiceAccountSettings(serviceAmazonAccountId string, ) *ServiceAccountSettings`

NewServiceAccountSettings instantiates a new ServiceAccountSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountSettingsWithDefaults

`func NewServiceAccountSettingsWithDefaults() *ServiceAccountSettings`

NewServiceAccountSettingsWithDefaults instantiates a new ServiceAccountSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAmazonAccountId

`func (o *ServiceAccountSettings) GetServiceAmazonAccountId() string`

GetServiceAmazonAccountId returns the ServiceAmazonAccountId field if non-nil, zero value otherwise.

### GetServiceAmazonAccountIdOk

`func (o *ServiceAccountSettings) GetServiceAmazonAccountIdOk() (*string, bool)`

GetServiceAmazonAccountIdOk returns a tuple with the ServiceAmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmazonAccountId

`func (o *ServiceAccountSettings) SetServiceAmazonAccountId(v string)`

SetServiceAmazonAccountId sets ServiceAmazonAccountId field to given value.


### GetServiceAmazonAccountName

`func (o *ServiceAccountSettings) GetServiceAmazonAccountName() string`

GetServiceAmazonAccountName returns the ServiceAmazonAccountName field if non-nil, zero value otherwise.

### GetServiceAmazonAccountNameOk

`func (o *ServiceAccountSettings) GetServiceAmazonAccountNameOk() (*string, bool)`

GetServiceAmazonAccountNameOk returns a tuple with the ServiceAmazonAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmazonAccountName

`func (o *ServiceAccountSettings) SetServiceAmazonAccountName(v string)`

SetServiceAmazonAccountName sets ServiceAmazonAccountName field to given value.

### HasServiceAmazonAccountName

`func (o *ServiceAccountSettings) HasServiceAmazonAccountName() bool`

HasServiceAmazonAccountName returns a boolean if a field has been set.

### GetServiceAmazonAccountRegionType

`func (o *ServiceAccountSettings) GetServiceAmazonAccountRegionType() RegionTypes`

GetServiceAmazonAccountRegionType returns the ServiceAmazonAccountRegionType field if non-nil, zero value otherwise.

### GetServiceAmazonAccountRegionTypeOk

`func (o *ServiceAccountSettings) GetServiceAmazonAccountRegionTypeOk() (*RegionTypes, bool)`

GetServiceAmazonAccountRegionTypeOk returns a tuple with the ServiceAmazonAccountRegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmazonAccountRegionType

`func (o *ServiceAccountSettings) SetServiceAmazonAccountRegionType(v RegionTypes)`

SetServiceAmazonAccountRegionType sets ServiceAmazonAccountRegionType field to given value.

### HasServiceAmazonAccountRegionType

`func (o *ServiceAccountSettings) HasServiceAmazonAccountRegionType() bool`

HasServiceAmazonAccountRegionType returns a boolean if a field has been set.

### GetServiceRegion

`func (o *ServiceAccountSettings) GetServiceRegion() string`

GetServiceRegion returns the ServiceRegion field if non-nil, zero value otherwise.

### GetServiceRegionOk

`func (o *ServiceAccountSettings) GetServiceRegionOk() (*string, bool)`

GetServiceRegionOk returns a tuple with the ServiceRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRegion

`func (o *ServiceAccountSettings) SetServiceRegion(v string)`

SetServiceRegion sets ServiceRegion field to given value.

### HasServiceRegion

`func (o *ServiceAccountSettings) HasServiceRegion() bool`

HasServiceRegion returns a boolean if a field has been set.

### GetLinks

`func (o *ServiceAccountSettings) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceAccountSettings) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceAccountSettings) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceAccountSettings) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


