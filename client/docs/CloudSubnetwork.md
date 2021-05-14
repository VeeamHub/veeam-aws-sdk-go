# CloudSubnetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** |  | 
**Name** | **string** |  | 
**CloudNetworkResourceId** | **string** |  | 
**CidrBlock** | **string** |  | 
**AvailabilityZone** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudSubnetwork

`func NewCloudSubnetwork(resourceId string, name string, cloudNetworkResourceId string, cidrBlock string, ) *CloudSubnetwork`

NewCloudSubnetwork instantiates a new CloudSubnetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSubnetworkWithDefaults

`func NewCloudSubnetworkWithDefaults() *CloudSubnetwork`

NewCloudSubnetworkWithDefaults instantiates a new CloudSubnetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *CloudSubnetwork) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CloudSubnetwork) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CloudSubnetwork) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetName

`func (o *CloudSubnetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSubnetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSubnetwork) SetName(v string)`

SetName sets Name field to given value.


### GetCloudNetworkResourceId

`func (o *CloudSubnetwork) GetCloudNetworkResourceId() string`

GetCloudNetworkResourceId returns the CloudNetworkResourceId field if non-nil, zero value otherwise.

### GetCloudNetworkResourceIdOk

`func (o *CloudSubnetwork) GetCloudNetworkResourceIdOk() (*string, bool)`

GetCloudNetworkResourceIdOk returns a tuple with the CloudNetworkResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkResourceId

`func (o *CloudSubnetwork) SetCloudNetworkResourceId(v string)`

SetCloudNetworkResourceId sets CloudNetworkResourceId field to given value.


### GetCidrBlock

`func (o *CloudSubnetwork) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *CloudSubnetwork) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *CloudSubnetwork) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetAvailabilityZone

`func (o *CloudSubnetwork) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CloudSubnetwork) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CloudSubnetwork) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CloudSubnetwork) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


