# CloudNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**CidrBlock** | **string** |  | 
**State** | [**CloudNetworkStates**](CloudNetworkStates.md) |  | 
**IsDefault** | **bool** |  | 

## Methods

### NewCloudNetwork

`func NewCloudNetwork(resourceId string, cidrBlock string, state CloudNetworkStates, isDefault bool, ) *CloudNetwork`

NewCloudNetwork instantiates a new CloudNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNetworkWithDefaults

`func NewCloudNetworkWithDefaults() *CloudNetwork`

NewCloudNetworkWithDefaults instantiates a new CloudNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *CloudNetwork) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CloudNetwork) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CloudNetwork) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetName

`func (o *CloudNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCidrBlock

`func (o *CloudNetwork) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *CloudNetwork) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *CloudNetwork) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetState

`func (o *CloudNetwork) GetState() CloudNetworkStates`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudNetwork) GetStateOk() (*CloudNetworkStates, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudNetwork) SetState(v CloudNetworkStates)`

SetState sets State field to given value.


### GetIsDefault

`func (o *CloudNetwork) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CloudNetwork) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CloudNetwork) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


