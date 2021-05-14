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

// BucketFolder struct for BucketFolder
type BucketFolder struct {
	Name string `json:"name"`
	FolderId string `json:"folderId"`
}

// NewBucketFolder instantiates a new BucketFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketFolder(name string, folderId string) *BucketFolder {
	this := BucketFolder{}
	this.Name = name
	this.FolderId = folderId
	return &this
}

// NewBucketFolderWithDefaults instantiates a new BucketFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketFolderWithDefaults() *BucketFolder {
	this := BucketFolder{}
	return &this
}

// GetName returns the Name field value
func (o *BucketFolder) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BucketFolder) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BucketFolder) SetName(v string) {
	o.Name = v
}

// GetFolderId returns the FolderId field value
func (o *BucketFolder) GetFolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
func (o *BucketFolder) GetFolderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FolderId, true
}

// SetFolderId sets field value
func (o *BucketFolder) SetFolderId(v string) {
	o.FolderId = v
}

func (o BucketFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["folderId"] = o.FolderId
	}
	return json.Marshal(toSerialize)
}

type NullableBucketFolder struct {
	value *BucketFolder
	isSet bool
}

func (v NullableBucketFolder) Get() *BucketFolder {
	return v.value
}

func (v *NullableBucketFolder) Set(val *BucketFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketFolder(val *BucketFolder) *NullableBucketFolder {
	return &NullableBucketFolder{value: val, isSet: true}
}

func (v NullableBucketFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


