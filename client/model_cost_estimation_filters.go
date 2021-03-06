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

// CostEstimationFilters struct for CostEstimationFilters
type CostEstimationFilters struct {
	VirtualMachineNameFilter *string `json:"VirtualMachineNameFilter,omitempty"`
	Offset *int32 `json:"Offset,omitempty"`
	Limit *int32 `json:"Limit,omitempty"`
	Sort *[]CostEstimationSortColumns `json:"Sort,omitempty"`
}

// NewCostEstimationFilters instantiates a new CostEstimationFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCostEstimationFilters() *CostEstimationFilters {
	this := CostEstimationFilters{}
	return &this
}

// NewCostEstimationFiltersWithDefaults instantiates a new CostEstimationFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostEstimationFiltersWithDefaults() *CostEstimationFilters {
	this := CostEstimationFilters{}
	return &this
}

// GetVirtualMachineNameFilter returns the VirtualMachineNameFilter field value if set, zero value otherwise.
func (o *CostEstimationFilters) GetVirtualMachineNameFilter() string {
	if o == nil || o.VirtualMachineNameFilter == nil {
		var ret string
		return ret
	}
	return *o.VirtualMachineNameFilter
}

// GetVirtualMachineNameFilterOk returns a tuple with the VirtualMachineNameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimationFilters) GetVirtualMachineNameFilterOk() (*string, bool) {
	if o == nil || o.VirtualMachineNameFilter == nil {
		return nil, false
	}
	return o.VirtualMachineNameFilter, true
}

// HasVirtualMachineNameFilter returns a boolean if a field has been set.
func (o *CostEstimationFilters) HasVirtualMachineNameFilter() bool {
	if o != nil && o.VirtualMachineNameFilter != nil {
		return true
	}

	return false
}

// SetVirtualMachineNameFilter gets a reference to the given string and assigns it to the VirtualMachineNameFilter field.
func (o *CostEstimationFilters) SetVirtualMachineNameFilter(v string) {
	o.VirtualMachineNameFilter = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *CostEstimationFilters) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimationFilters) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *CostEstimationFilters) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *CostEstimationFilters) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CostEstimationFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimationFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CostEstimationFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CostEstimationFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *CostEstimationFilters) GetSort() []CostEstimationSortColumns {
	if o == nil || o.Sort == nil {
		var ret []CostEstimationSortColumns
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimationFilters) GetSortOk() (*[]CostEstimationSortColumns, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *CostEstimationFilters) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given []CostEstimationSortColumns and assigns it to the Sort field.
func (o *CostEstimationFilters) SetSort(v []CostEstimationSortColumns) {
	o.Sort = &v
}

func (o CostEstimationFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtualMachineNameFilter != nil {
		toSerialize["VirtualMachineNameFilter"] = o.VirtualMachineNameFilter
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

type NullableCostEstimationFilters struct {
	value *CostEstimationFilters
	isSet bool
}

func (v NullableCostEstimationFilters) Get() *CostEstimationFilters {
	return v.value
}

func (v *NullableCostEstimationFilters) Set(val *CostEstimationFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableCostEstimationFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableCostEstimationFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCostEstimationFilters(val *CostEstimationFilters) *NullableCostEstimationFilters {
	return &NullableCostEstimationFilters{value: val, isSet: true}
}

func (v NullableCostEstimationFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCostEstimationFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


