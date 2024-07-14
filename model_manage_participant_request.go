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

// checks if the ManageParticipantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManageParticipantRequest{}

// ManageParticipantRequest struct for ManageParticipantRequest
type ManageParticipantRequest struct {
	GroupId *string `json:"group_id,omitempty"`
	Participants []string `json:"participants,omitempty"`
}

// NewManageParticipantRequest instantiates a new ManageParticipantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageParticipantRequest() *ManageParticipantRequest {
	this := ManageParticipantRequest{}
	return &this
}

// NewManageParticipantRequestWithDefaults instantiates a new ManageParticipantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageParticipantRequestWithDefaults() *ManageParticipantRequest {
	this := ManageParticipantRequest{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ManageParticipantRequest) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageParticipantRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ManageParticipantRequest) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ManageParticipantRequest) SetGroupId(v string) {
	o.GroupId = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *ManageParticipantRequest) GetParticipants() []string {
	if o == nil || IsNil(o.Participants) {
		var ret []string
		return ret
	}
	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageParticipantRequest) GetParticipantsOk() ([]string, bool) {
	if o == nil || IsNil(o.Participants) {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *ManageParticipantRequest) HasParticipants() bool {
	if o != nil && !IsNil(o.Participants) {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []string and assigns it to the Participants field.
func (o *ManageParticipantRequest) SetParticipants(v []string) {
	o.Participants = v
}

func (o ManageParticipantRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManageParticipantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.Participants) {
		toSerialize["participants"] = o.Participants
	}
	return toSerialize, nil
}

type NullableManageParticipantRequest struct {
	value *ManageParticipantRequest
	isSet bool
}

func (v NullableManageParticipantRequest) Get() *ManageParticipantRequest {
	return v.value
}

func (v *NullableManageParticipantRequest) Set(val *ManageParticipantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableManageParticipantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableManageParticipantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageParticipantRequest(val *ManageParticipantRequest) *NullableManageParticipantRequest {
	return &NullableManageParticipantRequest{value: val, isSet: true}
}

func (v NullableManageParticipantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageParticipantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


