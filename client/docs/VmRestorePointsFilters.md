# VmRestorePointsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualMachineId** | Pointer to **string** |  | [optional] 
**OnlyLatest** | Pointer to **bool** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to [**[]VmRestorePointsSortColumns**](VmRestorePointsSortColumns.md) |  | [optional] 

## Methods

### NewVmRestorePointsFilters

`func NewVmRestorePointsFilters() *VmRestorePointsFilters`

NewVmRestorePointsFilters instantiates a new VmRestorePointsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRestorePointsFiltersWithDefaults

`func NewVmRestorePointsFiltersWithDefaults() *VmRestorePointsFilters`

NewVmRestorePointsFiltersWithDefaults instantiates a new VmRestorePointsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualMachineId

`func (o *VmRestorePointsFilters) GetVirtualMachineId() string`

GetVirtualMachineId returns the VirtualMachineId field if non-nil, zero value otherwise.

### GetVirtualMachineIdOk

`func (o *VmRestorePointsFilters) GetVirtualMachineIdOk() (*string, bool)`

GetVirtualMachineIdOk returns a tuple with the VirtualMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineId

`func (o *VmRestorePointsFilters) SetVirtualMachineId(v string)`

SetVirtualMachineId sets VirtualMachineId field to given value.

### HasVirtualMachineId

`func (o *VmRestorePointsFilters) HasVirtualMachineId() bool`

HasVirtualMachineId returns a boolean if a field has been set.

### GetOnlyLatest

`func (o *VmRestorePointsFilters) GetOnlyLatest() bool`

GetOnlyLatest returns the OnlyLatest field if non-nil, zero value otherwise.

### GetOnlyLatestOk

`func (o *VmRestorePointsFilters) GetOnlyLatestOk() (*bool, bool)`

GetOnlyLatestOk returns a tuple with the OnlyLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyLatest

`func (o *VmRestorePointsFilters) SetOnlyLatest(v bool)`

SetOnlyLatest sets OnlyLatest field to given value.

### HasOnlyLatest

`func (o *VmRestorePointsFilters) HasOnlyLatest() bool`

HasOnlyLatest returns a boolean if a field has been set.

### GetOffset

`func (o *VmRestorePointsFilters) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *VmRestorePointsFilters) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *VmRestorePointsFilters) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *VmRestorePointsFilters) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *VmRestorePointsFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VmRestorePointsFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VmRestorePointsFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *VmRestorePointsFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *VmRestorePointsFilters) GetSort() []VmRestorePointsSortColumns`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *VmRestorePointsFilters) GetSortOk() (*[]VmRestorePointsSortColumns, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *VmRestorePointsFilters) SetSort(v []VmRestorePointsSortColumns)`

SetSort sets Sort field to given value.

### HasSort

`func (o *VmRestorePointsFilters) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


