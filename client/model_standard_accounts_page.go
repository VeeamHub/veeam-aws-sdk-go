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

// StandardAccountsPage struct for StandardAccountsPage
type StandardAccountsPage struct {
	Results *[]StandardAccount `json:"results,omitempty"`
	TotalCount int32 `json:"totalCount"`
	Links *[]Link `json:"_links,omitempty"`
}

// NewStandardAccountsPage instantiates a new StandardAccountsPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardAccountsPage(totalCount int32) *StandardAccountsPage {
	this := StandardAccountsPage{}
	this.TotalCount = totalCount
	return &this
}

// NewStandardAccountsPageWithDefaults instantiates a new StandardAccountsPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardAccountsPageWithDefaults() *StandardAccountsPage {
	this := StandardAccountsPage{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *StandardAccountsPage) GetResults() []StandardAccount {
	if o == nil || o.Results == nil {
		var ret []StandardAccount
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardAccountsPage) GetResultsOk() (*[]StandardAccount, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *StandardAccountsPage) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []StandardAccount and assigns it to the Results field.
func (o *StandardAccountsPage) SetResults(v []StandardAccount) {
	o.Results = &v
}

// GetTotalCount returns the TotalCount field value
func (o *StandardAccountsPage) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *StandardAccountsPage) GetTotalCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *StandardAccountsPage) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *StandardAccountsPage) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardAccountsPage) GetLinksOk() (*[]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StandardAccountsPage) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StandardAccountsPage) SetLinks(v []Link) {
	o.Links = &v
}

func (o StandardAccountsPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if true {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableStandardAccountsPage struct {
	value *StandardAccountsPage
	isSet bool
}

func (v NullableStandardAccountsPage) Get() *StandardAccountsPage {
	return v.value
}

func (v *NullableStandardAccountsPage) Set(val *StandardAccountsPage) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardAccountsPage) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardAccountsPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardAccountsPage(val *StandardAccountsPage) *NullableStandardAccountsPage {
	return &NullableStandardAccountsPage{value: val, isSet: true}
}

func (v NullableStandardAccountsPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardAccountsPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

