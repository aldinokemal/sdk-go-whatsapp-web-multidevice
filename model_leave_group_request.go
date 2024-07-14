/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 4.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk_go_whatsapp_web_multidevice

import (
	"encoding/json"
)

// checks if the LeaveGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LeaveGroupRequest{}

// LeaveGroupRequest struct for LeaveGroupRequest
type LeaveGroupRequest struct {
	GroupId *string `json:"group_id,omitempty"`
}

// NewLeaveGroupRequest instantiates a new LeaveGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLeaveGroupRequest() *LeaveGroupRequest {
	this := LeaveGroupRequest{}
	return &this
}

// NewLeaveGroupRequestWithDefaults instantiates a new LeaveGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeaveGroupRequestWithDefaults() *LeaveGroupRequest {
	this := LeaveGroupRequest{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *LeaveGroupRequest) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeaveGroupRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *LeaveGroupRequest) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *LeaveGroupRequest) SetGroupId(v string) {
	o.GroupId = &v
}

func (o LeaveGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LeaveGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	return toSerialize, nil
}

type NullableLeaveGroupRequest struct {
	value *LeaveGroupRequest
	isSet bool
}

func (v NullableLeaveGroupRequest) Get() *LeaveGroupRequest {
	return v.value
}

func (v *NullableLeaveGroupRequest) Set(val *LeaveGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLeaveGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLeaveGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeaveGroupRequest(val *LeaveGroupRequest) *NullableLeaveGroupRequest {
	return &NullableLeaveGroupRequest{value: val, isSet: true}
}

func (v NullableLeaveGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeaveGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


