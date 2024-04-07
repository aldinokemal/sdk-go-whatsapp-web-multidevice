/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 3.10.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk_go_whatsapp_web_multidevice

import (
	"encoding/json"
)

// checks if the CreateGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGroupRequest{}

// CreateGroupRequest struct for CreateGroupRequest
type CreateGroupRequest struct {
	Title *string `json:"title,omitempty"`
	Participants []string `json:"participants,omitempty"`
}

// NewCreateGroupRequest instantiates a new CreateGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupRequest() *CreateGroupRequest {
	this := CreateGroupRequest{}
	return &this
}

// NewCreateGroupRequestWithDefaults instantiates a new CreateGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupRequestWithDefaults() *CreateGroupRequest {
	this := CreateGroupRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreateGroupRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateGroupRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreateGroupRequest) SetTitle(v string) {
	o.Title = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *CreateGroupRequest) GetParticipants() []string {
	if o == nil || IsNil(o.Participants) {
		var ret []string
		return ret
	}
	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupRequest) GetParticipantsOk() ([]string, bool) {
	if o == nil || IsNil(o.Participants) {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *CreateGroupRequest) HasParticipants() bool {
	if o != nil && !IsNil(o.Participants) {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []string and assigns it to the Participants field.
func (o *CreateGroupRequest) SetParticipants(v []string) {
	o.Participants = v
}

func (o CreateGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Participants) {
		toSerialize["participants"] = o.Participants
	}
	return toSerialize, nil
}

type NullableCreateGroupRequest struct {
	value *CreateGroupRequest
	isSet bool
}

func (v NullableCreateGroupRequest) Get() *CreateGroupRequest {
	return v.value
}

func (v *NullableCreateGroupRequest) Set(val *CreateGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupRequest(val *CreateGroupRequest) *NullableCreateGroupRequest {
	return &NullableCreateGroupRequest{value: val, isSet: true}
}

func (v NullableCreateGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


