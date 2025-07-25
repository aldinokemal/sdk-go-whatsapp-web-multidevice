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

// checks if the SendChatPresenceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendChatPresenceRequest{}

// SendChatPresenceRequest struct for SendChatPresenceRequest
type SendChatPresenceRequest struct {
	// Phone number with country code
	Phone string `json:"phone"`
	// Action to perform - \"start\" to begin typing indicator, \"stop\" to end typing indicator
	Action string `json:"action"`
}

type _SendChatPresenceRequest SendChatPresenceRequest

// NewSendChatPresenceRequest instantiates a new SendChatPresenceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendChatPresenceRequest(phone string, action string) *SendChatPresenceRequest {
	this := SendChatPresenceRequest{}
	this.Phone = phone
	this.Action = action
	return &this
}

// NewSendChatPresenceRequestWithDefaults instantiates a new SendChatPresenceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendChatPresenceRequestWithDefaults() *SendChatPresenceRequest {
	this := SendChatPresenceRequest{}
	return &this
}

// GetPhone returns the Phone field value
func (o *SendChatPresenceRequest) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *SendChatPresenceRequest) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *SendChatPresenceRequest) SetPhone(v string) {
	o.Phone = v
}

// GetAction returns the Action field value
func (o *SendChatPresenceRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *SendChatPresenceRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *SendChatPresenceRequest) SetAction(v string) {
	o.Action = v
}

func (o SendChatPresenceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendChatPresenceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phone"] = o.Phone
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

func (o *SendChatPresenceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phone",
		"action",
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

	varSendChatPresenceRequest := _SendChatPresenceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendChatPresenceRequest)

	if err != nil {
		return err
	}

	*o = SendChatPresenceRequest(varSendChatPresenceRequest)

	return err
}

type NullableSendChatPresenceRequest struct {
	value *SendChatPresenceRequest
	isSet bool
}

func (v NullableSendChatPresenceRequest) Get() *SendChatPresenceRequest {
	return v.value
}

func (v *NullableSendChatPresenceRequest) Set(val *SendChatPresenceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendChatPresenceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendChatPresenceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendChatPresenceRequest(val *SendChatPresenceRequest) *NullableSendChatPresenceRequest {
	return &NullableSendChatPresenceRequest{value: val, isSet: true}
}

func (v NullableSendChatPresenceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendChatPresenceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


