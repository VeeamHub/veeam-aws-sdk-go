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

// WeeklyReplicaScheduleSettings struct for WeeklyReplicaScheduleSettings
type WeeklyReplicaScheduleSettings struct {
	Retention PointsRetentionOptions `json:"retention"`
	Schedule WeeklySchedule `json:"schedule"`
}

// NewWeeklyReplicaScheduleSettings instantiates a new WeeklyReplicaScheduleSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeeklyReplicaScheduleSettings(retention PointsRetentionOptions, schedule WeeklySchedule) *WeeklyReplicaScheduleSettings {
	this := WeeklyReplicaScheduleSettings{}
	this.Retention = retention
	this.Schedule = schedule
	return &this
}

// NewWeeklyReplicaScheduleSettingsWithDefaults instantiates a new WeeklyReplicaScheduleSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeeklyReplicaScheduleSettingsWithDefaults() *WeeklyReplicaScheduleSettings {
	this := WeeklyReplicaScheduleSettings{}
	return &this
}

// GetRetention returns the Retention field value
func (o *WeeklyReplicaScheduleSettings) GetRetention() PointsRetentionOptions {
	if o == nil {
		var ret PointsRetentionOptions
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *WeeklyReplicaScheduleSettings) GetRetentionOk() (*PointsRetentionOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *WeeklyReplicaScheduleSettings) SetRetention(v PointsRetentionOptions) {
	o.Retention = v
}

// GetSchedule returns the Schedule field value
func (o *WeeklyReplicaScheduleSettings) GetSchedule() WeeklySchedule {
	if o == nil {
		var ret WeeklySchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *WeeklyReplicaScheduleSettings) GetScheduleOk() (*WeeklySchedule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *WeeklyReplicaScheduleSettings) SetSchedule(v WeeklySchedule) {
	o.Schedule = v
}

func (o WeeklyReplicaScheduleSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["retention"] = o.Retention
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableWeeklyReplicaScheduleSettings struct {
	value *WeeklyReplicaScheduleSettings
	isSet bool
}

func (v NullableWeeklyReplicaScheduleSettings) Get() *WeeklyReplicaScheduleSettings {
	return v.value
}

func (v *NullableWeeklyReplicaScheduleSettings) Set(val *WeeklyReplicaScheduleSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableWeeklyReplicaScheduleSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableWeeklyReplicaScheduleSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeeklyReplicaScheduleSettings(val *WeeklyReplicaScheduleSettings) *NullableWeeklyReplicaScheduleSettings {
	return &NullableWeeklyReplicaScheduleSettings{value: val, isSet: true}
}

func (v NullableWeeklyReplicaScheduleSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeeklyReplicaScheduleSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


