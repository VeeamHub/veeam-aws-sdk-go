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

// VirtualMachineLocation struct for VirtualMachineLocation
type VirtualMachineLocation struct {
	Name string `json:"name"`
	Id string `json:"id"`
}

// NewVirtualMachineLocation instantiates a new VirtualMachineLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMachineLocation(name string, id string) *VirtualMachineLocation {
	this := VirtualMachineLocation{}
	this.Name = name
	this.Id = id
	return &this
}

// NewVirtualMachineLocationWithDefaults instantiates a new VirtualMachineLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMachineLocationWithDefaults() *VirtualMachineLocation {
	this := VirtualMachineLocation{}
	return &this
}

// GetName returns the Name field value
func (o *VirtualMachineLocation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineLocation) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VirtualMachineLocation) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *VirtualMachineLocation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineLocation) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VirtualMachineLocation) SetId(v string) {
	o.Id = v
}

func (o VirtualMachineLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualMachineLocation struct {
	value *VirtualMachineLocation
	isSet bool
}

func (v NullableVirtualMachineLocation) Get() *VirtualMachineLocation {
	return v.value
}

func (v *NullableVirtualMachineLocation) Set(val *VirtualMachineLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineLocation(val *VirtualMachineLocation) *NullableVirtualMachineLocation {
	return &NullableVirtualMachineLocation{value: val, isSet: true}
}

func (v NullableVirtualMachineLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

