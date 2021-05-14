# PoliciesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchPattern** | Pointer to **string** |  | [optional] 
**VirtualMachineId** | Pointer to **string** |  | [optional] 
**LastPolicySessionStatus** | Pointer to [**[]SessionStatuses**](SessionStatuses.md) |  | [optional] 
**PolicyStatus** | Pointer to [**[]PolicyStatuses**](PolicyStatuses.md) |  | [optional] 
**Sort** | Pointer to [**[]PoliciesSortColumns**](PoliciesSortColumns.md) |  | [optional] 
**Usn** | Pointer to **int64** |  | [optional] 
**RepositoryId** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewPoliciesFilters

`func NewPoliciesFilters() *PoliciesFilters`

NewPoliciesFilters instantiates a new PoliciesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoliciesFiltersWithDefaults

`func NewPoliciesFiltersWithDefaults() *PoliciesFilters`

NewPoliciesFiltersWithDefaults instantiates a new PoliciesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchPattern

`func (o *PoliciesFilters) GetSearchPattern() string`

GetSearchPattern returns the SearchPattern field if non-nil, zero value otherwise.

### GetSearchPatternOk

`func (o *PoliciesFilters) GetSearchPatternOk() (*string, bool)`

GetSearchPatternOk returns a tuple with the SearchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPattern

`func (o *PoliciesFilters) SetSearchPattern(v string)`

SetSearchPattern sets SearchPattern field to given value.

### HasSearchPattern

`func (o *PoliciesFilters) HasSearchPattern() bool`

HasSearchPattern returns a boolean if a field has been set.

### GetVirtualMachineId

`func (o *PoliciesFilters) GetVirtualMachineId() string`

GetVirtualMachineId returns the VirtualMachineId field if non-nil, zero value otherwise.

### GetVirtualMachineIdOk

`func (o *PoliciesFilters) GetVirtualMachineIdOk() (*string, bool)`

GetVirtualMachineIdOk returns a tuple with the VirtualMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineId

`func (o *PoliciesFilters) SetVirtualMachineId(v string)`

SetVirtualMachineId sets VirtualMachineId field to given value.

### HasVirtualMachineId

`func (o *PoliciesFilters) HasVirtualMachineId() bool`

HasVirtualMachineId returns a boolean if a field has been set.

### GetLastPolicySessionStatus

`func (o *PoliciesFilters) GetLastPolicySessionStatus() []SessionStatuses`

GetLastPolicySessionStatus returns the LastPolicySessionStatus field if non-nil, zero value otherwise.

### GetLastPolicySessionStatusOk

`func (o *PoliciesFilters) GetLastPolicySessionStatusOk() (*[]SessionStatuses, bool)`

GetLastPolicySessionStatusOk returns a tuple with the LastPolicySessionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPolicySessionStatus

`func (o *PoliciesFilters) SetLastPolicySessionStatus(v []SessionStatuses)`

SetLastPolicySessionStatus sets LastPolicySessionStatus field to given value.

### HasLastPolicySessionStatus

`func (o *PoliciesFilters) HasLastPolicySessionStatus() bool`

HasLastPolicySessionStatus returns a boolean if a field has been set.

### GetPolicyStatus

`func (o *PoliciesFilters) GetPolicyStatus() []PolicyStatuses`

GetPolicyStatus returns the PolicyStatus field if non-nil, zero value otherwise.

### GetPolicyStatusOk

`func (o *PoliciesFilters) GetPolicyStatusOk() (*[]PolicyStatuses, bool)`

GetPolicyStatusOk returns a tuple with the PolicyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyStatus

`func (o *PoliciesFilters) SetPolicyStatus(v []PolicyStatuses)`

SetPolicyStatus sets PolicyStatus field to given value.

### HasPolicyStatus

`func (o *PoliciesFilters) HasPolicyStatus() bool`

HasPolicyStatus returns a boolean if a field has been set.

### GetSort

`func (o *PoliciesFilters) GetSort() []PoliciesSortColumns`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PoliciesFilters) GetSortOk() (*[]PoliciesSortColumns, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PoliciesFilters) SetSort(v []PoliciesSortColumns)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PoliciesFilters) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetUsn

`func (o *PoliciesFilters) GetUsn() int64`

GetUsn returns the Usn field if non-nil, zero value otherwise.

### GetUsnOk

`func (o *PoliciesFilters) GetUsnOk() (*int64, bool)`

GetUsnOk returns a tuple with the Usn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsn

`func (o *PoliciesFilters) SetUsn(v int64)`

SetUsn sets Usn field to given value.

### HasUsn

`func (o *PoliciesFilters) HasUsn() bool`

HasUsn returns a boolean if a field has been set.

### GetRepositoryId

`func (o *PoliciesFilters) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *PoliciesFilters) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *PoliciesFilters) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *PoliciesFilters) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetOffset

`func (o *PoliciesFilters) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PoliciesFilters) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PoliciesFilters) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PoliciesFilters) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *PoliciesFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PoliciesFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PoliciesFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PoliciesFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


