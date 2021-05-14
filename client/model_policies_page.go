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

// PoliciesPage struct for PoliciesPage
type PoliciesPage struct {
	TotalCount int32 `json:"totalCount"`
	Results []Policy `json:"results"`
	Links *[]Link `json:"_links,omitempty"`
}

// NewPoliciesPage instantiates a new PoliciesPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoliciesPage(totalCount int32, results []Policy) *PoliciesPage {
	this := PoliciesPage{}
	this.TotalCount = totalCount
	this.Results = results
	return &this
}

// NewPoliciesPageWithDefaults instantiates a new PoliciesPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesPageWithDefaults() *PoliciesPage {
	this := PoliciesPage{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *PoliciesPage) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *PoliciesPage) GetTotalCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *PoliciesPage) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetResults returns the Results field value
func (o *PoliciesPage) GetResults() []Policy {
	if o == nil {
		var ret []Policy
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PoliciesPage) GetResultsOk() (*[]Policy, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PoliciesPage) SetResults(v []Policy) {
	o.Results = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PoliciesPage) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesPage) GetLinksOk() (*[]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PoliciesPage) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PoliciesPage) SetLinks(v []Link) {
	o.Links = &v
}

func (o PoliciesPage) MarshalJSON() ([]byte, error) {
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

type NullablePoliciesPage struct {
	value *PoliciesPage
	isSet bool
}

func (v NullablePoliciesPage) Get() *PoliciesPage {
	return v.value
}

func (v *NullablePoliciesPage) Set(val *PoliciesPage) {
	v.value = val
	v.isSet = true
}

func (v NullablePoliciesPage) IsSet() bool {
	return v.isSet
}

func (v *NullablePoliciesPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoliciesPage(val *PoliciesPage) *NullablePoliciesPage {
	return &NullablePoliciesPage{value: val, isSet: true}
}

func (v NullablePoliciesPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoliciesPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


