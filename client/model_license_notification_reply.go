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

// LicenseNotificationReply struct for LicenseNotificationReply
type LicenseNotificationReply struct {
	ReplyId string `json:"replyId"`
	NotificationType LicenseNotificationTypes `json:"notificationType"`
	ReplyText string `json:"replyText"`
}

// NewLicenseNotificationReply instantiates a new LicenseNotificationReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseNotificationReply(replyId string, notificationType LicenseNotificationTypes, replyText string) *LicenseNotificationReply {
	this := LicenseNotificationReply{}
	this.ReplyId = replyId
	this.NotificationType = notificationType
	this.ReplyText = replyText
	return &this
}

// NewLicenseNotificationReplyWithDefaults instantiates a new LicenseNotificationReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseNotificationReplyWithDefaults() *LicenseNotificationReply {
	this := LicenseNotificationReply{}
	return &this
}

// GetReplyId returns the ReplyId field value
func (o *LicenseNotificationReply) GetReplyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReplyId
}

// GetReplyIdOk returns a tuple with the ReplyId field value
// and a boolean to check if the value has been set.
func (o *LicenseNotificationReply) GetReplyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReplyId, true
}

// SetReplyId sets field value
func (o *LicenseNotificationReply) SetReplyId(v string) {
	o.ReplyId = v
}

// GetNotificationType returns the NotificationType field value
func (o *LicenseNotificationReply) GetNotificationType() LicenseNotificationTypes {
	if o == nil {
		var ret LicenseNotificationTypes
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *LicenseNotificationReply) GetNotificationTypeOk() (*LicenseNotificationTypes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *LicenseNotificationReply) SetNotificationType(v LicenseNotificationTypes) {
	o.NotificationType = v
}

// GetReplyText returns the ReplyText field value
func (o *LicenseNotificationReply) GetReplyText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReplyText
}

// GetReplyTextOk returns a tuple with the ReplyText field value
// and a boolean to check if the value has been set.
func (o *LicenseNotificationReply) GetReplyTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReplyText, true
}

// SetReplyText sets field value
func (o *LicenseNotificationReply) SetReplyText(v string) {
	o.ReplyText = v
}

func (o LicenseNotificationReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["replyId"] = o.ReplyId
	}
	if true {
		toSerialize["notificationType"] = o.NotificationType
	}
	if true {
		toSerialize["replyText"] = o.ReplyText
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseNotificationReply struct {
	value *LicenseNotificationReply
	isSet bool
}

func (v NullableLicenseNotificationReply) Get() *LicenseNotificationReply {
	return v.value
}

func (v *NullableLicenseNotificationReply) Set(val *LicenseNotificationReply) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseNotificationReply) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseNotificationReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseNotificationReply(val *LicenseNotificationReply) *NullableLicenseNotificationReply {
	return &NullableLicenseNotificationReply{value: val, isSet: true}
}

func (v NullableLicenseNotificationReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseNotificationReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

