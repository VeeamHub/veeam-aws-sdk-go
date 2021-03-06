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

// AmazonAccountsPage struct for AmazonAccountsPage
type AmazonAccountsPage struct {
	Results []AmazonAccount `json:"results"`
	TotalCount int32 `json:"totalCount"`
	Links *[]Link `json:"_links,omitempty"`
}

// NewAmazonAccountsPage instantiates a new AmazonAccountsPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonAccountsPage(results []AmazonAccount, totalCount int32) *AmazonAccountsPage {
	this := AmazonAccountsPage{}
	this.Results = results
	this.TotalCount = totalCount
	return &this
}

// NewAmazonAccountsPageWithDefaults instantiates a new AmazonAccountsPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonAccountsPageWithDefaults() *AmazonAccountsPage {
	this := AmazonAccountsPage{}
	return &this
}

// GetResults returns the Results field value
func (o *AmazonAccountsPage) GetResults() []AmazonAccount {
	if o == nil {
		var ret []AmazonAccount
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountsPage) GetResultsOk() (*[]AmazonAccount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *AmazonAccountsPage) SetResults(v []AmazonAccount) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value
func (o *AmazonAccountsPage) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountsPage) GetTotalCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *AmazonAccountsPage) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AmazonAccountsPage) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAccountsPage) GetLinksOk() (*[]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AmazonAccountsPage) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *AmazonAccountsPage) SetLinks(v []Link) {
	o.Links = &v
}

func (o AmazonAccountsPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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

type NullableAmazonAccountsPage struct {
	value *AmazonAccountsPage
	isSet bool
}

func (v NullableAmazonAccountsPage) Get() *AmazonAccountsPage {
	return v.value
}

func (v *NullableAmazonAccountsPage) Set(val *AmazonAccountsPage) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonAccountsPage) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonAccountsPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonAccountsPage(val *AmazonAccountsPage) *NullableAmazonAccountsPage {
	return &NullableAmazonAccountsPage{value: val, isSet: true}
}

func (v NullableAmazonAccountsPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonAccountsPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


