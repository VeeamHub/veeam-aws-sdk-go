/*
 * Veeam Backup for AWS public API 1.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0-rev0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// VmRestorePointsFilters struct for VmRestorePointsFilters
type VmRestorePointsFilters struct {
	VirtualMachineId *string `json:"VirtualMachineId,omitempty"`
	OnlyLatest *bool `json:"OnlyLatest,omitempty"`
	Offset *int32 `json:"Offset,omitempty"`
	Limit *int32 `json:"Limit,omitempty"`
	Sort *[]VmRestorePointsSortColumns `json:"Sort,omitempty"`
}

// NewVmRestorePointsFilters instantiates a new VmRestorePointsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmRestorePointsFilters() *VmRestorePointsFilters {
	this := VmRestorePointsFilters{}
	return &this
}

// NewVmRestorePointsFiltersWithDefaults instantiates a new VmRestorePointsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmRestorePointsFiltersWithDefaults() *VmRestorePointsFilters {
	this := VmRestorePointsFilters{}
	return &this
}

// GetVirtualMachineId returns the VirtualMachineId field value if set, zero value otherwise.
func (o *VmRestorePointsFilters) GetVirtualMachineId() string {
	if o == nil || o.VirtualMachineId == nil {
		var ret string
		return ret
	}
	return *o.VirtualMachineId
}

// GetVirtualMachineIdOk returns a tuple with the VirtualMachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmRestorePointsFilters) GetVirtualMachineIdOk() (*string, bool) {
	if o == nil || o.VirtualMachineId == nil {
		return nil, false
	}
	return o.VirtualMachineId, true
}

// HasVirtualMachineId returns a boolean if a field has been set.
func (o *VmRestorePointsFilters) HasVirtualMachineId() bool {
	if o != nil && o.VirtualMachineId != nil {
		return true
	}

	return false
}

// SetVirtualMachineId gets a reference to the given string and assigns it to the VirtualMachineId field.
func (o *VmRestorePointsFilters) SetVirtualMachineId(v string) {
	o.VirtualMachineId = &v
}

// GetOnlyLatest returns the OnlyLatest field value if set, zero value otherwise.
func (o *VmRestorePointsFilters) GetOnlyLatest() bool {
	if o == nil || o.OnlyLatest == nil {
		var ret bool
		return ret
	}
	return *o.OnlyLatest
}

// GetOnlyLatestOk returns a tuple with the OnlyLatest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmRestorePointsFilters) GetOnlyLatestOk() (*bool, bool) {
	if o == nil || o.OnlyLatest == nil {
		return nil, false
	}
	return o.OnlyLatest, true
}

// HasOnlyLatest returns a boolean if a field has been set.
func (o *VmRestorePointsFilters) HasOnlyLatest() bool {
	if o != nil && o.OnlyLatest != nil {
		return true
	}

	return false
}

// SetOnlyLatest gets a reference to the given bool and assigns it to the OnlyLatest field.
func (o *VmRestorePointsFilters) SetOnlyLatest(v bool) {
	o.OnlyLatest = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *VmRestorePointsFilters) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmRestorePointsFilters) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *VmRestorePointsFilters) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *VmRestorePointsFilters) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *VmRestorePointsFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmRestorePointsFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *VmRestorePointsFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *VmRestorePointsFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *VmRestorePointsFilters) GetSort() []VmRestorePointsSortColumns {
	if o == nil || o.Sort == nil {
		var ret []VmRestorePointsSortColumns
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmRestorePointsFilters) GetSortOk() (*[]VmRestorePointsSortColumns, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *VmRestorePointsFilters) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given []VmRestorePointsSortColumns and assigns it to the Sort field.
func (o *VmRestorePointsFilters) SetSort(v []VmRestorePointsSortColumns) {
	o.Sort = &v
}

func (o VmRestorePointsFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtualMachineId != nil {
		toSerialize["VirtualMachineId"] = o.VirtualMachineId
	}
	if o.OnlyLatest != nil {
		toSerialize["OnlyLatest"] = o.OnlyLatest
	}
	if o.Offset != nil {
		toSerialize["Offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["Limit"] = o.Limit
	}
	if o.Sort != nil {
		toSerialize["Sort"] = o.Sort
	}
	return json.Marshal(toSerialize)
}

type NullableVmRestorePointsFilters struct {
	value *VmRestorePointsFilters
	isSet bool
}

func (v NullableVmRestorePointsFilters) Get() *VmRestorePointsFilters {
	return v.value
}

func (v *NullableVmRestorePointsFilters) Set(val *VmRestorePointsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableVmRestorePointsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableVmRestorePointsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmRestorePointsFilters(val *VmRestorePointsFilters) *NullableVmRestorePointsFilters {
	return &NullableVmRestorePointsFilters{value: val, isSet: true}
}

func (v NullableVmRestorePointsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmRestorePointsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


