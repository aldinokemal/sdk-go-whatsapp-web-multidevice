/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 3.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk_go_whatsapp_web_multidevice

import (
	"encoding/json"
)

// checks if the SendResponseResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendResponseResults{}

// SendResponseResults struct for SendResponseResults
type SendResponseResults struct {
	MessageId *string `json:"message_id,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewSendResponseResults instantiates a new SendResponseResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendResponseResults() *SendResponseResults {
	this := SendResponseResults{}
	return &this
}

// NewSendResponseResultsWithDefaults instantiates a new SendResponseResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendResponseResultsWithDefaults() *SendResponseResults {
	this := SendResponseResults{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *SendResponseResults) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendResponseResults) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *SendResponseResults) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *SendResponseResults) SetMessageId(v string) {
	o.MessageId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SendResponseResults) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendResponseResults) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SendResponseResults) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SendResponseResults) SetStatus(v string) {
	o.Status = &v
}

func (o SendResponseResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendResponseResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageId) {
		toSerialize["message_id"] = o.MessageId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableSendResponseResults struct {
	value *SendResponseResults
	isSet bool
}

func (v NullableSendResponseResults) Get() *SendResponseResults {
	return v.value
}

func (v *NullableSendResponseResults) Set(val *SendResponseResults) {
	v.value = val
	v.isSet = true
}

func (v NullableSendResponseResults) IsSet() bool {
	return v.isSet
}

func (v *NullableSendResponseResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendResponseResults(val *SendResponseResults) *NullableSendResponseResults {
	return &NullableSendResponseResults{value: val, isSet: true}
}

func (v NullableSendResponseResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendResponseResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


