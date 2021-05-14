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

// TimeZones struct for TimeZones
type TimeZones struct {
	Items []TimeZone `json:"items"`
}

// NewTimeZones instantiates a new TimeZones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeZones(items []TimeZone) *TimeZones {
	this := TimeZones{}
	this.Items = items
	return &this
}

// NewTimeZonesWithDefaults instantiates a new TimeZones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeZonesWithDefaults() *TimeZones {
	this := TimeZones{}
	return &this
}

// GetItems returns the Items field value
func (o *TimeZones) GetItems() []TimeZone {
	if o == nil {
		var ret []TimeZone
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *TimeZones) GetItemsOk() (*[]TimeZone, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *TimeZones) SetItems(v []TimeZone) {
	o.Items = v
}

func (o TimeZones) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableTimeZones struct {
	value *TimeZones
	isSet bool
}

func (v NullableTimeZones) Get() *TimeZones {
	return v.value
}

func (v *NullableTimeZones) Set(val *TimeZones) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeZones) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeZones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeZones(val *TimeZones) *NullableTimeZones {
	return &NullableTimeZones{value: val, isSet: true}
}

func (v NullableTimeZones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeZones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


