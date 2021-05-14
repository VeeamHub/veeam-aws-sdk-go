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

// LogItem struct for LogItem
type LogItem struct {
	LogTime            Time    `json:"logTime"`
	Status             string  `json:"status"`
	Message            *string `json:"message,omitempty"`
	ExecutionStartTime *Time   `json:"executionStartTime,omitempty"`
	ExecutionDuration  *int64  `json:"executionDuration,omitempty"`
}

// NewLogItem instantiates a new LogItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogItem(logTime Time, status string) *LogItem {
	this := LogItem{}
	this.LogTime = logTime
	this.Status = status
	return &this
}

// NewLogItemWithDefaults instantiates a new LogItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogItemWithDefaults() *LogItem {
	this := LogItem{}
	return &this
}

// GetLogTime returns the LogTime field value
func (o *LogItem) GetLogTime() Time {
	if o == nil {
		var ret Time
		return ret
	}

	return o.LogTime
}

// GetLogTimeOk returns a tuple with the LogTime field value
// and a boolean to check if the value has been set.
func (o *LogItem) GetLogTimeOk() (*Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogTime, true
}

// SetLogTime sets field value
func (o *LogItem) SetLogTime(v Time) {
	o.LogTime = v
}

// GetStatus returns the Status field value
func (o *LogItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LogItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LogItem) SetStatus(v string) {
	o.Status = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LogItem) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogItem) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LogItem) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LogItem) SetMessage(v string) {
	o.Message = &v
}

// GetExecutionStartTime returns the ExecutionStartTime field value if set, zero value otherwise.
func (o *LogItem) GetExecutionStartTime() Time {
	if o == nil || o.ExecutionStartTime == nil {
		var ret Time
		return ret
	}
	return *o.ExecutionStartTime
}

// GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogItem) GetExecutionStartTimeOk() (*Time, bool) {
	if o == nil || o.ExecutionStartTime == nil {
		return nil, false
	}
	return o.ExecutionStartTime, true
}

// HasExecutionStartTime returns a boolean if a field has been set.
func (o *LogItem) HasExecutionStartTime() bool {
	if o != nil && o.ExecutionStartTime != nil {
		return true
	}

	return false
}

// SetExecutionStartTime gets a reference to the given Time and assigns it to the ExecutionStartTime field.
func (o *LogItem) SetExecutionStartTime(v Time) {
	o.ExecutionStartTime = &v
}

// GetExecutionDuration returns the ExecutionDuration field value if set, zero value otherwise.
func (o *LogItem) GetExecutionDuration() int64 {
	if o == nil || o.ExecutionDuration == nil {
		var ret int64
		return ret
	}
	return *o.ExecutionDuration
}

// GetExecutionDurationOk returns a tuple with the ExecutionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogItem) GetExecutionDurationOk() (*int64, bool) {
	if o == nil || o.ExecutionDuration == nil {
		return nil, false
	}
	return o.ExecutionDuration, true
}

// HasExecutionDuration returns a boolean if a field has been set.
func (o *LogItem) HasExecutionDuration() bool {
	if o != nil && o.ExecutionDuration != nil {
		return true
	}

	return false
}

// SetExecutionDuration gets a reference to the given int64 and assigns it to the ExecutionDuration field.
func (o *LogItem) SetExecutionDuration(v int64) {
	o.ExecutionDuration = &v
}

func (o LogItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["logTime"] = o.LogTime
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ExecutionStartTime != nil {
		toSerialize["executionStartTime"] = o.ExecutionStartTime
	}
	if o.ExecutionDuration != nil {
		toSerialize["executionDuration"] = o.ExecutionDuration
	}
	return json.Marshal(toSerialize)
}

type NullableLogItem struct {
	value *LogItem
	isSet bool
}

func (v NullableLogItem) Get() *LogItem {
	return v.value
}

func (v *NullableLogItem) Set(val *LogItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLogItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLogItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogItem(val *LogItem) *NullableLogItem {
	return &NullableLogItem{value: val, isSet: true}
}

func (v NullableLogItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
