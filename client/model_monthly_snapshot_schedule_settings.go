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

// MonthlySnapshotScheduleSettings struct for MonthlySnapshotScheduleSettings
type MonthlySnapshotScheduleSettings struct {
	Retention PointsRetentionOptions `json:"retention"`
	Schedule MonthlySchedule `json:"schedule"`
}

// NewMonthlySnapshotScheduleSettings instantiates a new MonthlySnapshotScheduleSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonthlySnapshotScheduleSettings(retention PointsRetentionOptions, schedule MonthlySchedule) *MonthlySnapshotScheduleSettings {
	this := MonthlySnapshotScheduleSettings{}
	this.Retention = retention
	this.Schedule = schedule
	return &this
}

// NewMonthlySnapshotScheduleSettingsWithDefaults instantiates a new MonthlySnapshotScheduleSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonthlySnapshotScheduleSettingsWithDefaults() *MonthlySnapshotScheduleSettings {
	this := MonthlySnapshotScheduleSettings{}
	return &this
}

// GetRetention returns the Retention field value
func (o *MonthlySnapshotScheduleSettings) GetRetention() PointsRetentionOptions {
	if o == nil {
		var ret PointsRetentionOptions
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *MonthlySnapshotScheduleSettings) GetRetentionOk() (*PointsRetentionOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *MonthlySnapshotScheduleSettings) SetRetention(v PointsRetentionOptions) {
	o.Retention = v
}

// GetSchedule returns the Schedule field value
func (o *MonthlySnapshotScheduleSettings) GetSchedule() MonthlySchedule {
	if o == nil {
		var ret MonthlySchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *MonthlySnapshotScheduleSettings) GetScheduleOk() (*MonthlySchedule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *MonthlySnapshotScheduleSettings) SetSchedule(v MonthlySchedule) {
	o.Schedule = v
}

func (o MonthlySnapshotScheduleSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["retention"] = o.Retention
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableMonthlySnapshotScheduleSettings struct {
	value *MonthlySnapshotScheduleSettings
	isSet bool
}

func (v NullableMonthlySnapshotScheduleSettings) Get() *MonthlySnapshotScheduleSettings {
	return v.value
}

func (v *NullableMonthlySnapshotScheduleSettings) Set(val *MonthlySnapshotScheduleSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMonthlySnapshotScheduleSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMonthlySnapshotScheduleSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonthlySnapshotScheduleSettings(val *MonthlySnapshotScheduleSettings) *NullableMonthlySnapshotScheduleSettings {
	return &NullableMonthlySnapshotScheduleSettings{value: val, isSet: true}
}

func (v NullableMonthlySnapshotScheduleSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonthlySnapshotScheduleSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

