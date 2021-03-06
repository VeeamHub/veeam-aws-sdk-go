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

// PolicyEmbeddedResources struct for PolicyEmbeddedResources
type PolicyEmbeddedResources struct {
	LastBackupSession *Link `json:"lastBackupSession,omitempty"`
	LastFinishedBackupSession *Link `json:"lastFinishedBackupSession,omitempty"`
	TargetRepository *Link `json:"targetRepository,omitempty"`
	ProtectedResources *Link `json:"protectedResources,omitempty"`
	ExcludedResources *Link `json:"excludedResources,omitempty"`
	Regions *Link `json:"regions,omitempty"`
}

// NewPolicyEmbeddedResources instantiates a new PolicyEmbeddedResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEmbeddedResources() *PolicyEmbeddedResources {
	this := PolicyEmbeddedResources{}
	return &this
}

// NewPolicyEmbeddedResourcesWithDefaults instantiates a new PolicyEmbeddedResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEmbeddedResourcesWithDefaults() *PolicyEmbeddedResources {
	this := PolicyEmbeddedResources{}
	return &this
}

// GetLastBackupSession returns the LastBackupSession field value if set, zero value otherwise.
func (o *PolicyEmbeddedResources) GetLastBackupSession() Link {
	if o == nil || o.LastBackupSession == nil {
		var ret Link
		return ret
	}
	return *o.LastBackupSession
}

// GetLastBackupSessionOk returns a tuple with the LastBackupSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEmbeddedResources) GetLastBackupSessionOk() (*Link, bool) {
	if o == nil || o.LastBackupSession == nil {
		return nil, false
	}
	return o.LastBackupSession, true
}

// HasLastBackupSession returns a boolean if a field has been set.
func (o *PolicyEmbeddedResources) HasLastBackupSession() bool {
	if o != nil && o.LastBackupSession != nil {
		return true
	}

	return false
}

// SetLastBackupSession gets a reference to the given Link and assigns it to the LastBackupSession field.
func (o *PolicyEmbeddedResources) SetLastBackupSession(v Link) {
	o.LastBackupSession = &v
}

// GetLastFinishedBackupSession returns the LastFinishedBackupSession field value if set, zero value otherwise.
func (o *PolicyEmbeddedResources) GetLastFinishedBackupSession() Link {
	if o == nil || o.LastFinishedBackupSession == nil {
		var ret Link
		return ret
	}
	return *o.LastFinishedBackupSession
}

// GetLastFinishedBackupSessionOk returns a tuple with the LastFinishedBackupSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEmbeddedResources) GetLastFinishedBackupSessionOk() (*Link, bool) {
	if o == nil || o.LastFinishedBackupSession == nil {
		return nil, false
	}
	return o.LastFinishedBackupSession, true
}

// HasLastFinishedBackupSession returns a boolean if a field has been set.
func (o *PolicyEmbeddedResources) HasLastFinishedBackupSession() bool {
	if o != nil && o.LastFinishedBackupSession != nil {
		return true
	}

	return false
}

// SetLastFinishedBackupSession gets a reference to the given Link and assigns it to the LastFinishedBackupSession field.
func (o *PolicyEmbeddedResources) SetLastFinishedBackupSession(v Link) {
	o.LastFinishedBackupSession = &v
}

// GetTargetRepository returns the TargetRepository field value if set, zero value otherwise.
func (o *PolicyEmbeddedResources) GetTargetRepository() Link {
	if o == nil || o.TargetRepository == nil {
		var ret Link
		return ret
	}
	return *o.TargetRepository
}

// GetTargetRepositoryOk returns a tuple with the TargetRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEmbeddedResources) GetTargetRepositoryOk() (*Link, bool) {
	if o == nil || o.TargetRepository == nil {
		return nil, false
	}
	return o.TargetRepository, true
}

// HasTargetRepository returns a boolean if a field has been set.
func (o *PolicyEmbeddedResources) HasTargetRepository() bool {
	if o != nil && o.TargetRepository != nil {
		return true
	}

	return false
}

// SetTargetRepository gets a reference to the given Link and assigns it to the TargetRepository field.
func (o *PolicyEmbeddedResources) SetTargetRepository(v Link) {
	o.TargetRepository = &v
}

// GetProtectedResources returns the ProtectedResources field value if set, zero value otherwise.
func (o *PolicyEmbeddedResources) GetProtectedResources() Link {
	if o == nil || o.ProtectedResources == nil {
		var ret Link
		return ret
	}
	return *o.ProtectedResources
}

// GetProtectedResourcesOk returns a tuple with the ProtectedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEmbeddedResources) GetProtectedResourcesOk() (*Link, bool) {
	if o == nil || o.ProtectedResources == nil {
		return nil, false
	}
	return o.ProtectedResources, true
}

// HasProtectedResources returns a boolean if a field has been set.
func (o *PolicyEmbeddedResources) HasProtectedResources() bool {
	if o != nil && o.ProtectedResources != nil {
		return true
	}

	return false
}

// SetProtectedResources gets a reference to the given Link and assigns it to the ProtectedResources field.
func (o *PolicyEmbeddedResources) SetProtectedResources(v Link) {
	o.ProtectedResources = &v
}

// GetExcludedResources returns the ExcludedResources field value if set, zero value otherwise.
func (o *PolicyEmbeddedResources) GetExcludedResources() Link {
	if o == nil || o.ExcludedResources == nil {
		var ret Link
		return ret
	}
	return *o.ExcludedResources
}

// GetExcludedResourcesOk returns a tuple with the ExcludedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEmbeddedResources) GetExcludedResourcesOk() (*Link, bool) {
	if o == nil || o.ExcludedResources == nil {
		return nil, false
	}
	return o.ExcludedResources, true
}

// HasExcludedResources returns a boolean if a field has been set.
func (o *PolicyEmbeddedResources) HasExcludedResources() bool {
	if o != nil && o.ExcludedResources != nil {
		return true
	}

	return false
}

// SetExcludedResources gets a reference to the given Link and assigns it to the ExcludedResources field.
func (o *PolicyEmbeddedResources) SetExcludedResources(v Link) {
	o.ExcludedResources = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *PolicyEmbeddedResources) GetRegions() Link {
	if o == nil || o.Regions == nil {
		var ret Link
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEmbeddedResources) GetRegionsOk() (*Link, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *PolicyEmbeddedResources) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given Link and assigns it to the Regions field.
func (o *PolicyEmbeddedResources) SetRegions(v Link) {
	o.Regions = &v
}

func (o PolicyEmbeddedResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastBackupSession != nil {
		toSerialize["lastBackupSession"] = o.LastBackupSession
	}
	if o.LastFinishedBackupSession != nil {
		toSerialize["lastFinishedBackupSession"] = o.LastFinishedBackupSession
	}
	if o.TargetRepository != nil {
		toSerialize["targetRepository"] = o.TargetRepository
	}
	if o.ProtectedResources != nil {
		toSerialize["protectedResources"] = o.ProtectedResources
	}
	if o.ExcludedResources != nil {
		toSerialize["excludedResources"] = o.ExcludedResources
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyEmbeddedResources struct {
	value *PolicyEmbeddedResources
	isSet bool
}

func (v NullablePolicyEmbeddedResources) Get() *PolicyEmbeddedResources {
	return v.value
}

func (v *NullablePolicyEmbeddedResources) Set(val *PolicyEmbeddedResources) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEmbeddedResources) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEmbeddedResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEmbeddedResources(val *PolicyEmbeddedResources) *NullablePolicyEmbeddedResources {
	return &NullablePolicyEmbeddedResources{value: val, isSet: true}
}

func (v NullablePolicyEmbeddedResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEmbeddedResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


