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

// BucketsFilters struct for BucketsFilters
type BucketsFilters struct {
	AmazonAccountId *string `json:"AmazonAccountId,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty"`
	Offset *int32 `json:"Offset,omitempty"`
	Limit *int32 `json:"Limit,omitempty"`
	Sort *[]BucketsSortColumns `json:"Sort,omitempty"`
}

// NewBucketsFilters instantiates a new BucketsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketsFilters() *BucketsFilters {
	this := BucketsFilters{}
	return &this
}

// NewBucketsFiltersWithDefaults instantiates a new BucketsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketsFiltersWithDefaults() *BucketsFilters {
	this := BucketsFilters{}
	return &this
}

// GetAmazonAccountId returns the AmazonAccountId field value if set, zero value otherwise.
func (o *BucketsFilters) GetAmazonAccountId() string {
	if o == nil || o.AmazonAccountId == nil {
		var ret string
		return ret
	}
	return *o.AmazonAccountId
}

// GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketsFilters) GetAmazonAccountIdOk() (*string, bool) {
	if o == nil || o.AmazonAccountId == nil {
		return nil, false
	}
	return o.AmazonAccountId, true
}

// HasAmazonAccountId returns a boolean if a field has been set.
func (o *BucketsFilters) HasAmazonAccountId() bool {
	if o != nil && o.AmazonAccountId != nil {
		return true
	}

	return false
}

// SetAmazonAccountId gets a reference to the given string and assigns it to the AmazonAccountId field.
func (o *BucketsFilters) SetAmazonAccountId(v string) {
	o.AmazonAccountId = &v
}

// GetSearchPattern returns the SearchPattern field value if set, zero value otherwise.
func (o *BucketsFilters) GetSearchPattern() string {
	if o == nil || o.SearchPattern == nil {
		var ret string
		return ret
	}
	return *o.SearchPattern
}

// GetSearchPatternOk returns a tuple with the SearchPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketsFilters) GetSearchPatternOk() (*string, bool) {
	if o == nil || o.SearchPattern == nil {
		return nil, false
	}
	return o.SearchPattern, true
}

// HasSearchPattern returns a boolean if a field has been set.
func (o *BucketsFilters) HasSearchPattern() bool {
	if o != nil && o.SearchPattern != nil {
		return true
	}

	return false
}

// SetSearchPattern gets a reference to the given string and assigns it to the SearchPattern field.
func (o *BucketsFilters) SetSearchPattern(v string) {
	o.SearchPattern = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *BucketsFilters) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketsFilters) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *BucketsFilters) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *BucketsFilters) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *BucketsFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketsFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *BucketsFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *BucketsFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *BucketsFilters) GetSort() []BucketsSortColumns {
	if o == nil || o.Sort == nil {
		var ret []BucketsSortColumns
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketsFilters) GetSortOk() (*[]BucketsSortColumns, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *BucketsFilters) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given []BucketsSortColumns and assigns it to the Sort field.
func (o *BucketsFilters) SetSort(v []BucketsSortColumns) {
	o.Sort = &v
}

func (o BucketsFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AmazonAccountId != nil {
		toSerialize["AmazonAccountId"] = o.AmazonAccountId
	}
	if o.SearchPattern != nil {
		toSerialize["SearchPattern"] = o.SearchPattern
	}
	if o.Offset != nil {
		toSerialize["Offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["Limit"] = o.Limit
	}
	if o.Sort != nil {
		toSerialize["Sort"] = o.Sort
	}
	return json.Marshal(toSerialize)
}

type NullableBucketsFilters struct {
	value *BucketsFilters
	isSet bool
}

func (v NullableBucketsFilters) Get() *BucketsFilters {
	return v.value
}

func (v *NullableBucketsFilters) Set(val *BucketsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketsFilters(val *BucketsFilters) *NullableBucketsFilters {
	return &NullableBucketsFilters{value: val, isSet: true}
}

func (v NullableBucketsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


