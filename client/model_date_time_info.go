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

// DateTimeInfo struct for DateTimeInfo
type DateTimeInfo struct {
	UtcDateTime Time `json:"utcDateTime"`
}

// NewDateTimeInfo instantiates a new DateTimeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateTimeInfo(utcDateTime Time) *DateTimeInfo {
	this := DateTimeInfo{}
	this.UtcDateTime = utcDateTime
	return &this
}

// NewDateTimeInfoWithDefaults instantiates a new DateTimeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateTimeInfoWithDefaults() *DateTimeInfo {
	this := DateTimeInfo{}
	return &this
}

// GetUtcDateTime returns the UtcDateTime field value
func (o *DateTimeInfo) GetUtcDateTime() Time {
	if o == nil {
		var ret Time
		return ret
	}

	return o.UtcDateTime
}

// GetUtcDateTimeOk returns a tuple with the UtcDateTime field value
// and a boolean to check if the value has been set.
func (o *DateTimeInfo) GetUtcDateTimeOk() (*Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UtcDateTime, true
}

// SetUtcDateTime sets field value
func (o *DateTimeInfo) SetUtcDateTime(v Time) {
	o.UtcDateTime = v
}

func (o DateTimeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["utcDateTime"] = o.UtcDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableDateTimeInfo struct {
	value *DateTimeInfo
	isSet bool
}

func (v NullableDateTimeInfo) Get() *DateTimeInfo {
	return v.value
}

func (v *NullableDateTimeInfo) Set(val *DateTimeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDateTimeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDateTimeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateTimeInfo(val *DateTimeInfo) *NullableDateTimeInfo {
	return &NullableDateTimeInfo{value: val, isSet: true}
}

func (v NullableDateTimeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateTimeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}