# CostEstimationFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualMachineNameFilter** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to [**[]CostEstimationSortColumns**](CostEstimationSortColumns.md) |  | [optional] 

## Methods

### NewCostEstimationFilters

`func NewCostEstimationFilters() *CostEstimationFilters`

NewCostEstimationFilters instantiates a new CostEstimationFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostEstimationFiltersWithDefaults

`func NewCostEstimationFiltersWithDefaults() *CostEstimationFilters`

NewCostEstimationFiltersWithDefaults instantiates a new CostEstimationFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualMachineNameFilter

`func (o *CostEstimationFilters) GetVirtualMachineNameFilter() string`

GetVirtualMachineNameFilter returns the VirtualMachineNameFilter field if non-nil, zero value otherwise.

### GetVirtualMachineNameFilterOk

`func (o *CostEstimationFilters) GetVirtualMachineNameFilterOk() (*string, bool)`

GetVirtualMachineNameFilterOk returns a tuple with the VirtualMachineNameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineNameFilter

`func (o *CostEstimationFilters) SetVirtualMachineNameFilter(v string)`

SetVirtualMachineNameFilter sets VirtualMachineNameFilter field to given value.

### HasVirtualMachineNameFilter

`func (o *CostEstimationFilters) HasVirtualMachineNameFilter() bool`

HasVirtualMachineNameFilter returns a boolean if a field has been set.

### GetOffset

`func (o *CostEstimationFilters) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CostEstimationFilters) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CostEstimationFilters) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CostEstimationFilters) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *CostEstimationFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CostEstimationFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CostEstimationFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CostEstimationFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *CostEstimationFilters) GetSort() []CostEstimationSortColumns`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *CostEstimationFilters) GetSortOk() (*[]CostEstimationSortColumns, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *CostEstimationFilters) SetSort(v []CostEstimationSortColumns)`

SetSort sets Sort field to given value.

### HasSort

`func (o *CostEstimationFilters) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


