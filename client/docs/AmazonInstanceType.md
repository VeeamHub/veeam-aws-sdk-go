# AmazonInstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CpuCount** | **int32** |  | 
**Ram** | **int32** |  | 
**IsArm64** | Pointer to **bool** |  | [optional] 
**VirtualizationType** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAmazonInstanceType

`func NewAmazonInstanceType(name string, cpuCount int32, ram int32, ) *AmazonInstanceType`

NewAmazonInstanceType instantiates a new AmazonInstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonInstanceTypeWithDefaults

`func NewAmazonInstanceTypeWithDefaults() *AmazonInstanceType`

NewAmazonInstanceTypeWithDefaults instantiates a new AmazonInstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AmazonInstanceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmazonInstanceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmazonInstanceType) SetName(v string)`

SetName sets Name field to given value.


### GetCpuCount

`func (o *AmazonInstanceType) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *AmazonInstanceType) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *AmazonInstanceType) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.


### GetRam

`func (o *AmazonInstanceType) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *AmazonInstanceType) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *AmazonInstanceType) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetIsArm64

`func (o *AmazonInstanceType) GetIsArm64() bool`

GetIsArm64 returns the IsArm64 field if non-nil, zero value otherwise.

### GetIsArm64Ok

`func (o *AmazonInstanceType) GetIsArm64Ok() (*bool, bool)`

GetIsArm64Ok returns a tuple with the IsArm64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArm64

`func (o *AmazonInstanceType) SetIsArm64(v bool)`

SetIsArm64 sets IsArm64 field to given value.

### HasIsArm64

`func (o *AmazonInstanceType) HasIsArm64() bool`

HasIsArm64 returns a boolean if a field has been set.

### GetVirtualizationType

`func (o *AmazonInstanceType) GetVirtualizationType() []string`

GetVirtualizationType returns the VirtualizationType field if non-nil, zero value otherwise.

### GetVirtualizationTypeOk

`func (o *AmazonInstanceType) GetVirtualizationTypeOk() (*[]string, bool)`

GetVirtualizationTypeOk returns a tuple with the VirtualizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationType

`func (o *AmazonInstanceType) SetVirtualizationType(v []string)`

SetVirtualizationType sets VirtualizationType field to given value.

### HasVirtualizationType

`func (o *AmazonInstanceType) HasVirtualizationType() bool`

HasVirtualizationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


