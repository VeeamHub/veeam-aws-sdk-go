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

// SessionLink struct for SessionLink
type SessionLink struct {
	SessionId string `json:"sessionId"`
	Links *[]Link `json:"_links,omitempty"`
}

// NewSessionLink instantiates a new SessionLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionLink(sessionId string) *SessionLink {
	this := SessionLink{}
	this.SessionId = sessionId
	return &this
}

// NewSessionLinkWithDefaults instantiates a new SessionLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionLinkWithDefaults() *SessionLink {
	this := SessionLink{}
	return &this
}

// GetSessionId returns the SessionId field value
func (o *SessionLink) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *SessionLink) GetSessionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *SessionLink) SetSessionId(v string) {
	o.SessionId = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SessionLink) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionLink) GetLinksOk() (*[]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SessionLink) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *SessionLink) SetLinks(v []Link) {
	o.Links = &v
}

func (o SessionLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sessionId"] = o.SessionId
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableSessionLink struct {
	value *SessionLink
	isSet bool
}

func (v NullableSessionLink) Get() *SessionLink {
	return v.value
}

func (v *NullableSessionLink) Set(val *SessionLink) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionLink) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionLink(val *SessionLink) *NullableSessionLink {
	return &NullableSessionLink{value: val, isSet: true}
}

func (v NullableSessionLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

