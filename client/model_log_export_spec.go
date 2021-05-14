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

// LogExportSpec struct for LogExportSpec
type LogExportSpec struct {
	ExportLogsType *string `json:"exportLogsType,omitempty"`
	Days           *int32  `json:"days,omitempty"`
	FromDateUtc    *Time   `json:"fromDateUtc,omitempty"`
	ToDateUtc      *Time   `json:"toDateUtc,omitempty"`
}

// NewLogExportSpec instantiates a new LogExportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogExportSpec() *LogExportSpec {
	this := LogExportSpec{}
	return &this
}

// NewLogExportSpecWithDefaults instantiates a new LogExportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogExportSpecWithDefaults() *LogExportSpec {
	this := LogExportSpec{}
	return &this
}

// GetExportLogsType returns the ExportLogsType field value if set, zero value otherwise.
func (o *LogExportSpec) GetExportLogsType() string {
	if o == nil || o.ExportLogsType == nil {
		var ret string
		return ret
	}
	return *o.ExportLogsType
}

// GetExportLogsTypeOk returns a tuple with the ExportLogsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExportSpec) GetExportLogsTypeOk() (*string, bool) {
	if o == nil || o.ExportLogsType == nil {
		return nil, false
	}
	return o.ExportLogsType, true
}

// HasExportLogsType returns a boolean if a field has been set.
func (o *LogExportSpec) HasExportLogsType() bool {
	if o != nil && o.ExportLogsType != nil {
		return true
	}

	return false
}

// SetExportLogsType gets a reference to the given string and assigns it to the ExportLogsType field.
func (o *LogExportSpec) SetExportLogsType(v string) {
	o.ExportLogsType = &v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *LogExportSpec) GetDays() int32 {
	if o == nil || o.Days == nil {
		var ret int32
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExportSpec) GetDaysOk() (*int32, bool) {
	if o == nil || o.Days == nil {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *LogExportSpec) HasDays() bool {
	if o != nil && o.Days != nil {
		return true
	}

	return false
}

// SetDays gets a reference to the given int32 and assigns it to the Days field.
func (o *LogExportSpec) SetDays(v int32) {
	o.Days = &v
}

// GetFromDateUtc returns the FromDateUtc field value if set, zero value otherwise.
func (o *LogExportSpec) GetFromDateUtc() Time {
	if o == nil || o.FromDateUtc == nil {
		var ret Time
		return ret
	}
	return *o.FromDateUtc
}

// GetFromDateUtcOk returns a tuple with the FromDateUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExportSpec) GetFromDateUtcOk() (*Time, bool) {
	if o == nil || o.FromDateUtc == nil {
		return nil, false
	}
	return o.FromDateUtc, true
}

// HasFromDateUtc returns a boolean if a field has been set.
func (o *LogExportSpec) HasFromDateUtc() bool {
	if o != nil && o.FromDateUtc != nil {
		return true
	}

	return false
}

// SetFromDateUtc gets a reference to the given Time and assigns it to the FromDateUtc field.
func (o *LogExportSpec) SetFromDateUtc(v Time) {
	o.FromDateUtc = &v
}

// GetToDateUtc returns the ToDateUtc field value if set, zero value otherwise.
func (o *LogExportSpec) GetToDateUtc() Time {
	if o == nil || o.ToDateUtc == nil {
		var ret Time
		return ret
	}
	return *o.ToDateUtc
}

// GetToDateUtcOk returns a tuple with the ToDateUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExportSpec) GetToDateUtcOk() (*Time, bool) {
	if o == nil || o.ToDateUtc == nil {
		return nil, false
	}
	return o.ToDateUtc, true
}

// HasToDateUtc returns a boolean if a field has been set.
func (o *LogExportSpec) HasToDateUtc() bool {
	if o != nil && o.ToDateUtc != nil {
		return true
	}

	return false
}

// SetToDateUtc gets a reference to the given Time and assigns it to the ToDateUtc field.
func (o *LogExportSpec) SetToDateUtc(v Time) {
	o.ToDateUtc = &v
}

func (o LogExportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExportLogsType != nil {
		toSerialize["exportLogsType"] = o.ExportLogsType
	}
	if o.Days != nil {
		toSerialize["days"] = o.Days
	}
	if o.FromDateUtc != nil {
		toSerialize["fromDateUtc"] = o.FromDateUtc
	}
	if o.ToDateUtc != nil {
		toSerialize["toDateUtc"] = o.ToDateUtc
	}
	return json.Marshal(toSerialize)
}

type NullableLogExportSpec struct {
	value *LogExportSpec
	isSet bool
}

func (v NullableLogExportSpec) Get() *LogExportSpec {
	return v.value
}

func (v *NullableLogExportSpec) Set(val *LogExportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableLogExportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableLogExportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogExportSpec(val *LogExportSpec) *NullableLogExportSpec {
	return &NullableLogExportSpec{value: val, isSet: true}
}

func (v NullableLogExportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogExportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}