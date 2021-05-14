# RegionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | **string** |  | 
**RegionName** | Pointer to **string** |  | [optional] 
**ParentReiognId** | Pointer to **string** |  | [optional] 
**ParentRegionName** | Pointer to **string** |  | [optional] 
**CloudNetworkId** | **string** |  | 
**CloudNetworkName** | Pointer to **string** |  | [optional] 
**CloudSubnetworkId** | **string** |  | 
**CloudSubnetworkName** | Pointer to **string** |  | [optional] 
**CloudSecurityGroupId** | **string** |  | 
**CloudSecurityGroupName** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewRegionSettings

`func NewRegionSettings(regionId string, cloudNetworkId string, cloudSubnetworkId string, cloudSecurityGroupId string, ) *RegionSettings`

NewRegionSettings instantiates a new RegionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionSettingsWithDefaults

`func NewRegionSettingsWithDefaults() *RegionSettings`

NewRegionSettingsWithDefaults instantiates a new RegionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *RegionSettings) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *RegionSettings) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *RegionSettings) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetRegionName

`func (o *RegionSettings) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *RegionSettings) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *RegionSettings) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *RegionSettings) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetParentReiognId

`func (o *RegionSettings) GetParentReiognId() string`

GetParentReiognId returns the ParentReiognId field if non-nil, zero value otherwise.

### GetParentReiognIdOk

`func (o *RegionSettings) GetParentReiognIdOk() (*string, bool)`

GetParentReiognIdOk returns a tuple with the ParentReiognId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReiognId

`func (o *RegionSettings) SetParentReiognId(v string)`

SetParentReiognId sets ParentReiognId field to given value.

### HasParentReiognId

`func (o *RegionSettings) HasParentReiognId() bool`

HasParentReiognId returns a boolean if a field has been set.

### GetParentRegionName

`func (o *RegionSettings) GetParentRegionName() string`

GetParentRegionName returns the ParentRegionName field if non-nil, zero value otherwise.

### GetParentRegionNameOk

`func (o *RegionSettings) GetParentRegionNameOk() (*string, bool)`

GetParentRegionNameOk returns a tuple with the ParentRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRegionName

`func (o *RegionSettings) SetParentRegionName(v string)`

SetParentRegionName sets ParentRegionName field to given value.

### HasParentRegionName

`func (o *RegionSettings) HasParentRegionName() bool`

HasParentRegionName returns a boolean if a field has been set.

### GetCloudNetworkId

`func (o *RegionSettings) GetCloudNetworkId() string`

GetCloudNetworkId returns the CloudNetworkId field if non-nil, zero value otherwise.

### GetCloudNetworkIdOk

`func (o *RegionSettings) GetCloudNetworkIdOk() (*string, bool)`

GetCloudNetworkIdOk returns a tuple with the CloudNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkId

`func (o *RegionSettings) SetCloudNetworkId(v string)`

SetCloudNetworkId sets CloudNetworkId field to given value.


### GetCloudNetworkName

`func (o *RegionSettings) GetCloudNetworkName() string`

GetCloudNetworkName returns the CloudNetworkName field if non-nil, zero value otherwise.

### GetCloudNetworkNameOk

`func (o *RegionSettings) GetCloudNetworkNameOk() (*string, bool)`

GetCloudNetworkNameOk returns a tuple with the CloudNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkName

`func (o *RegionSettings) SetCloudNetworkName(v string)`

SetCloudNetworkName sets CloudNetworkName field to given value.

### HasCloudNetworkName

`func (o *RegionSettings) HasCloudNetworkName() bool`

HasCloudNetworkName returns a boolean if a field has been set.

### GetCloudSubnetworkId

`func (o *RegionSettings) GetCloudSubnetworkId() string`

GetCloudSubnetworkId returns the CloudSubnetworkId field if non-nil, zero value otherwise.

### GetCloudSubnetworkIdOk

`func (o *RegionSettings) GetCloudSubnetworkIdOk() (*string, bool)`

GetCloudSubnetworkIdOk returns a tuple with the CloudSubnetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSubnetworkId

`func (o *RegionSettings) SetCloudSubnetworkId(v string)`

SetCloudSubnetworkId sets CloudSubnetworkId field to given value.


### GetCloudSubnetworkName

`func (o *RegionSettings) GetCloudSubnetworkName() string`

GetCloudSubnetworkName returns the CloudSubnetworkName field if non-nil, zero value otherwise.

### GetCloudSubnetworkNameOk

`func (o *RegionSettings) GetCloudSubnetworkNameOk() (*string, bool)`

GetCloudSubnetworkNameOk returns a tuple with the CloudSubnetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSubnetworkName

`func (o *RegionSettings) SetCloudSubnetworkName(v string)`

SetCloudSubnetworkName sets CloudSubnetworkName field to given value.

### HasCloudSubnetworkName

`func (o *RegionSettings) HasCloudSubnetworkName() bool`

HasCloudSubnetworkName returns a boolean if a field has been set.

### GetCloudSecurityGroupId

`func (o *RegionSettings) GetCloudSecurityGroupId() string`

GetCloudSecurityGroupId returns the CloudSecurityGroupId field if non-nil, zero value otherwise.

### GetCloudSecurityGroupIdOk

`func (o *RegionSettings) GetCloudSecurityGroupIdOk() (*string, bool)`

GetCloudSecurityGroupIdOk returns a tuple with the CloudSecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSecurityGroupId

`func (o *RegionSettings) SetCloudSecurityGroupId(v string)`

SetCloudSecurityGroupId sets CloudSecurityGroupId field to given value.


### GetCloudSecurityGroupName

`func (o *RegionSettings) GetCloudSecurityGroupName() string`

GetCloudSecurityGroupName returns the CloudSecurityGroupName field if non-nil, zero value otherwise.

### GetCloudSecurityGroupNameOk

`func (o *RegionSettings) GetCloudSecurityGroupNameOk() (*string, bool)`

GetCloudSecurityGroupNameOk returns a tuple with the CloudSecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSecurityGroupName

`func (o *RegionSettings) SetCloudSecurityGroupName(v string)`

SetCloudSecurityGroupName sets CloudSecurityGroupName field to given value.

### HasCloudSecurityGroupName

`func (o *RegionSettings) HasCloudSecurityGroupName() bool`

HasCloudSecurityGroupName returns a boolean if a field has been set.

### GetLinks

`func (o *RegionSettings) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RegionSettings) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RegionSettings) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RegionSettings) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


