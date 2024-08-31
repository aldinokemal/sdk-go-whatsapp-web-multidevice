/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 4.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk_go_whatsapp_web_multidevice

import (
	"encoding/json"
)

// checks if the ManageParticipantResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManageParticipantResponseResultsInner{}

// ManageParticipantResponseResultsInner struct for ManageParticipantResponseResultsInner
type ManageParticipantResponseResultsInner struct {
	Participant *string `json:"participant,omitempty"`
	Status *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewManageParticipantResponseResultsInner instantiates a new ManageParticipantResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageParticipantResponseResultsInner() *ManageParticipantResponseResultsInner {
	this := ManageParticipantResponseResultsInner{}
	return &this
}

// NewManageParticipantResponseResultsInnerWithDefaults instantiates a new ManageParticipantResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageParticipantResponseResultsInnerWithDefaults() *ManageParticipantResponseResultsInner {
	this := ManageParticipantResponseResultsInner{}
	return &this
}

// GetParticipant returns the Participant field value if set, zero value otherwise.
func (o *ManageParticipantResponseResultsInner) GetParticipant() string {
	if o == nil || IsNil(o.Participant) {
		var ret string
		return ret
	}
	return *o.Participant
}

// GetParticipantOk returns a tuple with the Participant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageParticipantResponseResultsInner) GetParticipantOk() (*string, bool) {
	if o == nil || IsNil(o.Participant) {
		return nil, false
	}
	return o.Participant, true
}

// HasParticipant returns a boolean if a field has been set.
func (o *ManageParticipantResponseResultsInner) HasParticipant() bool {
	if o != nil && !IsNil(o.Participant) {
		return true
	}

	return false
}

// SetParticipant gets a reference to the given string and assigns it to the Participant field.
func (o *ManageParticipantResponseResultsInner) SetParticipant(v string) {
	o.Participant = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ManageParticipantResponseResultsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageParticipantResponseResultsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ManageParticipantResponseResultsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ManageParticipantResponseResultsInner) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ManageParticipantResponseResultsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageParticipantResponseResultsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ManageParticipantResponseResultsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ManageParticipantResponseResultsInner) SetMessage(v string) {
	o.Message = &v
}

func (o ManageParticipantResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManageParticipantResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Participant) {
		toSerialize["participant"] = o.Participant
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableManageParticipantResponseResultsInner struct {
	value *ManageParticipantResponseResultsInner
	isSet bool
}

func (v NullableManageParticipantResponseResultsInner) Get() *ManageParticipantResponseResultsInner {
	return v.value
}

func (v *NullableManageParticipantResponseResultsInner) Set(val *ManageParticipantResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableManageParticipantResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableManageParticipantResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageParticipantResponseResultsInner(val *ManageParticipantResponseResultsInner) *NullableManageParticipantResponseResultsInner {
	return &NullableManageParticipantResponseResultsInner{value: val, isSet: true}
}

func (v NullableManageParticipantResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageParticipantResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


