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

// WeeklyScheduleSettings struct for WeeklyScheduleSettings
type WeeklyScheduleSettings struct {
	TimeLocal       Time                           `json:"timeLocal"`
	SnapshotOptions WeeklySnapshotScheduleSettings `json:"snapshotOptions"`
	BackupOptions   *WeeklyBackupScheduleSettings  `json:"backupOptions,omitempty"`
	ReplicaOptions  *WeeklyReplicaScheduleSettings `json:"replicaOptions,omitempty"`
}

// NewWeeklyScheduleSettings instantiates a new WeeklyScheduleSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeeklyScheduleSettings(timeLocal Time, snapshotOptions WeeklySnapshotScheduleSettings) *WeeklyScheduleSettings {
	this := WeeklyScheduleSettings{}
	this.TimeLocal = timeLocal
	this.SnapshotOptions = snapshotOptions
	return &this
}

// NewWeeklyScheduleSettingsWithDefaults instantiates a new WeeklyScheduleSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeeklyScheduleSettingsWithDefaults() *WeeklyScheduleSettings {
	this := WeeklyScheduleSettings{}
	return &this
}

// GetTimeLocal returns the TimeLocal field value
func (o *WeeklyScheduleSettings) GetTimeLocal() Time {
	if o == nil {
		var ret Time
		return ret
	}

	return o.TimeLocal
}

// GetTimeLocalOk returns a tuple with the TimeLocal field value
// and a boolean to check if the value has been set.
func (o *WeeklyScheduleSettings) GetTimeLocalOk() (*Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeLocal, true
}

// SetTimeLocal sets field value
func (o *WeeklyScheduleSettings) SetTimeLocal(v Time) {
	o.TimeLocal = v
}

// GetSnapshotOptions returns the SnapshotOptions field value
func (o *WeeklyScheduleSettings) GetSnapshotOptions() WeeklySnapshotScheduleSettings {
	if o == nil {
		var ret WeeklySnapshotScheduleSettings
		return ret
	}

	return o.SnapshotOptions
}

// GetSnapshotOptionsOk returns a tuple with the SnapshotOptions field value
// and a boolean to check if the value has been set.
func (o *WeeklyScheduleSettings) GetSnapshotOptionsOk() (*WeeklySnapshotScheduleSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotOptions, true
}

// SetSnapshotOptions sets field value
func (o *WeeklyScheduleSettings) SetSnapshotOptions(v WeeklySnapshotScheduleSettings) {
	o.SnapshotOptions = v
}

// GetBackupOptions returns the BackupOptions field value if set, zero value otherwise.
func (o *WeeklyScheduleSettings) GetBackupOptions() WeeklyBackupScheduleSettings {
	if o == nil || o.BackupOptions == nil {
		var ret WeeklyBackupScheduleSettings
		return ret
	}
	return *o.BackupOptions
}

// GetBackupOptionsOk returns a tuple with the BackupOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeeklyScheduleSettings) GetBackupOptionsOk() (*WeeklyBackupScheduleSettings, bool) {
	if o == nil || o.BackupOptions == nil {
		return nil, false
	}
	return o.BackupOptions, true
}

// HasBackupOptions returns a boolean if a field has been set.
func (o *WeeklyScheduleSettings) HasBackupOptions() bool {
	if o != nil && o.BackupOptions != nil {
		return true
	}

	return false
}

// SetBackupOptions gets a reference to the given WeeklyBackupScheduleSettings and assigns it to the BackupOptions field.
func (o *WeeklyScheduleSettings) SetBackupOptions(v WeeklyBackupScheduleSettings) {
	o.BackupOptions = &v
}

// GetReplicaOptions returns the ReplicaOptions field value if set, zero value otherwise.
func (o *WeeklyScheduleSettings) GetReplicaOptions() WeeklyReplicaScheduleSettings {
	if o == nil || o.ReplicaOptions == nil {
		var ret WeeklyReplicaScheduleSettings
		return ret
	}
	return *o.ReplicaOptions
}

// GetReplicaOptionsOk returns a tuple with the ReplicaOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeeklyScheduleSettings) GetReplicaOptionsOk() (*WeeklyReplicaScheduleSettings, bool) {
	if o == nil || o.ReplicaOptions == nil {
		return nil, false
	}
	return o.ReplicaOptions, true
}

// HasReplicaOptions returns a boolean if a field has been set.
func (o *WeeklyScheduleSettings) HasReplicaOptions() bool {
	if o != nil && o.ReplicaOptions != nil {
		return true
	}

	return false
}

// SetReplicaOptions gets a reference to the given WeeklyReplicaScheduleSettings and assigns it to the ReplicaOptions field.
func (o *WeeklyScheduleSettings) SetReplicaOptions(v WeeklyReplicaScheduleSettings) {
	o.ReplicaOptions = &v
}

func (o WeeklyScheduleSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timeLocal"] = o.TimeLocal
	}
	if true {
		toSerialize["snapshotOptions"] = o.SnapshotOptions
	}
	if o.BackupOptions != nil {
		toSerialize["backupOptions"] = o.BackupOptions
	}
	if o.ReplicaOptions != nil {
		toSerialize["replicaOptions"] = o.ReplicaOptions
	}
	return json.Marshal(toSerialize)
}

type NullableWeeklyScheduleSettings struct {
	value *WeeklyScheduleSettings
	isSet bool
}

func (v NullableWeeklyScheduleSettings) Get() *WeeklyScheduleSettings {
	return v.value
}

func (v *NullableWeeklyScheduleSettings) Set(val *WeeklyScheduleSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableWeeklyScheduleSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableWeeklyScheduleSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeeklyScheduleSettings(val *WeeklyScheduleSettings) *NullableWeeklyScheduleSettings {
	return &NullableWeeklyScheduleSettings{value: val, isSet: true}
}

func (v NullableWeeklyScheduleSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeeklyScheduleSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
