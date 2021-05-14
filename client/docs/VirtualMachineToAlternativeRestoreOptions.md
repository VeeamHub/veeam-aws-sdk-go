# VirtualMachineToAlternativeRestoreOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**RegionId** | **string** |  | 
**VmType** | **string** |  | 
**SubnetId** | **string** |  | 
**NetworkSecurityGroupId** | **string** |  | 
**PreserveEncryptionForVolumes** | **bool** |  | 
**EncryptionKeyId** | Pointer to **string** |  | [optional] 

## Methods

### NewVirtualMachineToAlternativeRestoreOptions

`func NewVirtualMachineToAlternativeRestoreOptions(name string, regionId string, vmType string, subnetId string, networkSecurityGroupId string, preserveEncryptionForVolumes bool, ) *VirtualMachineToAlternativeRestoreOptions`

NewVirtualMachineToAlternativeRestoreOptions instantiates a new VirtualMachineToAlternativeRestoreOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineToAlternativeRestoreOptionsWithDefaults

`func NewVirtualMachineToAlternativeRestoreOptionsWithDefaults() *VirtualMachineToAlternativeRestoreOptions`

NewVirtualMachineToAlternativeRestoreOptionsWithDefaults instantiates a new VirtualMachineToAlternativeRestoreOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualMachineToAlternativeRestoreOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualMachineToAlternativeRestoreOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualMachineToAlternativeRestoreOptions) SetName(v string)`

SetName sets Name field to given value.


### GetRegionId

`func (o *VirtualMachineToAlternativeRestoreOptions) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *VirtualMachineToAlternativeRestoreOptions) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *VirtualMachineToAlternativeRestoreOptions) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetVmType

`func (o *VirtualMachineToAlternativeRestoreOptions) GetVmType() string`

GetVmType returns the VmType field if non-nil, zero value otherwise.

### GetVmTypeOk

`func (o *VirtualMachineToAlternativeRestoreOptions) GetVmTypeOk() (*string, bool)`

GetVmTypeOk returns a tuple with the VmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmType

`func (o *VirtualMachineToAlternativeRestoreOptions) SetVmType(v string)`

SetVmType sets VmType field to given value.


### GetSubnetId

`func (o *VirtualMachineToAlternativeRestoreOptions) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *VirtualMachineToAlternativeRestoreOptions) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *VirtualMachineToAlternativeRestoreOptions) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetNetworkSecurityGroupId

`func (o *VirtualMachineToAlternativeRestoreOptions) GetNetworkSecurityGroupId() string`

GetNetworkSecurityGroupId returns the NetworkSecurityGroupId field if non-nil, zero value otherwise.

### GetNetworkSecurityGroupIdOk

`func (o *VirtualMachineToAlternativeRestoreOptions) GetNetworkSecurityGroupIdOk() (*string, bool)`

GetNetworkSecurityGroupIdOk returns a tuple with the NetworkSecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSecurityGroupId

`func (o *VirtualMachineToAlternativeRestoreOptions) SetNetworkSecurityGroupId(v string)`

SetNetworkSecurityGroupId sets NetworkSecurityGroupId field to given value.


### GetPreserveEncryptionForVolumes

`func (o *VirtualMachineToAlternativeRestoreOptions) GetPreserveEncryptionForVolumes() bool`

GetPreserveEncryptionForVolumes returns the PreserveEncryptionForVolumes field if non-nil, zero value otherwise.

### GetPreserveEncryptionForVolumesOk

`func (o *VirtualMachineToAlternativeRestoreOptions) GetPreserveEncryptionForVolumesOk() (*bool, bool)`

GetPreserveEncryptionForVolumesOk returns a tuple with the PreserveEncryptionForVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveEncryptionForVolumes

`func (o *VirtualMachineToAlternativeRestoreOptions) SetPreserveEncryptionForVolumes(v bool)`

SetPreserveEncryptionForVolumes sets PreserveEncryptionForVolumes field to given value.


### GetEncryptionKeyId

`func (o *VirtualMachineToAlternativeRestoreOptions) GetEncryptionKeyId() string`

GetEncryptionKeyId returns the EncryptionKeyId field if non-nil, zero value otherwise.

### GetEncryptionKeyIdOk

`func (o *VirtualMachineToAlternativeRestoreOptions) GetEncryptionKeyIdOk() (*string, bool)`

GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyId

`func (o *VirtualMachineToAlternativeRestoreOptions) SetEncryptionKeyId(v string)`

SetEncryptionKeyId sets EncryptionKeyId field to given value.

### HasEncryptionKeyId

`func (o *VirtualMachineToAlternativeRestoreOptions) HasEncryptionKeyId() bool`

HasEncryptionKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


