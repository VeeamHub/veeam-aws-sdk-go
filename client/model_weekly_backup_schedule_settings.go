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

// WeeklyBackupScheduleSettings struct for WeeklyBackupScheduleSettings
type WeeklyBackupScheduleSettings struct {
	Retention PeriodRetentionOptions `json:"retention"`
	Schedule WeeklySchedule `json:"schedule"`
}

// NewWeeklyBackupScheduleSettings instantiates a new WeeklyBackupScheduleSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeeklyBackupScheduleSettings(retention PeriodRetentionOptions, schedule WeeklySchedule) *WeeklyBackupScheduleSettings {
	this := WeeklyBackupScheduleSettings{}
	this.Retention = retention
	this.Schedule = schedule
	return &this
}

// NewWeeklyBackupScheduleSettingsWithDefaults instantiates a new WeeklyBackupScheduleSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeeklyBackupScheduleSettingsWithDefaults() *WeeklyBackupScheduleSettings {
	this := WeeklyBackupScheduleSettings{}
	return &this
}

// GetRetention returns the Retention field value
func (o *WeeklyBackupScheduleSettings) GetRetention() PeriodRetentionOptions {
	if o == nil {
		var ret PeriodRetentionOptions
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *WeeklyBackupScheduleSettings) GetRetentionOk() (*PeriodRetentionOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *WeeklyBackupScheduleSettings) SetRetention(v PeriodRetentionOptions) {
	o.Retention = v
}

// GetSchedule returns the Schedule field value
func (o *WeeklyBackupScheduleSettings) GetSchedule() WeeklySchedule {
	if o == nil {
		var ret WeeklySchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *WeeklyBackupScheduleSettings) GetScheduleOk() (*WeeklySchedule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *WeeklyBackupScheduleSettings) SetSchedule(v WeeklySchedule) {
	o.Schedule = v
}

func (o WeeklyBackupScheduleSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["retention"] = o.Retention
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableWeeklyBackupScheduleSettings struct {
	value *WeeklyBackupScheduleSettings
	isSet bool
}

func (v NullableWeeklyBackupScheduleSettings) Get() *WeeklyBackupScheduleSettings {
	return v.value
}

func (v *NullableWeeklyBackupScheduleSettings) Set(val *WeeklyBackupScheduleSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableWeeklyBackupScheduleSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableWeeklyBackupScheduleSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeeklyBackupScheduleSettings(val *WeeklyBackupScheduleSettings) *NullableWeeklyBackupScheduleSettings {
	return &NullableWeeklyBackupScheduleSettings{value: val, isSet: true}
}

func (v NullableWeeklyBackupScheduleSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeeklyBackupScheduleSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

