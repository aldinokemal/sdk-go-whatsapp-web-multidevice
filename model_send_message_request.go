/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 5.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SdkWhatsappWebMultiDevice

import (
	"encoding/json"
)

// checks if the SendMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendMessageRequest{}

// SendMessageRequest struct for SendMessageRequest
type SendMessageRequest struct {
	// Phone number with country code
	Phone *string `json:"phone,omitempty"`
	// Message to send
	Message *string `json:"message,omitempty"`
	// Message ID that you want reply
	ReplyMessageId *string `json:"reply_message_id,omitempty"`
}

// NewSendMessageRequest instantiates a new SendMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendMessageRequest() *SendMessageRequest {
	this := SendMessageRequest{}
	return &this
}

// NewSendMessageRequestWithDefaults instantiates a new SendMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendMessageRequestWithDefaults() *SendMessageRequest {
	this := SendMessageRequest{}
	return &this
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *SendMessageRequest) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageRequest) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *SendMessageRequest) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *SendMessageRequest) SetPhone(v string) {
	o.Phone = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SendMessageRequest) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageRequest) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SendMessageRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SendMessageRequest) SetMessage(v string) {
	o.Message = &v
}

// GetReplyMessageId returns the ReplyMessageId field value if set, zero value otherwise.
func (o *SendMessageRequest) GetReplyMessageId() string {
	if o == nil || IsNil(o.ReplyMessageId) {
		var ret string
		return ret
	}
	return *o.ReplyMessageId
}

// GetReplyMessageIdOk returns a tuple with the ReplyMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessageRequest) GetReplyMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyMessageId) {
		return nil, false
	}
	return o.ReplyMessageId, true
}

// HasReplyMessageId returns a boolean if a field has been set.
func (o *SendMessageRequest) HasReplyMessageId() bool {
	if o != nil && !IsNil(o.ReplyMessageId) {
		return true
	}

	return false
}

// SetReplyMessageId gets a reference to the given string and assigns it to the ReplyMessageId field.
func (o *SendMessageRequest) SetReplyMessageId(v string) {
	o.ReplyMessageId = &v
}

func (o SendMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.ReplyMessageId) {
		toSerialize["reply_message_id"] = o.ReplyMessageId
	}
	return toSerialize, nil
}

type NullableSendMessageRequest struct {
	value *SendMessageRequest
	isSet bool
}

func (v NullableSendMessageRequest) Get() *SendMessageRequest {
	return v.value
}

func (v *NullableSendMessageRequest) Set(val *SendMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendMessageRequest(val *SendMessageRequest) *NullableSendMessageRequest {
	return &NullableSendMessageRequest{value: val, isSet: true}
}

func (v NullableSendMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


