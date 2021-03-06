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

// RegionSettingsFilters struct for RegionSettingsFilters
type RegionSettingsFilters struct {
	Offset *int32 `json:"offset,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
}

// NewRegionSettingsFilters instantiates a new RegionSettingsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionSettingsFilters() *RegionSettingsFilters {
	this := RegionSettingsFilters{}
	return &this
}

// NewRegionSettingsFiltersWithDefaults instantiates a new RegionSettingsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionSettingsFiltersWithDefaults() *RegionSettingsFilters {
	this := RegionSettingsFilters{}
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *RegionSettingsFilters) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSettingsFilters) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *RegionSettingsFilters) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *RegionSettingsFilters) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *RegionSettingsFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSettingsFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *RegionSettingsFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *RegionSettingsFilters) SetLimit(v int32) {
	o.Limit = &v
}

func (o RegionSettingsFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullableRegionSettingsFilters struct {
	value *RegionSettingsFilters
	isSet bool
}

func (v NullableRegionSettingsFilters) Get() *RegionSettingsFilters {
	return v.value
}

func (v *NullableRegionSettingsFilters) Set(val *RegionSettingsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionSettingsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionSettingsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionSettingsFilters(val *RegionSettingsFilters) *NullableRegionSettingsFilters {
	return &NullableRegionSettingsFilters{value: val, isSet: true}
}

func (v NullableRegionSettingsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionSettingsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


