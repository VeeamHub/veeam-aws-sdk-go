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

// RetentionJobDeletedRestorePointEmbedded struct for RetentionJobDeletedRestorePointEmbedded
type RetentionJobDeletedRestorePointEmbedded struct {
	Region *Link `json:"region,omitempty"`
}

// NewRetentionJobDeletedRestorePointEmbedded instantiates a new RetentionJobDeletedRestorePointEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetentionJobDeletedRestorePointEmbedded() *RetentionJobDeletedRestorePointEmbedded {
	this := RetentionJobDeletedRestorePointEmbedded{}
	return &this
}

// NewRetentionJobDeletedRestorePointEmbeddedWithDefaults instantiates a new RetentionJobDeletedRestorePointEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetentionJobDeletedRestorePointEmbeddedWithDefaults() *RetentionJobDeletedRestorePointEmbedded {
	this := RetentionJobDeletedRestorePointEmbedded{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *RetentionJobDeletedRestorePointEmbedded) GetRegion() Link {
	if o == nil || o.Region == nil {
		var ret Link
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionJobDeletedRestorePointEmbedded) GetRegionOk() (*Link, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *RetentionJobDeletedRestorePointEmbedded) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given Link and assigns it to the Region field.
func (o *RetentionJobDeletedRestorePointEmbedded) SetRegion(v Link) {
	o.Region = &v
}

func (o RetentionJobDeletedRestorePointEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableRetentionJobDeletedRestorePointEmbedded struct {
	value *RetentionJobDeletedRestorePointEmbedded
	isSet bool
}

func (v NullableRetentionJobDeletedRestorePointEmbedded) Get() *RetentionJobDeletedRestorePointEmbedded {
	return v.value
}

func (v *NullableRetentionJobDeletedRestorePointEmbedded) Set(val *RetentionJobDeletedRestorePointEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableRetentionJobDeletedRestorePointEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableRetentionJobDeletedRestorePointEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetentionJobDeletedRestorePointEmbedded(val *RetentionJobDeletedRestorePointEmbedded) *NullableRetentionJobDeletedRestorePointEmbedded {
	return &NullableRetentionJobDeletedRestorePointEmbedded{value: val, isSet: true}
}

func (v NullableRetentionJobDeletedRestorePointEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetentionJobDeletedRestorePointEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


