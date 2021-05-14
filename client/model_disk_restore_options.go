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

// DiskRestoreOptions struct for DiskRestoreOptions
type DiskRestoreOptions struct {
	DiskId string `json:"diskId"`
	Name *string `json:"name,omitempty"`
}

// NewDiskRestoreOptions instantiates a new DiskRestoreOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskRestoreOptions(diskId string) *DiskRestoreOptions {
	this := DiskRestoreOptions{}
	this.DiskId = diskId
	return &this
}

// NewDiskRestoreOptionsWithDefaults instantiates a new DiskRestoreOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskRestoreOptionsWithDefaults() *DiskRestoreOptions {
	this := DiskRestoreOptions{}
	return &this
}

// GetDiskId returns the DiskId field value
func (o *DiskRestoreOptions) GetDiskId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiskId
}

// GetDiskIdOk returns a tuple with the DiskId field value
// and a boolean to check if the value has been set.
func (o *DiskRestoreOptions) GetDiskIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DiskId, true
}

// SetDiskId sets field value
func (o *DiskRestoreOptions) SetDiskId(v string) {
	o.DiskId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DiskRestoreOptions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskRestoreOptions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DiskRestoreOptions) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DiskRestoreOptions) SetName(v string) {
	o.Name = &v
}

func (o DiskRestoreOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["diskId"] = o.DiskId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDiskRestoreOptions struct {
	value *DiskRestoreOptions
	isSet bool
}

func (v NullableDiskRestoreOptions) Get() *DiskRestoreOptions {
	return v.value
}

func (v *NullableDiskRestoreOptions) Set(val *DiskRestoreOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskRestoreOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskRestoreOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskRestoreOptions(val *DiskRestoreOptions) *NullableDiskRestoreOptions {
	return &NullableDiskRestoreOptions{value: val, isSet: true}
}

func (v NullableDiskRestoreOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskRestoreOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


