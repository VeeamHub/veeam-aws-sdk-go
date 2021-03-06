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

// Session struct for Session
type Session struct {
	Id                 string                    `json:"id"`
	Type               SessionTypes              `json:"type"`
	ExecutionStartTime *Time                     `json:"executionStartTime,omitempty"`
	ExecutionStopTime  *Time                     `json:"executionStopTime,omitempty"`
	ExecutionDuration  *int64                    `json:"executionDuration,omitempty"`
	Status             SessionStatuses           `json:"status"`
	Result             *SessionResults           `json:"result,omitempty"`
	Reason             *string                   `json:"reason,omitempty"`
	Usn                int64                     `json:"usn"`
	Links              *[]Link                   `json:"_links,omitempty"`
	Embedded           *SessionEmbeddedResources `json:"_embedded,omitempty"`
}

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession(id string, type_ SessionTypes, status SessionStatuses, usn int64) *Session {
	this := Session{}
	this.Id = id
	this.Type = type_
	this.Status = status
	this.Usn = usn
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	return &this
}

// GetId returns the Id field value
func (o *Session) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Session) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Session) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *Session) GetType() SessionTypes {
	if o == nil {
		var ret SessionTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Session) GetTypeOk() (*SessionTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Session) SetType(v SessionTypes) {
	o.Type = v
}

// GetExecutionStartTime returns the ExecutionStartTime field value if set, zero value otherwise.
func (o *Session) GetExecutionStartTime() Time {
	if o == nil || o.ExecutionStartTime == nil {
		var ret Time
		return ret
	}
	return *o.ExecutionStartTime
}

// GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetExecutionStartTimeOk() (*Time, bool) {
	if o == nil || o.ExecutionStartTime == nil {
		return nil, false
	}
	return o.ExecutionStartTime, true
}

// HasExecutionStartTime returns a boolean if a field has been set.
func (o *Session) HasExecutionStartTime() bool {
	if o != nil && o.ExecutionStartTime != nil {
		return true
	}

	return false
}

// SetExecutionStartTime gets a reference to the given time.Time and assigns it to the ExecutionStartTime field.
func (o *Session) SetExecutionStartTime(v Time) {
	o.ExecutionStartTime = &v
}

// GetExecutionStopTime returns the ExecutionStopTime field value if set, zero value otherwise.
func (o *Session) GetExecutionStopTime() Time {
	if o == nil || o.ExecutionStopTime == nil {
		var ret Time
		return ret
	}
	return *o.ExecutionStopTime
}

// GetExecutionStopTimeOk returns a tuple with the ExecutionStopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetExecutionStopTimeOk() (*Time, bool) {
	if o == nil || o.ExecutionStopTime == nil {
		return nil, false
	}
	return o.ExecutionStopTime, true
}

// HasExecutionStopTime returns a boolean if a field has been set.
func (o *Session) HasExecutionStopTime() bool {
	if o != nil && o.ExecutionStopTime != nil {
		return true
	}

	return false
}

// SetExecutionStopTime gets a reference to the given time.Time and assigns it to the ExecutionStopTime field.
func (o *Session) SetExecutionStopTime(v Time) {
	o.ExecutionStopTime = &v
}

// GetExecutionDuration returns the ExecutionDuration field value if set, zero value otherwise.
func (o *Session) GetExecutionDuration() int64 {
	if o == nil || o.ExecutionDuration == nil {
		var ret int64
		return ret
	}
	return *o.ExecutionDuration
}

// GetExecutionDurationOk returns a tuple with the ExecutionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetExecutionDurationOk() (*int64, bool) {
	if o == nil || o.ExecutionDuration == nil {
		return nil, false
	}
	return o.ExecutionDuration, true
}

// HasExecutionDuration returns a boolean if a field has been set.
func (o *Session) HasExecutionDuration() bool {
	if o != nil && o.ExecutionDuration != nil {
		return true
	}

	return false
}

// SetExecutionDuration gets a reference to the given int64 and assigns it to the ExecutionDuration field.
func (o *Session) SetExecutionDuration(v int64) {
	o.ExecutionDuration = &v
}

// GetStatus returns the Status field value
func (o *Session) GetStatus() SessionStatuses {
	if o == nil {
		var ret SessionStatuses
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Session) GetStatusOk() (*SessionStatuses, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Session) SetStatus(v SessionStatuses) {
	o.Status = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *Session) GetResult() SessionResults {
	if o == nil || o.Result == nil {
		var ret SessionResults
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetResultOk() (*SessionResults, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *Session) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given SessionResults and assigns it to the Result field.
func (o *Session) SetResult(v SessionResults) {
	o.Result = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Session) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Session) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Session) SetReason(v string) {
	o.Reason = &v
}

// GetUsn returns the Usn field value
func (o *Session) GetUsn() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Usn
}

// GetUsnOk returns a tuple with the Usn field value
// and a boolean to check if the value has been set.
func (o *Session) GetUsnOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usn, true
}

// SetUsn sets field value
func (o *Session) SetUsn(v int64) {
	o.Usn = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Session) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetLinksOk() (*[]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Session) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *Session) SetLinks(v []Link) {
	o.Links = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *Session) GetEmbedded() SessionEmbeddedResources {
	if o == nil || o.Embedded == nil {
		var ret SessionEmbeddedResources
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetEmbeddedOk() (*SessionEmbeddedResources, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *Session) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given SessionEmbeddedResources and assigns it to the Embedded field.
func (o *Session) SetEmbedded(v SessionEmbeddedResources) {
	o.Embedded = &v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.ExecutionStartTime != nil {
		toSerialize["executionStartTime"] = o.ExecutionStartTime
	}
	if o.ExecutionStopTime != nil {
		toSerialize["executionStopTime"] = o.ExecutionStopTime
	}
	if o.ExecutionDuration != nil {
		toSerialize["executionDuration"] = o.ExecutionDuration
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["usn"] = o.Usn
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}
	return json.Marshal(toSerialize)
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
