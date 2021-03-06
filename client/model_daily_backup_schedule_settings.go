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

// DailyBackupScheduleSettings struct for DailyBackupScheduleSettings
type DailyBackupScheduleSettings struct {
	Retention PeriodRetentionOptions `json:"retention"`
	Schedule DailySchedule `json:"schedule"`
}

// NewDailyBackupScheduleSettings instantiates a new DailyBackupScheduleSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyBackupScheduleSettings(retention PeriodRetentionOptions, schedule DailySchedule) *DailyBackupScheduleSettings {
	this := DailyBackupScheduleSettings{}
	this.Retention = retention
	this.Schedule = schedule
	return &this
}

// NewDailyBackupScheduleSettingsWithDefaults instantiates a new DailyBackupScheduleSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyBackupScheduleSettingsWithDefaults() *DailyBackupScheduleSettings {
	this := DailyBackupScheduleSettings{}
	return &this
}

// GetRetention returns the Retention field value
func (o *DailyBackupScheduleSettings) GetRetention() PeriodRetentionOptions {
	if o == nil {
		var ret PeriodRetentionOptions
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *DailyBackupScheduleSettings) GetRetentionOk() (*PeriodRetentionOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *DailyBackupScheduleSettings) SetRetention(v PeriodRetentionOptions) {
	o.Retention = v
}

// GetSchedule returns the Schedule field value
func (o *DailyBackupScheduleSettings) GetSchedule() DailySchedule {
	if o == nil {
		var ret DailySchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *DailyBackupScheduleSettings) GetScheduleOk() (*DailySchedule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *DailyBackupScheduleSettings) SetSchedule(v DailySchedule) {
	o.Schedule = v
}

func (o DailyBackupScheduleSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["retention"] = o.Retention
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableDailyBackupScheduleSettings struct {
	value *DailyBackupScheduleSettings
	isSet bool
}

func (v NullableDailyBackupScheduleSettings) Get() *DailyBackupScheduleSettings {
	return v.value
}

func (v *NullableDailyBackupScheduleSettings) Set(val *DailyBackupScheduleSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyBackupScheduleSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyBackupScheduleSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyBackupScheduleSettings(val *DailyBackupScheduleSettings) *NullableDailyBackupScheduleSettings {
	return &NullableDailyBackupScheduleSettings{value: val, isSet: true}
}

func (v NullableDailyBackupScheduleSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyBackupScheduleSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


