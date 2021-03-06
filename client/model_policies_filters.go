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

// PoliciesFilters struct for PoliciesFilters
type PoliciesFilters struct {
	SearchPattern *string `json:"SearchPattern,omitempty"`
	VirtualMachineId *string `json:"VirtualMachineId,omitempty"`
	LastPolicySessionStatus *[]SessionStatuses `json:"LastPolicySessionStatus,omitempty"`
	PolicyStatus *[]PolicyStatuses `json:"PolicyStatus,omitempty"`
	Sort *[]PoliciesSortColumns `json:"Sort,omitempty"`
	Usn *int64 `json:"Usn,omitempty"`
	RepositoryId *string `json:"RepositoryId,omitempty"`
	Offset *int32 `json:"Offset,omitempty"`
	Limit *int32 `json:"Limit,omitempty"`
}

// NewPoliciesFilters instantiates a new PoliciesFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoliciesFilters() *PoliciesFilters {
	this := PoliciesFilters{}
	return &this
}

// NewPoliciesFiltersWithDefaults instantiates a new PoliciesFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesFiltersWithDefaults() *PoliciesFilters {
	this := PoliciesFilters{}
	return &this
}

// GetSearchPattern returns the SearchPattern field value if set, zero value otherwise.
func (o *PoliciesFilters) GetSearchPattern() string {
	if o == nil || o.SearchPattern == nil {
		var ret string
		return ret
	}
	return *o.SearchPattern
}

// GetSearchPatternOk returns a tuple with the SearchPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesFilters) GetSearchPatternOk() (*string, bool) {
	if o == nil || o.SearchPattern == nil {
		return nil, false
	}
	return o.SearchPattern, true
}

// HasSearchPattern returns a boolean if a field has been set.
func (o *PoliciesFilters) HasSearchPattern() bool {
	if o != nil && o.SearchPattern != nil {
		return true
	}

	return false
}

// SetSearchPattern gets a reference to the given string and assigns it to the SearchPattern field.
func (o *PoliciesFilters) SetSearchPattern(v string) {
	o.SearchPattern = &v
}

// GetVirtualMachineId returns the VirtualMachineId field value if set, zero value otherwise.
func (o *PoliciesFilters) GetVirtualMachineId() string {
	if o == nil || o.VirtualMachineId == nil {
		var ret string
		return ret
	}
	return *o.VirtualMachineId
}

// GetVirtualMachineIdOk returns a tuple with the VirtualMachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesFilters) GetVirtualMachineIdOk() (*string, bool) {
	if o == nil || o.VirtualMachineId == nil {
		return nil, false
	}
	return o.VirtualMachineId, true
}

// HasVirtualMachineId returns a boolean if a field has been set.
func (o *PoliciesFilters) HasVirtualMachineId() bool {
	if o != nil && o.VirtualMachineId != nil {
		return true
	}

	return false
}

// SetVirtualMachineId gets a reference to the given string and assigns it to the VirtualMachineId field.
func (o *PoliciesFilters) SetVirtualMachineId(v string) {
	o.VirtualMachineId = &v
}

// GetLastPolicySessionStatus returns the LastPolicySessionStatus field value if set, zero value otherwise.
func (o *PoliciesFilters) GetLastPolicySessionStatus() []SessionStatuses {
	if o == nil || o.LastPolicySessionStatus == nil {
		var ret []SessionStatuses
		return ret
	}
	return *o.LastPolicySessionStatus
}

// GetLastPolicySessionStatusOk returns a tuple with the LastPolicySessionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesFilters) GetLastPolicySessionStatusOk() (*[]SessionStatuses, bool) {
	if o == nil || o.LastPolicySessionStatus == nil {
		return nil, false
	}
	return o.LastPolicySessionStatus, true
}

// HasLastPolicySessionStatus returns a boolean if a field has been set.
func (o *PoliciesFilters) HasLastPolicySessionStatus() bool {
	if o != nil && o.LastPolicySessionStatus != nil {
		return true
	}

	return false
}

// SetLastPolicySessionStatus gets a reference to the given []SessionStatuses and assigns it to the LastPolicySessionStatus field.
func (o *PoliciesFilters) SetLastPolicySessionStatus(v []SessionStatuses) {
	o.LastPolicySessionStatus = &v
}

// GetPolicyStatus returns the PolicyStatus field value if set, zero value otherwise.
func (o *PoliciesFilters) GetPolicyStatus() []PolicyStatuses {
	if o == nil || o.PolicyStatus == nil {
		var ret []PolicyStatuses
		return ret
	}
	return *o.PolicyStatus
}

// GetPolicyStatusOk returns a tuple with the PolicyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesFilters) GetPolicyStatusOk() (*[]PolicyStatuses, bool) {
	if o == nil || o.PolicyStatus == nil {
		return nil, false
	}
	return o.PolicyStatus, true
}

// HasPolicyStatus returns a boolean if a field has been set.
func (o *PoliciesFilters) HasPolicyStatus() bool {
	if o != nil && o.PolicyStatus != nil {
		return true
	}

	return false
}

// SetPolicyStatus gets a reference to the given []PolicyStatuses and assigns it to the PolicyStatus field.
func (o *PoliciesFilters) SetPolicyStatus(v []PolicyStatuses) {
	o.PolicyStatus = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *PoliciesFilters) GetSort() []PoliciesSortColumns {
	if o == nil || o.Sort == nil {
		var ret []PoliciesSortColumns
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesFilters) GetSortOk() (*[]PoliciesSortColumns, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *PoliciesFilters) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given []PoliciesSortColumns and assigns it to the Sort field.
func (o *PoliciesFilters) SetSort(v []PoliciesSortColumns) {
	o.Sort = &v
}

// GetUsn returns the Usn field value if set, zero value otherwise.
func (o *PoliciesFilters) GetUsn() int64 {
	if o == nil || o.Usn == nil {
		var ret int64
		return ret
	}
	return *o.Usn
}

// GetUsnOk returns a tuple with the Usn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesFilters) GetUsnOk() (*int64, bool) {
	if o == nil || o.Usn == nil {
		return nil, false
	}
	return o.Usn, true
}

// HasUsn returns a boolean if a field has been set.
func (o *PoliciesFilters) HasUsn() bool {
	if o != nil && o.Usn != nil {
		return true
	}

	return false
}

// SetUsn gets a reference to the given int64 and assigns it to the Usn field.
func (o *PoliciesFilters) SetUsn(v int64) {
	o.Usn = &v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *PoliciesFilters) GetRepositoryId() string {
	if o == nil || o.RepositoryId == nil {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesFilters) GetRepositoryIdOk() (*string, bool) {
	if o == nil || o.RepositoryId == nil {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *PoliciesFilters) HasRepositoryId() bool {
	if o != nil && o.RepositoryId != nil {
		return true
	}

	return false
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *PoliciesFilters) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *PoliciesFilters) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesFilters) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *PoliciesFilters) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *PoliciesFilters) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PoliciesFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoliciesFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PoliciesFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PoliciesFilters) SetLimit(v int32) {
	o.Limit = &v
}

func (o PoliciesFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SearchPattern != nil {
		toSerialize["SearchPattern"] = o.SearchPattern
	}
	if o.VirtualMachineId != nil {
		toSerialize["VirtualMachineId"] = o.VirtualMachineId
	}
	if o.LastPolicySessionStatus != nil {
		toSerialize["LastPolicySessionStatus"] = o.LastPolicySessionStatus
	}
	if o.PolicyStatus != nil {
		toSerialize["PolicyStatus"] = o.PolicyStatus
	}
	if o.Sort != nil {
		toSerialize["Sort"] = o.Sort
	}
	if o.Usn != nil {
		toSerialize["Usn"] = o.Usn
	}
	if o.RepositoryId != nil {
		toSerialize["RepositoryId"] = o.RepositoryId
	}
	if o.Offset != nil {
		toSerialize["Offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["Limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullablePoliciesFilters struct {
	value *PoliciesFilters
	isSet bool
}

func (v NullablePoliciesFilters) Get() *PoliciesFilters {
	return v.value
}

func (v *NullablePoliciesFilters) Set(val *PoliciesFilters) {
	v.value = val
	v.isSet = true
}

func (v NullablePoliciesFilters) IsSet() bool {
	return v.isSet
}

func (v *NullablePoliciesFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoliciesFilters(val *PoliciesFilters) *NullablePoliciesFilters {
	return &NullablePoliciesFilters{value: val, isSet: true}
}

func (v NullablePoliciesFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoliciesFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


