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

// UsersFilters struct for UsersFilters
type UsersFilters struct {
	SearchPattern *string `json:"searchPattern,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Sort *[]UsersSortColumns `json:"sort,omitempty"`
}

// NewUsersFilters instantiates a new UsersFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersFilters() *UsersFilters {
	this := UsersFilters{}
	return &this
}

// NewUsersFiltersWithDefaults instantiates a new UsersFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersFiltersWithDefaults() *UsersFilters {
	this := UsersFilters{}
	return &this
}

// GetSearchPattern returns the SearchPattern field value if set, zero value otherwise.
func (o *UsersFilters) GetSearchPattern() string {
	if o == nil || o.SearchPattern == nil {
		var ret string
		return ret
	}
	return *o.SearchPattern
}

// GetSearchPatternOk returns a tuple with the SearchPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersFilters) GetSearchPatternOk() (*string, bool) {
	if o == nil || o.SearchPattern == nil {
		return nil, false
	}
	return o.SearchPattern, true
}

// HasSearchPattern returns a boolean if a field has been set.
func (o *UsersFilters) HasSearchPattern() bool {
	if o != nil && o.SearchPattern != nil {
		return true
	}

	return false
}

// SetSearchPattern gets a reference to the given string and assigns it to the SearchPattern field.
func (o *UsersFilters) SetSearchPattern(v string) {
	o.SearchPattern = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *UsersFilters) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersFilters) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *UsersFilters) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *UsersFilters) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *UsersFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *UsersFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *UsersFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *UsersFilters) GetSort() []UsersSortColumns {
	if o == nil || o.Sort == nil {
		var ret []UsersSortColumns
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersFilters) GetSortOk() (*[]UsersSortColumns, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *UsersFilters) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given []UsersSortColumns and assigns it to the Sort field.
func (o *UsersFilters) SetSort(v []UsersSortColumns) {
	o.Sort = &v
}

func (o UsersFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SearchPattern != nil {
		toSerialize["searchPattern"] = o.SearchPattern
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	return json.Marshal(toSerialize)
}

type NullableUsersFilters struct {
	value *UsersFilters
	isSet bool
}

func (v NullableUsersFilters) Get() *UsersFilters {
	return v.value
}

func (v *NullableUsersFilters) Set(val *UsersFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersFilters(val *UsersFilters) *NullableUsersFilters {
	return &NullableUsersFilters{value: val, isSet: true}
}

func (v NullableUsersFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


