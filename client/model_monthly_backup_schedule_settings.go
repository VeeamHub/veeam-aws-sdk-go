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

// MonthlyBackupScheduleSettings struct for MonthlyBackupScheduleSettings
type MonthlyBackupScheduleSettings struct {
	Retention PeriodRetentionOptions `json:"retention"`
	Schedule MonthlySchedule `json:"schedule"`
}

// NewMonthlyBackupScheduleSettings instantiates a new MonthlyBackupScheduleSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonthlyBackupScheduleSettings(retention PeriodRetentionOptions, schedule MonthlySchedule) *MonthlyBackupScheduleSettings {
	this := MonthlyBackupScheduleSettings{}
	this.Retention = retention
	this.Schedule = schedule
	return &this
}

// NewMonthlyBackupScheduleSettingsWithDefaults instantiates a new MonthlyBackupScheduleSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonthlyBackupScheduleSettingsWithDefaults() *MonthlyBackupScheduleSettings {
	this := MonthlyBackupScheduleSettings{}
	return &this
}

// GetRetention returns the Retention field value
func (o *MonthlyBackupScheduleSettings) GetRetention() PeriodRetentionOptions {
	if o == nil {
		var ret PeriodRetentionOptions
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *MonthlyBackupScheduleSettings) GetRetentionOk() (*PeriodRetentionOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *MonthlyBackupScheduleSettings) SetRetention(v PeriodRetentionOptions) {
	o.Retention = v
}

// GetSchedule returns the Schedule field value
func (o *MonthlyBackupScheduleSettings) GetSchedule() MonthlySchedule {
	if o == nil {
		var ret MonthlySchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *MonthlyBackupScheduleSettings) GetScheduleOk() (*MonthlySchedule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *MonthlyBackupScheduleSettings) SetSchedule(v MonthlySchedule) {
	o.Schedule = v
}

func (o MonthlyBackupScheduleSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["retention"] = o.Retention
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableMonthlyBackupScheduleSettings struct {
	value *MonthlyBackupScheduleSettings
	isSet bool
}

func (v NullableMonthlyBackupScheduleSettings) Get() *MonthlyBackupScheduleSettings {
	return v.value
}

func (v *NullableMonthlyBackupScheduleSettings) Set(val *MonthlyBackupScheduleSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMonthlyBackupScheduleSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMonthlyBackupScheduleSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonthlyBackupScheduleSettings(val *MonthlyBackupScheduleSettings) *NullableMonthlyBackupScheduleSettings {
	return &NullableMonthlyBackupScheduleSettings{value: val, isSet: true}
}

func (v NullableMonthlyBackupScheduleSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonthlyBackupScheduleSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

