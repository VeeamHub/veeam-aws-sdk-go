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

// RetentionSetting struct for RetentionSetting
type RetentionSetting struct {
	TimeRetentionDuration int32 `json:"timeRetentionDuration"`
	RetentionDurationType string `json:"retentionDurationType"`
}

// NewRetentionSetting instantiates a new RetentionSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetentionSetting(timeRetentionDuration int32, retentionDurationType string) *RetentionSetting {
	this := RetentionSetting{}
	this.TimeRetentionDuration = timeRetentionDuration
	this.RetentionDurationType = retentionDurationType
	return &this
}

// NewRetentionSettingWithDefaults instantiates a new RetentionSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetentionSettingWithDefaults() *RetentionSetting {
	this := RetentionSetting{}
	return &this
}

// GetTimeRetentionDuration returns the TimeRetentionDuration field value
func (o *RetentionSetting) GetTimeRetentionDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeRetentionDuration
}

// GetTimeRetentionDurationOk returns a tuple with the TimeRetentionDuration field value
// and a boolean to check if the value has been set.
func (o *RetentionSetting) GetTimeRetentionDurationOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimeRetentionDuration, true
}

// SetTimeRetentionDuration sets field value
func (o *RetentionSetting) SetTimeRetentionDuration(v int32) {
	o.TimeRetentionDuration = v
}

// GetRetentionDurationType returns the RetentionDurationType field value
func (o *RetentionSetting) GetRetentionDurationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RetentionDurationType
}

// GetRetentionDurationTypeOk returns a tuple with the RetentionDurationType field value
// and a boolean to check if the value has been set.
func (o *RetentionSetting) GetRetentionDurationTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RetentionDurationType, true
}

// SetRetentionDurationType sets field value
func (o *RetentionSetting) SetRetentionDurationType(v string) {
	o.RetentionDurationType = v
}

func (o RetentionSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timeRetentionDuration"] = o.TimeRetentionDuration
	}
	if true {
		toSerialize["retentionDurationType"] = o.RetentionDurationType
	}
	return json.Marshal(toSerialize)
}

type NullableRetentionSetting struct {
	value *RetentionSetting
	isSet bool
}

func (v NullableRetentionSetting) Get() *RetentionSetting {
	return v.value
}

func (v *NullableRetentionSetting) Set(val *RetentionSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableRetentionSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableRetentionSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetentionSetting(val *RetentionSetting) *NullableRetentionSetting {
	return &NullableRetentionSetting{value: val, isSet: true}
}

func (v NullableRetentionSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetentionSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

