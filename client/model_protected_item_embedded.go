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

// ProtectedItemEmbedded struct for ProtectedItemEmbedded
type ProtectedItemEmbedded struct {
	VirtualMachine *Link `json:"virtualMachine,omitempty"`
	Region *Link `json:"region,omitempty"`
}

// NewProtectedItemEmbedded instantiates a new ProtectedItemEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectedItemEmbedded() *ProtectedItemEmbedded {
	this := ProtectedItemEmbedded{}
	return &this
}

// NewProtectedItemEmbeddedWithDefaults instantiates a new ProtectedItemEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectedItemEmbeddedWithDefaults() *ProtectedItemEmbedded {
	this := ProtectedItemEmbedded{}
	return &this
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *ProtectedItemEmbedded) GetVirtualMachine() Link {
	if o == nil || o.VirtualMachine == nil {
		var ret Link
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectedItemEmbedded) GetVirtualMachineOk() (*Link, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *ProtectedItemEmbedded) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given Link and assigns it to the VirtualMachine field.
func (o *ProtectedItemEmbedded) SetVirtualMachine(v Link) {
	o.VirtualMachine = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ProtectedItemEmbedded) GetRegion() Link {
	if o == nil || o.Region == nil {
		var ret Link
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectedItemEmbedded) GetRegionOk() (*Link, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ProtectedItemEmbedded) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given Link and assigns it to the Region field.
func (o *ProtectedItemEmbedded) SetRegion(v Link) {
	o.Region = &v
}

func (o ProtectedItemEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtualMachine != nil {
		toSerialize["virtualMachine"] = o.VirtualMachine
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableProtectedItemEmbedded struct {
	value *ProtectedItemEmbedded
	isSet bool
}

func (v NullableProtectedItemEmbedded) Get() *ProtectedItemEmbedded {
	return v.value
}

func (v *NullableProtectedItemEmbedded) Set(val *ProtectedItemEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectedItemEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectedItemEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectedItemEmbedded(val *ProtectedItemEmbedded) *NullableProtectedItemEmbedded {
	return &NullableProtectedItemEmbedded{value: val, isSet: true}
}

func (v NullableProtectedItemEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectedItemEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

