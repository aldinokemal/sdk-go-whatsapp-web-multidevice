/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 6.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SdkWhatsappWebMultiDevice

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApproveGroupParticipantRequestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApproveGroupParticipantRequestRequest{}

// ApproveGroupParticipantRequestRequest struct for ApproveGroupParticipantRequestRequest
type ApproveGroupParticipantRequestRequest struct {
	// The group ID
	GroupId string `json:"group_id"`
	// Array of participant WhatsApp IDs to approve
	Participants []string `json:"participants"`
}

type _ApproveGroupParticipantRequestRequest ApproveGroupParticipantRequestRequest

// NewApproveGroupParticipantRequestRequest instantiates a new ApproveGroupParticipantRequestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproveGroupParticipantRequestRequest(groupId string, participants []string) *ApproveGroupParticipantRequestRequest {
	this := ApproveGroupParticipantRequestRequest{}
	this.GroupId = groupId
	this.Participants = participants
	return &this
}

// NewApproveGroupParticipantRequestRequestWithDefaults instantiates a new ApproveGroupParticipantRequestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproveGroupParticipantRequestRequestWithDefaults() *ApproveGroupParticipantRequestRequest {
	this := ApproveGroupParticipantRequestRequest{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *ApproveGroupParticipantRequestRequest) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ApproveGroupParticipantRequestRequest) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ApproveGroupParticipantRequestRequest) SetGroupId(v string) {
	o.GroupId = v
}

// GetParticipants returns the Participants field value
func (o *ApproveGroupParticipantRequestRequest) GetParticipants() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value
// and a boolean to check if the value has been set.
func (o *ApproveGroupParticipantRequestRequest) GetParticipantsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Participants, true
}

// SetParticipants sets field value
func (o *ApproveGroupParticipantRequestRequest) SetParticipants(v []string) {
	o.Participants = v
}

func (o ApproveGroupParticipantRequestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApproveGroupParticipantRequestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group_id"] = o.GroupId
	toSerialize["participants"] = o.Participants
	return toSerialize, nil
}

func (o *ApproveGroupParticipantRequestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group_id",
		"participants",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApproveGroupParticipantRequestRequest := _ApproveGroupParticipantRequestRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApproveGroupParticipantRequestRequest)

	if err != nil {
		return err
	}

	*o = ApproveGroupParticipantRequestRequest(varApproveGroupParticipantRequestRequest)

	return err
}

type NullableApproveGroupParticipantRequestRequest struct {
	value *ApproveGroupParticipantRequestRequest
	isSet bool
}

func (v NullableApproveGroupParticipantRequestRequest) Get() *ApproveGroupParticipantRequestRequest {
	return v.value
}

func (v *NullableApproveGroupParticipantRequestRequest) Set(val *ApproveGroupParticipantRequestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApproveGroupParticipantRequestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApproveGroupParticipantRequestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproveGroupParticipantRequestRequest(val *ApproveGroupParticipantRequestRequest) *NullableApproveGroupParticipantRequestRequest {
	return &NullableApproveGroupParticipantRequestRequest{value: val, isSet: true}
}

func (v NullableApproveGroupParticipantRequestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproveGroupParticipantRequestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


