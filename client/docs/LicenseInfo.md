# LicenseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseType** | Pointer to **string** |  | [optional] 
**InstancesUses** | Pointer to **float64** |  | [optional] 
**Instances** | Pointer to **int32** |  | [optional] 
**LicensedTo** | Pointer to **string** |  | [optional] 
**LicenseExpires** | Pointer to **time.Time** |  | [optional] 
**LicenseId** | Pointer to **string** |  | [optional] 
**SupportId** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewLicenseInfo

`func NewLicenseInfo() *LicenseInfo`

NewLicenseInfo instantiates a new LicenseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseInfoWithDefaults

`func NewLicenseInfoWithDefaults() *LicenseInfo`

NewLicenseInfoWithDefaults instantiates a new LicenseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseType

`func (o *LicenseInfo) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *LicenseInfo) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *LicenseInfo) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *LicenseInfo) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetInstancesUses

`func (o *LicenseInfo) GetInstancesUses() float64`

GetInstancesUses returns the InstancesUses field if non-nil, zero value otherwise.

### GetInstancesUsesOk

`func (o *LicenseInfo) GetInstancesUsesOk() (*float64, bool)`

GetInstancesUsesOk returns a tuple with the InstancesUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesUses

`func (o *LicenseInfo) SetInstancesUses(v float64)`

SetInstancesUses sets InstancesUses field to given value.

### HasInstancesUses

`func (o *LicenseInfo) HasInstancesUses() bool`

HasInstancesUses returns a boolean if a field has been set.

### GetInstances

`func (o *LicenseInfo) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *LicenseInfo) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *LicenseInfo) SetInstances(v int32)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *LicenseInfo) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetLicensedTo

`func (o *LicenseInfo) GetLicensedTo() string`

GetLicensedTo returns the LicensedTo field if non-nil, zero value otherwise.

### GetLicensedToOk

`func (o *LicenseInfo) GetLicensedToOk() (*string, bool)`

GetLicensedToOk returns a tuple with the LicensedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedTo

`func (o *LicenseInfo) SetLicensedTo(v string)`

SetLicensedTo sets LicensedTo field to given value.

### HasLicensedTo

`func (o *LicenseInfo) HasLicensedTo() bool`

HasLicensedTo returns a boolean if a field has been set.

### GetLicenseExpires

`func (o *LicenseInfo) GetLicenseExpires() time.Time`

GetLicenseExpires returns the LicenseExpires field if non-nil, zero value otherwise.

### GetLicenseExpiresOk

`func (o *LicenseInfo) GetLicenseExpiresOk() (*time.Time, bool)`

GetLicenseExpiresOk returns a tuple with the LicenseExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpires

`func (o *LicenseInfo) SetLicenseExpires(v time.Time)`

SetLicenseExpires sets LicenseExpires field to given value.

### HasLicenseExpires

`func (o *LicenseInfo) HasLicenseExpires() bool`

HasLicenseExpires returns a boolean if a field has been set.

### GetLicenseId

`func (o *LicenseInfo) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *LicenseInfo) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *LicenseInfo) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *LicenseInfo) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetSupportId

`func (o *LicenseInfo) GetSupportId() string`

GetSupportId returns the SupportId field if non-nil, zero value otherwise.

### GetSupportIdOk

`func (o *LicenseInfo) GetSupportIdOk() (*string, bool)`

GetSupportIdOk returns a tuple with the SupportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportId

`func (o *LicenseInfo) SetSupportId(v string)`

SetSupportId sets SupportId field to given value.

### HasSupportId

`func (o *LicenseInfo) HasSupportId() bool`

HasSupportId returns a boolean if a field has been set.

### GetLinks

`func (o *LicenseInfo) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LicenseInfo) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LicenseInfo) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LicenseInfo) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


