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

// SessionsPage struct for SessionsPage
type SessionsPage struct {
	TotalCount int32 `json:"totalCount"`
	Results []Session `json:"results"`
	Links *[]Link `json:"_links,omitempty"`
}

// NewSessionsPage instantiates a new SessionsPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionsPage(totalCount int32, results []Session) *SessionsPage {
	this := SessionsPage{}
	this.TotalCount = totalCount
	this.Results = results
	return &this
}

// NewSessionsPageWithDefaults instantiates a new SessionsPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionsPageWithDefaults() *SessionsPage {
	this := SessionsPage{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *SessionsPage) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *SessionsPage) GetTotalCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *SessionsPage) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetResults returns the Results field value
func (o *SessionsPage) GetResults() []Session {
	if o == nil {
		var ret []Session
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *SessionsPage) GetResultsOk() (*[]Session, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *SessionsPage) SetResults(v []Session) {
	o.Results = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SessionsPage) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsPage) GetLinksOk() (*[]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SessionsPage) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *SessionsPage) SetLinks(v []Link) {
	o.Links = &v
}

func (o SessionsPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totalCount"] = o.TotalCount
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableSessionsPage struct {
	value *SessionsPage
	isSet bool
}

func (v NullableSessionsPage) Get() *SessionsPage {
	return v.value
}

func (v *NullableSessionsPage) Set(val *SessionsPage) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionsPage) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionsPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionsPage(val *SessionsPage) *NullableSessionsPage {
	return &NullableSessionsPage{value: val, isSet: true}
}

func (v NullableSessionsPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionsPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


