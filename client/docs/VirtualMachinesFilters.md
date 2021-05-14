# VirtualMachinesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchPattern** | Pointer to **string** |  | [optional] 
**ProtectionStatus** | Pointer to [**[]VirtualMachineProtectionStatuses**](VirtualMachineProtectionStatuses.md) |  | [optional] 
**BackupDestination** | Pointer to [**[]BackupDestinations**](BackupDestinations.md) |  | [optional] 
**BackupState** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to [**[]VirtualMachinesSortColumns**](VirtualMachinesSortColumns.md) |  | [optional] 

## Methods

### NewVirtualMachinesFilters

`func NewVirtualMachinesFilters() *VirtualMachinesFilters`

NewVirtualMachinesFilters instantiates a new VirtualMachinesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachinesFiltersWithDefaults

`func NewVirtualMachinesFiltersWithDefaults() *VirtualMachinesFilters`

NewVirtualMachinesFiltersWithDefaults instantiates a new VirtualMachinesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchPattern

`func (o *VirtualMachinesFilters) GetSearchPattern() string`

GetSearchPattern returns the SearchPattern field if non-nil, zero value otherwise.

### GetSearchPatternOk

`func (o *VirtualMachinesFilters) GetSearchPatternOk() (*string, bool)`

GetSearchPatternOk returns a tuple with the SearchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPattern

`func (o *VirtualMachinesFilters) SetSearchPattern(v string)`

SetSearchPattern sets SearchPattern field to given value.

### HasSearchPattern

`func (o *VirtualMachinesFilters) HasSearchPattern() bool`

HasSearchPattern returns a boolean if a field has been set.

### GetProtectionStatus

`func (o *VirtualMachinesFilters) GetProtectionStatus() []VirtualMachineProtectionStatuses`

GetProtectionStatus returns the ProtectionStatus field if non-nil, zero value otherwise.

### GetProtectionStatusOk

`func (o *VirtualMachinesFilters) GetProtectionStatusOk() (*[]VirtualMachineProtectionStatuses, bool)`

GetProtectionStatusOk returns a tuple with the ProtectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionStatus

`func (o *VirtualMachinesFilters) SetProtectionStatus(v []VirtualMachineProtectionStatuses)`

SetProtectionStatus sets ProtectionStatus field to given value.

### HasProtectionStatus

`func (o *VirtualMachinesFilters) HasProtectionStatus() bool`

HasProtectionStatus returns a boolean if a field has been set.

### GetBackupDestination

`func (o *VirtualMachinesFilters) GetBackupDestination() []BackupDestinations`

GetBackupDestination returns the BackupDestination field if non-nil, zero value otherwise.

### GetBackupDestinationOk

`func (o *VirtualMachinesFilters) GetBackupDestinationOk() (*[]BackupDestinations, bool)`

GetBackupDestinationOk returns a tuple with the BackupDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDestination

`func (o *VirtualMachinesFilters) SetBackupDestination(v []BackupDestinations)`

SetBackupDestination sets BackupDestination field to given value.

### HasBackupDestination

`func (o *VirtualMachinesFilters) HasBackupDestination() bool`

HasBackupDestination returns a boolean if a field has been set.

### GetBackupState

`func (o *VirtualMachinesFilters) GetBackupState() string`

GetBackupState returns the BackupState field if non-nil, zero value otherwise.

### GetBackupStateOk

`func (o *VirtualMachinesFilters) GetBackupStateOk() (*string, bool)`

GetBackupStateOk returns a tuple with the BackupState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupState

`func (o *VirtualMachinesFilters) SetBackupState(v string)`

SetBackupState sets BackupState field to given value.

### HasBackupState

`func (o *VirtualMachinesFilters) HasBackupState() bool`

HasBackupState returns a boolean if a field has been set.

### GetOffset

`func (o *VirtualMachinesFilters) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *VirtualMachinesFilters) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *VirtualMachinesFilters) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *VirtualMachinesFilters) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *VirtualMachinesFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VirtualMachinesFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VirtualMachinesFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *VirtualMachinesFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *VirtualMachinesFilters) GetSort() []VirtualMachinesSortColumns`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *VirtualMachinesFilters) GetSortOk() (*[]VirtualMachinesSortColumns, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *VirtualMachinesFilters) SetSort(v []VirtualMachinesSortColumns)`

SetSort sets Sort field to given value.

### HasSort

`func (o *VirtualMachinesFilters) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


